package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	fmt.Println(Relu(7))
	fmt.Println(Relu(-1))
	fmt.Println(Relu(1.2))
	fmt.Println(Relu(time.February))

	m, err := NewMatrix[float64](10, 3)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println("m:", m)
	fmt.Println(m.At(3,2))

	fmt.Println(Max([]int{3,1,2})) // 3 <nil>
	fmt.Println(Max([]float64{3,1,2})) // 3 <nil>
	fmt.Println(Max[int](nil)) // 3 <nil>
}

func Max[T Number](values []T) (T, error) {
	if len(values) == 0 {
		var zero T // Return zero for any type. works for strings
		return zero, errors.New("Max of empty slice")
	}

	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}

	return m, nil
}

func (m *Matrix[T]) At(row, col int) T {
	i := (row * m.Cols) + col
	return m.data[i]
}

func NewMatrix[T Number](rows, cols int) (*Matrix[T], error) {
	if rows <= 0 || cols <= 0 {
		return nil, fmt.Errorf("bad dimensions: %d/%d", rows, cols)
	}

	m := Matrix[T]{
		Rows: rows,
		Cols: cols,
		data: make([]T, rows*cols),
	}

	return &m, nil
}

type Matrix[T Number]struct {
	Rows int
	Cols int

	data []T
}

type Number interface {
	~int | ~float64
}

// T is a "type constraint" (not a new type)
// ~ checks if type is the underlying type in memory then work with it
func Relu[T Number](i T) T {
	if i < 0 {
		return 0
	}

	return i
}

/*

func ReluInt(i int) int {
	if i < 0 {
		return 0
	}

	return i
}

func ReluFloat64(i float64) float64 {
	if i < 0 {
		return 0
	}

	return i
}

*/