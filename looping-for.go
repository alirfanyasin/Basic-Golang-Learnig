package main

func main() {

	i := 1
	for i <= 3 {
		println(i)
		i++
	}

	println("====================================")

	for j := 0; j < 3; j++ {
		println(j)
	}

	println("====================================")

	var angka = [5]int{16, 42, 52, 45, 10}

	// Index bisa di ganti menjadi underscore. _
	for index, k := range angka {
		println(index)
		println(k)
	}

	println("====================================")

	// Break Loop
	for {
		println("Loop")
		break
	}

	for l := range 10 {
		if l == 7 {
			break
		}
		println(l)
	}

	println("====================================")

	// Continue Loop
	for m := range 10 {
		if m%2 == 0 {
			continue
		}
		println(m)
	}

}