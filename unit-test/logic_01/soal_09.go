package logic_01

func Soal9(n int) []int {
	slice := make([]int, n)
	mid := n / 2
	val := 3

	for i := 0; i <= mid; i++ {
		slice[i] = val
		slice[n-1-i] = val
		val += 3
	}
	return slice
}
