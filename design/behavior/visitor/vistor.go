package main

import "fmt"

type IOrderVisitor interface {
	// 这里参数不能定义成 OrderService
	Visit(order *Order)
	Report()
}

type CityVisitor struct {
	cities map[string]int
}

func (cv *CityVisitor) Visit(o *Order) {
	n, ok := cv.cities[o.City]
	if ok {
		cv.cities[o.City] = n + o.Quantity
	} else {
		cv.cities[o.City] = o.Quantity
	}
}

func (cv *CityVisitor) Report() {
	for k, v := range cv.cities {
		fmt.Printf("city=%s, sum=%v\n", k, v)
	}
}

func NewCityVisitor() IOrderVisitor {
	return &CityVisitor{
		cities: make(map[string]int, 0),
	}
}

// 品类销售报表, 按产品汇总销售情况, 实现ISaleOrderVisitor接口
type ProductVisitor struct {
	products map[string]int
}

func (pv *ProductVisitor) Visit(it *Order) {
	n, ok := pv.products[it.Product]
	if ok {
		pv.products[it.Product] = n + it.Quantity
	} else {
		pv.products[it.Product] = it.Quantity
	}
}

func (pv *ProductVisitor) Report() {
	for k, v := range pv.products {
		fmt.Printf("product=%s, sum=%v\n", k, v)
	}
}

func NewProductVisitor() IOrderVisitor {
	return &ProductVisitor{
		products: make(map[string]int, 0),
	}
}
