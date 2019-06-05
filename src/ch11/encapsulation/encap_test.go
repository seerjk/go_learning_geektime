package encap

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

// 通常情况下为了了避免内存拷⻉贝我们使⽤用第⼆二种定义⽅方式
func (e *Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

// 第⼀一种定义⽅方式在实例例对应⽅方法被调⽤用时，实例例的成员会进⾏行行值复制
//func (e Employee) String() string {
//	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
//}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee) // 注意这⾥里里返回的引⽤用/指针，相当于 e := &Employee{}
	e2.Id = "2"         //与其他主要编程语⾔言的差异：通过实例例的指针访问成员不不需要使⽤用->
	e2.Age = 22
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)
}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
