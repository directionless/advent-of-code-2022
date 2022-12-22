package day21

import "fmt"

type monkeyMap map[string]monkeyRun

type monkeyRun interface {
	Run(monkeyMap) (int, error)
}

type monkeyConstant struct {
	val int
}

func (m *monkeyConstant) Run(_ monkeyMap) (int, error) {
	return m.val, nil
}

type monkeyAdd struct {
	a, b string
}

func (m *monkeyAdd) Run(mm monkeyMap) (int, error) {
	if a, ok := mm[m.a]; ok {
		if b, ok := mm[m.b]; ok {
			aVal, err := a.Run(mm)
			if err != nil {
				return 0, err
			}
			bVal, err := b.Run(mm)
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

func (m *monkeySub) Run(mm monkeyMap) (int, error) {
	if a, ok := mm[m.a]; ok {
		if b, ok := mm[m.b]; ok {
			aVal, err := a.Run(mm)
			if err != nil {
				return 0, err
			}
			bVal, err := b.Run(mm)
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

func (m *monkeyMul) Run(mm monkeyMap) (int, error) {
	if a, ok := mm[m.a]; ok {
		if b, ok := mm[m.b]; ok {
			aVal, err := a.Run(mm)
			if err != nil {
				return 0, err
			}
			bVal, err := b.Run(mm)
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

func (m *monkeyDiv) Run(mm monkeyMap) (int, error) {
	if a, ok := mm[m.a]; ok {
		if b, ok := mm[m.b]; ok {
			aVal, err := a.Run(mm)
			if err != nil {
				return 0, err
			}
			bVal, err := b.Run(mm)
			if err != nil {
				return 0, err
			}
			return aVal / bVal, nil
		}
		return 0, fmt.Errorf("missing %s", m.b)
	}
	return 0, fmt.Errorf("missing %s", m.a)
}
