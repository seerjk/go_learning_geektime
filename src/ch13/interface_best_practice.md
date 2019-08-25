# Go 接口最佳实践

## 1.倾向于使用小的接口定义，很多接口只包含一个方法

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```


## 2.较大的接口定义，可以由多个小接口定义组合而成

```go
type ReadWriter interface {
    Reader
    Writer
}
```


## 3.只依赖于必要功能的最小接口

```go
func StoreData(reader Reader) error {
    …
}
```
