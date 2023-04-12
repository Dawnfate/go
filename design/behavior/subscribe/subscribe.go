package subscribe

import "fmt"

type Subject interface {
	Subscribe(observer Observer)
	Notify(msg string)
}

type Observer interface {
	Update(msg string)
}

type SubjectImpl struct {
	observers []Observer
}

func (sub *SubjectImpl) Subscribe(observer Observer) {
	sub.observers = append(sub.observers, observer)
}

//Notify 发布通知
func (sub *SubjectImpl) Notify(msg string) {
	for _, o := range sub.observers {
		o.Update(msg)
	}
}

type Observer1 struct {
}

// Update 实现观察者接口
func (Observer1) Update(msg string) {
	fmt.Printf("Observer1: %s\n", msg)
}

type Observer2 struct {
}

// Update 实现观察者接口
func (Observer2) Update(msg string) {
	fmt.Printf("Observer2: %s\n", msg)
}

func main() {
	sub := &SubjectImpl{}
	sub.Subscribe(&Observer1{})
	sub.Subscribe(&Observer2{})
	sub.Notify("Hello")
}
