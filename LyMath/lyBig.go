package main

import (
	"fmt"
	//"math"
	"math/big"
	//"math/rand"
)

func main() {
	lyGCD()
}

func lyNewInt() {
	intBig := big.NewInt(1)
	fmt.Println(intBig)
}

func lyMulRange() {
	intBig := big.NewInt(1)
	mul := intBig.MulRange(1, 3)
	fmt.Println(mul)
}

func lyBinomial() {
	intBig := big.NewInt(1)
	bino := intBig.Binomial(200, 14)
	fmt.Println(bino)
}

func lyRand() {
	fmt.Println("rand")
}

func lySign() {
	b := big.NewInt(-100)
	s := b.Sign()
	fmt.Println(s)
}

//比较x和y的大小。x<y时返回-1；x>y时返回+1；否则返回0。
func lyCmp() {
	x := big.NewInt(1)
	y := big.NewInt(2)
	c := x.Cmp(y)
	fmt.Println(c)
}

func lyAbs() {
	z := big.NewInt(10)
	x := big.NewInt(-10)
	r := z.Neg(x)
	fmt.Println(r)
}

func lyAdd() {
	z := big.NewInt(0)
	x := big.NewInt(10)
	y := big.NewInt(20)
	r := z.Add(x, y)
	fmt.Println(r)
}

func lySub() {
	z := big.NewInt(0)
	x := big.NewInt(10)
	y := big.NewInt(2)
	r := z.Sub(x, y)
	fmt.Println(r)
}

/*
func (z *Int) GCD(x, y, a, b *Int) *Int
将z设为a和b的最大公约数并返回z。a或b为nil时会panic；a和b都>0时设置z为最大公约数；
如果任一个<=0方法就会设置z = x = y = 0。如果x和y都不是nil，会将x和y设置为满足a*x + b*y==z。
*/
func lyGCD() {
	z := big.NewInt(0)
	a := big.NewInt(1024)
	b := big.NewInt(768)
	x := big.NewInt(0)
	y := big.NewInt(0)
	r := z.GCD(x, y, a, b)
	fmt.Println(r)
}
