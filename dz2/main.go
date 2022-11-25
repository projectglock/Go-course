// Задача Чудные вхождения в массив
package main

import "fmt"

func main() {
	arr := []int{6, 2, 4, 1, 4, 6, 2}
	fmt.Println(Solution(arr))
}

//O(n) по времени, O(1) по памяти, скорее всего быстрее невозможно
func Solution(A []int) int {
	var result int
	for i := 0; i < len(A); i++ {
		result ^= A[i]
	}
	return result
}
