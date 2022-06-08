// Package slice
// @Time  : 2022/6/7 15:11
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package slice_test

import (
	"fmt"
	"github.com/jtyoui/gotool/slice"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxLen(t *testing.T) {
	s := [][]any{{1, 2, 3}, {"a"}, {true, false}, {}}
	size, data := slice.MaxLen(s...)
	assert.Equal(t, size, 3)
	assert.Equal(t, data, []any{1, 2, 3})
}

func TestMinLen(t *testing.T) {
	s := [][]any{{1, 2, 3}, {"a"}, {true, false}}
	size, data := slice.MinLen(s...)
	assert.Equal(t, size, 1)
	assert.Equal(t, data, []any{"a"})
}

func ExampleMaxLen() {
	s := [][]any{{1, 2, 3}, {"a"}, {true, false}}
	size, data := slice.MaxLen(s...)
	fmt.Println(size)
	fmt.Println(data)

	// Output:
	// 3
	// [1 2 3]
}

func ExampleMinLen() {
	s := [][]any{{1, 2, 3}, {"a"}, {true, false}}
	size, data := slice.MinLen(s...)
	fmt.Println(size)
	fmt.Println(data)

	// Output:
	// 1
	// [a]
}
