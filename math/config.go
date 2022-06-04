// Package math
// @Time  : 2022/6/4 23:19
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package math

// all integers
type number interface {
	pI | nI
}

// all positive integers.
type pI interface {
	int | int8 | int16 | int32 | int64
}

// all negative integers.
type nI interface {
	uint | uint8 | uint16 | uint32 | uint64
}
