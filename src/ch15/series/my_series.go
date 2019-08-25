package series

import (
	"errors"
	"fmt"
)

// 最佳实践1
// 定义不同的错误变量，以便判断错误类型
var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThanHundredError = errors.New("n should be not larger than 100")

func init() {
	fmt.Println("init1 in package series--my_series.go")
}

// 同一个源码文件中 可以有多个 init() 函数
func init() {
	fmt.Println("init2 in package series--my_series.go")
}

func GetFibonacci(n int) ([]int, error) {
	// 快速失败，把失败检测放在最前面
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThanHundredError
	}
	//if n < 2 || n > 100 {
	//	return nil, errors.New("n should bi in [2,100]")
	//}
	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

// 小写的不能被package以外调用
func square(n int) int {
	return n * n
}

// 大写 可以被package以外调用
func Square(n int) int {
	return n * n
}
