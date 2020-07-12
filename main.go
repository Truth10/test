package main

import (
	"fmt"
)

const UserNamePrefix = "user"

func main() {
	// 第一题
	fmt.Println("---------第一题-----------")
	a := [4]int{10, 40, 5, 280}
	b := [5]int{234, 5, 2, 148, 23}
	v := 42

	var c, d []int
	for i := 0; i < len(a); i++ {
		c = append(c, a[i])
	}

	for i := 0; i < len(b); i++ {
		d = append(d, b[i])
	}

	ok := test(c, d, v)
	if ok {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	// 第二题
	fmt.Println("---------第二题-----------")
	for i := 0; i < 10; i++ {
		fmt.Println("i=", i, " username=", CreateUserName(int32(i)))
	}
}

// 第一题
func test(a []int, b []int, v int) bool {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i]+b[j] == v {
				return true
			}
		}
	}
	return false
}

// 第二题：创建用户名
func CreateUserName(a int32) string {
	randomNum := pseudo_encrypt(a)
	return fmt.Sprintf("user%v", randomNum)
}

// pseudo_encrypt 加密算法
func pseudo_encrypt(a int32) int {
	var l1, l2, r1, r2 int32
	l1 = (a >> 16) & 0xffff
	r1 = a & 0xffff

	for i := 0; i <= 3; i++ {
		l2 = r1
		r2 = l1 ^ int32((((1366*r1+150889)%714025)/714025.0)*32767)
		l1 = l2
		r1 = r2
	}
	return int((r1 << 16) + l1)
}

// 第三题：改进后的用户名生成算法
