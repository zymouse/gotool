// Package array
// @Time  : 2022/5/25 15:33
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package array

/*************************************************************
  Two array/slice changes between positions.It contains:
	Reverse; Remove; Checked index; Contains element
	Convert; Sort; Join ......
*************************************************************/

// Reverse Any array/slice positions reverse.
func Reverse[T ~[]E, E any](arr T) {
	length := len(arr)
	for i := 0; i < length/2; i++ {
		j := length - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// CopyReverse copy Any array/slice and positions reverse
func CopyReverse[T ~[]E, E any](arr T) T {
	var cr = make(T, len(arr))
	for j, i := 0, len(arr)-1; i >= 0; i, j = i-1, j+1 {
		cr[j] = arr[i]
	}
	return cr
}
