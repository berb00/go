package base

import (
	"fmt"
)
func DefPoint () {
	//指针取值
	a := 10
	b := &a // 取地址符	取变量a的地址，将指针保存到b中
	c := *b // 取值符	指针取值（根据指针去内存取值）
	fmt.Printf("type of b:%T\n", b)
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
}

// 指针作为参数
func PointParams () {
	a := 10
	modify1(a)
	fmt.Println(a) // 10

	modify2(&a)
	fmt.Println(a) // 100
}

func modify1(x int) {
	x = 100
}
func modify2(x *int) {
	*x = 100
}


// new函数
func TestNew () {
	/*
	表达式new(T)将创建一个T类型的匿名变量，
	为T类型的新值分配并清零一块内存空间，然后将这块内存空间的地址作为结果返回，
	这个结果就是指向这个新的T类型值的指针值，返回的指针类型为*T。

	new创建的内存空间位于heap上，
	空间的默认值为数据类型默认值。如：new(int) 则 *p为0，new(bool) 则 *p为false。

	func new(Type) *Type
	*/
	var a *int
	a = new(int)		//	*int
	fmt.Printf("new(type) - %T\n", a)
	*a = 10
	fmt.Println(*a)
}

func TestMake () {
	/*
	只用于slice、map以及chan的内存创建，返回引用类型本身

	func make(Type, size IntegerType) Type
	*/
	var b map[string]int
	b = make(map[string]int, 10)
	b["沙河娜扎"] = 100
	fmt.Println(b)
}