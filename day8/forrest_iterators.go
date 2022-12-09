package day8

// Designed to be passed to the look function to abstract a little
type heightBlockComparison func(a, b byte) bool

func sameHeightOkay(a, b byte) bool {
	return a >= b
}

func sameHeightBlocks(a, b byte) bool {
	return a > b
}

type looperStruct struct {
	Row           int
	Col           int
	End           int
	step          incrOrDecr
	constRowOrCol rowOrCol
}

type incrOrDecr string

const (
	incr incrOrDecr = "incr"
	decr incrOrDecr = "decr"
)

type rowOrCol string

const (
	row rowOrCol = "row"
	col rowOrCol = "col"
)

type rowOrColComparison string

func (l *looperStruct) Next() {
	switch {
	case l.constRowOrCol == "col" && l.step == incr:
		l.Row++
	case l.constRowOrCol == "col" && l.step == decr:
		l.Row--
	case l.constRowOrCol == "row" && l.step == incr:
		l.Col++
	case l.constRowOrCol == "row" && l.step == decr:
		l.Col--
	default:
		panic("How did you get here")
	}
}

func (l *looperStruct) Done() bool {
	cur := l.Row
	if l.constRowOrCol == "row" {
		cur = l.Col
	}

	if l.step == incr {
		return cur > l.End
	}

	return cur < l.End
}
