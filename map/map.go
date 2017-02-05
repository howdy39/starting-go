package main

import "fmt"

func main() {
	{
		var m map[int]string
		m = make(map[int]string)
		m[1] = "ONE"
		m[2] = "TWO"
		fmt.Printf("%#v\n", m) // => map[int]string{1:"ONE", 2:"TWO"}
		m[2] = "上書き可能"
		fmt.Printf("%#v\n", m) // => map[int]string{1:"ONE", 2:"上書き可能"}
	}

	{
		// リテラルで生成
		m := map[int]string{1: "Taro", 2: "Jiro"}
		fmt.Println(m) // => map[1:Taro 2:Jiro]
	}

	{
		// 省略記法
		m := map[int][]int{
			1: []int{1},
			2: []int{1, 2},
			3: []int{1, 2, 3},
		}
		fmt.Println(m) // => map[1:[1] 2:[1 2] 3:[1 2 3]]

		n := map[int][]int{
			1: {1},
			2: {1, 2},
			3: {1, 2, 3},
		}
		fmt.Println(n) // => map[1:[1] 2:[1 2] 3:[1 2 3]]
	}

	{
		// キーがなくてもエラーにはならない
		m := map[int]string{1: "A", 2: "B", 3: "C"}
		fmt.Println(m[1]) // => "A"
		fmt.Println(m[9]) // => ""

		// ２番めにbool型で取れる
		if s, ok := m[1]; ok {
			fmt.Printf("m[1]:%s\n", s) // => m[1]:A
		}
		if s, ok := m[9]; ok == false {
			fmt.Printf("m[9]:%s\n", s) // => m[9]:
		}

		// for
		for k, v := range m {
			fmt.Printf("k, v= %d, %s\n", k, v) // => k, v = 1, A  k, v= 2, B  k, v= 3, C
		}

		// deleteで削除可能
		delete(m, 3)

		// for
		for k, v := range m {
			fmt.Printf("k, v= %d, %s\n", k, v) // => k, v = 1, A  k, v= 2, B
		}
	}
}
