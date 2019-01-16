package main
import "fmt"
func printSlice(x []int)  {
	//len()获取长度；cap（）测量切片最长可以达到多少
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
func main(){
	numbers := []int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers)
	fmt.Println("number==",numbers)
	numbers1 := numbers[1:4]
	printSlice(numbers1)
	numbers2 := numbers[4:]
	printSlice(numbers2)
	numbers3 :=make([]int,0,5)
	printSlice(numbers3)
	numbers4 :=numbers[:2]
	printSlice(numbers4)

}