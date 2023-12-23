/*
3 Задание
Максим пришёл в Тинькофф, чтобы взять кредит на покупку новогодних подарков. Ему предодобрен кредит в размере m бурлей, значит он может взять любое целое количество бурлей от 0 до m включительно.

У Максима есть заранее подготовленный список из n подарков. Он планирует идти по порядку вдоль списка и каждый раз, когда он видит подарок, на который у него хватает денег, он обязательно моментально его покупает.

Помогите ему посчитать, какое максимальное количество денег у него может остаться после закупки подарков, если он возьмёт кредит оптимального размера (то есть такого, чтобы у него осталось как можно больше денег после покупки подарков по алгоритму).

Формат входных данных

Первая строка содержит два целых числа n и m  — длина списка подарков и размер предодобренного кредита.

Вторая строка содержит n целых чисел — цены подарков ai​.

Формат выходных данных

Выведите одно число — максимальное количество бурлей, которое могло остаться у Максима после закупки подарков. 

Замечание 1

В первом примере Максим может взять в кредит 3 бурля, и тогда он купит только последний подарок. Во втором примере Максим может взять кредит на все 10 бурлей и купить все подарки. 

Замечание 2

Ошибок в примерах нет.

Примеры данных
Пример 1
Ввод:
3 10
5 4 1

Вывод:
2

Пример 1
Ввод:
3 10
1 2 3

Вывод:
4
*/

package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func abdra() {
    reader := bufio.NewReader(os.Stdin)

    line, _ := reader.ReadString('\n')
    firstLine := strings.Fields(strings.TrimSpace(line))
    n, _ := strconv.Atoi(firstLine[0])
    m, _ := strconv.Atoi(firstLine[1])

    line, _ = reader.ReadString('\n')
    gifts := strings.Fields(strings.TrimSpace(line))
    giftPrices := make([]int, n)
    for i := 0; i < n; i++ {
        giftPrices[i], _ = strconv.Atoi(gifts[i])
    }

    maxRemaining := 0
    for loan := 0; loan <= m; loan++ {
        moneyLeft := loan
        for _, price := range giftPrices {
            if moneyLeft >= price {
                moneyLeft -= price
            }
        }

        if moneyLeft > maxRemaining {
            maxRemaining = moneyLeft
        }
    }

    fmt.Println(maxRemaining)
}