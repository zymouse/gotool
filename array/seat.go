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
func Reverse[T ~[]E, E any](arr T) T {
	length := len(arr)
	for i := 0; i < length/2; i++ {
		j := length - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
