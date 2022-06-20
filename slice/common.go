// Package slice
// @Time  : 2022/6/7 15:06
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package slice

import (
	"reflect"
)

// MaxLen gets the maximum length of many slice.
// returns size and max slice.
func MaxLen[T ~[]E, E any](s ...T) (size int, value T) {
	num := len(s)

	if num == 0 {
		return
	}

	size, value = len(s[0]), s[0]

	for i := 1; i < num; i++ {
		length := len(s[i])

		if length > size {
			size = length
			value = s[i]
		}
	}
	return
}

// MinLen gets the min length of many slice.
// returns size and min slice.
func MinLen[T ~[]E, E any](s ...T) (size int, value T) {
	num := len(s)

	if num == 0 {
		return
	}

	size, value = len(s[0]), s[0]

	for i := 1; i < num; i++ {
		length := len(s[i])

		if length < size {
			size = length
			value = s[i]
		}
	}
	return
}

// Len count all elements number of many slice or ptr slice.
func Len(data any) (size int) {
	value := reflect.ValueOf(data)
	typ := value.Type()

	// if data is ptr
	if typ.Kind() == reflect.Pointer {
		value = value.Elem()
		typ = value.Type()
	}

	if typ.Kind() == reflect.Slice {
		for i := 0; i < value.Len(); i++ {
			next := value.Index(i).Interface()
			size += Len(next)
		}
		return
	}
	size++
	return
}
