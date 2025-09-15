package main

import "fmt"

func main() {

	var names [3]string

	names[0] = "Al"
	names[1] = "Irfan"
	names[2] = "Yasin"

	println(names[0])
	println(names[1])
	println(names[2])

	// names[3] = "Diubah" // ini error karena tdk ada index ke 3

	angkaTest := [...]int{1, 2, 3, 4, 5} // [...] jumlah array tidak di definisikan di awal
	for i := 0; i < len(angkaTest); i++ {
		println(angkaTest[i])
	}

	b := [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)



	var nilai = [3] int {23,34,45}
	fmt.Println(nilai)
	println(nilai[1])

	// Function Array
	/*
	- len(array) => Jumlah panjang array
	- array[index] => mendapatkan data di posisi index
	- array[index] = value => mengubah data dari posisi index
	*/

	println(len(nilai))

}