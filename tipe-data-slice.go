package main

import "fmt"

func main() {

	/*
		- Tipe data slice adalah potongan dari data Array
		- Slice mirip dengan Array, yang membedakan adalah ukuran Slice bisa berubah
		- Slice dan Array selalu terkoneksi, dimana Slice adalah data yang mengakses sebagian
			atau seluruh data di Array

		- Tipe data Slice memiliki 3 data, yaitu pointer, lenght dan capacity
		- Pointer adalah petunjuk data pertama di array pada Slice
		- Lenght adalah panjang dari slice
		- Capacity adalah kapasitas dari Slice, dimana lenght tidak boleh lebih dari capacity
	*/

	/*
		MEMBUAT SLICE DARI ARRAY
		- array[low:high]		=> Membuat slice dari array dimulai index low sampai index sebelum high
		- array[low:]				=> Membuat slice dari array dimulai index low sampai index akhir di array
		- array[:high]			=> Membuat slice dari array dimulai index 0 sampai index sebelum high
		- array[:]					=> Membuat slice dari array dimulai index 0 sampai index akhir di array
	*/

	names := [...]string{"Budi", "Rizal", "Paimin", "Joko", "Saipul", "Herman", "Galuh"}

	slice1 := names[2:6]
	fmt.Println(slice1) // pointernya  = 2, lenght = 4, capacity = 5
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))


	slice2 := names[:3] // mengambil semua dari awal sampai index ke 3
	fmt.Println(slice2)

	slice3 := names[3:] // mengambil semua dari index ke 3 sampai akhir
	fmt.Println(slice3)

	slice4 := names[:] // mengambil semua data 
	fmt.Println(slice4)


	// Cara deklarasi tipe data slice dari array
	var sliceExample []string = names[:]
	fmt.Println(sliceExample)
}