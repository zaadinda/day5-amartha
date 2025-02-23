package logic_01

func Soal6(n int) []int {
	slice := make([]int, n)

	num := 30

	for i := n * 3; i >= 3; i -= 3 {
		slice[i] = num
		num -= 3
	}
	return slice
}
