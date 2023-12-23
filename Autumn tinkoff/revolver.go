package main

func mostExpensiveRevolver(revolversAmount int, joeMoney int, revolversPrices []int) int {
	if revolversAmount <= 0 {
		return 0
	}

	mostExpensive := 0
	left := 0
	right := revolversAmount - 1
	for left < right {
		if joeMoney >= revolversPrices[left] && revolversPrices[left] > mostExpensive {
			mostExpensive = revolversPrices[left]
			left++
		} else if joeMoney >= revolversPrices[right] && revolversPrices[right] > mostExpensive {
			mostExpensive = revolversPrices[right]
			right--
		} else {
			break
		}
	}
	return mostExpensive
}
