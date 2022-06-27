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

func TestBtoI(t *testing.T) {
	assert.Equal(t, math.BtoI(true), 1)
	assert.Equal(t, math.BtoI(false), 0)
}

func ExampleBtoI() {
	fmt.Println(math.BtoI(true))

	fmt.Println(math.BtoI(false))
	// Output:
	// 1
	// 0
}

func TestMin(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 0, -1, 29}
	n := math.Min(numbers...)
	assert.Equal(t, n, -1)
}

func BenchmarkMin(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5, 0, -1, 29}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		math.Min(numbers...)
	}
}

func ExampleMin() {
	fmt.Println(math.Min(1, 2, 3, -9, -8, 10))
	// Output:
	// -9
}

func TestMax(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 0, -1, 29}
	n := math.Max(numbers...)
	assert.Equal(t, n, 29)
}

func BenchmarkMax(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5, 0, -1, 29}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		math.Max(numbers...)
	}
}

func ExampleMax() {
	fmt.Println(math.Max(1, 2, 3, -9, -8, 10))
	// Output:
	// 10
}
