package base
import "fmt"
/*
切片 动态数组

*/


func DefSlice()  {
	var slice1 []int					// 声明一个未指定大小的数组来定义切片
	if(slice1 == nil) {					// 未初始化之前默认为 nil，长度为 0
		fmt.Printf("空切片: len=%d cap=%d slice=%v\n",len(slice1),cap(slice1),slice1)
	}

	
	var slice2 []int = make([]int, 10)	// 使用make函数声明切片	默认元素值为0
	slice3 := make([]int, 10)			// 简写
	var slice4 = make([]int, 3, 5)		// 

	fmt.Println(slice1, slice2, slice3)
	// len():			返回切片长度
	// cap(slice4)		返回切片容量(切片最长长度)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(slice4),cap(slice4),slice4)


	s :=[] int {0,1,2,3,4,5,6,7,8} 				// 初始化切片
	fmt.Println(s)
}

func Slice () {
	   /* 创建切片 */
	   numbers := []int{0,1,2,3,4,5,6,7,8}  
	   fmt.Printf("len=%d cap=%d slice=%v\n",len(numbers),cap(numbers),numbers)
	
	   /* 打印原始切片 */
	   fmt.Println("numbers ==", numbers)
	
	   /* 打印子切片从索引1(包含) 到索引4(不包含)*/
	   fmt.Println("numbers[1:4] ==", numbers[1:4])
	
	   /* 默认下限为 0*/
	   fmt.Println("numbers[:3] ==", numbers[:3])
	
	   /* 默认上限为 len(s)*/
	   fmt.Println("numbers[4:] ==", numbers[4:])
	
	   number1 := make([]int,0,5)
	   fmt.Printf("len=%d cap=%d slice=%v\n",len(number1),cap(number1),number1)

	
	   /* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	   number2 := numbers[:2]
	   fmt.Printf("len=%d cap=%d slice=%v\n",len(number2),cap(number2),number2)

	
	   /* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	   number3 := numbers[2:5]
	   fmt.Printf("len=%d cap=%d slice=%v\n",len(number3),cap(number3),number3)

}