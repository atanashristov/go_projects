package main

import "fmt"

func main() {
	var a = 4
	var b = 5.6

	//za = b // compiler error "cannot use b ..."
	a = int(b) // but we can convert
	fmt.Println(a, b)

	//** ZERO VALUES *//

	var (
		value int
		price float64
		name  string
		done  bool
	)

	fmt.Println(value, price, name, done)

	x, y := 2., 5
	x = y
	fmt.Println(x, y)
}
