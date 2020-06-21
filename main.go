// 定义包名。
// 指明这个文件属于哪个包，package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
	"golang/base"
	_ "golang/math"
	// _ "fmt"		// _指匿名包
	// F "fmt"		// 给包起别名为
)

// 初始化函数，若存在会在main函数之前执行
func init() {
	// fmt.Println("init")
}

// 入口函数 自动调用
func main() { // { 不能另起一行

	// fmt.Println("外部包测试：", configor.Config{})

	// math
	// fmt.Println(math.Add(1, 1))
	// fmt.Println(math.Sub(1, 1))

	// var
	// base.VarDefault()
	// base.VarString()
	// base.VarInt()
	// base.VarFloat()
	// base.VarBool()
	// base.Printf()

	// slice
	// base.DefSlice()
	// base.Slice()
	// base.DefMap()

	// fmt
	// base.Print()

	// struct
	// base.InitStruct()
	base.TestStructFunc()

	//
	testPackage()
	testDecimal()
}

func testPackage() {
	const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

	value := gjson.Get(json, "name.last")
	fmt.Println(value.String())

}

func testDecimal() {

	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}

	quantity := decimal.NewFromInt(3)

	fee, _ := decimal.NewFromString(".035")
	taxRate, _ := decimal.NewFromString(".08875")

	subtotal := price.Mul(quantity)

	preTax := subtotal.Mul(fee.Add(decimal.NewFromFloat(1)))

	total := preTax.Mul(taxRate.Add(decimal.NewFromFloat(1)))

	fmt.Println("Subtotal:", subtotal)                      // Subtotal: 408.06
	fmt.Println("Pre-tax:", preTax)                         // Pre-tax: 422.3421
	fmt.Println("Taxes:", total.Sub(preTax))                // Taxes: 37.482861375
	fmt.Println("Total:", total)                            // Total: 459.824961375
	fmt.Println("Tax rate:", total.Sub(preTax).Div(preTax)) // Tax rate: 0.08875
}
