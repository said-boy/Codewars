package code6kyu

// Link codewars : https://www.codewars.com/kata/556deca17c58da83c00002db/train/go

func Tribonacci(signature [3]float64, n int) []float64 {
	a, b, c := signature[0], signature[1], signature[2]
	result := []float64{}
	for i := 0; i < n; i++ {
		result = append(result, a)
		a, b, c = b, c, a+b+c
	}
	return result
}
