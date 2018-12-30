package toolbox

// Problem Description
// Task. Given an integer n, find the nth Fibonacci number Fn
// Input Format. The input consists of a single integer Fn
// Constraints. 0 ≤ n ≤ 45
// Output Format. Output Fn

func nthfibnumberNaive(f int) int {
	if f <= 1 {
		return f
	}
	return nthfibnumberNaive(f-1) + nthfibnumberNaive(f-2)
}

func nthFibnumberFast(f int) int {
	fib := []int{0, 1}

	for i := 2; i <= f; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	return fib[f]
}

// Problem Description
// Task. Given an integer n, find the last digit of the nth Fibonacci number F(n) (that is, F(n)mod 10).
// Input Format. The input consists of a single integer n.
// Constraints. 0 ≤ n ≤ 10^7 .
// Output Format. Output the last digit of F(n)

func lastDigitofnthFibnumberNaive(f int) int {
	fib := []int{0, 1}

	//as i grows the ith iteration of the loop computes the sum of longer
	//and longer numbers which do not fit in Go standard int type
	for i := 2; i <= f; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}

	return fib[f] % 10
}

func lastDigitofnthFibnumberFast(f int) int {
	fib := []int{0, 1}

	for i := 2; i <= f; i++ {
		fib = append(fib, (fib[i-1]+fib[i-2])%10)
	}
	return fib[f]
}
