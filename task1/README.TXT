Массив A состоит из N целых чисел.
Ротация массива - это сдвиг каждого элемента вправо, все элементы с конца двигаются в начало. Например, ротация массива
A = [3, 8, 9, 7, 6] это [6, 3, 8, 9, 7] (все элементы сдвинуты вправо на 1 элемент и 6 сдвигается на первое место).
Цель - это сдвинуть массив A К раз. 

Необходимо написать функцию:

func Solution(A []int, K int) []int

К примеру для параметров

    A = [3, 8, 9, 7, 6]
    K = 3

функция должна вернуть [9, 7, 6, 3, 8]. Необходимо сделать 3 ротации:

    [3, 8, 9, 7, 6] -> [6, 3, 8, 9, 7]
    [6, 3, 8, 9, 7] -> [7, 6, 3, 8, 9]
    [7, 6, 3, 8, 9] -> [9, 7, 6, 3, 8]

Другой пример:

    A = [1, 2, 3, 4]
    K = 4

результат [1, 2, 3, 4]

Условия:

N и K целые числа в диапазоне [0..100];
каждый элемент массива A целые числа в диапазоне [−1,000..1,000].