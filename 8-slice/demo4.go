package main

import "fmt"

func main() {
	s := []int{1,2,3}
	s1 := s[:2]
	fmt.Printf("s1切片：长度为：%d, cap为：%d，内容为：%v\n", len(s1), cap(s1),s1)
	s[0] = 22
	fmt.Printf("s1切片：长度为：%d, cap为：%d，内容为：%v\n", len(s1), cap(s1),s1)

	s2 := make([]int,3)
	copy(s2,s1)
	fmt.Printf("s2切片：长度为：%d, cap为：%d，内容为：%v\n", len(s2), cap(s2),s2)

}
