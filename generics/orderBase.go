package main

import (
	"encoding/json"
	"fmt"
)

type OrderInfo1 struct{}

func (o *Order1) Create() error {
	// 数据结构 OrderInfo1 保存的是 Order1 订单信息
	var order *OrderInfo1
	// 插入msyql
	err := mysql.Insert(order)
	if err != nil {
		return err
	}
	// 序列化后打印
	b, err := json.Marshal(order)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

type OrderInfo2 struct{}

func (o *Order2) Create() error {
	// 数据结构 OrderInfo2 保存的是 Order2 订单信息
	var order *OrderInfo2
	// 后面操作同Order1
}

type OrderInfo3 struct{}

type Order3 struct{}

func (o *Order3) Create() error {
	// 数据结构 OrderInfo3 保存的是 Order3 订单信息
	var order *OrderInfo3
	// 后面操作同Order1
}
