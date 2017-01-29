package main

import "fmt"

var packageVariable = "package変数です"

func main() {
	fmt.Printf("パッケージ変数=%#v", packageVariable)

	// 明示的な定義
	var n int
	n = 100
	fmt.Println(n)

	var x, y, z int
	x, y, z = 1, 2, 3
	fmt.Printf("%#v %#v %#v\n", x, y, z)

	var (
		key   int
		value string
	)

	key = 10
	value = "hoge"
	fmt.Printf("key=%#v value=%#v\n", key, value)

	var (
		a = "A"
		b = "B"
		c = "C"
	)
	fmt.Printf("a=%#v b=%#v c=%#v\n", a, b, c)

	// 暗黙的な定義
	i := 1
	fmt.Printf("i=%#v\n", i)
}
