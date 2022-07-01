// Package slice
// @Time  : 2022/7/1 13:44
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package slice_test

import (
	"fmt"
	"github.com/jtyoui/gotool/slice"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoves(t *testing.T) {
	s := []string{"a", "b", "c", "d"}
	r := slice.Removes(s, "b", "c")
	assert.Equal(t, r, []string{"a", "d"})
	assert.Equal(t, s, []string{"a", "b", "c", "d"})
}

func BenchmarkRemoves(b *testing.B) {
	s := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	for i := 0; i < b.N; i++ {
		slice.Removes(s, "b", "d", "f", "h")
	}
}

func ExampleRemoves() {
	s := []string{"a", "b", "c", "d"}
	r := slice.Removes(s, "b", "c")
	fmt.Println(r)
	// Output: [a d]
}

func TestRemove(t *testing.T) {
	s := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	r := slice.Remove(s, "b")
	assert.Equal(t, r, []string{"a", "c", "d", "e", "f", "g", "h", "i", "j"})
	assert.Equal(t, s, []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"})
}

func BenchmarkRemove(b *testing.B) {
	s := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	for i := 0; i < b.N; i++ {
		slice.Remove(s, "b")
	}
}

func ExampleRemove() {
	s := []string{"a", "b", "c", "d"}
	r := slice.Remove(s, "b")
	fmt.Println(r)
	// Output: [a c d]
}

func TestRemoveIF(t *testing.T) {
	s := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	r := slice.RemoveIF(s, func(v string) bool { return v == "b" })
	assert.Equal(t, r, []string{"a", "c", "d", "e", "f", "g", "h", "i", "j"})

	r = slice.RemoveIF(s, func(v string) bool { return v != "b" })
	assert.Equal(t, r, []string{"b"})
}

func BenchmarkRemoveIF(b *testing.B) {
	s := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	for i := 0; i < b.N; i++ {
		slice.RemoveIF(s, func(v string) bool { return v == "b" })
	}
}

func ExampleRemoveIF() {
	s := []string{"a", "b", "c", "d", "b"}
	r := slice.RemoveIF(s, func(v string) bool { return v == "b" })
	fmt.Println(r)
	// Output: [a c d]
}

func TestClone(t *testing.T) {
	s := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	r := slice.Clone(s)
	assert.Equal(t, s, r)
	assert.NotEqual(t, fmt.Sprintf("%p", s), fmt.Sprintf("%p", r))
}

func ExampleClone() {
	s := []string{"a", "b", "c", "d"}
	r := slice.Clone(s)
	fmt.Println(r)
	// Output: [a b c d]
}
