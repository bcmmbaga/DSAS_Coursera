package greedy

// The goal  int this problem is to find the minimum number of coins
// Needed to change the input value (interger) into coins with
// Denominator 1, 5, and 10

// Sample
// Input = 28 Output = 6
// 28 = 10 + 10 + 5 + 1 + 1 + 1 = 6 coins (minimum)
func getChange(amount int) int {
	coins := []int{10, 5, 1}
	count := 0

	for _, coin := range coins {
		dividend, remainder := amount/coin, amount%coin
		count += dividend
		amount = remainder
	}
	return count
}
