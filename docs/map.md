

# map 数据结构

`map` 是一种存储键值对的哈希表。

## 1. 定义

键值类型 `KeyType` 是唯一的，且是[可比较的]{https://go.dev/ref/spec#Comparison_operators}类型，因此它不能是函数、切片和另一个 `map`。

值类型 `EleType` 是可以任意类型。

定义形式：

```go
map[KeyType]EleType


// 1. 定义一个 nil 的 map，会得到一个 nil 指针异常
var m map[Type]Type

// panic: assignment to entry in nil map
m[key] = value

// 2. 正确的方法

m := map[Type]Type{}
// OR 
m := make(map[Type]Type)

```

## 2. 特征

### 2.1 map 是引用类型，因此只有一个副本；

程序运行效率高，尤其是函数传递的场景。

### 2.2 map 查找的第二个值 `bool`，表示是否成功找到 `key`。

如以下返回值： `ok`

```go

// 空 map
dic := map[string]string{}

value, ok := dic["key"]
if !ok {
    fmt.Println("Don't find key")
}

```
