package calculator

import (
// "fmt"
)

type OperationInterface interface {
	Exe() int
}

type Operation struct {
	num1 int
	num2 int
}

func (o *Operation) SetNum1(num int) {
	o.num1 = num
}

func (o *Operation) SetNum2(num int) {
	o.num2 = num
}

func (o *Operation) GetNum1() int {
	return o.num1
}

func (o *Operation) GetNum2() int {
	return o.num2
}

type Add struct {
	oper *Operation
}

func NewAdd(num1 int, num2 int) *Add {
	return &Add{
		oper: &Operation{
			num1: num1,
			num2: num2,
		},
	}
}

func (a *Add) Exe() int {
	return a.oper.num1 + a.oper.num2
}

// type Calculator struct {
// 	number1  float32
// 	number2  float32
// 	results  float32
// 	operator string
// }

// func (c *Calculator) Desc() string {
// 	fmt.Println("----------------------+----------------------")
// 	fmt.Println("||    计算器功能      |         01         ||")

// 	return "01"
// 	// fmt.Println("----------------------+----------------------")
// 	// fmt.Println("||    信息查询功能    |         02         ||")
// }

// func (c *Calculator) next() {
// 	fmt.Println("输入运算符继续计算，输入 c 重新计算， 输入其他字符结束")
// 	var z string
// 	fmt.Scan(&z)
// 	if z == "+" || z == "-" || z == "*" || z == "/" {
// 		c.number1 = c.results
// 		c.operator = z
// 		fmt.Println("请输入一个数字，按回车确认")
// 		c.number2 = 0
// 		fmt.Scan(&c.number2)
// 		c.verify_num()
// 		c.result()
// 	} else if z == "c" {
// 		c.number1 = 0
// 		c.number2 = 0
// 		c.Start()
// 	} else {
// 		fmt.Print("感谢您的使用！！！")
// 	}
// }

// func (c *Calculator) result() {
// 	switch c.operator {
// 	case "+":
// 		c.results = c.number1 + c.number2
// 	case "-":
// 		c.results = c.number1 - c.number2
// 	case "*":
// 		c.results = c.number1 * c.number2
// 	case "/":
// 		c.results = c.number1 / c.number2
// 	}
// 	fmt.Println("计算结果为：", c.number1, c.operator, c.number2, "=", c.results)
// 	c.next()
// }

// func (c *Calculator) verify_num() {
// 	if c.number2 == 0 && c.operator == "/" {
// 		fmt.Println("除数不能为0，按回车确认")
// 		fmt.Println("请输入一个有效数字，按回车确认")
// 		fmt.Scan(&c.number2)
// 		c.verify_num()
// 	}
// }

// func (c *Calculator) verify_ope() {
// 	if c.operator != "+" && c.operator != "-" && c.operator != "*" && c.operator != "/" {
// 		fmt.Println("您输入的不是一个有效运算符，请输入任意运算符 + - * / ,按回车确认")
// 		fmt.Scan(&c.operator)
// 		c.verify_ope()
// 	}
// }

// func (c *Calculator) Start() {
// 	fmt.Println("请输入一个数字，按回车确认")
// 	fmt.Scan(&c.number1)
// 	fmt.Println("请输入任意运算符 + - * / ,按回车确认")
// 	fmt.Scan(&c.operator)
// 	c.verify_ope()
// 	fmt.Println("请输入一个数字，按回车确认")
// 	fmt.Scan(&c.number2)
// 	c.verify_num()
// 	c.result()
// }
