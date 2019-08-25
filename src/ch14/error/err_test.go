package err_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

// 最佳实践1
// 定义不同的错误变量，以便判断错误类型
var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThanHundredError = errors.New("n should be not larger than 100")

//func GetFibonacci(n int) []int {
//	fibList := []int{1, 1}
//
//	for i := 2; i < n; i++ {
//		fibList = append(fibList, fibList[i-2]+fibList[i-1])
//	}
//	return fibList
//}

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

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(1); err != nil {
		if err == LessThanTwoError {
			fmt.Println("It is less.")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
	//t.Log(GetFibonacci(-10))
}

// 最佳实践2
// 及早失败，避免嵌套

func GetFibonacci1(str string) {
	var (
		i    int
		err  error
		list []int
	)

	// 不好的实践：
	// 错误多级嵌套，正向考虑无错误的情况
	if i, err = strconv.Atoi(str); err == nil {
		if list, err = GetFibonacci(i); err == nil {
			fmt.Println(list)
		} else {
			fmt.Println("fib error: ", err)
		}
	} else {
		fmt.Println("atoi error: ", err)
	}
}

func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)
	// 最佳实际：及早失败，避免嵌套
	// 失败后return，避免多级嵌套造成代码结构复杂
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("atoi error:", err)
		return
	}
	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("fib error", err)
		return
	}

	fmt.Println(list)
}

func TestGetFibo(t *testing.T) {
	//GetFibonacci1("a")
	GetFibonacci2("a")
}
