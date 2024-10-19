package main

import (
	"fmt"
	"sort"
	"strconv"

	"strings"
	"testing"
)

func TestSoalEnam1(t *testing.T) {
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i % 3 == 0 {
			fmt.Println("fizz")
		} else if i % 5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func TestSoalEnam2(t *testing.T) {
	var asli string = "Kasur ini rusak"
	var kebalikan string = ""
	var panjang = len(asli)

	for i := panjang - 1; i >= 0; i-- {
		kebalikan = kebalikan + string(asli[i])
	}

	if strings.ToLower(asli) == strings.ToLower(kebalikan) {
		fmt.Println("Merupakan palindrom")
	} else {
		fmt.Println("Bukan palindrom")
	}
}

// Func soal 6 nomer 3
func FactorialRecursive(number int) int {
	if number == 1 {
		return 1
	} else {
		return number * FactorialRecursive(number - 1)
	}
}

func TestSoalEnam3(t *testing.T) {
	fmt.Println(FactorialRecursive(10))
}

func TestSoalEnam4(t *testing.T) {
	value := []int{
		0, -2, 3, 4, 7, 11, 12, 17, 19, 23, 26,
	}

	sort.Ints(value)
	minimum := value[0]
	maximum := value[len(value)-1]
	fmt.Println("Minimum:", minimum)
	fmt.Println("Maximum:", maximum)
}

func TestSoalEnam5(t *testing.T) {
	value := "belajar bahasa golang"
	valueint := 21

	fmt.Println(len(value) + len(strconv.Itoa(valueint)))
}
