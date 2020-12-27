package model

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

type Model interface {
	ToString() string
}

var (
	path   string                 = "D:\\phpstudy_pro\\WWW\\go\\src\\fangwenju\\study\\data\\" // 数据存储路径
	suffix string                 = ".sql"
	models map[string]interface{} //与下面的不会重复吗
)

func init() {
	models = make(map[string]interface{})
	models["user"] = NewUser
	userDatas = make(map[string]Model, 0)
	rfdata("user", "username", userDatas)
}

func rfdata(name string, pirmay string, datas map[string]Model) error {
	f, err := os.Open(path + name + suffix)
	if err != nil {
		fmt.Println("文件读取异常,", err)
		return errors.New("文件查询不到 error")
	}

	//延迟执行关闭资源
	defer f.Close()
	//创建文件读取缓冲区
	buf := bufio.NewReader(f)
	field := make([]string, 0) //记录字段名
	for {
		row, err := buf.ReadBytes('\n') //读取到指定前的信息返回切片
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("文件读取异常", err)
		}
		data := strings.Split(string(row), ",") //读取数据

		if len(field) == 0 { //第一行为字段名
			field = data
			// for k, v := range data {
			// field[k] = strings.TrimSpace(strings.TrimSuffix(v, "\n"))
			// }
		} else { //从第二行开始为数据
			toModel(name, pirmay, datas, data, field)
		}
	}
	return nil
}

//存储数据
func toModel(name string, pirmay string, datas map[string]Model, data, field []string) error {
	if models[name] == nil {
		return errors.New("模型不存在：" + name)
	}

	mode := reflect.ValueOf(models[name]).Call([]reflect.Value{})[0]
	fmt.Println(mode)
	fmt.Println("data:", data)
	fmt.Println("field", field)
	return nil
}
