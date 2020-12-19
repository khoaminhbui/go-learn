package domain

// TrialDivision salculate largest prime number under an integer
func TrialDivision(n int) int {
	for i := n - 1; i >= 2; i-- {
		var isPrime = true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			return i
		}
	}

	return 1
}
