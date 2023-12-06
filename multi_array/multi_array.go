package multiarray

import (
	"fmt"
	"strings"
)

type TwoDArray[T any] struct {
	data [][]T
}

func New[T any](rows, columns int) *TwoDArray[T] {
	twoDArray := &TwoDArray[T]{
		data: make([][]T, rows),
	}

	for i := range twoDArray.data {
		twoDArray.data[i] = make([]T, columns)
	}

	return twoDArray
}

func (t *TwoDArray[T]) Rows() int {
	return len(t.data)
}

func (t *TwoDArray[T]) Cols() int {
	return len(t.data[0])
}

func (t *TwoDArray[T]) Set(row, col int, value T) {
	t.data[row][col] = value
}

func (t *TwoDArray[T]) Get(row, col int) T {
	return t.data[row][col]
}

func (t *TwoDArray[T]) String() string {
	resultBuilder := strings.Builder{}

	for i := range t.data {
		for j := range t.data[i] {
			stringValue := fmt.Sprintf("%v ", t.data[i][j])
			resultBuilder.WriteString(stringValue)
		}
		resultBuilder.WriteString("\n")
	}

	return resultBuilder.String()
}

func (t *TwoDArray[T]) Fill(value T) {
	for row := range t.data {
		for col := range t.data[row] {
			t.Set(row, col, value)
		}
	}
}

func (t *TwoDArray[T]) FillRow(row int, value T) {
	for i := range t.data[row] {
		t.Set(row, i, value)
	}
}

func (t *TwoDArray[T]) FillCol(col int, value T) {
	for rowIndex := range t.data {
		t.Set(rowIndex, col, value)
	}
}

func (t *TwoDArray[T]) Row(row int) []T {
	return t.data[row]
}

func (t *TwoDArray[T]) Col(col int) []T {
	returnData := make([]T, len(t.data), len(t.data))
	for rowIndex := range t.data {
		returnData[rowIndex] = t.data[rowIndex][col]
	}

	return returnData
}
