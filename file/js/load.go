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

// LoadJson load json file to slice structs.
func LoadJson[T ~[]E, E any](path string) (t T, err error) {
	// Read file pointer
	var file *os.File

	if file, err = os.Open(path); err != nil {
		return
	}

	// create read bufio
	r := bufio.NewReader(file)

	d := json.NewDecoder(r)

	if _, err = d.Token(); err != nil {
		return
	}

	for d.More() {
		var data E
		if err = d.Decode(&data); err != nil {
			return
		}
		t = append(t, data)
	}

	if _, err = d.Token(); err != nil {
		return
	}

	err = file.Close()
	return
}
