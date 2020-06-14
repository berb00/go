package base		// 标明该文件属于base包
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




*/

func VarDefault  () {	// 模块中要导出的函数，必须首字母大写
	var strdefault string			// string 默认值
	fmt.Println("string default:	", strdefault)

	var intdefault int			// string 默认值
	fmt.Println("int default:		", intdefault)

	var float32default float32		// float32 默认值
	fmt.Println("float32 default:	", float32default)

	var booldefault bool			// bool 默认值
	fmt.Println("bool default:		", booldefault)
	
}
func VarString  () {	// 模块中要导出的函数，必须首字母大写
	var str1 string = "sss"	// 单个声明
	var (					// 批量声明
		str2 string = "str2"
		str3 string = "str3"
	)

	fmt.Println(str1, str2, str3)

}

func VarInt ()  {

}
func VarFloat ()  {
	// var a float32 = 3.22
    // fmt.Println(a)
}
func VarBool ()  {
	// var a bool = false
    // fmt.Println(a)
}

