package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("1")

	result := m.Run()
	fmt.Printf("result:%d", result)

	fmt.Println("this is the end")

	os.Exit(result)
}
