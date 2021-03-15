package calculator

import (
	"fmt"
	"reflect"
)

var opers map[string]interface{}

func init() {
	opers = make(map[string]interface{}, 0)
	opers["+"] = NewAdd
}

func OperationFactory(num1 int, num2 int, flag string) OperationInterface {
	oper := opers[flag]
	fmt.Println("oper:", oper)

	valueOper := reflect.ValueOf(oper)
	fmt.Println("reflect.ValueOf(oper)", valueOper)
	fmt.Printf("type:%T \n", valueOper)

	aegs := []reflect.Value{
		reflect.ValueOf(num1), reflect.ValueOf(num2),
	}

	arrValueOper := valueOper.Call(aegs)[0]
	fmt.Println("arrValueOper.Call()", arrValueOper)
	fmt.Printf("type:%T \n", arrValueOper)

	operation := arrValueOper.Interface().(OperationInterface)
	fmt.Printf("type:%T \n", operation)
	return operation
}
