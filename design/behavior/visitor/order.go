package main

// 订单服务接口
type IOrderService interface {
	Save(order *Order) error
	// 有的教程里把接收 visitor 实现的方法名定义成 Accept
	Accept(visitor IOrderVisitor)
}

// 订单实体类，实现IOrderService 接口
type Order struct {
	ID       int
	Customer string
	City     string
	Product  string
	Quantity int
}

type OrderService struct {
	orders []*Order
}

func NewOrderService() OrderService {
	return OrderService{}
}

func (mo *OrderService) Save(o *Order) error {
	mo.orders = append(mo.orders, o)
	mo.orders[o.ID] = o
	return nil
}

func (mo *OrderService) Accept(visitor IOrderVisitor) {
	for _, v := range mo.orders {
		visitor.Visit(v)
	}
}

func NewOrder(id int, customer string, city string, product string, quantity int) *Order {
	return &Order{
		id, customer, city, product, quantity,
	}
}
