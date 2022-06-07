// Package slice
// @Time  : 2022/6/5 19:52
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package slice

// SetIntersection many slices exist intersection.
func SetIntersection[T ~[]E, E comparable](many ...T) (set T) {
	num := len(many)
	if num == 0 {
		return
	}

	// get all slice number.
	size := 0
	for _, m := range many {
		size += len(m)
	}

	// all keys and count the number of many
	manyMap := make(map[E]int, size)

	for _, slice := range many {
		remove := RemoveDuplicates(slice)
		for _, value := range remove {
			manyMap[value]++
		}
	}

	// returns unique values add to set if manyMap[key]==num
	for key, value := range manyMap {
		if value == num {
			set = append(set, key)
		}
	}
	return
}
