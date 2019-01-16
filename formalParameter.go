package main
import "fmt"
// 声明全局变量
var a int = 20
func main(){
	// main函数中声明局部变量
	var a int = 10
	var b int = 20
	var c int =0
	fmt.Println("main()函数中a = %d\n",a)
	c=sum(a,b)
	fmt.Println("main()函数中a = %d\n",a)
	fmt.Println("main()函数中c = %d\n",c)

}
// 函数定义，两数相加
func sum(a,b int) int {
	a= a+1
	fmt.Println("sum()函数中a = %d\n",a)
	fmt.Println("sum()函数中b = %d\n",b)
	return a+b
}