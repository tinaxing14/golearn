package main

import "fmt"

func main() {
	//variables and constants
	const pi = 3.1415926
	var (
		a int = 1
		b int = 2
	)
	fmt.Println(a, b, pi)

	var c string = "yes"
	fmt.Println(len(c))

	x, y := 14, 15
	fmt.Println(x, ",", y)

	//pointers
	z := 10
	changeValue(&z)
	fmt.Println(z)

}

func changeValue(x *int) {
	*x = 7
}

// map[string]interface{} - decoding unknown JSON data to map[string]interface{}:

// Known Person object
type Person struct {
	Name    string
	Age     int
	Hobbies []string
}

var person1 Person

var person2 map[string]interface{}

// Do a type switch do check what type of value is in the unknown object

func check(p map[string]interface{}) {
	for k, v := range p {
		switch c := v.(type) {
		case string:
			fmt.Printf("item %q is a string containing %q\n, k, c")
		case float64:
			fmt.Printf("item %q is a number")
		default:
			fmt.Printf("not sure the type, might be %T\n", k, c)
		}
	}
}
