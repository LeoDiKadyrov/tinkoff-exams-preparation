/*
2 Задание
Тинькофф начал разрабатывать новый проект. Для этого было подобрано n разработчиков. У i-го разработчика есть порог социальности ai​, это значит, что он готов контактировать напрямую с не более чем ai​ другими разработчиками.

Определите, можно ли наладить контакт между какими-то парами разработчиков, так чтобы любые два контактировали либо напрямую, либо через других разработчиков.

Формат входных данных

Каждый тест состоит из нескольких наборов входных данных.

В первой строке находится одно целое число t (1≤t≤1000)(1≤t≤1000) — количество наборов входных данных. Далее следует описание наборов входных данных.

Первая строка каждого набора входных данных содержит одно число n (1≤n≤105)(1≤n≤105) — количество разработчиков. Вторая строка содержит n натуральных чисел ai​ (1≤ai≤109)(1≤ai​≤109) — пороги социальностей разработчиков. Гарантируется, что сумма значений n по всем наборам входных данных не превосходит 105.

Формат выходных данных

Для каждого набора входных данных, выведите «Yes», если можно наладить контакт между программистами, и «No» иначе. Вы можете выводить каждую букву в любом регистре (строчную или заглавную). Например, строки «yEs», «yes», «Yes» и «YES» будут приняты как положительный ответ. 

Примеры данных
Ввод:
4
1
1000000000
2
1 1
3
1 1 1
4
1 1 2 2

Вывод:
Yes
Yes
No
Yes
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abdrah() {
	reader := bufio.NewReader(os.Stdin)
	tStr, _ := reader.ReadString('\n')
    t, _ := strconv.Atoi(strings.TrimSpace(tStr))

	for i := 0; i < t; i++ {
        nStr, _ := reader.ReadString('\n')
        n, _ := strconv.Atoi(strings.TrimSpace(nStr))

        thresholdsStr, _ := reader.ReadString('\n')
        thresholdsStr = strings.TrimSpace(thresholdsStr)
        thresholdsSlice := strings.Split(thresholdsStr, " ")
        socialThresholds := make([]int, n)
        for j, s := range thresholdsSlice {
            socialThresholds[j], _ = strconv.Atoi(s)
        }

        graphConnectivity(n, socialThresholds)
    }
}

func graphConnectivity(n int, socialThresholds []int) {
    sum := 0
    for _, threshold := range socialThresholds {
        sum += threshold
    }
    if sum >= 2*(n-1) {
        fmt.Println("Yes")
    } else {
		fmt.Println("No")
	}
}