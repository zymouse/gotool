package array_test

import (
	"fmt"
	"github.com/jtyoui/gotool/array"
	"github.com/stretchr/testify/assert"
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
		array.Reverse(data)
		assert.Equal(t, predict[i], data)
	}
}

func ExampleReverse() {
	data := []string{"a", "b", "c"}
	array.Reverse(data)
	fmt.Println(data)

	value := []int{1, 2, 3}
	array.Reverse(value)
	fmt.Println(value)
	// Output:
	// [c b a]
	// [3 2 1]
}

func ExampleCopyReverse() {
	data := []string{"a", "b", "c"}
	value := array.CopyReverse(data)
	fmt.Println(data)
	fmt.Println(value)
	// Output:
	// [a b c]
	// [c b a]
}
