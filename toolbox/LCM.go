package toolbox

// Problem Description
// Task. Given two integers a and b, find their least common multiple.
// Input Format. The two integers a and b
// Constraints. 1 ≤ a, b ≤ 2 · 10 9 .
// Output Format. Output the least common multiple of a and b

func lcm(a, b int64) int64 {
	res := a * b

	for b != 0 {
		s := b
		b, a = a%b, s
	}
	return res / a
}
