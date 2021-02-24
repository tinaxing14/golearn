package xu

import "fmt"

func variables() {
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
