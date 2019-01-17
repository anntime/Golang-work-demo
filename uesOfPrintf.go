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
	//

}
