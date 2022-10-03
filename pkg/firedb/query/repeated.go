package query

import (
	"io"

	"github.com/segmentio/parquet-go"

	"github.com/grafana/fire/pkg/iter"
)

type RepeatedRow[T any] struct {
	Row    T
	Values []parquet.Value
}

type repeatedPageIterator[T any] struct {
	rows     iter.Iterator[T]
	column   int
	readSize int

	rgs                 []parquet.RowGroup
	startRowGroupRowNum int64

	currentPage     parquet.Page
	startPageRowNum int64

	pageNextRowNum int64

	currentPages parquet.Pages
	valueReader  parquet.ValueReader

	rowFinished    bool
	skipping       bool
	firstRow       bool
	err            error
	done           bool // because we advance the iterator to seek in advance we remember if we are done
	currentValue   *RepeatedRow[T]
	buffer         []parquet.Value
	originalBuffer []parquet.Value
}

// NewRepeatedPageIterator returns an iterator that iterates over the repeated values in a column.
// The iterator can only seek forward and so rows should be sorted by row number.
func NewRepeatedPageIterator[T any](
	rows iter.Iterator[T],
	rgs []parquet.RowGroup,
	column int,
	readSize int,
) iter.Iterator[*RepeatedRow[T]] {
	if readSize <= 0 {
		panic("readSize must be greater than 0")
	}
	buffer := make([]parquet.Value, readSize)
	done := !rows.Next()
	return &repeatedPageIterator[T]{
		rows:           rows,
		rgs:            rgs,
		column:         column,
		readSize:       readSize,
		buffer:         buffer[:0],
		originalBuffer: buffer,
		currentValue:   &RepeatedRow[T]{},
		done:           done,
		firstRow:       true,
		rowFinished:    true,
		skipping:       false,
	}
}

// seekRowNum the row num to seek to.
func (it *repeatedPageIterator[T]) seekRowNum() int64 {
	return any(it.rows.At()).(RowGetter).RowNumber()
}

func (it *repeatedPageIterator[T]) Next() bool {
Outer:
	for {
		if it.done {
			return false
		}
		for len(it.rgs) != 0 && (it.seekRowNum() >= (it.startRowGroupRowNum + it.rgs[0].NumRows())) {
			if !it.closeCurrentPages() {
				return false
			}
			it.startRowGroupRowNum += it.rgs[0].NumRows()
			it.rgs = it.rgs[1:]
		}
		if len(it.rgs) == 0 {
			return false
		}
		if it.currentPages == nil {
			it.currentPages = it.rgs[0].ColumnChunks()[it.column].Pages()
		}
		// read a new page.
		if it.currentPage == nil {
			// SeekToRow seek across and within pages. So the next position in the page will the be the row.
			seekTo := it.seekRowNum() - it.startRowGroupRowNum
			if err := it.currentPages.SeekToRow(seekTo); err != nil {
				it.err = err
				it.currentPages = nil // we can set it to nil since somehow it was closed.
				return false
			}
			it.startPageRowNum = it.seekRowNum()
			it.pageNextRowNum = 0
			it.buffer = it.buffer[:0]
			it.firstRow = true
			it.rowFinished = true
			var err error
			it.currentPage, err = it.currentPages.ReadPage()
			if err != nil {
				if err == io.EOF {
					continue
				}
				it.err = err
				return false
			}
			it.valueReader = it.currentPage.Values()
		}
		// if there's no more value in that page we can skip it.
		if it.seekRowNum() >= it.startPageRowNum+it.currentPage.NumRows() {
			it.currentPage = nil
			continue
		}

		// only read values if the buffer is empty
		if len(it.buffer) == 0 {
			// reading values....
			it.buffer = it.originalBuffer
			n, err := it.valueReader.ReadValues(it.buffer)
			if err != nil && err != io.EOF {
				it.err = err
				return false
			}
			it.buffer = it.buffer[:n]
			// no more buffer, move to next page
			if len(it.buffer) == 0 {
				it.done = !it.rows.Next() // if the page has no more data the current row is over.
				it.currentPage = nil
				continue
			}
		}

		// we have data in the buffer.
		it.currentValue.Row = it.rows.At()
		start, next, ok := it.readNextRow()
		if ok && it.rowFinished {
			if it.seekRowNum() > it.startPageRowNum+it.pageNextRowNum {
				it.pageNextRowNum++
				it.buffer = it.buffer[next:]
				continue Outer
			}
			it.pageNextRowNum++
			it.currentValue.Values = it.buffer[:next]
			it.buffer = it.buffer[next:] // consume the values.
			it.done = !it.rows.Next()
			return true
		}
		// we read a partial row or we're skipping a row.
		if it.rowFinished || it.skipping {
			it.rowFinished = false
			// skip until we find the next row.
			if it.seekRowNum() > it.startPageRowNum+it.pageNextRowNum {
				last := it.buffer[start].RepetitionLevel()
				if it.skipping && last == 0 {
					it.buffer = it.buffer[start:]
					it.pageNextRowNum++
					it.skipping = false
					it.rowFinished = true
				} else {
					if start != 0 {
						next = start + 1
					}
					it.buffer = it.buffer[next:]
					it.skipping = true
				}
				continue Outer
			}
			it.currentValue.Values = it.buffer[:next]
			it.buffer = it.buffer[next:] // consume the values.
			return true
		}
		// this is the start of a new row.
		if !it.rowFinished && it.buffer[start].RepetitionLevel() == 0 {
			// consume values up to the new start if there is
			if start >= 1 {
				it.currentValue.Values = it.buffer[:start]
				it.buffer = it.buffer[start:] // consume the values.
				return true
			}
			// or move to the next row.
			it.pageNextRowNum++
			it.done = !it.rows.Next()
			it.rowFinished = true
			continue Outer
		}
		it.currentValue.Values = it.buffer[:next]
		it.buffer = it.buffer[next:] // consume the values.
		return true
	}
}

func (it *repeatedPageIterator[T]) readNextRow() (int, int, bool) {
	start := 0
	foundStart := false
	for i, v := range it.buffer {
		if v.RepetitionLevel() == 0 && !foundStart {
			foundStart = true
			start = i
			continue
		}
		if v.RepetitionLevel() == 0 && foundStart {
			return start, i, true
		}
	}
	return start, len(it.buffer), false
}

func (it *repeatedPageIterator[T]) closeCurrentPages() bool {
	if it.currentPages != nil {
		if err := it.currentPages.Close(); err != nil {
			it.err = err
			it.currentPages = nil
			return false
		}
		it.currentPages = nil
	}
	return true
}

// At returns the current value.
// Only valid after a call to Next.
// The returned value is reused on the next call to Next and should not be retained.
func (it *repeatedPageIterator[T]) At() *RepeatedRow[T] {
	return it.currentValue
}

func (it *repeatedPageIterator[T]) Err() error {
	return it.err
}

func (it *repeatedPageIterator[T]) Close() error {
	if it.currentPages != nil {
		if err := it.currentPages.Close(); err != nil {
			return err
		}
		it.currentPages = nil
	}
	return nil
}
