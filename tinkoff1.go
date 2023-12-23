/*
1 Задание
Для вывески нового офиса Тинькофф были заказаны неоновые буквы.

В офис привезли какой-то набор из больших латинских букв. Проверьте, что из них действительно можно составить строку «TINKOFF» для вывески. Тинькофф не хочет платить за лишние буквы, поэтому должны быть использованы все привезённые буквы.

Формат входных данных

Каждый тест состоит из нескольких наборов входных данных. В первой строке находится одно целое число t (1≤t≤100)(1≤t≤100) — количество наборов входных данных. Далее следует описание наборов входных данных. Единственная строка каждого набора входных данных содержит одну непустую строку из больших латинских букв длиной не более 20 символов — привезённый набор букв. 

Формат выходных данных

Для каждого набора входных данных, в отдельной строке, выведите «Yes» (без кавычек), если из всех привезённых букв можно составить строку «TINKOFF», и «No» иначе. Вы можете выводить каждую букву в любом регистре (строчную или заглавную). Например, строки «yEs», «yes», «Yes» и «YES» будут приняты как положительный ответ.

Примеры данных
Ввод:
4
TINKOFF
TINKOFFF
AAAA
FFOKNIT

Вывод: 
Yes
No
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

func abdrahman() {
	reader := bufio.NewReader(os.Stdin)
	tStr, _ := reader.ReadString('\n')
    t, _ := strconv.Atoi(strings.TrimSpace(tStr))

	targetFreq := map[rune]int{
		'T': 1,
		'I': 1,
		'N': 1,
		'K': 1,
		'O': 1,
		'F': 2,
	}

	for i := 0; i < t; i++ {
		line, _ := reader.ReadString('\n')
		word := strings.TrimSpace(line)
		tinkoffAssemble(word, targetFreq)
	}
}

func tinkoffAssemble(word string, targetFreq map[rune]int) {
	wordFreq := make(map[rune]int)
	for _, char := range word {
		wordFreq[char]++
	}

	if len(word) != len("TINKOFF") {
		fmt.Println("No")
		return
	}

	for char, freq := range targetFreq {
		if wordFreq[char] != freq {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}