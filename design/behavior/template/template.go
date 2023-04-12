package template

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type BankBusinessHandler interface {
	// TakeRowNumber 排队拿号
	TakeRowNumber()
	// WaitHead 等位
	WaitHead()
	// HandleBusiness 处理业务
	HandleBusiness()
	// Commentate 对服务进行评价
	Commentate()
	// CheckVipIdentity 钩子方法，判断是不是VIP，vip不需要等位
	CheckVipIdentity() bool
}

type BankBusinessExecutor struct {
	handler BankBusinessHandler
}

func (b *BankBusinessExecutor) ExecuteBankBusiness() {
	b.handler.TakeRowNumber()
	if !b.handler.CheckVipIdentity() {
		b.handler.WaitHead()
	}
	b.handler.HandleBusiness()
	b.handler.Commentate()
}

type DepositBusinessHandler struct {
	*DefaultBusinessHandler
	userVip bool
}

func (*DepositBusinessHandler) HandleBusiness() {
	fmt.Println("账户存储很多万人民币...")
}

func (dh *DepositBusinessHandler) CheckVipIdentity() bool {
	return dh.userVip
}

type DefaultBusinessHandler struct {
}

func (*DefaultBusinessHandler) TakeRowNumber() {
	fmt.Println("请拿好您的取件码：" + strconv.Itoa(rand.Intn(100)) +
		" ，注意排队情况，过号后顺延三个安排")
}

func (dbh *DefaultBusinessHandler) WaitHead() {
	fmt.Println("排队等号中...")
	time.Sleep(5 * time.Second)
	fmt.Println("请去窗口xxx...")
}

func (*DefaultBusinessHandler) Commentate() {

	fmt.Println("请对我的服务作出评价，满意请按0，满意请按0，(～￣▽￣)～")
}

func (*DefaultBusinessHandler) CheckVipIdentity() bool {
	// 留给具体实现类实现
	return false
}

func NewBankBusinessExecutor(businessHandler BankBusinessHandler) *BankBusinessExecutor {
	return &BankBusinessExecutor{handler: businessHandler}
}

func main() {
	dh := &DepositBusinessHandler{userVip: false}
	bbe := NewBankBusinessExecutor(dh)
	bbe.ExecuteBankBusiness()
}
