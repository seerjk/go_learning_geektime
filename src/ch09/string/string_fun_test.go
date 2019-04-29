package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	// 分割
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	// 连接
	t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {
	// int --> string
	s := strconv.Itoa(10)
	t.Log("str" + s)

	// string --> int
	// func Atoi(s string) (int, error)
	// t.Log(10 + strconv.Atoi("10"))
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}

}
