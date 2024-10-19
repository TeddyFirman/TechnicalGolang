package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestSoal2(t *testing.T) {
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
