package slice_test

import (
	"fmt"
	"github.com/jtyoui/gotool/slice"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestReverse(t *testing.T) {
	value := [][]any{
		{"a", "b", "c"},
		{1, 2, 3},
		{true, 1},
		{},
	}

	predict := [][]any{
		{"c", "b", "a"},
		{3, 2, 1},
		{1, true},
		{},
	}

	for i := 0; i < len(value); i++ {
		data := value[i]
		slice.Reverse(data)
		assert.Equal(t, predict[i], data)
	}
}

func ExampleReverse() {
	data := []string{"a", "b", "c"}
	slice.Reverse(data)
	fmt.Println(data)

	value := []int{1, 2, 3}
	slice.Reverse(value)
	fmt.Println(value)
	// Output:
	// [c b a]
	// [3 2 1]
}

func ExampleCopyReverse() {
	data := []string{"a", "b", "c"}
	value := slice.CopyReverse(data)
	fmt.Println(data)
	fmt.Println(value)
	// Output:
	// [a b c]
	// [c b a]
}

func TestRemoveElement(t *testing.T) {
	value := []int{1, 2, 3, 4}
	remove := slice.RemoveElement(value, 1)
	assert.Equal(t, remove, []int{2, 3, 4})

	var elements []string
	element := slice.RemoveElement(elements, "hello")
	assert.Equal(t, element, elements)

	num := []int{1, 1, 2, 3}
	nums := slice.RemoveElement(num, 1)
	assert.Equal(t, nums, []int{2, 3})
}

func BenchmarkRemoveElement(b *testing.B) {
	s := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	for i := 0; i < b.N; i++ {
		slice.RemoveElement(s, "a")
	}
}

func ExampleRemoveElement() {
	value := []int{1, 2, 3, 4}
	remove := slice.RemoveElement(value, 1)
	fmt.Println(remove)
	// Output:
	// [2 3 4]
}

func TestGetElementByRandom(t *testing.T) {
	rand.Seed(100)
	s := []int{1, 2, 3, 4}
	ele := slice.GetElementByRandom(s)
	assert.Equal(t, ele, 4)

	var p []string
	v := slice.GetElementByRandom(p)
	assert.Equal(t, v, "")
}

func ExampleGetElementByRandom() {
	rand.Seed(100)

	s := []int{1, 2, 3, 4}
	ele := slice.GetElementByRandom(s)
	fmt.Println(ele)

	values := []string{"A", "B", "C"}
	value := slice.GetElementByRandom(values)
	fmt.Println(value)

	// Output:
	// 4
	// C
}

func ExampleRemoveDuplicates() {
	data := []int{1, 1, 2, 3, 4, 5, 5, 1}
	value := slice.RemoveDuplicates(data)
	fmt.Println(value)
	// Output:
	// [1 2 3 4 5]
}

func TestRemoveDuplicates(t *testing.T) {
	data := []int{1, 1, 2, 3, 4, 5, 5, 1}
	value := slice.RemoveDuplicates(data)
	assert.Equal(t, value, []int{1, 2, 3, 4, 5})
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	s := make([]int64, 10000)

	for i := 0; i < 10000; i++ {
		s[i] = rand.Int63()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice.RemoveDuplicates(s)
	}
}
