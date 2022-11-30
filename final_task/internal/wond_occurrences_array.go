// Задача Чудные вхождения в массив
// Копия для использования в финальной задаче
package internal

//основная идея - делаем XOR по всем элементам, результатом будет то число без пары
//O(n) по времени, O(1) по памяти, скорее всего быстрее невозможно
func WondOccurrencesArraySolve(A []int) int {
	var result int
	for i := 0; i < len(A); i++ {
		result ^= A[i]
	}
	return result
}
