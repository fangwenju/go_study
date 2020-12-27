package calculator

import (
	"fmt"
	"testing"
)

func TestArr(t *testing.T) {
	arr := [2]int{1, 2}
	fmt.Println(arr)
}

func TestSelip(t *testing.T) {
	selip := []int{3, 4}
	fmt.Println(selip)
}

func TestMap(t *testing.T) {
	mp := make(map[string]string)
	mp["a"] = "a"
	fmt.Println(mp)
}

func aa() func() {
	i := 0
	// fmt.Println(&i)
	return func() {
		i++
		// fmt.Println(i)
		// fmt.Println("地址", &i)
	}
}

func TestFun(t *testing.T) {
	a := aa()

	fmt.Println(a)
	fmt.Println(a)
	fmt.Println(a)
}
