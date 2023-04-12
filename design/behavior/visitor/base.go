package main

func main() {
	orderService := NewOrderService()
	orderService.Save(NewOrder(0, "李生", "广州", "电视", 10))
	orderService.Save(NewOrder(1, "张三", "广州", "电视", 10))
	orderService.Save(NewOrder(2, "李四", "深圳", "冰箱", 20))
	orderService.Save(NewOrder(3, "王五", "东莞", "空调", 30))
	orderService.Save(NewOrder(4, "张三三", "广州", "空调", 10))
	orderService.Save(NewOrder(5, "李四四", "深圳", "电视", 20))
	orderService.Save(NewOrder(6, "王五五", "东莞", "冰箱", 30))

	cv := NewCityVisitor()
	orderService.Accept(cv)
	cv.Report()

	pv := NewProductVisitor()
	orderService.Accept(pv)
	pv.Report()
}
