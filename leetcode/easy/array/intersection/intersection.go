package intersection

import "sort"

/*
	Intersection of Two Arrays II

	Given two integer arrays nums1 and nums2, return an array of their intersection.
	Each element in the result must appear as many times as it shows in both arrays,
	and you may return the result in any order.

	What if the given array is already sorted? How would you optimize your algorithm?
	What if nums1's size is small compared to nums2's size? Which algorithm is better?
	What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?
*/

func intersectV1(nums1 []int, nums2 []int) []int {
	var res []int

	m := make(map[int]int)

	// Перекладываем числа из первого массива в мапу с указанием количества
	for _, num := range nums1 {
		m[num]++
	}

	// Если в мапе есть числа из второго массива, то добавляем их результирующий массив
	for _, num := range nums2 {
		if m[num] > 0 {
			res = append(res, num)
			m[num]--
		}
	}
	return res
}

func intersectV2(nums1 []int, nums2 []int) []int {
	var res []int

	sort.Ints(nums1)
	sort.Ints(nums2)

	ptr1, ptr2 := 0, 0

	for ptr1 < len(nums1) && ptr2 < len(nums2) {
		if nums1[ptr1] < nums2[ptr2] {
			ptr1++
		}

		if nums2[ptr2] < nums1[ptr1] {
			ptr2++
		}

		if nums1[ptr1] == nums2[ptr2] {
			res = append(res, nums1[ptr1])
			ptr1++
			ptr2++
		}
	}

	return res
}
