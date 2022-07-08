// Package tool
// @Time  : 2022/6/27 16:05
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package tool

// IF  returns trueVal is cond is true, otherwise falseVal.
func IF[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}
