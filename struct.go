// collecton types that can have different types of data
package xu

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func structs() {
	aDocter := Doctor{
		number:    3,
		actorName: "dave",
		companions: []string{
			"person1", "person2",
		},
	}
	// note that field names(number, actorName, companions) are optional but not encouraged.

	fmt.Println(aDocter.companions[0])

	// fields in struct with lowercase will not be exposed to outside world even if struct name is cap

	type DoctorExpose struct {
		Number     int
		ActorName  string
		Companions []string
	}

	// struct is not like map, when we pass struct around we are passing the copy of it.
	aDoctor := struct{ name string }{name: "Tina"}
	aDoctor.name = "Peter"
	fmt.Println(aDoctor)

	anotherDoctor := aDoctor
	anotherDoctor.name = "New Name"
	fmt.Println(anotherDoctor)
	fmt.Println(aDoctor)

	changeDoctor := &aDoctor
	changeDoctor.name = "change original name"
	fmt.Println(aDoctor)
	fmt.Println(changeDoctor)
	fmt.Println(*changeDoctor)

	// tags for struct

	type Animal struct {
		Name string `required max: "100"`
	}

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println((field.Tag))
}
