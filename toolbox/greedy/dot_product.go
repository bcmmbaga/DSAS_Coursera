package greedy

import "sort"

// You have n ads to place on a popular Internet page. For each ad, you know how
// much is the advertiser willing to pay for one click on this ad. You have set up n
// slots on your page and estimated the expected number of clicks per day for each
// slot. Now, your goal is to distribute the ads among the slots to maximize the
// total revenue
func maxDotProduct(a, b []int) int {
	sort.Ints(a)
	sort.Ints(b)

	res := 0
	for i, _ := range a {
		res += a[i] * b[i]
	}
	return res
}
