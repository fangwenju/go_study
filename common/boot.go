package common

import (
	// calc "fangwenju/go_study/calc"
	calculator "fangwenju/study/calculator"
	"fmt"
	"reflect"
)

func Welcome() {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>> START <<<<<<<<<<<<<<<<<<<<<<")
	fmt.Println("欢迎使用")
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>> END <<<<<<<<<<<<<<<<<<<<<<")
}

type Boot struct {
}

func (b *Boot) Project_list() {
	fmt.Println("=============================================")
	fmt.Println("||      可选功能      |         编号       ||")

	var list = make(map[string]interface{})
	calc := calculator.Calculator{}
	list[calc.Desc()] = &calc

	fmt.Println("=============================================")

	fmt.Println("输入功能编号进入对应功能，输入x退出程序")
	var num string = "01"
	// fmt.Scan(&num)
	// .Start()
	// args := []reflect.Value{
	// 	reflect.ValueOf(1.1), reflect.ValueOf(2.2), reflect.ValueOf(3.0), reflect.ValueOf("+"),
	// }

	valueOper := reflect.ValueOf(list[num]) // reflect.value
	// aa := reflect.Zero(valueOper)           // reflect.value

	arrValueOper := valueOper.Elem()
	fmt.Printf("type2:%T \n", valueOper)
	fmt.Printf("type2:%T \n", arrValueOper)
	// fmt.Printf("type2:%T \n", aa)

	// v := valueOper.Interface().(calculator.Calculator)
	// fmt.Println("valueOper.Call", v)
	// reflect.Value => 解析为原来的对象 OperationInterface
	// fmt.Printf("type:%T \n", v)
}
