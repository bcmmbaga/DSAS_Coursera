package greedy

type item struct {
	value, weight, ratio float32
}

func getOptimalValue(n, capacity int, values, weights []float32) float32 {
	items := make([]item, len(values))
	var value float32
	for i := 0; i < len(values); i++ {
		items[i] = item{
			value:  values[i],
			weight: weights[i],
			ratio:  values[i] / weights[i],
		}
	}
	sort.SliceStable(items, func(i, j int) bool { return items[i].ratio > items[j].ratio })

	for _, itm := range items {
		if capacity-int(itm.weight) >= 0 {
			value += itm.value
			capacity -= int(itm.weight)
		} else {
			value += (float32(capacity) / itm.weight) * itm.value
		}
	}
	return value
}
