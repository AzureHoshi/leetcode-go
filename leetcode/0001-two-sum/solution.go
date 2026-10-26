package twosum

// TwoSum finds two indices in nums such that their values sum up to the given target.
//
// Time Complexity: O(n) - Single pass through the array with O(1) hash map lookups.
// Space Complexity: O(n) - Hash map stores up to n elements in the worst case.
func TwoSum(nums []int, target int) []int {
	numToIndex := make(map[int]int, len(nums))

	for i, num := range nums {
		complement := target - num
		if idx, found := numToIndex[complement]; found {
			return []int{idx, i}
		}
		numToIndex[num] = i
	}

	return nil
}
