package calculator

import (
	"testing"
)

func TestLink(t *testing.T) {
	ret := OperationFactory(1, 2, "+").Exe()
	t.Log(ret)
	// try(func() {
	// 	fmt.Println("kk")
	// 	panic("error this try catch")
	// }, func(err interface{}) {
	// 	fmt.Println("捕获异常", err)
	// })
}
