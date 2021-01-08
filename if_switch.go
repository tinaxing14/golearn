package main

import "fmt"

func ifandswitch() {
	// simple if statement
	if true {
		fmt.Println("true")
	}

	// if statement with an initializer
	statePopulations := map[string]int{
		"PA": 12802503,
	}
	// variables in initializer are local to if block
	if pop, ok := statePopulations["PA"]; ok {
		fmt.Println("PA is in the map", pop)
	}

	// tag syntax
	switch i := 2 + 8; i {
	case 10:
		fmt.Println("its 10")
	case 20:
		fmt.Println("its 20")
	default:
		fmt.Println("nothing")
	}

	//without tag syntax
	x := 10
	switch {
	case x == 10:
		fmt.Println("its 10")
	case x == 20:
		fmt.Println("its 20")
	default:
		fmt.Println("nothing")
	}

	// break keyword is implicit. Use fallthrough
	a := 10
	switch {
	case a == 10:
		fmt.Println("its 10")
		fallthrough
	case a == 20:
		fmt.Println("its 20")
	default:
		fmt.Println("nothing")
	}

	// type switch
	var ii interface{} = 1
	switch ii.(type) {
	case int:
		fmt.Println(" ii is an int")
		if ii == 1 {
			break
		}
		fmt.Println("I don't want this to print")
	case float32:
		fmt.Println(" ii is a float32")
	}
}
