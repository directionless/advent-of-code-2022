package cpu

type CPU struct {
	cycle int
	X     int
	taps  []tapInterface
}

func New() *CPU {
	return &CPU{
		cycle: 0,
		X:     1,
	}
}

func (c *CPU) tick() {
	c.cycle++
	c.RunTaps()
}

func (c *CPU) AddTap(tap tapInterface) {
	c.taps = append(c.taps, tap)
}

func (c *CPU) RunTaps() {
	for _, tap := range c.taps {
		tap.Examine(c.cycle, c.X)
	}
}

// noop takes one cycle to complete. It has no other effect.
func (c *CPU) ExecNoop() {
	c.tick()
}

// addx V takes two cycles to complete. After two cycles, the X register is increased by the value V. (V can be negative.)
func (c *CPU) ExecAddx(x int) {
	c.tick()
	c.tick()
	c.X += x
}

type tapInterface interface {
	Examine(cycle int, x int)
}
