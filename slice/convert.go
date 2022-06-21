// Package slice
// @Time  : 2022/6/8 11:02
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package slice

// ToSet convert slice to set
func ToSet[T ~[]E, E comparable](slice T) map[E]struct{} {
	// if data are not elements. return nil.
	if len(slice) == 0 {
		return nil
	}

	// all value from data stored in the set map.
	m := make(map[E]struct{}, len(slice))
	for _, value := range slice {
		m[value] = struct{}{}
	}
	return m
}
