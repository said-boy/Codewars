package code6kyu

// Link codewars : https://www.codewars.com/kata/54da5a58ea159efa38000836/train/go

func FindOdd(seq []int) int {
	for _, num := range seq {
		count := 0
		for _, elem := range seq {
			if num == elem {
				count++
			}
		}
		if count%2 == 1 {
			return num
		}
	}
	return 0
}
