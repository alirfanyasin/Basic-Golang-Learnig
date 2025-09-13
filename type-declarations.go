package main

func main() {

	type NoKTP string // type data alias
	var ktp NoKTP = "122323232323232"

	// Conversi Type Data
	var contoh string = "3333333333333333"
	var contoh_no_ktp NoKTP = NoKTP(contoh)

	println(ktp)
	println(NoKTP("22222222222"))
	println(contoh_no_ktp)

}