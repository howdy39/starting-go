package main

import (
	"fmt"

	"./foo"
)

func main() {
	fmt.Printf("%#v \n", foo.A)

	a := "hoge"
	{
		a := "piyo"
		fmt.Printf("%#v \n", a) // => "piyo"
	}
	fmt.Printf("%#v \n", a) // => "hoge"
}
