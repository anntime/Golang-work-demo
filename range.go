package main

import "fmt"

func main() {
	//用range去求一个slice的和
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	// 在数组上使用range将传入index和值
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//range也可以用在map的键值对上
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)

	}
	// range用来枚举Unicode字符串
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
