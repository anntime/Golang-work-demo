package main
import "fmt"
func printSlice(x []int)  {
	//len()获取长度；cap（）测量切片最长可以达到多少
	// 切片实际是获取数组的某一部分，len切片<=cap<=len数组
	//对于底层数组容量是k的切片，slice[i：j]：长度是j-i，容量是k-i
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