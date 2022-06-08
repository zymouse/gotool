// Package slice
// @Time  : 2022/6/7 15:53
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package slice_test

import (
	"fmt"
	"github.com/jtyoui/gotool/slice"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetIntersection(t *testing.T) {
	data := [][]string{{"a", "b", "c"}, {"a", "d"}, {"a", "f"}}
	set := slice.SetIntersection(data...)
	assert.Equal(t, set, []string{"a"})

	data = [][]string{{"a", "b", "c"}, {"a", "d"}, {"a", "f"}, {"g"}}
	set = slice.SetIntersection(data...)
	assert.Equal(t, set, []string(nil))
}

func ExampleSetIntersection() {
	data := [][]string{{"a", "b", "c"}, {"a", "d"}, {"a", "f"}}
	set := slice.SetIntersection(data...)
	fmt.Println(set)
	// Output:
	// [a]
}

func TestContainsAny(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	flag := slice.ContainsAny(data, 1, 2, 13)
	assert.Equal(t, flag, true)
	flag = slice.ContainsAny(data, 7, 8, 9)
	assert.Equal(t, flag, false)
}

func TestContainsAll(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	flag := slice.ContainsAll(data, 1, 2, 3)
	assert.Equal(t, flag, true)
	flag = slice.ContainsAll(data, 1, 2, 6)
	assert.Equal(t, flag, false)
}

func ExampleContainsAny() {
	data := []int{1, 2, 3, 4, 5}
	flag := slice.ContainsAny(data, 1, 13)
	fmt.Println(flag)

	flag = slice.ContainsAny(data, 6, 13)
	fmt.Println(flag)
	// Output:
	// true
	// false
}

func ExampleContainsAll() {
	data := []int{1, 2, 3, 4, 5}
	flag := slice.ContainsAll(data, 1, 13)
	fmt.Println(flag)

	flag = slice.ContainsAll(data, 1, 2)
	fmt.Println(flag)
	// Output:
	// false
	// true
}
