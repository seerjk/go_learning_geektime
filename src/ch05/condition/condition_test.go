package condition_test

import "testing"

func TestIfMultiSec(t *testing.T) {
	/*
		if v, err := someFunc(); err == nil {
			// do
		} else {
			// do
		}
	*/

	if a := 1 == 1; a {
		t.Log("1==1")
	}
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
			// 末尾自带break
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("unknown")
		}
	}
}
