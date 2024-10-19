package main

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestSoal3(t *testing.T) {
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		fmt.Println(err.Error())
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(responseData))
}
