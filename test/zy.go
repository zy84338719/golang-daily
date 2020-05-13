package main

import (
	"errors"
	"fmt"
)

func main() {
	testError()
	afterErrorfunc()
}

func testError() {
	defer catch()

	go func() {
		defer catch()
		panic(" \"panic 错误\"")

	}()
	fmt.Println("抛出一个错误后继续执行代码")
}
func catch() {
	if r := recover(); r != nil {
		fmt.Println("testError() 遇到错误:", r)
		var err error
		switch x := r.(type) {
		case string:
			err = errors.New(x)
		case error:
			err = x
		default:
			err = errors.New("")
		}
		if err != nil {
			fmt.Println("recover后的错误:", err)
		}
	}
}

func afterErrorfunc() {
	fmt.Println("遇到错误之后 func ")
}
