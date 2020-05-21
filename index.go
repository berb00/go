// 定义包名。指明这个文件属于哪个包，package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
package main

// 引入fmt包(实现了格式化 IO（输入/输出）的函数)
import (
	"fmt"
	"./math"	//math 目录
	)	

func main() {
   fmt.Println("Hello, World!")
   fmt.Println(math.Add(1,1))
   fmt.Println(math.Sub(1,1))
}

func baseType () {
	// 指定变量类型但不初始化
	var i int			// 数值类型默认0
    var f float64
    var b bool			// 布尔类型默认false
	var s string		// 字符串默认''
	
	var b1 = true		// 不指定数据类型(自动判断)

	intVal,intVal1 := 1,2 // := 声明新的变量
	f := "Runoob"
	var intVal int 
	intVal :=1 // 这时候会产生编译错误
    fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
