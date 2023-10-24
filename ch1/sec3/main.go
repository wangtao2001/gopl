package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileIn2()
}

func stdIn() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { // 每次读入一行，并移除行尾的换行符，不再有输入时返回false
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

var file = "e:/goproject/src/gopl/ch1/sec3/data.txt"

func fileIn() {
	f, err := os.Open(file)
	if err == nil { // 错误不存在的情况下
		input := bufio.NewScanner(f) // 现在的scanner是从文件读取
		for input.Scan() {
			fmt.Println(input.Text())
		}
	}
	_ = f.Close()
}

func fileIn2() {
	data, err := ioutil.ReadFile(file) // 整个文件读入内存，返回的是字节切片
	if err == nil {
		fmt.Print(string(data))
	}
}
