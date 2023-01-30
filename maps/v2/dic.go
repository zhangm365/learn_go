package main

import "errors"

type Dic map[string]string

var ErrNotFound = errors.New("could not find the word")

func (d Dic) Search(word string) (string, error) {
	// map 查找的特性：第二个返回值是 bool，表示是否成功找到 key
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}
