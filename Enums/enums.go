package main

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
	StateRandom
)

var StateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

//I don't fully understand the diff between
//Creating a method for a type and simply defning a function
func (ss ServerState) String() string {
	return StateName[ss]
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state %s", s))
	}
}

func main() {
	fmt.Println(StateIdle)
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(StateRandom)
	fmt.Println(ns2)
}
