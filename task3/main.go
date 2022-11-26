// Задача Проверка последовательности
package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(Solution(arr))
}

// O(n) по времени, O(1) по памяти
func Solution(A []int) int {
	max_num := A[0]
	for i := 0; i < len(A); i++ {
		if A[i] > max_num {
			max_num = A[i]
		}
	}
	if max_num > len(A) {
		return 0
	}
	return 1
}

// это первое решение, не самое оптимальное
// сортировка тут скорее всего за O(nlog(n)) поэтому такая сложность по времени и будет
func Solution2(A []int) int {
	sort.Ints(A)
	for i := 0; i < len(A); i++ {
		if i+1 != A[i] {
			return 0
		}
	}
	return 1
}
