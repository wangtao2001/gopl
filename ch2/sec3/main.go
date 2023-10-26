package main

import "fmt"

func main() {
	x := 1
	p := &x
	fmt.Println(p)
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)

	v := &x
	fmt.Println(v == p)

	fmt.Println(f() == f())

	a := new(int) // 返回的是指针类型
	println(a)
	println(*a)
	println(newInt())
}

func f() *int { // 返回的是一个指向int类型变量的地址
	var p int
	return &p
}

func newInt() int {
	return *new(int)
}
