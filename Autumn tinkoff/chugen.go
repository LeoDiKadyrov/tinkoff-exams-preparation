package main

func findOptimalX(n, m int, roads [][3]int) int {
	var result int
	low, high := 0, roads[0][2]

	for low <= high {
		mid := (low + high) / 2
		connectedCities := make([]int, n)

		var findParent func(city int) int
		findParent = func(city int) int {
			if connectedCities[city] != city {
				connectedCities[city] = findParent(connectedCities[city])
			}
			return connectedCities[city]
		}

		merge := func(city1, city2 int) {
			parent1 := findParent(city1)
			parent2 := findParent(city2)
			if parent1 != parent2 {
				connectedCities[parent1] = parent2
			}
		}

		for j := 0; j < m && roads[j][2] > mid; j++ {
			merge(roads[j][0]-1, roads[j][1]-1)
		}

		if isConnected(connectedCities) {
			result = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return result
}

func isConnected(connectedCities []int) bool {
	state := connectedCities[0]
	for i := 1; i < len(connectedCities); i++ {
		if connectedCities[i] != state {
			return false
		}
	}
	return true
}