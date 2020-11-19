package myJson

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	IDCard string
	Age int
}

func getJson() {
	person := &Person{"lai", "44011213213", 18}
	bytes, err := json.Marshal(person)
	if err != nil {
		return
	}

	newPerson := new(Person)
	err2 := json.Unmarshal(bytes, newPerson)
	if err2 != nil {
		return
	}

	fmt.Printf("person is %v", newPerson)
}
