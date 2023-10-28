package main

import "fmt"

var x complex128 = complex(1, 2) // 1+2i
var y complex128 = complex(3, 4) // 3+4i

func main() {
	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))
	fmt.Println(1 + 2i)
}
