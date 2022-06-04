// Package math
// @Time  : 2022/6/4 23:25
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package math_test

import (
	"fmt"
	"github.com/jtyoui/gotool/math"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsOdd(t *testing.T) {
	var i = 11
	var j int8 = 13
	var p int16 = -2321
	var k int32 = 34322
	var h int64 = 873215

	assert.Equal(t, math.IsOdd(i), true)
	assert.Equal(t, math.IsOdd(j), true)
	assert.Equal(t, math.IsOdd(p), true)
	assert.Equal(t, math.IsOdd(k), false)
	assert.Equal(t, math.IsOdd(h), true)
}

func ExampleIsOdd() {
	var i = 11
	fmt.Print(math.IsOdd(i))
	// Output: true
}
