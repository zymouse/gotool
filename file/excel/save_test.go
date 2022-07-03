// Package excel
// @Time  : 2022/7/3 22:17
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package excel_test

import (
	"fmt"
	"github.com/jtyoui/gotool/file/excel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveXlsx(t *testing.T) {
	values := []Test{{Name: "张三", Age: 17, Sex: "男"}, {Name: "李四", Age: 18, Sex: "女"}}
	err := excel.SaveExcel("test.xlsx", values)
	assert.NoError(t, err)

	data, _ := excel.LoadXlsx[Test]("test.xlsx")
	assert.Equal(t, data, values)
}

func ExampleSaveExcel() {
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

	values := []Test{{Name: "张三", Age: 17, Sex: "男"}, {Name: "李四", Age: 18, Sex: "女"}}
	/***
	name	age	sex
	张三		17	男
	李四		18	女
	*/
	err := excel.SaveExcel("test.xlsx", values)
	fmt.Println(err)
	// Output:
	// <nil>
}
