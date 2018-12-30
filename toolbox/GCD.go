package toolbox

// Problem Description
// Task: Given two integers a and b, find their greatest common divisor
// Input Format. The two integers a, b
// Constraints. 1 ≤ a, b ≤ 2 · 10^9
// Output Format. Output GCD(a, b).

// Euclidean algorithm for computing
// the greatest common divisor
func gcd(a, b int64) int64 {
	for b != 0 {
		s := b
		b = a % b
		a = s
	}

	return a
}
