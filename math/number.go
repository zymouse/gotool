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

// BtoI converts a bool to an integer.
func BtoI(b bool) int {
	if b {
		return 1
	}
	return 0
}

//Min returns the minimum value of the given numbers.
func Min(numbers ...int) int {
	if len(numbers) == 0 {
		panic("numbers must not be empty")
	}
	x := numbers[0]
	for i := 1; i < len(numbers); i++ {
		y := numbers[i]
		x = y ^ ((x ^ y) & -BtoI(x < y)) // y^x^y=x if x < y else y^0=y
	}
	return x
}

// Max returns the maximum value of the given numbers
func Max(numbers ...int) int {
	if len(numbers) == 0 {
		panic("numbers must not be empty")
	}
	x := numbers[0]
	for i := 1; i < len(numbers); i++ {
		y := numbers[i]
		x = x ^ ((x ^ y) & -BtoI(x < y)) // x^x^y=y if x < y else x^0=x
	}
	return x
}

// Abs returns the absolute value of x.
func Abs[T f.RealNumber](x T) T {
	if x < 0 {
		return -x
	}
	return x
	//  x ^ (x >> 31)
}
