package main

import (
	"slices"
	"sort"
)

func joeCards(n int, nums []int, win []int) string {
	l, r := 0, n-1

	for l <= r && nums[l] == win[l] {
		l++
	}
	for r >= 0 && nums[r] == win[r] {
		r--
	}

	if l >= r {
		return "YES"
	}
	sortPartly(nums, l, r)
	if slices.Equal(nums, win) {
		return "YES"
	}
	return "NO"
}

func sortPartly(nums []int, start, end int) {
	subSlice := nums[start : end+1]
	sort.Ints(subSlice)
}
