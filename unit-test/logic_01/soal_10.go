package logic_01

import "fmt"

func Soal10(n int) []int {
	slice := make([]int, n)
	num := 2

	for i := 1; i < n; i++ {
		slice[i] = num

		if num%2 != 0 {
			fmt.Print("buzz ", "\t")
		} else {
			fmt.Print(num, "\t")
		}
		num++ // iterasi ke angka berikutnya
	}
	return slice
}
