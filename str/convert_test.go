// Package str_test
// @Time  : 2022/6/21 11:12
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package str_test

import (
	"fmt"
	"github.com/jtyoui/gotool/str"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTo(t *testing.T) {
	s1, _ := str.To[int]("65")
	assert.Equal(t, s1, 65)

	s2, _ := str.To[int8]("65")
	assert.Equal(t, s2, int8(65))

	s3, _ := str.To[int16]("65")
	assert.Equal(t, s3, int16(65))

	s4, _ := str.To[int32]("65")
	assert.Equal(t, s4, int32(65))

	s5, _ := str.To[int64]("65")
	assert.Equal(t, s5, int64(65))

	s6, _ := str.To[uint]("65")
	assert.Equal(t, s6, uint(65))

	s7, _ := str.To[uint8]("65")
	assert.Equal(t, s7, uint8(65))

	s8, _ := str.To[uint16]("65")
	assert.Equal(t, s8, uint16(65))

	s9, _ := str.To[uint32]("65")
	assert.Equal(t, s9, uint32(65))

	s10, _ := str.To[uint64]("65")
	assert.Equal(t, s10, uint64(65))

	s11, _ := str.To[float32]("65")
	assert.Equal(t, s11, float32(65))

	s12, _ := str.To[float64]("65")
	assert.Equal(t, s12, float64(65))

	s13, _ := str.To[string]("65")
	assert.Equal(t, s13, "65")

	s14, _ := str.To[bool]("true")
	assert.Equal(t, s14, true)

	s15, _ := str.To[bool]("false")
	assert.Equal(t, s15, false)

	s16, _ := str.To[bool]("1")
	assert.Equal(t, s16, true)

	s17, _ := str.To[bool]("0")
	assert.Equal(t, s17, false)
}

func ExampleTo() {
	s1, _ := str.To[int]("65")
	fmt.Println(s1)

	s2, _ := str.To[bool]("1")
	fmt.Println(s2)

	// Output:
	// 65
	// true
}
