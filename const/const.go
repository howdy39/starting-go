package main

import "fmt"

const PACKAGE_CONST = "これはパッケージ定数"

func main() {
	const LOCAL_CONST = "これはローカル定数"
	fmt.Printf("PACKAGE_CONST=%s\n", PACKAGE_CONST) // => PACKAGE_CONST=これはパッケージ定数
	fmt.Printf("LOCAL_CONST=%s\n", LOCAL_CONST)     // => LOCAL_CONST=これはローカル定数

	// 値の省略が可能（直前の定数が流れて設定される）
	const (
		X = 10
		Y
		Z = 100
	)
	fmt.Printf("X=%d, Y=%d, Z=%d \n", X, Y, Z) // => X=10, Y=10, Z=100

	// 定数値の式
	const (
		XYZ = X + Y + Z
	)
	fmt.Printf("XYZ=%d %T \n", XYZ, XYZ) // => XYZ=120 int

	const (
		A = iota
		B = iota
		C = iota
	)
	fmt.Printf("A=%d, B=%d, C=%d \n", A, B, C) // => A=0, B=1, C=2

}
