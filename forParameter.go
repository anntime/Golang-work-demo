package main
import "fmt"
func main(){
	var a int =0;
	fmt.Println("for start")
	for a := 0; a < 10; a++ {
	//for 循环的initialize(a:=0)中，此时的啊与外层的a不是同一个变量，这里的a是for循环中的局部变量
		fmt.Println(a);
	}

	fmt.Println("for end")
	fmt.Println(a)
}