package services

// IsPrime returns true if n is a prime number, false otherwise.
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= n/i; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// LoopAndCheckPrimes iterates over numbers 1 through 10, calling IsPrime on
// each, and returns a map of each number to whether it is prime.
func LoopAndCheckPrimes() map[int]bool {
	results := make(map[int]bool)
	for i := 1; i <= 10; i++ {
		results[i] = IsPrime(i)
	}
	return results
}
