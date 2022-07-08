// Package js
// @Time  : 2022/7/8 14:06
// @Email: jtyoui@qq.com
// @Author: ZhangWei
package js

import (
	"encoding/json"
	"os"
)

type entry[T any] struct {
	Error error // if stream is fails. an error will be present.
	Data  T     // read struct data
}

type bigJsonStream[T any] struct {
	stream chan entry[T]
	path   string
}

// NewBigJsonStream create a big json stream.
func NewBigJsonStream[T any](path string, cache int) *bigJsonStream[T] {
	return &bigJsonStream[T]{
		stream: make(chan entry[T], cache),
		path:   path,
	}
}

// Watch watches the json file streams. Each stream entry will either have an error
// or data object.
// if error present as the `Start` method after. will close channel automatically.
func (s *bigJsonStream[T]) Watch() <-chan entry[T] {
	return s.stream
}

// Start starts the json file stream. read line by line.
// if an error occurs, the channel will be closed.
func (s *bigJsonStream[T]) Start() {
	defer close(s.stream)
	file, err := os.Open(s.path)
	if err != nil {
		s.stream <- entry[T]{Error: err}
		return
	}

	decoder := json.NewDecoder(file)

	if _, err = decoder.Token(); err != nil {
		s.stream <- entry[T]{Error: err}
		return
	}

	for decoder.More() {
		var data T
		if err = decoder.Decode(&data); err != nil {
			s.stream <- entry[T]{Error: err}
			return
		}

		s.stream <- entry[T]{Data: data}
	}

	if _, err = decoder.Token(); err != nil {
		s.stream <- entry[T]{Error: err}
		return
	}

	if err = file.Close(); err != nil {
		s.stream <- entry[T]{Error: err}
		return
	}

}
