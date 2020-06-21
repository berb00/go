package base // 标明该文件属于base包
import (
	"fmt"
)

/*
变量声明

布尔类型
	bool	只有 true 和 false，默认为 false。
数字类型
	int							默认 0
		int8 	uint8
		int16 	uint16
		int32 	uint32
		int64 	uint64
		int 	uint	具体长度取决于 CPU 位数
	float
		float32 	float64
		complex64 	complex128
字符串类型
	string		""、`` 		不能使用''

派生类型



值类型:
	int			0
	float		0
	string		""
	bool		false
	array		nil
	struct
引用类型:
	slice		nil
	function	nil
	map			nil
	pointer		nil
	channel		nil
	interface	nil






*/

var PublicVar = 10  // 公有变量
var privateVar = 20 // 私有变量

func Printf() {
	var num int = 10
	fmt.Printf("%T\n", num) // 变量类型
	fmt.Printf("%v\n", num) // 变量值
	fmt.Printf("%d\n", num) // 十进制
	fmt.Printf("%b\n", num) // 二进制
	fmt.Printf("%o\n", num) // 八进制
	fmt.Printf("%x\n", num) // 十六进制

	var str string = "sss"
	fmt.Printf("%s\n", str)  // 字符串
	fmt.Printf("%v\n", str)  // 字符串
	fmt.Printf("%#v\n", str) // 加双引号的字符串

}

func VarDefault() { // 模块中要导出的函数，必须首字母大写
	var booldefault bool // bool 默认值
	fmt.Println("bool default:		", booldefault)

	var strdefault string // string 默认值
	fmt.Println("string default:	", strdefault)

	var intdefault int // string 默认值
	fmt.Println("int default:		", intdefault)

	var float32default float32 // float32 默认值
	fmt.Println("float32 default:	", float32default)

	var slicedefault []int // slice 默认值
	fmt.Println("slice default:		", slicedefault)

}

func VarString() { // 模块中要导出的函数，必须首字母大写
	str0 := 's'             // 声明初始化字符 ascII:115
	var str1 string = "sss" // 单个声明初始化
	str4 := "sss"           // 简短声明初始化 只能在函数内使用
	var (                   // 批量声明初始化
		str2 string = "str2"
		str3 string = "str3"
	)
	fmt.Println(str0, str1, str2, str3, str4)

	// 字符串拼接
	str6 := str1 + str2
	str7 := fmt.Sprintf("%s%s", str1, str2) // 返回拼接字符串
	fmt.Println(str6)
	fmt.Println(str7)

	// 保持字符串格式
	var str5 string = `		
	大江东去
		浪淘尽
			故垒西边
	`
	fmt.Printf(str5)       // 字符串长度
	fmt.Println(len(str5)) // 字符串长度
}

func VarInt() {

}
func VarFloat() {
	// var a float32 = 3.22
	// fmt.Println(a)
}
func VarBool() {
	// var a bool = false
	// fmt.Println(a)
}

func VarConst() {
	const (
		a = 10
		b // 只声明不赋值 默认和上一个变量值相同 10
		c
	)

	const (
		n = iota
		m // 只声明不赋值 默认和上一个变量值相同 10
		j
	)

	const (
		a1 = iota //	0	iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)
		b1        //	1	const中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)
		c1        //	2
		d  = "ha" //	独立值，iota += 1
		e         //	"ha"   iota += 1
		f  = 100  //	iota +=1
		g         //	100  iota +=1
		h  = iota //	7,恢复计数
		i         //	8
	)
	fmt.Println(a1, b1, c1, d, e, f, g, h, i)

	const (
		i1 = 1 << iota // 1<<0		1左移0位	00000001	1
		j1 = 3 << iota // 3<<1		3左移1位	00000110	6
		k              // 3<<2		3左移2位	00001100	12
		l              // 3<<3		3左移3位	00011000	24
	)

	// 定义数量级
	const (
		_  = iota
		KB = 1 << (10 * iota) // 2e10		1024	1左移十位
		MB = 1 << (10 * iota) // 2e20
		GB = 1 << (10 * iota) // 2e30
		TB = 1 << (10 * iota) // 2e40
		PB = 1 << (10 * iota) // 2e50
	)

	const (
		d1, d2 = iota + 1, iota + 2 // iota: 0		d1: 1		d2: 2
		d3, d4 = iota + 1, iota + 2 // iota: 1		d3: 1		d4: 2
	)
}
