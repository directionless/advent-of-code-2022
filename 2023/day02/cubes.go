package day02

import (
	"fmt"
	"strconv"
	"strings"
)

type pixelTyp struct {
	red   int
	green int
	blue  int
}

func (c pixelTyp) Red() int   { return c.red }
func (c pixelTyp) Green() int { return c.green }
func (c pixelTyp) Blue() int  { return c.blue }

// PixelFromString returns set of cubes from a string. Strings should be in the form
// of "1 green, 3 red, 6 blue" note that not all colors are present, and may be in
// differnt orders.
func PixelFromString(s string) (pixelTyp, error) {
	pixel := pixelTyp{}

	//fmt.Println("<<", s, ">>")

	for _, chunk := range strings.Split(s, ",") {
		chunk = strings.Trim(chunk, " ")
		fields := strings.Fields(chunk)
		if len(fields) > 2 {
			return pixel, fmt.Errorf(`too many words in "%s"`, chunk)
		}

		if len(fields) < 2 {
			return pixel, fmt.Errorf(`too few words in "%s"`, chunk)
		}

		num, err := strconv.Atoi(fields[0])
		if err != nil {
			return pixel, fmt.Errorf(`unable to parse number from "%s": %w`, fields[0], err)
		}

		switch fields[1] {
		case "red":
			pixel.red = num
		case "green":
			pixel.green = num
		case "blue":
			pixel.blue = num
		default:
			return pixel, fmt.Errorf(`unknown color "%s"`, fields[1])
		}
	}

	return pixel, nil
}

func (outer pixelTyp) Contains(p pixelTyp) bool {
	if outer.red >= p.Red() && outer.green >= p.Green() && outer.blue >= p.Blue() {
		return true
	}
	return false
}

func (p pixelTyp) Power() int {
	return p.red * p.green * p.blue
}
