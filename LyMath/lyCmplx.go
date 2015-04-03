package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	lyNaN()
}

func lyNaN() {
	c := cmplx.NaN()
	fmt.Println(c)
	abs := cmplx.Abs(-12.3)
	fmt.Println(abs)
}
