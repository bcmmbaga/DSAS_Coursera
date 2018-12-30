package toolbox

import "testing"

func TestGcd(t *testing.T) {
	tests := []struct {
		a, b, expected int64
	}{
		{18, 35, 1},
		{28851538, 1183019, 17657},
		{1344, 217, 7},
	}

	for _, v := range tests {
		if actual := gcd(v.a, v.b); actual != v.expected {
			t.Errorf("Ecpected %d got %d", v.expected, actual)
		}
	}

}
