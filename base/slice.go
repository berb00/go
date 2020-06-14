package base
import "fmt"
/*
切片 动态数组

*/


func DefSlice()  {
	var slice1 []int					// 声明一个未指定大小的数组来定义切片	
	var slice2 []int = make([]int, 10)	// 使用make函数声明切片	默认元素值为0
	slice3 := make([]int, 10)			// 简写
	fmt.Println(slice1, slice2, slice3)

	s :=[] int {1,2,3 } 				// 初始化切片
	fmt.Println(s)
}