package base

import "fmt"

func Print() {
	fmt.Print("A")           // 打印不换行
	fmt.Print("A", "B", "C") // 打印无空格间隔

	fmt.Println("===============") // 打印无空格间隔
	fmt.Println("A")               // 打印自动换行
	fmt.Println("A", "B", "C")     // 打印有空格间隔

	fmt.Println("===============") // 打印无空格间隔
	a := 10                        // 打印格式化
	fmt.Printf("a=%v\n", a)        // %v	相应值的默认格式。在打印结构体时，“加号”标记（%+v）会添加字段名
	fmt.Printf("a=%#v\n", a)       // %#v	%#v	相应值的Go语法表示
	fmt.Printf("a=%% b=%T\n", a)   // %%	%%	字面上的百分号，并非值的占位符
	fmt.Printf("a=%T\n\n", a)      // %T	相应值的类型的Go语法表示

	fmt.Printf("二进制 	= %b\n", a) // %b	二进制表示
	fmt.Printf("十进制 	= %d\n", a) // %d	十进制表示
	fmt.Printf("八进制 	= %o\n", a) // %o	八进制表示
	fmt.Printf("十六进制 = %x\n", a) // %x	十六进制表示，字母形式为小写 a-f
	fmt.Printf("十六进制 = %X\n", a) // %X	十六进制表示，字母形式为大写 A-F
	fmt.Printf("字面值 = %q\n", a)  // %q	单引号围绕的字符字面值，由Go语法安全地转义
	fmt.Printf("a = %U\n", a)    // %U	Unicode格式：U+1234，等同于 "U+%04X"
	fmt.Printf("a = %c\n", a)    // %c	相应Unicode码点所表示的字符

}
