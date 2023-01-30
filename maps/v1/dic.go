package main

// 哈希表
type Dic map[string]string

func (d Dic) Search(word string) string {
	return d[word]
}
