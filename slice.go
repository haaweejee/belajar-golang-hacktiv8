package main

import (
	"fmt"
	"strings"
)

func implementSlice() {
	fruits := make([]string, 3)

	_ = fruits
	fmt.Printf("%#v", fruits)
}

func accessIndexSlice() {
	fruits := make([]string, 3)

	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "grape"

	fmt.Printf("%#v", fruits)
}

func appendSlice() {
	fruits := make([]string, 0)

	fruits = append(fruits, "apple", "banana", "grape")

	fmt.Printf("%#v", fruits)
}

func appendSliceWithEllipsis() {
	fruits1 := []string{"apple", "banana", "grape"}
	fruits2 := []string{"durian", "pinaple", "starfruit"}

	fruits1 = append(fruits1, fruits2...)
	fmt.Printf("%#v", fruits1)
}

func sliceWithCopy() {
	fruits1 := []string{"apple", "banana", "grape"}
	fruits2 := []string{"durian", "pinaple", "starfruit"}

	copyElement := copy(fruits1, fruits2)

	fmt.Println("Fruits 1 =>", fruits1)
	fmt.Println("Fruits 2 =>", fruits2)
	fmt.Println("Copied Elements =>", copyElement)
}

func slicingSlice() {
	var fruits1 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	fruits2 := fruits1[1:4]
	fmt.Printf("%#v\n", fruits2)

	fruits3 := fruits1[0:]
	fmt.Printf("%#v\n", fruits3)

	fruits4 := fruits1[:3]
	fmt.Printf("%#v\n", fruits4)

	fruits5 := fruits1[:]
	fmt.Printf("%#v\n", fruits5)
}

func combSliceWithAppend() {
	var fruits1 = []string{"apple", "banana", "mango", "durian", "pineapple"}
	fruits1 = append(fruits1[:3], "rambutan")
	fmt.Printf("%#v\n", fruits1)
}

func sliceBackingArray() {
	var fruits1 = []string{"apple", "mango", "durian", "banana", "starfruit"}

	var fruits2 = fruits1[2:4]

	fruits2[0] = "rambutan"

	fmt.Println("fruits1 => ", fruits1)
	fmt.Println("fruits2 => ", fruits2)
}

func sliceCheckCapAndLen() {
	var fruits1 = []string{"apple", "mango", "durian", "banana", "starfruit"}
	fmt.Println("Fruits1 cap:", cap(fruits1))
	fmt.Println("Fruits1 len:", len(fruits1))

	fmt.Println(strings.Repeat("#", 20))

	var fruits2 = fruits1[0:3]
	fmt.Println("Fruits2 cap:", cap(fruits2))
	fmt.Println("Fruits2 len:", len(fruits2))

	fmt.Println(strings.Repeat("#", 20))

	var fruits3 = fruits1[1:]
	fmt.Println("Fruits3 cap:", cap(fruits3))
	fmt.Println("Fruits3 len:", len(fruits3))
}

func sliceNewBackingArray() {
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"
	fmt.Println("cars:", cars)
	fmt.Println("newCars:", newCars)

}
