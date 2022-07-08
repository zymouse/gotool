// Package json
// @Time  : 2022/7/8 9:58
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package js

import (
	"bufio"
	"encoding/json"
	"os"
)

// SaveJson Big file save file to json format.
func SaveJson(path string, data any) (err error) {
	// create new file or truncate old file
	var file *os.File
	if file, err = os.Create(path); err != nil {
		return
	}

	// add bufio io.Writer
	writer := bufio.NewWriter(file)

	// encode to data
	encoder := json.NewEncoder(writer)

	if err = encoder.Encode(data); err != nil {
		return
	}

	// flush to file
	if err = writer.Flush(); err != nil {
		return
	}
	err = file.Close()
	return
}
