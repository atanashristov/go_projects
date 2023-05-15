package main

import "fmt"

func main() {

	//** DECLARING VARIABLES *//
	var name string = "Tony Hristov"
	name = "Alex Hristov"
	fmt.Printf("Name: %s \n", name)

	var age = 37 // Go can infer the type from the value
	fmt.Printf("Age: %d \n", age)

	// Short declaration - use when we know the value
	name1 := "Someone"
	_ = name1

	s := "Learning Golang!"
	fmt.Println(s)

	//** TYPE INFERENCE **//

	// Go deduces automatically the type of the variable by looking at the initial value (bool, int, string etc)

	var k int = 6 // not necessary to say the type (int). It is inferred from the literal on the right side of =
	var i = 5     // type int
	var j = 5.6   // type float64

	_, _, _ = k, i, j

	// Short declaration of multiple variables
	car, cost := "Audi", 50000
	fmt.Println(car, cost)

	car, cost = "BMW", 65000
	fmt.Println(car, cost)

	// Short re-declaration - at least one variable has to be new
	// car, cost := "Subaru", 48000 // at least one variable has to be new
	car, year := "Subaru", 2022
	fmt.Println(car, year)

	var opened = false
	opened, file := true, "a.txt"

	_, _ = opened, file

	// Multiple variable declaration
	// Preferred for better readability
	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	// Declare multiple variables of same type
	var a, b, c int

	fmt.Println(a, b, c)

	var ii, jj int
	ii, jj = 5, 8 // Still compiler error, "i declared and not used"

	_, _ = ii, jj

	// Swap variables
	fmt.Println(ii, jj)
	jj, ii = ii, jj
	fmt.Println(ii, jj)

	// Initialize with expression
	sum := 5 + 2.3
	fmt.Println(sum)
}
