// Package slice
// @Time  : 2022/5/25 15:33
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package slice

import (
	"math/rand"
)

// Reverse Any slice positions reverse.
func Reverse[T ~[]E, E any](arr T) {
	length := len(arr)
	for i := 0; i < length/2; i++ {
		j := length - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// CopyReverse copy Any slice and positions reverse
func CopyReverse[T ~[]E, E any](arr T) T {
	cr := make(T, len(arr))
	for j, i := 0, len(arr)-1; i >= 0; i, j = i-1, j+1 {
		cr[j] = arr[i]
	}
	return cr
}

// RemoveElement a value from a slice for removing elements.
// slice is comparable type.
func RemoveElement[T ~[]E, E comparable](s T, ele E) (t T) {
	if s == nil || len(s) == 0 {
		return
	}

	t = make(T, len(s)-1)
	index := 0

	for _, element := range s {
		if element != ele {
			t[index] = element
			index++
		}
	}

	return
}

// GetElementByRandom a value from a slice and random selection.
func GetElementByRandom[T ~[]E, E any](s T) (e E) {
	if s == nil || len(s) == 0 {
		return
	}
	length := len(s)
	index := rand.Intn(length)
	e = s[index]
	return
}
