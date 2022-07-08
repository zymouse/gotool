// Package json
// @Time  : 2022/7/8 9:59
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package js_test

import (
	"fmt"
	"github.com/jtyoui/gotool/file/js"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestSaveJSON(t *testing.T) {
	users := make([]User, 3)
	for i := 0; i < 3; i++ {
		users[i] = User{Name: "ZhangWei", Age: 18}
	}
	err := js.SaveJson("test.json", users)
	assert.NoError(t, err)
}

func BenchmarkSaveJson(b *testing.B) {
	users := make([]User, 3)
	for i := 0; i < 3; i++ {
		users[i] = User{Name: "ZhangWei", Age: 18}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = js.SaveJson("test.json", users)
	}
}

func ExampleSaveJson() {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	users := make([]User, 3)
	for i := 0; i < 3; i++ {
		users[i] = User{Name: "ZhangWei", Age: 18}
	}
	err := js.SaveJson("test.json", users)
	fmt.Println(err)

	// Output:
	// <nil>
}
