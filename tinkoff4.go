/*
4 Задание
Боб начинает свой путь в Тинькофф Инвестициях. Сейчас его интересуют k компаний, и он обязательно хочет купить акции всех этих компаний, чтобы у каждой компании была хотя бы одна акция.

Тинькофф Инвестиции предложили ему корневое дерево из nn вершин, в каждой из которых лежит пакет акций какой-то из интересующих его компаний, также задана стоимость каждого пакета. Вершины корневого дерева пронумерованы целыми числами от 1 до n.

Функционал нынешнего приложения позволяет Бобу купить поддерево этого дерева, на это он потратит столько денег, сколько суммарно стоят пакеты в поддереве и получит все акции из этого поддерева. В результате покупки Боб хочет, чтобы у него были акции всех интересующих его компаний. Поскольку Боб ещё студент, он хочет потратить минимальное количество денег.

Определите может ли Боб выкупить какое-то поддерево так, чтобы у него оказались все нужные ему акции, и если да, то какое минимальное количество денег он для этого должен потратить. 

Формат входных данных

В первой строке заданы два целых числа n, k — размер дерева и количество интересных для Боба компаний.

Следующие k строк содержат различные строки длиной не более 10 символов из латинских букв — названия компаний.

Следующие n строк содержат описание дерева. В i-ой строке находится описание i-й вершины дерева pi​, ai, ci​ :

    pi​ — номер родителя i-ой вершины или 0, если вершина является корнем
    ai​ — стоимость пакета акций в i-ой вершине
    ci​ — название компании, пакет акции которой лежит в i-ой вершине

Гарантируется, что компании, акции которых лежат в вершинах, интересуют Боба. Гарантируется, что входные данные задают корректное корневое дерево.

Формат выходных данных

Выведите единственное число — минимальное количество денег, которое должен потратить Боб, или −1, если выкупить акции всех компаний невозможно.

Примеры данных
Ввод:
5 2
A
B
0 1 A
1 2 A
1 2 B
1 1 B
4 2 A

Вывод:
3
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Children []*Node
	Price int
	Company string
}

func abdrahm() {
    reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	firstLine := strings.Fields(strings.TrimSpace(line))
	n, _ := strconv.Atoi(firstLine[0])
	k, _ := strconv.Atoi(firstLine[1])

	companies := make([]string, k)
	for i := 0; i < k; i++ {
		company, _ := reader.ReadString('\n')
		companies[i] = strings.TrimSpace(company)
	}

	nodes := make([]*Node, n+1)
	for i := range nodes {
		nodes[i] = &Node{}
	}

	for i := 1; i <= n; i++ {
		nodeLine, _ := reader.ReadString('\n')
		nodeDetails := strings.Split(strings.TrimSpace(nodeLine), " ")
		p, _ := strconv.Atoi(nodeDetails[0])
		a, _ := strconv.Atoi(nodeDetails[1])
		c := nodeDetails[2]

		nodes[i].Price = a
		nodes[i].Company = c

		if p != 0 {
			nodes[p].Children = append(nodes[p].Children, nodes[i])
		}
	}

	var root *Node
	for i := 1; i <= n; i++ {
		if len(nodes[i].Children) > 0 || nodes[i].Company != "" {
			root = nodes[i]
			break
		}
	}

	minCost := TreeTraversalForBob(root, companies)
	fmt.Println(minCost)

}

func TreeTraversalForBob(root *Node, companies []string) int {
    companyToBit := make(map[string]int)
    allCompaniesBit := 0
    for i, company := range companies {
        bit := 1 << i
        companyToBit[company] = bit
        allCompaniesBit |= bit
    }

    minCost := -1

    var dfs func(node *Node) int
    dfs = func(node *Node) int {
        if node == nil {
            return 0
        }

        currentBit := companyToBit[node.Company]
        subtreeCost := node.Price
        subtreeBit := currentBit

        for _, child := range node.Children {
            childCost := dfs(child)
            subtreeCost += childCost
            subtreeBit |= companyToBit[child.Company]
        }

        if subtreeBit == allCompaniesBit {
            if minCost == -1 || subtreeCost < minCost {
                minCost = subtreeCost
            }
        }

        return subtreeCost
    }

    dfs(root)
    return minCost
}

