// Package excel
// @Time  : 2022/7/3 9:24
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package excel_test

import (
	"fmt"
	"github.com/jtyoui/gotool/excel"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	Name string `xlsx:"name"`
	Age  int    `xlsx:"age"`
	Sex  string `xlsx:"sex"`
}

func (t Test) GetXLSXSheetName() string {
	return "Sheet1"
}

func TestLoadExcel(t *testing.T) {
	data := excel.LoadExcel[Test]("test.xlsx")
	test := []Test{{Name: "张三", Age: 17, Sex: "男"}, {Name: "李四", Age: 18, Sex: "女"}}
	assert.Equal(t, data, test)
}

func ExampleLoadExcel() {
	// excel file struct
	/***
	type Test struct {
		Name string `xlsx:"name"`
		Age  int    `xlsx:"age"`
		Sex  string `xlsx:"sex"`
	}

	func (t Test) GetXLSXSheetName() string {
		return "Sheet1"
	}
	*/

	// Excel File view
	/***
	name	age	sex
	张三		17	男
	李四		18	女
	*/
	data := excel.LoadExcel[Test]("test.xlsx")
	fmt.Println(data)
	// Output:
	// [{张三 17 男} {李四 18 女}]
}
