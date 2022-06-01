package contains_duplicate

/*
	Contains Duplicate

	Given an integer array nums, return true if any value appears at least twice in the array,
	and return false if every element is distinct.
*/

// Time: O(n)
// Space O(n)
func containsDuplicate(nums []int) bool {
	unique := make(map[int]int)

	for _, num := range nums {
		unique[num]++
	}

	for _, count := range unique {
		if count > 1 {
			return true
		}
	}

	return false
}
