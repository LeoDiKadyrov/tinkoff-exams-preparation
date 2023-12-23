/*
6 Задание
Поздравляю, вы дошли до серьёзной задачи, больше никаких легенд, только хардкор.

Дан массив a из n целых чисел. Требуется выполнить q запросов такого вида (1≤l≤r≤n,0≤k,b,x≤10^9)(1≤l≤r≤n,0≤k,b,x≤10^9):

    + l r x — прибавить xx ко всем ai​ на отрезке i∈[l,r]i∈[l,r]
    ? l r k b — вывести maxl≤i≤rmin(ai,k⋅i+b)maxl≤i≤r​min(ai​,k⋅i+b)

Формат входных данных

В первой строке заданы два числа n, q (1≤n≤2⋅10^5,1≤q≤5⋅10^5)(1≤n≤2⋅10^5,1≤q≤5⋅10^5).

Во второй строке задан массив aa (0≤ai≤10^9)(0≤ai​≤10^9).

Следующие q строк содержат запросы в заданном формате. Гарантируется, что будет хотя бы один запрос типа ?.

Формат выходных данных

Для каждого запроса типа ? выведите ответ в отдельной строке.

Примеры данных
Ввод:
6 3
2 4 6 8 10 12
? 2 5 3 0
+ 2 3 6
? 2 5 3 2

Вывод:
10
11
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abd() {
	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	firstLine := strings.Fields(strings.TrimSpace(line))
	n, _ := strconv.Atoi(firstLine[0])
	q, _ := strconv.Atoi(firstLine[1])

	line, _ = reader.ReadString('\n')
	aStr := strings.Fields(strings.TrimSpace(line))
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(aStr[i])
	}

	for i := 0; i < q; i++ {
		line, _ = reader.ReadString('\n')
		query := strings.Fields(strings.TrimSpace(line))
		if query[0] == "+" {
			l, _ := strconv.Atoi(query[1])
			r, _ := strconv.Atoi(query[2])
			x, _ := strconv.Atoi(query[3])
			for j := l - 1; j < r; j++ {
				a[j] += x
			}
		} else if query[0] == "?" {
			l, _ := strconv.Atoi(query[1])
			r, _ := strconv.Atoi(query[2])
			k, _ := strconv.Atoi(query[3])
			b, _ := strconv.Atoi(query[4])
			maxVal := 0
			for j := l - 1; j < r; j++ {
				minVal := min(a[j], k*(j+1)+b)
				if minVal > maxVal {
					maxVal = minVal
				}
			}
			fmt.Println(maxVal)
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
