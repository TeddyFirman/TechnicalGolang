package main

import (
	"fmt"
	"testing"

)

func TestSoal1(t *testing.T) {
	jumlah := 0

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		jumlah += i
	}

	angka := 5

	fmt.Println("Total angka pada file:", angka)
	fmt.Println("Jumlah semua angka:", jumlah)
}
