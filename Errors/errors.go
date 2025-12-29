package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Don't fw 42, pick another #")
	}
	return arg + 3, nil
}

var ErrorOutOFTea = errors.New("NO MORE TEA")
var ErrorPower = errors.New("Can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrorOutOFTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrorPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil {
			fmt.Println("f failed: ", e)
		} else {
			fmt.Println("f worked: ", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrorOutOFTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrorPower) {
				fmt.Println("Now it is dark. (we out of power)")
			} else {
				fmt.Printf("Unknown error: %s\n", err)
			}
			continue
		}
		fmt.Println("Tea is ready!")
	}
}
