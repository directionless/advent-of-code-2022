package day15

import "fmt"

type location struct {
	X, Y int
}

func (l location) String() string {
	return fmt.Sprintf("(%d, %d)", l.X, l.Y)
}

func manhattenDistance(a, b location) int {
	return Abs(a.X-b.X) + Abs(a.Y-b.Y)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
