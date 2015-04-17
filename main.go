package main

import (
	//"NativePackage/LyCrypto"
	"NativePackage/LyHeap"
	//"fmt"
)

type user struct {
	name string
	age  int
}

func main() {
	/*
		u1 := new(user)
		u2 := &user{}
		u3 := &user{"a", 1}

		fmt.Println(u1)
		fmt.Println(u2)
		fmt.Println(*u3)
	*/
	LyHeap.Example_intHeap()
}
