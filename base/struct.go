package base

import "fmt"

/*
结构体是值类型
*/
// 自定义类型
type myInt int
type myFun func(int, int) int

func (m myInt) MyIntFunc1() {
	fmt.Printf("this is custum func %T", m)
}

// 类型别名
type myFloat = float64

func MyType() {
	var a myInt = 10
	fmt.Printf("myInt: %T", a) // main.myInt

	var b myFloat = 10.1
	fmt.Printf("myFloat: %T", b) // float64
}

// 定义结构体
type Person struct {
	Name string // 公有属性
	sex  string
	age  int // 私有属性
}

// 定义结构体方法
func (p Person) SayName() {
	fmt.Printf("my name is %v\n", p.Name)
}

// 修改结构体属性
func (p *Person) SetInfo(Name string, age int) {
	p.Name = Name
	p.age = age
}

func TestStructFunc() {
	var p = Person{
		Name: "赵六",
		sex:  "女",
		age:  28,
	}
	p.SayName()
	p.SetInfo("wangwu", 27)
	p.SayName()
}

// 初始化结构体
func InitStruct() {
	// 方式一：基本的实例化形式
	var p1 Person
	p1.Name = "张三"
	p1.sex = "男"
	p1.age = 20
	fmt.Printf("p1:	%#v %T\n", p1, p1) // base.Person

	// 方式二：创建指针类型的结构体
	var p2 = new(Person)
	p2.Name = "李四"
	p2.sex = "男"
	p2.age = 24
	(*p2).age = 25                     // 使用指针改变属性值
	fmt.Printf("p2:	%#v %T\n", p2, p2) // *base.Person

	// 方式三：取结构体的地址实例化(对结构体进行&取地址操作时，视为对该类型进行一次 new 的实例化操作)
	var p3 = &Person{}
	p3.Name = "王五"
	p3.sex = "女"
	p3.age = 28
	fmt.Printf("p3:	%#v %T\n", p3, p3) // *base.Person

	// 方式四：
	var p4 = Person{
		Name: "赵六",
		sex:  "女",
		age:  28,
	}
	fmt.Printf("p4:	%#v %T\n", p4, p4) // base.Person

	// 方式五：
	var p5 = &Person{
		Name: "麻七", // 不初始化属性则为类型默认值
		sex:  "女",
		age:  28,
	}
	fmt.Printf("p5:	%#v %T\n", p5, p5) // *base.Person

	// 方式六：
	var p6 = &Person{
		"花八", // 不填键名时需按顺序
		"女",
		28,
	}
	fmt.Printf("p6:	%#v %T\n", p6, p6) // *base.Person

}
