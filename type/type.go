package main

import (
	"fmt"
)

func main() {
	// 論理値型
	var b bool
	b = true
	fmt.Printf("b=%#v\n", b) // => b=true

	// 整数型
	var i8 int8
	i8 = -100
	fmt.Printf("i8=%#v\n", i8) // => i8=-100

	var u8 uint8
	u8 = 100
	fmt.Printf("u8=%#v\n", u8) // => u8=0x64

	var by byte
	by = u8                    // byteはuint8のエイリアスなのでキャストが要らない
	fmt.Printf("by=%#v\n", by) // => by=0x64

	var i int
	i = 300
	fmt.Printf("i=%#v\n", i) // => i=300

	// キャスト
	i = int(i8)
	fmt.Printf("i=%#v\n", i) // => i=-100

	// 浮動小数点型
	var f32 float32
	f32 = 1.1111
	fmt.Printf("f32=%T\n", f32) // => f32=float32

	f64 := 1.1111
	fmt.Printf("f64=%T\n", f64) // => f64=float64

	// 複素数型
	c128 := 1.0 + 3i
	fmt.Printf("c128=%T\n", c128) // => c128=complex128

	// ルーン型
	var r rune
	r = '\n'
	fmt.Printf("r=%T r=%#v\n", r, r) // => r=int32 r=10

	var i32 int32
	i32 = 20
	r = i32                          // runeはint32のエイリアスなのでキャストが要らない
	fmt.Printf("r=%T r=%#v\n", r, r) // => r=int32 r=20

	// 文字列型
	var str string
	str = "hoge"
	fmt.Printf("str=%#v\n", str) // => str="hoge"

	rawstr := `
  １行目
  ２行目
  `
	fmt.Printf("rawstr=%#v\n", rawstr) // => awstr="\n  １行目\n  ２行目\n  "

	// 配列型
	ia := [5]int{1, 2, 3, 4}
	fmt.Printf("ia=%T ia=%#v\n", ia, ia) // => ia=[5]int ia=[5]int{1, 2, 3, 4, 0}

	ia2 := [...]int{1, 2, 3, 4}              // 要素数の省略
	fmt.Printf("ia2=%T ia2=%#v\n", ia2, ia2) // => ia2=[4]int ia2=[4]int{1, 2, 3, 4}

	// interface{}型
	var iface interface{}
	fmt.Printf("iface=%#v\n", iface) // => iface=<nil>
	iface = "hoge"
	fmt.Printf("iface=%#v\n", iface) // => iface="hoge"
}
