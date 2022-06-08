// Package slice
// @Time  : 2022/6/8 11:06
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package slice_test

import (
	"fmt"
	"github.com/jtyoui/gotool/slice"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToSet(t *testing.T) {
	num := []int{1, 2, 3, 4}
	set := slice.ToSet(num)
	assert.Equal(t, set, map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}})

	var str []string
	value := slice.ToSet(str)
	assert.Equal(t, value, map[string]struct{}(nil))
}

func ExampleToSet() {
	num := []int{1, 2, 3}
	set := slice.ToSet(num)
	fmt.Println(set)
	// Output:
	// map[1:{} 2:{} 3:{}]
}
