package main

import "fmt"

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "エラーが発生しました", ErrCode: 1234}
}

func main() {
	error := RaiseError()
	fmt.Println(error.Error())

	e, ok := error.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
	}
}
