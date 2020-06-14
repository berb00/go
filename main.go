// 定义包名。
// 指明这个文件属于哪个包，package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
package main

// 引入包
import (
	"./base" // base 目录
	"./math" // math 目录
	"fmt"    // 格式化 IO（输入/输出）
)

// 初始化函数，若存在会在main函数之前执行
func init() {
	// fmt.Println("init")
}

// 入口函数 自动调用
func main() { // { 不能另起一行
	// math
	fmt.Println(math.Add(1, 1))
	// fmt.Println(math.Sub(1,1))

	// var
	// base.VarDefault()
	// base.VarString()
	// base.VarInt()
	// base.VarFloat()
	// base.VarBool()
	// base.Printf()

	// slice
	base.DefSlice()

}
