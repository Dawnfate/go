package main

import "syscall/js"

func main() {
	alert := js.Global().Get("alert")
	alert.Invoke("脑子进煎鱼了！")
}
