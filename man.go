package main

import (
	// common "fangwenju/study/common"
	// cal "fangwenju/study/calculator"
	// socket "fangwenju/study/socket"
	"fmt"
)

// interface{} 可以接收所有的异常
// fun 正常的方法
// catch 异常处理的方法
func Try(fun func(), catch func(err interface{})) {
	defer func() {
		fmt.Println("捕获的异常, 进行异常处理")
		err := recover()
		if err != nil {
			catch(err)
		}
	}()
	fun() // 正常的方法
}
