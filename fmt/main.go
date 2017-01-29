package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Golang!")    // => Hello, Golang!
	fmt.Println("Hello,", "Golang!") // => Hello, Golang!

	fmt.Printf("10進数=%d 2進数=%b 8進数=%o 16進数=%x\n", 10, 10, 10, 10) // => 10進数=10 2進数=1010 8進数=12 16進数=a

	fmt.Printf("数値=%v 文字列=%v 配列=%v\n", 10, "文字列", [...]string{"hoge", "piyo", "fuga"})    // => 数値=10 文字列=文字列 配列=[hoge piyo fuga]
	fmt.Printf("数値=%#v 文字列=%#v 配列=%#v\n", 10, "文字列", [...]string{"hoge", "piyo", "fuga"}) // => 数値=10 文字列="文字列" 配列=[3]string{"hoge", "piyo", "fuga"}
	fmt.Printf("数値=%T 文字列=%T 配列=%T\n", 10, "文字列", [...]string{"hoge", "piyo", "fuga"})    // => 数値=int 文字列=string 配列=[3]string

	print("Hello, World!\n") // => Hello, World!
	println("Hello, World!") // => Hello, World!
}
