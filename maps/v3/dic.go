package main

import "errors"

// map 是引用类型
type Dic map[string]string

var ErrNotFound = errors.New("could not find the word")

func (d Dic) Search(word string) (string, error) {
	// map 查找的特性：第二个返回值是 bool，表示是否成功找到 key
	definition, ok := d[word]
	// 没有找到
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// 向 map 添加 [key, value]
func (d Dic) Add(key, value string) {
	d[key] = value
}
