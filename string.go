package main

func main() {
	println("String di Golang")
	println(len("Hello, World!")) // Menghitung panjang string
	println("Hello, " + "World!") // Menggabungkan string
	println("Hello, World!"[7])   // Mengakses karakter pada index tertentu
	println("Hello World"[0:6])   // Mengakses substring
	println("Hello" == "World")   // Membandingkan dua string
}

// Note
/*
- len("string") -> untuk menghitung panjang string
- string1 + string2 -> untuk menggabungkan string
- string[index] -> untuk mengakses karakter pada index tertentu
- string1 == string2 -> untuk membandingkan dua string
- string1 != string2 -> untuk membandingkan dua string
- string1 < string2 -> untuk membandingkan dua string
- string1 <= string2 -> untuk membandingkan dua string
- string1 > string2 -> untuk membandingkan dua string
- string1 >= string2 -> untuk membandingkan dua string
- string[index] = 'a' -> untuk mengubah karakter pada index tertentu (tidak bisa karena string bersifat immutable)

*/