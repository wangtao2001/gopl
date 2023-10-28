package main

type Weekday int

func main() {
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
	)

	const (
		a int = 0
	)

	println(Tuesday)
}
