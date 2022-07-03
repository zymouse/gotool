// Package excel
// @Time  : 2022/7/3 9:11
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package excel

import "github.com/szyhf/go-excel"

// Xlsx loading Excel must be implementation interface.
type Xlsx interface {
	// GetXLSXSheetName return need loading sheet name in the Excel.
	GetXLSXSheetName() string
}

// LoadExcel loading Excel file.
//
// must be implemented by Xlsx interface.
func LoadExcel[T Xlsx](filePath string) (t []T) {
	if err := excel.UnmarshalXLSX(filePath, &t); err != nil {
		panic(err)
	}
	return
}
