// Package tool
// @Time  : 2022/6/27 16:05
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package tool_test

import (
	"fmt"
	"github.com/jtyoui/gotool/tool"
	"testing"
)

func TestIF(t *testing.T) {
	if tool.IF(true, 1, 0) != 1 {
		t.Error("IF(true, 1, 0) != 1")
	}
	if tool.IF(false, 1, 0) != 0 {
		t.Error("IF(false, 1, 0) != 0")
	}
}

func ExampleIF() {
	fmt.Println(tool.IF(true, 1, 0))
	fmt.Println(tool.IF(false, 1, 0))
	// Output:
	// 1
	// 0
}
