package driedel

import "testing"

func TestSpin(t *testing.T) {
	cases := []struct {
		name     string
		seed     int64
		expected SpinResult
	}{
		{
			name:     "hey",
			seed:     0,
			expected: Hey,
		},
		{
			name:     "gimel",
			seed:     1,
			expected: Gimel,
		},
		{
			name:     "shin",
			seed:     15,
			expected: Shin,
		},
		{
			name:     "nun",
			seed:     3,
			expected: Nun,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			driedel := NewDriedelWithSeed(tc.seed)
			result := driedel.Spin()
			if result != tc.expected {
				t.Errorf("expected driedel to spin %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestSpinRandomization(t *testing.T) {
	driedel := NewDriedelWithSeed(1)
	results := map[SpinResult]int{}
	for i := 0; i < 100; i++ {
		result := driedel.Spin()
		newVal := results[result] + 1
		results[result] = newVal
	}
	for _, val := range []SpinResult{Nun, Gimel, Hey, Shin} {
		if count := results[val]; count < 15 {
			t.Errorf("count for %s is %d", val, count)
		}
	}
}
