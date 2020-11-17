package main

import "fmt"

func arrays() {
	a := [3]int{1, 2, 3}
	fmt.Println(a)
	b := [...]int{1, 2, 3}
	fmt.Println(b)
}

func slices() {
	// create sklices
	// slice existing array
	a := []string{"1", "2", "3", "4", "5"}
	fmt.Println(a)
	b := a[:2]
	fmt.Println(b)
	// via make function
	c := make([]string, 10)      //create a slice with capacity and length ==10
	d := make([]string, 10, 100) // slice with length == 10 and capacity ==100
	fmt.Println(c, d)

	//append function to add elements to a slice. - may cause expensive copy operation if underlying array is too small
	e := append(a, "7")
	fmt.Println(e)

	//copies of slice refers to the same underlying array
	array := [...]int{1, 2, 3, 4, 5, 6}
	slice1 := array[:3]
	slice2 := slice1
	slice2[1] = 0
	fmt.Println(array, slice1, slice2)
}
