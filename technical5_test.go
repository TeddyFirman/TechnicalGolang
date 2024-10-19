package main

import (
	"fmt"
	"testing"
	"time"
)

func TestSoal5(t *testing.T) {

	waktu, err := time.Parse("03:04:05PM", "06:30:00PM")

	//! Error
	// waktu, err := time.Parse("13:04:05PM", "06:30:00PM")

	if err != nil {
		fmt.Println("Invalid input format")
	} else {
		fmt.Println(waktu.Format("15:04:05"))
	}

}
