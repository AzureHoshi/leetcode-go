package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "standard positive numbers",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "contains negative numbers",
			nums:     []int{-3, 4, 3, 90},
			target:   0,
			expected: []int{0, 2},
		},
		{
			name:     "duplicate elements summing to target",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			name:     "target at end of array",
			nums:     []int{1, 5, 8, 12, 19, 21},
			target:   40,
			expected: []int{4, 5},
		},
		{
			name:     "no matching pair",
			nums:     []int{1, 2, 3},
			target:   7,
			expected: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := TwoSum(tc.nums, tc.target)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("TwoSum(%v, %d) = %v; want %v", tc.nums, tc.target, got, tc.expected)
			}
		})
	}
}

func BenchmarkTwoSum(b *testing.B) {
	nums := []int{2, 7, 11, 15, 23, 34, 45, 56, 67, 78, 89, 90, 101, 112, 123}
	target := 235

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		TwoSum(nums, target)
	}
}
