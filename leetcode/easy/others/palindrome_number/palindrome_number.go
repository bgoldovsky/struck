package main

import "fmt"

/*
	Palindrome Number

	Given an integer x, return true if x is palindrome integer.
	An integer is a palindrome when it reads the same backward as forward.
	For example, 121 is palindrome while 123 is not.
*/

func main() {
	res := IsPalindrome(1234)
	fmt.Println(res)
}

// IsPalindrome Определяет, является ли число палиндромом
// Time complexity: O(log10 n), мы делим входное значение на 10 каждую итерацию
// Space complexity: O(1)
func IsPalindrome(x int) bool {
	// Отрицательное число не может быть палиндромом
	if x < 0 {
		return false
	}

	// Если последняя цифра ноль, то число не может быть палиндромом, если само число не ноль
	if x%10 == 0 && x != 0 {
		return false
	}

	// Разворачиваем правую половину числа x
	// и отрезаем развернутую часть
	reverted := 0
	for x > reverted {
		// Добавляем последний разряд x в первый разряд reversed
		reverted = reverted*10 + x%10
		// Удаляем последний разряд у x
		x /= 10
	}

	// Число палиндром, если
	// 1. правая развернутая часть равна оставшейся левой
	// 2. Если количество цифр нечетное надо отрезать цифру, которая была в середине, а теперь в младшем разряде reversed
	return x == reverted || x == reverted/10
}
