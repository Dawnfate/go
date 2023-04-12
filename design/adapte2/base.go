package adapte2

// Target 目标接口
type Target interface {
	Request() string
}

// Adaptee 被适配的接口
type Adaptee interface {
	SpecificRequest() string
}

type AdapteeImpl struct {
}

func (a AdapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

func NewAdaptee() Adaptee {
	return &AdapteeImpl{}
}

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{Adaptee: adaptee}
}

type adapter struct {
	Adaptee
}

//Request 实现Target接口
func (a *adapter) Request() string {
	return a.SpecificRequest()
}
