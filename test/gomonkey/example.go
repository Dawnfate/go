package gomonkey

import (
	"fmt"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/smartystreets/goconvey/convey"
)

// 假设networkFunc是一个网络调用
func networkFunc(a, b int) int {
	return a + b
}

// 本地单测一般不会进行网络调用，所以要mock住networkFunc
func Test_MockNetworkFunc(t *testing.T) {
	convey.Convey("123", t, func() {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFunc(networkFunc, func(a, b int) int {
			fmt.Println("in mock function")
			return a + b
		})
		_ = networkFunc(10, 20)
	})
}
