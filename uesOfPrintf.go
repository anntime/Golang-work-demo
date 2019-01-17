package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	// 打印point结构体的一个实例
	p := point{1, 2}
	fmt.Printf("%v\n", p)
	// 如果值是一个结构体，%+v的格式化输出内容将包括结构体的字段名
	fmt.Printf("%+v\n", p)
	// 值的运行源代码片段
	fmt.Printf("%#v\n", p)
	//需要打印值的类型
	fmt.Printf("%T\n", p)
	// 格式化布尔值
	fmt.Printf("%t\n", true)
	// 格式化整型数，标准的十进制格式化
	fmt.Printf("%d\n", 123)
	// 二进制
	fmt.Printf("%n\n", 14)
	//输出给定证书的对应字符
	fmt.Printf("%c\n", 33)
	//十六进制编码
	fmt.Printf("%x\n", 456)
	// 对于浮点型也有很多的格式化选项，下面进行最基本的十进制格式化
	fmt.Printf("%f\n", 78.9)
	// 科学计数法表示形式
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)
	// 基本的字符串输出
	fmt.Printf("%s\n", "\"string\"")
	// 像Go源代码中那样带有双引号的输出，
	fmt.Printf("%q\n", "\"string\"")
	// 输出使用base-16编码的字符串
	fmt.Printf("%x\n", " i love you")
	// 输出指针的值
	fmt.Printf("%p\n", &p)
	// 当输出数字的时候，有时需要控制输出结果的宽度和精度，可以使用在%后面使用数字来控制出宽度。默认结果使用右对齐并且通过空格来填充空白部分
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	// 也可以指定浮点型的输出宽度，同时也可以通过（宽度.精度）的语法来指定输出的精度
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	// 要由对齐
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	//控制字符串输出时的宽度
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	// 左对齐
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

}
