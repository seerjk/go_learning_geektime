package panic_recover

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	defer func() {
		fmt.Println("Finally! defer run")
	}()

	fmt.Println("Start")

	os.Exit(-1)
	//panic(errors.New("Something wrong!"))
}

func TestRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			// 恢复错误
			// err 是 panic 传递进去的
			fmt.Println("recovered from ", err)
		}
	}()

	fmt.Println("Start")
	panic(errors.New("something wrong"))
}
