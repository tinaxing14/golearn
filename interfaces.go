package main

import (
	"fmt"
)

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

func main() {

	var aw Writer = new(anotherWriter)
	aw.Write([]byte("hellooooo"))
	var w Writer = new(consoleWriter)
	w.Write([]byte("helooooo"))

	var x Incrementer = &integer{value: 1}
	x.Increment()
	var y Incrementer = &float{value: 1, value2: 2}
	y.Increment()

}

// struct store data, interface store method definitions
// naming convention - if the interface has only one method, it will add +er to that method as the interface name

// interfaces can be used to group methods that has different receivers
type Writer interface {
	Write([]byte) (int, error)
}

type consoleWriter struct{}

func (cw consoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type anotherWriter struct{}

func (a anotherWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// any type that has methods associated with it can have interfaces

type Incrementer interface {
	Increment() int
}

type integer struct {
	value int
}

type float struct {
	value  float32
	value2 float64
}

func (int *integer) Increment() int {
	int.value++
	fmt.Println(int.value)
	return int.value
}

func (float *float) Increment() int {
	float.value++
	float.value2++
	fmt.Println(float.value, float.value2)
	return int(float.value)
}

// best practices of interfaces
// use small interfaces
// design functions and methods to receive interfaces whenever possible
// if you need access to the underlying data field, it's not possible. just take in the concrete types
