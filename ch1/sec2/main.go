package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string
	for _, arg := range os.Args { // range 同时返回索引和索引值
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(strings.Join(os.Args[1:], " ")) // 字符串拼接
}
