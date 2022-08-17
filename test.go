package main

import (
	"fmt"
	"testing"

	"go.uber.org/multierr"
)

func TestHelloEmpty(t *testing.T) error {
	msg := "123"
	if msg != "" {
		err1 := fmt.Errorf("error 1")
		err2 := fmt.Errorf("error 2")
		return multierr.Combine(err1, err2)
	}

	return nil
}
