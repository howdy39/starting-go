package main

import "fmt"

func main() {
	{
		// 普通の配列
		var s [3]int
		s[0] = 10
		s[2] = 30
		fmt.Printf("%#v\n", s) // => [3]int{10, 0, 30}
	}

	{
		// スライス
		var s []int
		s = make([]int, 5)     // 要素数と容量が5
		fmt.Printf("%#v\n", s) // => []int{0, 0, 0, 0, 0}

		// len
		fmt.Println(len(s)) // => 5
		// cap
		fmt.Println(cap(s)) // => 5
	}

	{
		// スライス
		var s []int
		s = make([]int, 5, 10) // 要素数が5容量が10
		fmt.Printf("%#v\n", s) // => []int{0, 0, 0, 0, 0}

		// len
		fmt.Println(len(s)) // => 5
		// cap
		fmt.Println(cap(s)) // => 10
	}

	{
		// スライスを生成するリテラル
		s := []int{3, 4, 5, 6}
		fmt.Printf("%#v\n", s) // => []int{3, 4, 5, 6}

		// 簡易スライス
		ss := s[:2]
		fmt.Printf("%#v\n", ss) // => []int{3, 4}

		// 文字列と簡易スライス
		st := "abcde"[1:3]
		fmt.Printf("%#v\n", st) // => "bc"
	}

	{
		// append
		s := []int{1, 2, 3}
		s = append(s, 4)
		fmt.Printf("%#v\n", s) // => []int{1, 2, 3, 4}
		s = append(s, 4, 5, 6, 7)
		fmt.Printf("%#v\n", s) // => []int{1, 2, 3, 4, 4, 5, 6, 7}
	}

	{
		// copy
		s1 := []int{1, 2, 3, 4, 5}
		s2 := []int{10, 11}
		n := copy(s1, s2)
		fmt.Printf("%#v %#v\n", s1, n) // => []int{10, 11, 3, 4, 5} 2
	}

	{
		// スライスとfor
		s := []string{"Apple", "Banana", "Cherry"}
		for i, v := range s {
			fmt.Printf("%#v %#v\n", i, v) // => 0 "Apple"   1 "Banana"   2 "Cherry"
		}

		for i := 0; i < len(s); i++ {
			fmt.Printf("%#v %#v\n", i, s[i]) // => 0 "Apple"   1 "Banana"   2 "Cherry"
		}
	}

	{
		// スライスと可変長引数
		sum := func(s ...int) int {
			n := 0
			for _, v := range s {
				n = n + v
			}
			return n
		}
		fmt.Println(sum(1, 2, 3))
	}

}
