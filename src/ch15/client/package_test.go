package client

import (
	"testing"

	"github.com/seerjk/go_learning_geektime/src/ch15/series"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacci(5))

	//series.square(5)
	t.Log(series.Square(5))
}
