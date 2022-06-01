package main

/*
	Palindrome Linked List

	Given the head of a singly linked list, return true if it is a palindrome.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int, next *ListNode) *ListNode {
	return &ListNode{
		Val:  val,
		Next: next,
	}
}

func IsPalindrome(head *ListNode) bool {
	// 1. Пустой список или список из 1 элемента всегда является палиндромом
	if head == nil || head.Next == nil {
		return true
	}

	isPalindrome := true

	// 2. Получаем длину связанного списка
	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}

	// 3. Разворачиваем первую половину списка
	step := length / 2
	var prev *ListNode
	cur := head
	for i := 1; i <= step; i++ {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	mid := cur

	// 4. Уравниваем длину, если количество элементов списка нечетное
	var left, right *ListNode = prev, nil
	if length%2 == 0 {
		right = mid
	} else {
		right = mid.Next
	}

	// 5. Идем от середины списка, сравнивая левую и правую части
	for left != nil && right != nil {
		// Если значения не равны, то список не палиндром
		if left.Val != right.Val {
			isPalindrome = false
			break
		}
		// Идем дальше
		left = left.Next
		right = right.Next
	}

	// 6. Восстанавливаем первую половину списка
	cur = prev
	prev = mid
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}

	return isPalindrome
}
