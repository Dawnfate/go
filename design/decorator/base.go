package main

import "fmt"

// PS5 产品接口
type PS5 interface {
	StartGPUEngine()
	GetPrice() int64
}

// CD 版 PS5主机   PS5行为实现的对象
type PS5WithCD struct{}

func (p PS5WithCD) StartGPUEngine() {
	fmt.Println("start engine")
}
func (p PS5WithCD) GetPrice() int64 {
	return 5000
}

// PS5 数字版主机  PS5行为实现的对象
type PS5WithDigital struct{}

func (p PS5WithDigital) StartGPUEngine() {
	fmt.Println("start normal gpu engine")
}

func (p PS5WithDigital) GetPrice() int64 {
	return 3600
}

// PS5MachinePlus struct可以接收interface
type PS5MachinePlus struct {
	ps5Machine PS5
}

// SetPS5Machine Plus 版的装饰器 plus针对PS5的特定行为进行升级
//所以需要绑定原定的PS5，所以这里可以防止修改原本的代码结构（通常原本的代码结构已经稳定）
func (p *PS5MachinePlus) SetPS5Machine(ps5 PS5) {
	p.ps5Machine = ps5
}

func (p PS5MachinePlus) StartGPUEngine() {
	p.ps5Machine.StartGPUEngine() //这里其实可以加GPUEngine的各种参数
	fmt.Println("start plus plugin")
}

func (p PS5MachinePlus) GetPrice() int64 {
	return p.ps5Machine.GetPrice() + 500
}

// 主题色版的装饰器
type PS5WithTopicColor struct {
	ps5Machine PS5
}

func (p *PS5WithTopicColor) SetPS5Machine(ps5 PS5) {
	p.ps5Machine = ps5
}

func (p PS5WithTopicColor) StartGPUEngine() {
	p.ps5Machine.StartGPUEngine()
	fmt.Println("尊贵的主题色主机，GPU启动")
}
func (p PS5WithTopicColor) GetPrice() int64 {
	return p.ps5Machine.GetPrice() + 200
}

func main() {
	//装饰器的本质是在原本的结构上，可以加上其他的属性或者是额外的动作
	//还有一点是可以无限装饰
	ps5MachinePlus := PS5MachinePlus{}
	//参数是接口类型的情况下，只要是实现了接口的结构体都可以作为传参的对象
	ps5MachinePlus.SetPS5Machine(PS5WithCD{})
	// ps5MachinePlus.SetPS5Machine(PS5WithDigital{}) // 可以在更换主机
	ps5MachinePlus.StartGPUEngine()
	price := ps5MachinePlus.GetPrice()
	fmt.Printf("PS5 CD 豪华Plus版，价格: %d 元\n\n", price)

	ps5WithTopicColor := PS5WithTopicColor{}
	//参数是接口类型的情况下，只要是实现了接口的结构体都可以作为传参的对象
	ps5WithTopicColor.SetPS5Machine(ps5MachinePlus)
	ps5WithTopicColor.StartGPUEngine()
	price = ps5WithTopicColor.GetPrice()
	fmt.Printf("PS5 CD 豪华Plus 经典主题配色版，价格: %d 元\n", price)
}
