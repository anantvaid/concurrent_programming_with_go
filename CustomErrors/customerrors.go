package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type NetworkError struct {
	Code      int
	Msg       string
	Retryable bool
}

// Implement Error
func (ne *NetworkError) Error() string {
	return fmt.Sprintf("Network Error %d: %s (Retryable %t )", ne.Code, ne.Msg, ne.Retryable)
}

func checkError() error {
	outcome := rand.Intn(3)

	if outcome == 0 {
		return nil
	} else if outcome == 1 {
		return &NetworkError{
			Code:      500,
			Msg:       "Internal Server Error",
			Retryable: true,
		}
	} else {
		return &NetworkError{
			Code:      404,
			Msg:       "Page Not Found",
			Retryable: false,
		}
	}
}

func main() {
	err := checkError()
	if err != nil {
		var netErr *NetworkError

		if errors.As(err, &netErr) {
			if netErr.Retryable {
				fmt.Printf("Temporary Error, reason: %s", netErr.Msg)
			} else {
				fmt.Printf("Fatal Error, reason: %s", netErr.Msg)
			}
		} else {
			fmt.Println("Generic Error")
		}
	} else {
		fmt.Println("No Errors")
	}
}
