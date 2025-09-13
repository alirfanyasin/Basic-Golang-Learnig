package main

import "fmt"

// =========================================
// Tipe Data
// =========================================
// int8				: -128 s/d 127= 8 bit
// int16			: -32.768 s/d 32.767 = 16 bit
// int32			: -2.147.483.648 s/d 2.147.483.647 = 32 bit
// int64			: -9.223.372.036.854.775.808 s/d 9.223.372.036.854.775.807 = 64 bit

// uint8			: 0 s/d 255 = 8 bit
// uint16			: 0 s/d 65.535 = 16 bit
// uint32			: 0 s/d 4.294.967.295 = 32 bit
// uint64			: 0 s/d 18.446.744.073.709.551.615 = 64 bit

// float32			: -3.402823466E+38 s/d 3.402823466E+38 = 32 bit
// float64			: -1.797693134862315708145274237317043567981E+308 s/d 1.797693134862315708145274237317043567981E+308 = 64 bit
// complex64		: 32 bit untuk real dan 32 bit untuk imaginer
// complex128		: 64 bit untuk real dan 64 bit untuk imaginer

// byte			: alias dari uint8
// rune			: alias dari int32, digunakan untuk karakter unicode
// int			: tipe data integer yang sesuai dengan arsitektur sistem (32 bit atau 64 bit)
// uint			: tipe data unsigned integer yang sesuai dengan arsitektur sistem (32 bit atau 64 bit)
// string			: kumpulan karakter
// bool			: true atau false

func main() {
	fmt.Println("Tipe Data di Golang")
	fmt.Println("int8:", int8(-128), "s/d", int8(127))
	fmt.Println("int16:", int16(-32768), "s/d", int16(32767))
	fmt.Println("int32:", int32(-2147483648), "s/d", int32(2147483647))
	fmt.Println("int64:", int64(-9223372036854775808), "s/d", int64(9223372036854775807))
}