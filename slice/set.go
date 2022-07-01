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
	size := Len(many)

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

// ContainsAny there are one element in the data, will return true. otherwise false.
func ContainsAny[T ~[]E, E comparable](data T, elements ...E) (ok bool) {
	// get set from data slice
	m := ToSet(data)

	if m == nil {
		return
	}

	for i := 0; i < len(elements); i++ {
		if _, ok = m[elements[i]]; ok {
			return
		}
	}

	return
}

// ContainsAll all element are in the data, will return true. otherwise false.
func ContainsAll[T ~[]E, E comparable](data T, elements ...E) (ok bool) {
	// get set from data slice
	m := ToSet(data)

	if m == nil {
		return
	}

	for i := 0; i < len(elements); i++ {
		if _, ok = m[elements[i]]; !ok {
			return
		}
	}

	return
}

// In an element in the data, will return true. otherwise false.
//
// slice.In comparable slice.ContainsAll more fast.
func In[T ~[]E, E comparable](data T, e E) (ok bool) {
	length := len(data)

	for i := 0; i < length; i++ {
		if data[i] == e {
			ok = true
			return
		}
	}
	return
}

// Contains returns true if the element is in the data.
func Contains[T ~[]E, E comparable](data T, e E) (ok bool) {
	return In(data, e)
}
