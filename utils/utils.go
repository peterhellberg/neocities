package utils

import (
	"errors"
	"fmt"
	"os"
)

func Check(err error) {
	if err != nil {
		fmt.Println("Error:", err)

		os.Exit(1)
	}
}

func Error(message string) error {
	return errors.New(message)
}
