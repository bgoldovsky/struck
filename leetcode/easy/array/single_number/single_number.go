package single_number

/*
	Single Number

	Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

	You must implement a solution with a linear runtime complexity and use only constant extra space.
*/

// Time: O(n)
// Space: O(n)
func singleNumberNaive(nums []int) int {
	unique := make(map[int]int)

	for _, num := range nums {
		unique[num]++
	}

	for digit, count := range unique {
		if count == 1 {
			return digit
		}
	}

	return -1
}

// Time: O(n)
// Space: O(1)
func singleNumber(nums []int) int {
	res := 0

	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}

	return res
}
