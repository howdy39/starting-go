package main

import "fmt"

type Callback func(i int) int

func Sum(ints []int, callback Callback) int {
	var sum int
	for _, i := range ints {
		sum += i
	}
	return callback(sum)
}

type Point struct {
	X, Y int
}

func (p *Point) Render() {
	fmt.Printf("<%d, %d>\n", p.X, p.Y)
}

func main() {
	{
		type MyInt int
		var n1 MyInt = 5
		fmt.Println(n1)
	}

	{
		n := Sum(
			[]int{1, 2, 3, 4, 5},
			func(i int) int {
				return i * 2
			},
		)
		fmt.Println(n) // => 30
	}

	{
		type Point struct {
			X int
			Y int
		}

		var pt Point
		pt.X = 10
		pt.Y = 5
		fmt.Printf("X=%v, Y=%v\n", pt.X, pt.Y) // => X=10, Y=5

		// 複合リテラル
		pt2 := Point{20, 10}
		fmt.Printf("X=%v, Y=%v\n", pt2.X, pt2.Y) // => X=20, Y=10

		pt3 := Point{Y: 20, X: 10}
		fmt.Printf("X=%v, Y=%v\n", pt3.X, pt3.Y) // => X=10, Y=20
	}

	{
		// 構造体を含む構造体
		type Feed struct {
			Name   string
			Amount int
		}
		type Animal struct {
			Name string
			Feed Feed
		}

		a := Animal{
			Name: "Monkey",
			Feed: Feed{
				Name:   "Banana",
				Amount: 10,
			},
		}

		fmt.Printf("%#v\n", a) // => main.Animal{Name:"Monkey", Feed:main.Feed{Name:"Banana", Amount:10}}

		// new
		b := new(Animal)
		b.Name = "Monkey"
		b.Feed.Name = "Banana"
		b.Feed.Amount = 100
		fmt.Printf("%#v\n", a) // => main.Animal{Name:"Monkey", Feed:main.Feed{Name:"Banana", Amount:10}}
	}

	{
		// メソッド
		x := &Point{X: 5, Y: 12}
		x.Render()
	}

}
