package main

import "fmt"
import "runtime"

func init() {
	fmt.Println("init()")
}

func main() {
	// if
	x := 1
	if x == 1 {
		fmt.Println("xは1です")
	}

	// 簡易文付きif
	a, b := 3, 5
	if n := a * b; n%2 == 0 {
		fmt.Printf("n(%d) is even \n", n)
	} else {
		fmt.Printf("n(%d) is odd \n", n)
	}

	// forは無限ループ、breakが使える
	i := 0
	for {
		fmt.Println(i)
		i++
		if i > 10 {
			break
		}
	}

	// continue break
	for i := 0; i < 20; i++ {
		if i < 10 {
			fmt.Printf("continue:%d\n", i)
			continue
		} else {
			fmt.Printf("break:%d\n", i)
			break
		}
	}

	// range
	fruits := [...]string{"Apple", "Banana", "Cherry"}
	for i, s := range fruits {
		fmt.Printf("i:%d s:%s\n", i, s)
	}

	// 簡易文付きswitch
	{
		switch n := 3; n {
		case 1, 2:
			fmt.Println("1 or 2")
		case 3, 4:
			fmt.Println("3 or 4")
		default:
			fmt.Println("unkhown")
		}
	}

	// defer
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("done")
	/*
	  done
	  3
	  2
	  1
	*/

	// goroutine
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("Version: %s\n", runtime.Version())

	// recover
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("reocver1")
		}
	}()
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("reocver2") // こっちが呼ばれる
		}
	}()

	// panic
	panic("runtime error")
	// fmt.Println("ここにはこない")

}
