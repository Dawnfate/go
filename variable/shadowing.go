package main

import (
	"errors"
	"fmt"
)

func case1() {
	x := 0

	// value - 0
	fmt.Println(x)

	if true {
		// variable shadowing
		x := 1
		x++
		// value - 2
		fmt.Println(x)
	}

	// value - 0
	fmt.Println(x)
}

func main() {
	f, err := foo()
	_ = f

	// variable shadowing
	// 函数多变量返回的时候，第一个变量之后的函数如果已经被定义了
	//那么就不会被赋值，而是采用上面的值
	if b, err := bar(); err == nil {
		_ = b
	}
	/*	if b, errB := bar(); errB == nil {
		_ = b
		_ = errB
		fmt.Println(errB)
	}*/
	/*	var b string
		if b, err = bar(); err == nil {
			_ = b
		}
		fmt.Println(err)*/

	// nil
	fmt.Println(err)
}

func foo() (string, error) {
	return "", nil
}

func bar() (string, error) {
	return "", errors.New("bar")
}

func acc() {
	// 变量的生命周期 - 仅在括号内
	// 重命名
	/*	if b, errB := bar(); errB == nil {
		_ = b
		_ = errB
	}*/

	// 变量的生命周期 - 括号内外
	// 移除:=
	/*	var b string
		if b, err = bar(); err == nil {
			_ = b
		}
		fmt.Println(err)*/
}
