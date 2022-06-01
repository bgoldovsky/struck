package main

import "fmt"

/*
	Reverse Integer

	Given a signed 32-bit integer x, return x with its digits reversed.
	If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
*/

func main() {
	fmt.Println(Reverse(-12345))
}

// Reverse Возвращает число с инвертированным порядком чисел
// Time complexity: O(log10 n), мы делим входное значение на 10 каждую итерацию
// Space complexity: O(1)
func Reverse(x int) int {
	reverted := 0
	for x != 0 {
		// Добавляем последний разряд x в первый разряд reversed
		reverted = reverted*10 + x%10
		// Удаляем последний разряд у x
		x /= 10
	}

	return reverted
}
