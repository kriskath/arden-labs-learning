package main

import "fmt"

func main() {
	fmt.Println(safeDiv(7, 3))
	fmt.Println(safeDiv(7, 0))

}

/* Using named return values good for:
- defer/recover to change return error value
- Documentation
*/

func safeDiv(a, b int) (q int, err error) {
	// q & err are variables inside safeDiv just like a & b
	defer func() {
		if e := recover(); e != nil {
			// fmt.Println("Error:", e)
			err = fmt.Errorf("%v", e)

		}
	}()

	/* Valid code, but hard to follow
	q = div(a,b)
	return
	*/

	return div(a, b), nil
}

func div(a, b int) int {
	/*
		if b == 0 {
			panic("division by zero")
		}
	*/
	return a / b
}
