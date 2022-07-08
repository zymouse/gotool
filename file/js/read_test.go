// Package js
// @Time  : 2022/7/8 14:18
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package js_test

import (
	"fmt"
	"github.com/jtyoui/gotool/file/js"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBigJsonStream(t *testing.T) {
	stream := js.NewBigJsonStream[User]("test.json", 100)
	var users []User

	stop := make(chan struct{})

	go func() {
		for entry := range stream.Watch() {
			assert.NoError(t, entry.Error)
			users = append(users, entry.Data)
		}

		stop <- struct{}{}
	}()
	stream.Start()
	<-stop

	assert.Equal(t, len(users), 3)
	assert.Equal(t, users[0], User{Name: "ZhangWei", Age: 18})
}

func BenchmarkNewBigJsonStream(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stream := js.NewBigJsonStream[User]("test.json", 100)
		var users []User
		stop := make(chan struct{})

		go func() {
			for entry := range stream.Watch() {
				users = append(users, entry.Data)
			}

			stop <- struct{}{}
		}()
		stream.Start()
		<-stop
	}
}

func ExampleNewBigJsonStream() {
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

	var users []User

	stream := js.NewBigJsonStream[User]("test.json", 100)
	stop := make(chan struct{})

	go func() {
		for entry := range stream.Watch() {
			if entry.Error != nil {
				fmt.Println(entry.Error)
				continue
			}
			users = append(users, entry.Data)
		}
		stop <- struct{}{}
	}()
	stream.Start()
	<-stop
	fmt.Println(users)
	// Output:
	// [{ZhangWei 18} {ZhangWei 18} {ZhangWei 18}]
}
