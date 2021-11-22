package main

import "fmt"

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLookBack
	FlagPointToPoint
	FlagMulticast
)
const format = "%b %t\n"

func IsUp(v Flags) bool {
	return FlagUp&v == FlagUp
}

func SetBroadcast(v *Flags) {
	*v |= FlagBroadcast
}

func TurnDown(v *Flags) {
	*v &^= FlagUp
}

func IsCast(v Flags) bool {
	return v&(FlagBroadcast|FlagMulticast) != 0
}

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf(format, v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf(format, v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf(format, v, IsUp(v))   // "10010 false"
	fmt.Printf(format, v, IsCast(v)) // "10010 true"
}
