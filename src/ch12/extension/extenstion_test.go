/*
扩展和复用

Go 不支持继承，但可以通过复合的方式来复用
*/
package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

/*
// 复合，子类的方法都需要重新定义
type Dog struct {
	p *Pet
}

func (d *Dog) Speak() {
	//d.p.Speak()
	fmt.Println("Wang!")
}

func (d *Dog) SpeakTo(host string) {
	//d.p.SpeakTo(host)
	d.Speak()
	fmt.Println(" ", host)
}
*/

// 匿名类型嵌入
/*
匿名类型嵌入

它不是继承，如果把"内部struct"看作弗雷，把"外部struct"看作子类，会发现如下问题：
1.不支持子类替换
2.子类并不是真正继承了父类的方法
	* 父类的定义的方法无法访问子类的数据和方法

*/
type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	fmt.Println("Wang!")
}

func TestDog(t *testing.T) {
	dog := new(Dog)

	// go不支持显示类型转换，不支持继承
	//var dog *Pet
	//dog = new(Dog)

	dog.Speak()
	// 重载不成功  成功应该重装父类Pet的 Speak()  Wang! Chao
	dog.SpeakTo("Chao")
}
