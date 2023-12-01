package day21

import (
	"errors"
	"fmt"
)

const maxDepth = 5000

type monkeyTyp struct {
	runs []monkeyRun
}

func (m *monkeyTyp) AddRun(run monkeyRun) {
	if m.runs == nil {
		m.runs = make([]monkeyRun, 0)
	}
	m.runs = append(m.runs, run)
}

// ConvertRoot will convert the root monkey to the part2 def
func (m *monkeyTyp) ExtractRoot() (string, string, error) {
	if len(m.runs) != 1 {
		return "", "", errors.New("root monkey should have exactly 1 run")
	}

	ma, ok := m.runs[0].(*monkeyAdd)
	if !ok {
		return "", "", errors.New("root monkey failed run type conversion")
	}

	return ma.A(), ma.B(), nil
}

func (m *monkeyTyp) GetHuman() (int, error) {
	if len(m.runs) != 1 {
		return 0, errors.New("human should have exactly 1 run")
	}

	mc, ok := m.runs[0].(*monkeyConstant)
	if !ok {
		return 0, errors.New("human failed run type conversion")
	}

	return mc.C(), nil
}

func (m *monkeyTyp) OverrideHuman(num int) error {
	if len(m.runs) != 1 {
		return errors.New("human should have exactly 1 run")
	}

	_, ok := m.runs[0].(*monkeyConstant)
	if !ok {
		return errors.New("human failed run type conversion")
	}

	m.runs[0] = &monkeyConstant{num}
	return nil
}

// Run is basically a DFS on monkey, which contains other monkeys. Unfortunaetly,
// it doesn't work. There are loops in this graph. And this wasn't written with
// that in mind.
func (m *monkeyTyp) Run(depth int, mm monkeyMap) (int, error) {
	if len(m.runs) == 0 {
		return 0, fmt.Errorf("no runs in this monkey")
	}

	if depth > maxDepth {
		return 0, errors.New("too deep")
	}

	for _, run := range m.runs {
		result, err := run.Run(depth+1, mm)
		if err == nil {
			return result, nil
		}
	}

	return 0, fmt.Errorf("errors from all runs")
}

type monkeyMap map[string]monkeyRun

type monkeyRun interface {
	Run(depth int, mm monkeyMap) (int, error)
}

type monkeyConstant struct {
	val int
}

func (m *monkeyConstant) Run(_ int, _ monkeyMap) (int, error) {
	return m.val, nil
}

func (m *monkeyConstant) C() int {
	return m.val
}

type monkeyAdd struct {
	a, b string
}

func (m *monkeyAdd) A() string {
	return m.a
}

func (m *monkeyAdd) B() string {
	return m.b
}

func (m *monkeyAdd) Run(depth int, mm monkeyMap) (int, error) {
	if depth > maxDepth {
		return 0, errors.New("too deep")
	}

	if a, ok := mm[m.a]; ok {
		if b, ok := mm[m.b]; ok {
			aVal, err := a.Run(depth+1, mm)
			if err != nil {
				return 0, err
			}
			bVal, err := b.Run(depth+1, mm)
			if err != nil {
				return 0, err
			}
			return aVal + bVal, nil
		}
		return 0, fmt.Errorf("missing %s", m.b)
	}
	return 0, fmt.Errorf("missing %s", m.a)
}

type monkeySub struct {
	a, b string
}

func (m *monkeySub) Run(depth int, mm monkeyMap) (int, error) {
	if depth > maxDepth {
		return 0, errors.New("too deep")
	}

	if a, ok := mm[m.a]; ok {
		if b, ok := mm[m.b]; ok {
			aVal, err := a.Run(depth+1, mm)
			if err != nil {
				return 0, err
			}
			bVal, err := b.Run(depth+1, mm)
			if err != nil {
				return 0, err
			}
			return aVal - bVal, nil
		}
		return 0, fmt.Errorf("missing %s", m.b)
	}
	return 0, fmt.Errorf("missing %s", m.a)
}

type monkeyMul struct {
	a, b string
}

func (m *monkeyMul) Run(depth int, mm monkeyMap) (int, error) {
	if depth > maxDepth {
		return 0, errors.New("too deep")
	}

	if a, ok := mm[m.a]; ok {
		if b, ok := mm[m.b]; ok {
			aVal, err := a.Run(depth+1, mm)
			if err != nil {
				return 0, err
			}
			bVal, err := b.Run(depth+1, mm)
			if err != nil {
				return 0, err
			}
			return aVal * bVal, nil
		}
		return 0, fmt.Errorf("missing %s", m.b)
	}
	return 0, fmt.Errorf("missing %s", m.a)
}

type monkeyDiv struct {
	a, b string
}

func (m *monkeyDiv) Run(depth int, mm monkeyMap) (int, error) {
	if depth > maxDepth {
		return 0, errors.New("too deep")
	}

	if a, ok := mm[m.a]; ok {
		if b, ok := mm[m.b]; ok {
			aVal, err := a.Run(depth+1, mm)
			if err != nil {
				return 0, err
			}
			bVal, err := b.Run(depth+1, mm)
			if err != nil {
				return 0, err
			}
			return aVal / bVal, nil
		}
		return 0, fmt.Errorf("missing %s", m.b)
	}
	return 0, fmt.Errorf("missing %s", m.a)
}
