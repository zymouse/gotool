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
	s := [][]any{{"a"}, {true, false}, {}, {1, 2, 3}}
	size, data := slice.MaxLen(s...)
	assert.Equal(t, size, 3)
	assert.Equal(t, data, []any{1, 2, 3})

	var i []int
	a, b := slice.MaxLen(i)
	assert.Equal(t, a, 0)
	assert.Equal(t, b, []int(nil))
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

func TestMinLen(t *testing.T) {
	s := [][]any{{1, 2, 3}, {"a"}, {true, false}}
	size, data := slice.MinLen(s...)
	assert.Equal(t, size, 1)
	assert.Equal(t, data, []any{"a"})

	var i []any
	size, data = slice.MinLen(i)
	assert.Equal(t, size, 0)
	assert.Equal(t, data, []interface{}(nil))
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

func TestLen(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	s := slice.Len(data)
	assert.Equal(t, s, len(data))

	other := [][]int{{1, 2}, {4, 5, 6}, {1, 5, 8, 3}}
	s1 := slice.Len(&other)
	assert.Equal(t, s1, 9)
}

func ExampleLen() {
	other := []any{1, []any{1, 2, 3}, []string{"a", "b"}}
	s := slice.Len(&other)
	fmt.Println(s)
	// Output:
	// 6
}

func TestCloneSlice(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	s := slice.CloneSlice(data)
	assert.Equal(t, s, data)

	data[1] = 100
	assert.Equal(t, s, []int{1, 2, 3, 4, 5})
	assert.Equal(t, data, []int{1, 100, 3, 4, 5})
}

func ExampleCloneSlice() {
	data := []int{1, 2, 3, 4, 5}
	s := slice.CloneSlice(data)
	fmt.Println(s)
	// Output:
	// [1 2 3 4 5]
}
