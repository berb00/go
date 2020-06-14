package base		// 标明该文件属于base包
import (
	"fmt"
)

/*
变量声明

布尔型
	bool
数字类型
	int
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

bool

只有 true 和 false，默认为 false。


*/


var aaa string = "s"
var bbb string = "dsa"


func var_string ()  {
	var a string = "Runoob"
    fmt.Println(a)
}

func var_int ()  {
	var a string = "Runoob"
    fmt.Println(a)
}
func var_float ()  {
	var a string = "Runoob"
    fmt.Println(a)
}
func var_bool ()  {
	var a string = "Runoob"
    fmt.Println(a)
}

func baseType () {
	// 指定变量类型但不初始化
	var i int			// 数值类型默认0
    var f float64
    var b bool			// 布尔类型默认false
	var s string		// 字符串默认''
	
	var b1 = true		// 不指定数据类型(自动判断)

	intVal,intVal1 := 1,2 // := 声明新的变量
	// f := "Runoob"
	// var intVal int
	// intVal :=1 // 这时候会产生编译错误
    fmt.Printf("%v %v %v %q\n", i, f, b, s)
}