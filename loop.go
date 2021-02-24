package main

import "fmt"

func LoopFunc() {
	// i is available in global scope in this example
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
	fmt.Println(i)

	// i is scoped to the for loop
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}
	// breaking out the outer loop with a tag
Loop:
	for x := 1; x <= 3; x++ {
		for y := 1; y <= 3; y++ {
			fmt.Println(x * y)
			if x*y >= 3 {
				break Loop
			}
		}
	}

	// for range loop. looping through slices with unknown size
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}
	st := "hello go!"
	for k, v := range st {
		fmt.Println(k, string(v))
	}
}
