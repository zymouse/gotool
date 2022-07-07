// Package slice
// @Time  : 2022/7/1 13:33
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package slice

// Removes all elements from the slice that satisfy the predicate.
func Removes[T ~[]E, E comparable](s T, r ...E) T {
	set := ToSet(r)
	return RemoveIF(s, func(v E) bool {
		_, ok := set[v]
		return ok
	})
}

// Remove removes the first element from the slice that satisfies the predicate.
func Remove[T ~[]E, E comparable](s T, e E) T {
	return RemoveIF(s, func(v E) bool { return v == e })
}

// RemoveIF removes the first element from the slice that if filtered.Func == true
// filter is a function that returns true if the element should be removed.
func RemoveIF[T ~[]E, E comparable](s T, filter func(E) bool) T {
	if s == nil || filter == nil {
		return s
	}
	n := make(T, 0, len(s))
	for i := 0; i < len(s); i++ {
		if !filter(s[i]) {
			n = append(n, s[i])
		}
	}
	return n
}

// Clone returns a copy of the slice.
func Clone[T ~[]E, E comparable](s T) T {
	return append(T(nil), s...)
}
