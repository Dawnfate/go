package main

import "fmt"

type Printer interface {
	Print(name string) string
}

type CnPrint struct {
}

func (*CnPrint) Print(name string) string {
	return fmt.Sprintf("你好，%s", name)
}

func (*EnPrint) Print(name string) string {
	return fmt.Sprintf("Hello，%s", name)
}

type EnPrint struct {
}

func NewPrinter(lang string) Printer {
	switch lang {
	case "cn":
		return new(CnPrint)
	case "en":
		return new(EnPrint)
	default:
		return new(CnPrint)
	}
}

func main() {
	printer := NewPrinter("en")
	fmt.Println(printer.Print("Bob"))
}
