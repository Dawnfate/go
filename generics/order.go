package main

import (
	"encoding/json"
	"fmt"
)

// 将公共逻辑抽象到这个泛型函数中
func Create[OrderInfo interface{}](order OrderInfo) error {
	err := mysql.Insert(order)
	if err != nil {
		return err
	}
	b, err := json.Marshal(order)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

type OrderInfo1 Order1

// 三个Order的Create方法就非常清晰了
func (o *Order1) Create() error {
	var order *OrderInfo1
	return Create[OrderInfo1](order)
}

func (o *Order2) Create() error {
	var order *OrderInfo2
	return Create[OrderInfo2](order)
}

func (o *Order3) Create() error {
	var order *OrderInfo3
	return Create[OrderInfo3](order)
}
