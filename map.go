package xu

import "fmt"

func main() {
	var m = map[string]int{
		"yes": 12,
		"no":  13,
	}
	fmt.Println(m)

	makeMap := make(map[string]int)
	makeMap = map[string]int{
		"yes": 12,
		"no":  13,
	}
	fmt.Println(makeMap)
	fmt.Println(makeMap["yes"])

	delete(makeMap, "yes")
	fmt.Println(makeMap)
	//note that after deleting a key, the value shows 0 when you try to print out the key
	fmt.Println(makeMap["yes"])

	// if not sure if the key is in the map or not, use ok syntax
	_, ok := makeMap["yes"]
	fmt.Println(ok)

	// checking the length of a map
	fmt.Println((len(makeMap)))

	// maping are passing reference.
}

// the return order of a map is not guranteed
