package cal

func SumOfFirst(n int) int {
	result := 0

	for i := 0; i < n; i++ {
		result += n - i
	}

	return result
}