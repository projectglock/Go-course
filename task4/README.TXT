Представлен массив А состоящий из N целых чисел. 
В массиве представлены числа в диапазоне [1..(N + 1)], что значит, что в массиве отсутствует 1 элемент.

Цель - найти этот элемент

Напишите функцию:

func Solution(A []int) int

Функция возвращает элемент - который не представлен в массиве.

Например для массива А:

  A[0] = 2
  A[1] = 3
  A[2] = 1
  A[3] = 5

Функция должна вернуть 4.

Ограничения:
N это целое число для диапазона [0..100,000];
все элементы массива А уникальные;
каждый элемент А это число в  диапазоне [1..(N + 1)].