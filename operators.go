package main

import "fmt"

/**
* file ini akan membahas tentang penggunaan Operators
* setiap penjelasan akan dibahas per function
**/

func arithmathicsOperators() {
	var variableA = 20
	var variableB = 30

	add := variableA + variableB
	substract := variableA - variableB
	multiplies := variableA * variableB
	divides := variableA / variableB
	modulus := variableB % variableA

	fmt.Printf("Hasil dari VariableA + VariableB = %d \n", add)
	fmt.Printf("Hasil dari VariableA - VariableB = %d \n", substract)
	fmt.Printf("Hasil dari VariableA * VariableB = %d \n", multiplies)
	fmt.Printf("Hasil dari VariableA / VariableB = %d \n", divides)
	fmt.Printf("Hasil dari VariableA Modulus VariableB = %d \n", modulus)

	variableA++
	fmt.Printf("Hasil Increment dari VariableA = %d \n", variableA)
	variableB--
	fmt.Printf("Hasil Decrement dari VariableB = %d \n", variableB)

}

func relationalOperators() {
	var firstCondition bool = 2 < 3
	var secondCondition bool = "Joey" == "joey"
	var thirdCondition bool = 10 != 11
	var fourthCondition bool = 11 <= 11

	fmt.Println("first condition:", firstCondition)
	fmt.Println("second condition:", secondCondition)
	fmt.Println("third condition:", thirdCondition)
	fmt.Println("fourth condition:", fourthCondition)
}

func logicalOperators() {
	var right = true
	var wrong = false

	var wrongAndRight = wrong && right
	var wrongOrRight = wrong || right
	var wrongReverse = !wrong
	fmt.Printf("wrong && right \t(%t) \n", wrongAndRight)
	fmt.Printf("wrong || right \t(%t) \n", wrongOrRight)
	fmt.Printf("!wrong \t(%t) \n", wrongReverse)
}
