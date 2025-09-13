package main

func main() {

	var a int = 10
	b := 2

	if a%b == 0 {
		println("10 is Genap")
	} else {
		println("10 is not Genap")
	}

	if a < b || a == b {
		println("True")
	} else {
		println("False")
	}

	if 8%2 == 0 || 7%2 == 0 {
		println("either 8 or 7 are even")
	}

	if num := 9; num < 0 {
		println("is Negatif")
	} else if num < 10 {
		println("has 1 digit")
	} else {
		println("has multiple digits")
	}

}