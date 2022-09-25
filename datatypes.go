package main

import "fmt"

/**
* file ini akan membahas tentang penggunaan data types
* setiap penjelasan akan dibahas per function
**/

func usingIntegers() {

	/**
	* int untuk bilangan bulat(bilangan positif dan bilangan negatif)
	* uint untuk bilangan cacah (bilangan positif)
	**/

	var firstUInt uint8 = 255
	var secondUInt uint16 = 65535
	var thirdUInt uint32 = 4294967295
	var fourthUInt uint64 = 18446744073709551615
	var firstInt int8 = -128
	var secondInt int16 = -32768
	var thirdInt int32 = -2147483648
	var fourthInt int64 = -9223372036854775808

	fmt.Printf("cek type variable first Uint %T \n", firstUInt)
	fmt.Printf("cek type variable second Uint %T \n", secondUInt)
	fmt.Printf("cek type variable third Uint %T \n", thirdUInt)
	fmt.Printf("cek type variable fourth Uint %T \n", fourthUInt)
	fmt.Printf("cek type variable first Int %T \n", firstInt)
	fmt.Printf("cek type variable second Int %T \n", secondInt)
	fmt.Printf("cek type variable third Int %T \n", thirdInt)
	fmt.Printf("cek type variable fourth Int %T \n", fourthInt)
}

func usingFloat() {

	/**
	* Secara umum float dibagi menjadi 2 jenis yaitufloat32 dan float64
	**/

	var float64 float64 = 3.63
	var float32 float32 = 3.14
	fmt.Printf("cek type variable float64 %T \n", float64)
	fmt.Printf("cek type variable float32 %T \n", float32)
}

func usingBool() {

	/**
	* Tipe data bool pada bahasa Go hanya terdiri dari 2 nilai yaitu false, dan true.
	* Tipe data ini biasanya digunakan dalam perkondisian atau perulangan.
	**/

	var condition bool = true
	fmt.Printf("is it permitted? %t \n", condition)
	condition = false
	fmt.Printf("okay, is it permitted? %t \n", condition)
}

func usingStringWithQoute() {
	var message string = "Hendrawan Ganteng"
	fmt.Println(message)
}

func usingStringWithBackticks() {
	var message = `Selamat Datang di "Kanbaw University". 
	Salam kenal :).
	Mari mencuci gelas dan piring di "Pak Aji".`

	fmt.Println(message)
}
