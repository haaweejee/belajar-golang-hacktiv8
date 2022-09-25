package main

import "fmt"

/**
* file ini akan membahas tentang penggunaan variable
* setiap penjelasan akan dibahas per function
**/
func variableWithDataType() {

	var fullName string = "Hendrawan Wibowo Jatmiko"
	var age int = 22

	fmt.Println("Nama saya adalah", fullName)
	fmt.Println("Umur saya adalah", age)

	fullName = "Jatmiko Wibowo Hendrawan"
	fmt.Println("oh ternyata bukan, terus nama nya siapa?", fullName)
}

func variableWithoutDataTypes() {
	var address = "Jl. Periuk Jaya Permai 2 No.6 RT/RW 07/05, Periuk Jaya, Periuk"
	var city = "Kota Tangerang"
	var postalCode = 15131

	fmt.Printf("sebelum dibaca coba cek type data variable nya %T, %T, %T", address, city, postalCode)
}

func shortDeclarationVariables() {
	job := "Android Developer"
	company := "PT.Kompas Media Nusantara"

	fmt.Printf("Saya bekerja sebagai %s di perusahaan %s", job, company)
}

func multipleVariableDeclarations() {
	var studentOne, studentTwo, studentThree = "Hendra", "Budi", "Syarif"
	var numberOne, numberTwo, numberThree int

	numberOne, numberTwo, numberThree = 1, 2, 3

	fmt.Printf("Daftar Mahasiswa: %s, %s, %s \n", studentOne, studentTwo, studentThree)
	fmt.Printf("Nomer Absensi: %d, %d, %d", numberOne, numberTwo, numberThree)
}

func underscoreVariableDeclarations() {
	var firstVariable string

	var name, age, address = "Hendra", 22, "Tangerang"

	_, _, _, _ = firstVariable, name, age, address
}

func createConstantVariable() {
	/**
	* constant adalah variable yang nilai nya tidak bisa berubah
	* jika nilainya dirubah akan menimbulkan error
	**/
	const URL = "google.com"
	// URL = "facebook.com"

	println(URL)
}
