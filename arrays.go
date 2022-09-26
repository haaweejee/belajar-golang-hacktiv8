package main

import (
	"fmt"
	"strings"
)

func simpleArrays() {
	var numbers [5]int
	numbers = [5]int{5, 4, 3, 2, 1}

	var strings = [2]string{"Hendra", "irwan"}

	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", strings)
}

func modifyArrays() {
	var fruits = [3]string{"Banana", "Grape", "Strawberry"}
	fruits[0] = "Apple"
	fruits[1] = "Watermelon"
	fruits[2] = "Mango"

	fmt.Printf("%#v\n", fruits)
}

func loopArrays() {
	var fruits = [3]string{"Banana", "Grape", "Strawberry"}

	//cara pertama
	for i, v := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}

	fmt.Println(strings.Repeat("#", 25))

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, fruits[i])
	}
}

func multidimensionalArrays() {

	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}
