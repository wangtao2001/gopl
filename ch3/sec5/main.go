package main

func main() {
	s := `
	多行原生字符串字面量
`
	println(s)
	a := "\xe4\xb8\x96\xe7\x95\x8c"
	b := "\u4e16\u754c"
	c := "\U00004e16\U0000754c"
	println(a, b, c)

	d := []rune(a)
	println(d)
}
