package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg     int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't wirk with it"}
	}
	return arg + 3, nil
}

func main() {
	_, err := f(42)
	var ae *argError //Im confused as fuck as to waht this actually holds
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err dosn't match argError")
	}
}
