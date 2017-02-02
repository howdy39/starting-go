package main

import (
	"fmt"
)

func main() {
	plus := plus(1, 5)
	fmt.Printf("%#v\n", plus) // => 6

	// 複数の戻り値
	div1, div2 := div(10, 3)
	fmt.Printf("%#v %#v\n", div1, div2) // => 3 1

	doSomething1, doSomething2 := doSomething()
	fmt.Printf("%#v %#v\n", doSomething1, doSomething2) // => 0 10

	// 無名関数
	f := func(x, y int) int { return x + y }
	fmt.Printf("%T %#v\n", f, f(2, 3)) // => func(int, int) int 5

	// 即時関数
	fmt.Printf("%#v\n", func(x, y int) int { return x + y }(1, 3)) // => 4

	// 関数を返す関数
	returnFunc()() // I'm a function.
}

func plus(x, y int) int {
	return x + y
}

func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

func doSomething() (x, y int) {
	y = 10
	return
}

func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function.")
	}
}
