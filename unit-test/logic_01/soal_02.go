package logic_01

func Soal2(n int) []int {
	slice := make([]int, n)

	num := 2

	for i := 0; i < n; i++ {
		slice[i] = num
		num += 2
	}
	return slice
}
