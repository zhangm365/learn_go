package main

const (
	ErrNotFound     = DicError("could not find the key")
	ErrWordExists   = DicError("can't add the key because it already exists")
	ErrWordNotExist = DicError("can't update the key because it doesn't exist")
)

type DicError string

func (e DicError) Error() string {
	return string(e)
}

// map 是引用类型
type Dic map[string]string

// 查找给定 key 所对应的 value
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
func (d Dic) Add(key, value string) error {

	_, err := d.Search(key)
	switch err {
	case ErrNotFound: // 不存在则添加
		d[key] = value
	case nil: // 存在则返回提示
		return ErrWordExists
	default:
		return nil
	}

	return nil
}

// 更新 key 所对应的 value
func (d Dic) Update(key, value string) error {

	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		return ErrWordNotExist
	case nil: // 存在则更新
		d[key] = value
	default:
		return nil
	}

	return nil
}

// 删除
func (d Dic) Delete(key string) {
	// map 中有个内置 delete 函数，对不存在的 key 或 nil 不进行操作
	delete(d, key)
}
