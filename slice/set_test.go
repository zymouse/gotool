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
