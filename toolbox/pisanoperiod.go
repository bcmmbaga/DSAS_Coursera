package toolbox

// Problem Description
// Task. Given two integers n and m, output F(n) mod m
// Input Format:: Twointergers a and m
// Constraints. 1 ≤ n ≤ 10^18 , 2 ≤ m ≤ 10^3
// Output Format. Output F(n) mod m

// Pisano period
func fiboModulo(n, m int) int {
	fib := []int{0, 1}

	if n <= 2 {
		return n
	}

	for i := 2; ; i++ {
		fib = append(fib, (fib[i-1]+fib[i-2])%m)
		if fib[i] == 1 && fib[i-1] == 0 {
			fib = fib[:len(fib)-2]
			break
		}
	}
	return fib[n%(len(fib))]
}
