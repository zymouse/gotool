// Package math
// @Time  : 2022/6/4 23:13
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package math

import "github.com/jtyoui/gotool/f"

// IsOdd returns true if the n is an odd number,otherwise false.
func IsOdd[E f.Number](n E) bool {
	return n&1 == 1
}
