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

func TestLoadJson(t *testing.T) {
	users, err := js.LoadJson[[]User]("test.json")
	assert.NoError(t, err)
	assert.Equal(t, len(users), 3)
	assert.Equal(t, users[0], User{Name: "ZhangWei", Age: 18})
}

func BenchmarkLoadJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = js.LoadJson[[]User]("test.json")
	}
}

func ExampleLoadJson() {
	/**
	 test.json

	[{"name":"ZhangWei","age":18},
	{"name":"ZhangWei","age":18},
	{"name":"ZhangWei","age":18}]
	*/

	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	users, _ := js.LoadJson[[]User]("test.json")
	fmt.Println(users)
	// Output:
	// [{ZhangWei 18} {ZhangWei 18} {ZhangWei 18}]
}
