// Задача Поиск отсутствующего элемента
package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{5, 2, 4, 3}
	fmt.Println(Solution(arr))
}

// сортировка тут скорее всего за O(nlog(n)) поэтому такая сложность по времени и будет
func Solution(A []int) int {
	sort.Ints(A)
	for i := 1; i <= len(A); i++ {
		if i != A[i-1] {
			return i
		}
	}
	return 0
}
