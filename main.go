package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

// RahulUser is a constant
const RahulUser = "Rahul's user"

// RahulAge is a constant
const RahulAge = "Rahul's Age"

func main() {

	s := person{name: "Sean", age: 50}

	setIntoMap(s, RahulUser)
	user := getFromMap(RahulUser)
	if u, ok := user.(person); ok {
		fmt.Println(u)
	}

	var age int64 = 26
	setIntoMap(age, RahulAge)
	_age := getFromMap(RahulAge)
	if u, ok := _age.(int64); ok {
		fmt.Println(u)
	}

}

var hashMap = make(map[string]interface{})

func setIntoMap(t interface{}, name string) {
	if i, ok := hashMap[name]; ok {
		fmt.Printf("Already exists %v\n", i)
		return
	}
	hashMap[name] = t
}

func getFromMap(name string) (value interface{}) {
	if value, ok := hashMap[name]; ok {
		return value
	}
	return nil
}
