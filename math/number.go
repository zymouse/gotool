// Package math
// @Time  : 2022/6/4 23:13
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package math

// IsOdd returns true if the n is an odd number,otherwise false.
func IsOdd[E number](n E) bool {
	return n&1 == 1
}
