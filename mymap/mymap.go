package mymap

import (
	"fmt"
	"sort"
)

func week() {
	myMap := make(map[int]string)
	myMap[0] = "Monday"
	myMap[1] = "Tuesday"
	myMap[2] = "Wednesday"
	myMap[3] = "Thursday"
	myMap[4] = "Friday"
	myMap[5] = "Saturday"
	myMap[6] = "Sunday"

	for _, value := range myMap {
		fmt.Printf("value:%s \n", value)
	}
}

var babyMap = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
	"delta": 87, "echo": 56, "foxtrot": 12,
	"golf": 34, "hotel": 16, "indio": 87,
	"juliet": 65, "kili": 43, "lima": 98}

func sortedMap() {
	fmt.Println("unsorted")
	for k, v := range babyMap{
		fmt.Printf("key: %s, value: %d", k, v)
	}

	keys := make([]string, len(babyMap))

	i:= 0
	for k, _ := range babyMap {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	fmt.Println("sorted")

	for _, k := range keys {
		fmt.Printf("key: %s, value:%d", k, babyMap[k])
	}
}
