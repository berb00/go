package base

import (
	"fmt"
)

func DefineUnidimensionalArray() {
	// declare an array
	var arr1 [5]int     // 声明空数组 [0 0 0 0 0]
	var arr2 [5]float32 // 声明空数组 [0 0 0 0 0]
	arr3 := [5]string{"a", "b"}
	arr4 := [...]int{1, 2, 3, 4, 5, 6} // 不指定具体长度， go自动判定
	arr5 := [5]int{0: 3, 1: 5, 4: 6}   // 初始化特定下标

	/* 为数组 n 初始化元素 */
	for i := 0; i < 5; i++ {
		arr1[i] = i + 103 /* 设置元素为 i + 100 */
	}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)

	/* 数组作为参数传递给函数 */
	var avg = getAverage(arr1, 5)
	fmt.Printf("平均值为: %f\n ", avg)

}

func DefinedMultidimensionedArray() {
	var arr1 = [3][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {3, 4, 5, 6, 7}}
	arr2 := [3][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {3, 4, 5, 6, 7}}
	arr3 := [...][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {0: 3, 1: 5, 4: 6}}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}

func getAverage(arr [5]int, size int) float32 { // 函数接收数组
	var sum int
	var avg float32
	for i := 0; i < size; i += 1 {
		sum += arr[i]
	}
	avg = float32(sum) / float32(size)

	return avg
}
