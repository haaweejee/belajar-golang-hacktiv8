package main

import "fmt"

func ifElse() {
	var birthYears = 2000

	if birthYears == 2000 {
		fmt.Println("Umur ku 22 Tahun")
	} else {
		fmt.Println("Umur ku bukan 22 tahun")
	}
}

func temporaryVariable() {
	var currentYear = 2021
	if age := currentYear - 2000; age < 17 {
		fmt.Println("kamu tidak boleh membuat ktp")
	} else {
		fmt.Println("kamu boleh membuat ktp")
	}
}

func switchCondition() {
	score := 6

	switch score {
	case 9:
		fmt.Println("Selamat Kamu lulus nilai kamu A")
	case 8:
		fmt.Println("Selamat Kamu lulus nilai kamu B")
	case 7:
		fmt.Println("Selamat Kamu lulus nilai kamu C")
	default:
		fmt.Println("Maaf kamu tidak lulus.")
	}
}

func switchWithRelationalCondition() {
	var score = 6

	switch {
	case score == 9:
		fmt.Println("Selamat Kamu lulus nilai kamu A")
	case (score < 8) && (score >= 7):
		fmt.Println("Selamat Kamu lulus nilai kamu B")
	case (score < 7) && (score >= 6):
		fmt.Println("Selamat Kamu lulus nilai kamu C")
	default:
		fmt.Println("Maaf kamu tidak lulus.")
	}
}

func switchFallthroughKeyword() {
	score := 7
	switch {
	case score == 9:
		fmt.Println("Selamat Kamu lulus nilai kamu A")
	case (score < 8) && (score >= 7):
		fmt.Println("Selamat Kamu lulus nilai kamu B")
		fallthrough
	case (score < 7) && (score >= 6):
		fmt.Println("Selamat Kamu lulus nilai kamu C")
	default:
		fmt.Println("Maaf kamu tidak lulus.")
	}
}

func nestedCondition() {
	score := 10

	if score > 7 {
		switch score {
		case 10:
			fmt.Println("Perfect")
		default:
			fmt.Println("Nice")
		}
	} else {
		if score == 5 {
			fmt.Println("Not Bad")
		} else if score == 3 {
			fmt.Println("Keep Trying")
		} else {
			fmt.Println("You Can Do It")
		}
	}
}
