package main
import "fmt"
func main(){
	sliceTest()
	twoDimensionArray()
}
func sliceTest()  {
	arr := []int{1,2,3,4,5}
	s := arr[:]
	for e := range s {
		fmt.Println(s[e])
	}
	s1 := make([]int, 3)
	for e := range s1 {
		fmt.Println(s1[e])
	}
}
func twoDimensionArray(){
	var a = [][]int{{0,0},{1,2},{2},{3,6},{4,8}}
	var i, j int
	for i = 0; i < len(a); i++ {
		for j = 0; j < len(a[i]); j++ {
			//Println:打印字符串和变量
			//Printf:打印格式化的字符串，不可以输出整型变量和整型
			fmt.Printf("a[%d][%d] = %d\n",i,j,a[i][j])


		}
	}
}
