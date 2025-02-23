package logic_01

func Soal7(n int) []int {
	slice := make([]int, n)
	mid := n / 2 // cari indeks tengah
	val := 1     // mulai dr angka ganjil pertama

	for i := 0; i <= mid; i++ {
		slice[i] = val     // isi angka sebelah kiri (dr indeks 0 -- mid)
		slice[n-1-i] = val // isi sisi kanan dr index n-1 ke mid
		val += 2           // supaya ganjil
	}
	return slice
}
