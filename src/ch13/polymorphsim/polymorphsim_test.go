/* 面向对象：多态 polymorphsim */
package polymorphsim

import (
	"fmt"
	"testing"
)

type Code string
type Programmer interface {
	WriteHelloWorld() Code
}

type GoPragrammer struct {
}

func (p *GoPragrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World!\")"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World!\")"
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphsim(t *testing.T) {
	// 因为 GoPragrammer的 WriteHelloWorld() method has pointer receiver
	goProg := new(GoPragrammer) // 指针

	// GoPragrammer does not implement Programmer (WriteHelloWorld method has pointer receiver)
	// func (p GoPragrammer) WriteHelloWorld() Code  // 如果receiver不是指针类型即可
	//goProg := GoPragrammer{}

	javaProg := new(JavaProgrammer)

	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
}
