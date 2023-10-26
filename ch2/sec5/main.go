package main

type A string
type B string

var a A = "ssss"

// var b B = "ffff"

type p *int

var pp p

func main() {
	println(a)
	x := 10
	pp = &x
	println(*pp)
	// println(a == b) // 不同类型不能直接比较
}
