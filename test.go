package main
import "fmt"
// func main(){
// 	fmt.Println("hello word");
// }
// func main(){
// 	const LENGTH int = 10
// 	const WIDTH int = 5
// 	var area int
// 	const a,b,c = 1,false,"str"
// 	area = LENGTH * WIDTH
// 	fmt.Println("面积为 : %d",area)
// 	println()
// 	println(a,b,c)
// }

func max(num1,num2 int) int {
	var result int 
	if(num1> num2){
		result =num1
	}else{
		result = num2
	}
	return result
}
func swap(x,y string)(string ,string){
	return y,x
}
func main(){
	var a int = 100
	var b int = 200
	var ret int
	ret=max(a,b)
	fmt.Println("最大值是 : %d\n",ret)
	f,g := swap("abc","def");
	fmt.Println(f,g);
}