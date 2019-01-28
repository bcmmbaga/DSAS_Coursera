package greedy

// You are responsible for collecting signatures from all tenants of a certain build-
// ing. For each tenant, you know a period of time when he or she is at home.
// You would like to collect all signatures by visiting the building as few times as
// possible.
// The mathematical model for this problem is the following. You are given a set
// of segments on a line and your goal is to mark as few points on a line as possible
// so that each segment contains at least one marked point

type segment struct {
	start, end int
}

func optimalPoints(s []segment) []int {
	sort.SliceStable(s, func(i, j int) bool { return s[i].end < s[j].end })
	point := s[0].end

	points := make([]int, 0)
	points = append(points, point)
	for _, seg := range s {
		if seg.start > point {
			points = append(points, seg.end)
			point = seg.end
		}
	}
	return points
}
