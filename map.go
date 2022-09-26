package main

import "fmt"

func mapPersonOne() {
	var person map[string]string // Deklarasi

	person = map[string]string{}

	person["name"] = "Airrell"

	person["age"] = "23"

	person["address"] = "Jl. Sudirman"

	fmt.Println("name:", person["name"])
	fmt.Println("age:", person["age"])
	fmt.Println("address:", person["address"])
}

func mapPersonTwo() {
	var person = map[string]string{
		"name":    "Hendra",
		"age":     "22",
		"address": "Tangerang",
	}

	fmt.Println("name: ", person["name"])
	fmt.Println("age: ", person["age"])
	fmt.Println("address: ", person["address"])
}

func mapPersonThree() {
	var person = map[string]string{
		"name":    "Hendra",
		"age":     "22",
		"address": "Tangerang",
		"job":     "Android Developer",
	}

	for key, value := range person {
		fmt.Println(key, ":", value)
	}
}

func deleteMapValue() {
	var person = map[string]string{
		"name":    "Hendra",
		"age":     "22",
		"address": "Tangerang",
		"job":     "Android Developer",
	}

	fmt.Println("Before Deleting:", person)

	delete(person, "age")

	fmt.Println("After Deleting:", person)
}

func detectingMapValue() {
	var person = map[string]string{
		"name":    "Hendra",
		"age":     "22",
		"address": "Tangerang",
		"job":     "Android Developer",
	}

	value, exist := person["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value Don't exist")
	}

	delete(person, "age")

	value, exist = person["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value Don't exist")
	}
}

func combiningMapWithSlice() {
	var people = []map[string]string{
		{"name": "Airell", "age": "23"},
		{"name": "Nanda", "age": "23"},
		{"name": "Nando", "age": "23"},
	}

	for i, person := range people {
		fmt.Printf("Index: %d, name: %s, age: %s\n", i, person["name"], person["age"])
	}
}
