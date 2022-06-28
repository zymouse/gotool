// Package f
// @Time  : 2022/6/21 11:05
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package f

type Integer interface {
	int | int8 | int16 | int32 | int64
}

type UInteger interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Floats interface {
	float32 | float64
}

type TypeConvert interface {
	string | bool | Integer | UInteger | Floats
}

type Number interface {
	Integer | UInteger
}

type RealNumber interface {
	Number | Floats
}
