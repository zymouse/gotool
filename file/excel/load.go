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

// LoadXlsx loading Excel file.only supports loading xlsx format file.
//
// must be implemented Xlsx interface.
func LoadXlsx[T Xlsx](filePath string) (t []T, err error) {
	err = excel.UnmarshalXLSX(filePath, &t)
	return
}
