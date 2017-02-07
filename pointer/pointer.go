package main

import "fmt"

func main() {
	{
		var p *int
		fmt.Printf("%#v\n", p) // => (*int)(nil)
	}

	{
		// アドレス演算子 &
		var i int
		p := &i
		fmt.Printf("%#v\n", p) // => (*int)(0xc420070038)

		// デリファレンス *
		*p = 10
		fmt.Printf("%#v\n", i) // => 10

	}

	{
		p := &[3]int{1, 2, 3}
		func(p *[3]int) {
			i := 0
			for i < 3 {
				(*p)[i] = (*p)[i] * (*p)[i]
				i++
			}
		}(p)
		fmt.Println(p) // => &[1 4 9]
	}

	{
		// 組み込み関数や[x]でのアクセスなら、デリファレンスは省略可能
		p := &[3]int{1, 2, 3}
		func(p *[3]int) {
			i := 0
			for i < 3 {
				p[i] = p[i] * p[i]
				i++
			}
		}(p)
		fmt.Println(p) // => &[1 4 9]
	}

	{
		// 配列は参照型ではない
		p := [3]int{1, 2, 3}
		func(p [3]int) {
			i := 0
			for i < 3 {
				p[i] = p[i] * p[i]
				i++
			}
		}(p)
		fmt.Println(p) // => [1 2 3]
	}

	{
		// スライスは参照型
		p := []int{1, 2, 3}
		func(p []int) {
			i := 0
			for i < 3 {
				p[i] = p[i] * p[i]
				i++
			}
		}(p)
		fmt.Println(p) // => [1 4 9]
	}

	{
		s := "hoge"
		func(s string) {
			s += "piyo"
		}(s)
		fmt.Printf("%#v", s)
	}

}
