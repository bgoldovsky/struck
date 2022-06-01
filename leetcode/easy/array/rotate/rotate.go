package rotate

/*
	Rotate Array

	Given an array, rotate the array to the right by k steps, where k is non-negative.

	Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
	Could you do it in-place with O(1) extra space?
*/

/*
	Time: O(n)
	Space: O(r)

	       k        r
	[1][2][3][4][5][6][7]
	[6][7][1][2][3][4][5]
	k = 2
	r = len(nums) - 2 = 5

	1. Рассчитываем индекс r - это граница между двумя группами элементов:
	- элементы из начала массива которые будут перенесены в конец.
	- элементы из конца массива, которые будут перенесены в начало.
	2. Копируем кусок массива до r в новый массив tmp
	3. Переносим элементы с r до конца массива в начало
	4. Переносим элементы из массива tmp обратно записывая их с k до конца массива
*/
func rotateV1(nums []int, k int) {
	if len(nums) == 0 {
		return
	}

	k = k % len(nums)
	r := len(nums) - k

	tmp := make([]int, r)
	for i := 0; i < r; i++ {
		tmp[i] = nums[i]
	}

	for i, j := 0, r; i < k; i, j = i+1, j+1 {
		nums[i] = nums[j]
	}

	for i, j := k, 0; i < len(nums); i, j = i+1, j+1 {
		nums[i] = tmp[j]
	}
}

/*
	Time: O(n^2)
	Space: O(1)
*/
func rotateV2(nums []int, k int) {
	if len(nums) == 0 || k == 0 {
		return
	}

	for i := 0; i < k; i++ {
		rotateByOne(nums)
	}
}

func rotateByOne(nums []int) {
	last := nums[len(nums)-1]

	for i := len(nums) - 2; i >= 0; i-- {
		nums[i+1] = nums[i]
	}

	nums[0] = last
}
