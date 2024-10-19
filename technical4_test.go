package main

import (
	"math/rand"
	"testing"
	"time"
)

func shuffleRecursively(arr []int, n int) []int {
	if n == 1 {
		return arr
	}

	randIndex := rand.Intn(n)

	arr[n-1], arr[randIndex] = arr[randIndex], arr[n-1]

	return shuffleRecursively(arr, n-1)
}

func TestSoal4(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	shuffled := shuffleRecursively(arr, len(arr))

	if isSame(arr, shuffled) {
		t.Errorf("Hasil acak seharusnya berbeda, tetapi tetap sama: %v", shuffled) // untuk melihat hasilnya
	}

}

func isSame(arr1, arr2 []int) bool {
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return true
		}
	}
	return false
}
