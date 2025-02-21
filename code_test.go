package main

import (
	"math/rand/v2"
	"slices"
	"sort"
	"testing"
)

func generateRandomArray(maxLen, maxVal int) []int {
	arr := make([]int, rand.IntN(maxLen)+1)
	for i := range arr {
		arr[i] = rand.IntN(maxVal) - rand.IntN(maxVal)
	}
	sort.Ints(arr)
	return arr
}

func TestMaxCordCover(t *testing.T) {
	// 二分法+贪心
	var validate = func(points []int, k int) int {
		var res int
		for i, point := range points {
			idx, _ := slices.BinarySearch(points[:i+1], point-k)
			res = max(res, i-idx+1)
		}
		return res
	}

	for i := 0; i < 10000; i++ {
		points := generateRandomArray(10000, 100000)
		k := rand.IntN(100)
		expected := validate(points, k)
		actual := maxCordCover(points, k)
		if actual != expected {
			t.Logf("points: %v", points)
			t.Logf("k: %v", k)
			t.Fatalf("got %d, want %d", actual, expected)
		}
	}
}
