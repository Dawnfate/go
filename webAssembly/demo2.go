package main

import (
	figure "github.com/common-nighthawk/go-figure"
)

//export HelloWorld
func HelloWorld() {
	myFigure := figure.NewFigure("Hello World", "", true)
	myFigure.Print()
}

func main() {}
