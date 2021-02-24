package xu

import "fmt"

func Pointers() error {
	a := 1
	b := a
	fmt.Println(a, b)

	//  c is a pointer to the integer to the address of a
	// the * here before a type means this is a pointer to the integer pointer type
	var c *int = &a
	fmt.Println(c)

	// the * here is a deference,deference * should be before a pointer value, giving the value of an address(pointer)
	fmt.Println(*c)

	a = 2
	fmt.Println(a, b, *c)

	*c = 3
	fmt.Println(a, b, *c)

	// every variable we initialize in Go has an initialization value

	// variable ms is a pointer to type myStruct. since myStruct isn't assigned a value, mystruct has an initialization value of 0. a pointer to that value would be nil
	var ms *myStruct
	fmt.Println(ms)
	// even though ms is a pointer, if we want to give a value to the pointing struct, we don't need to dereference it first
	ms = new(myStruct)
	fmt.Println(ms)
	ms.foo = 5 // is the same as (*ms.foo) = 5
	fmt.Println(*ms)

	//slices are passing by reference, however array are passing by value
	x := [2]int{1, 2}
	y := x
	x[0] = 0
	fmt.Println(x, y)

	m := []int{1, 2}
	n := m
	m[0] = 0
	fmt.Println(m, n)

	// map is passing by reference as well

	o := map[string]int{"1": 1, "2": 2}
	p := o
	o["1"] = 0
	fmt.Println(o, p)

	return nil
}

type myStruct struct {
	foo int
}
