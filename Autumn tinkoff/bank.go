package main

func robBank(n int, m int, arr []int) (int, []int) {
	all_notes := duplicateArrayByTwo(arr)
  
	var res []int
	for i := len(all_notes) - 1; i >= 0; i-- {
		if n >= all_notes[i] {
			res = append(res, all_notes[i])
			n -= all_notes[i]
		}
	}
  
	if n == 0 {
	  	return len(res), res
	} else {
	  	return -1, nil
	}
  }

func duplicateArrayByTwo(numbers []int) []int {
	result := make([]int, len(numbers)*2)
	index := 0
	for _, num := range numbers {
		result[index] = num
		index++
		result[index] = num
		index++
	}
	return result
}