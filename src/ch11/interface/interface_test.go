package interface_test

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {

}

// duck type 方法签名看起来是一样的
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}