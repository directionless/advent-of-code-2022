package day11

import (
	"fmt"
	"sort"
	"strings"
)

type Item struct {
	Worry   uint64
	NextHop int
}

func (i Item) String() string {
	return fmt.Sprintf("%d", i.Worry)
}

func NewItem(w int) *Item {
	return &Item{
		Worry:   uint64(w),
		NextHop: -1,
	}
}

type itemTestFunc func(*Item) bool

func genTestDivisible(n int) itemTestFunc {
	return func(i *Item) bool {
		return i.Worry%uint64(n) == 0
	}
}

type itemManipulatorFunc func(*Item)

func genOperateAdd(n int) itemManipulatorFunc {
	return func(i *Item) {
		i.Worry += uint64(n)
	}
}

func genOperateMultiply(n int) itemManipulatorFunc {
	return func(i *Item) {
		i.Worry = i.Worry * uint64(n)
	}
}

func genOperateSquare() itemManipulatorFunc {
	return func(i *Item) {
		i.Worry = i.Worry * i.Worry
	}
}

type Monkey struct {
	Pos          int
	InspectCount int
	Items        []*Item
	Inspect      itemManipulatorFunc
	TestFn       itemTestFunc
	TrueDest     int
	FalseDest    int
	noRelief     bool
}

func NewMonkey(pos int, noRelief bool) *Monkey {
	m := &Monkey{
		Pos:      pos,
		noRelief: noRelief,
	}

	return m
}

func (m *Monkey) String() string {
	var sb strings.Builder

	fmt.Fprintf(&sb, "Monkey %d: ", m.Pos)
	for _, i := range m.Items {
		fmt.Fprintf(&sb, "%d ", i.Worry)
	}
	fmt.Fprintf(&sb, "\n")

	return sb.String()
}

func (m *Monkey) Push(item *Item) {
	item.NextHop = -1
	m.Items = append(m.Items, item)
}

func (m *Monkey) Pop() *Item {
	if len(m.Items) == 0 {
		return nil
	}

	m.InspectCount++

	item := m.Items[0]
	m.Items = m.Items[1:]

	// After each monkey inspects an item but before it tests your worry level,
	// your relief that the monkey's inspection didn't damage the item causes
	// your worry level to be divided by three and rounded down to the nearest integer.
	m.Inspect(item)
	if !m.noRelief {
		item.Worry = item.Worry / 3
	}
	//fmt.Printf("item worry: %d\n", item.Worry)

	//fmt.Printf("Monkey %d is inspecting item %d (%d remaining)\n", m.Pos, item.Worry, len(m.Items))

	if m.TestFn(item) {
		item.NextHop = m.TrueDest
	} else {
		item.NextHop = m.FalseDest
	}

	return item
}

type monkeyNetwork struct {
	round             int
	monkeys           []*Monkey
	commonWorryFactor int
}

// RunRound runs the monkey network for a round. This is implemented with logic
// down on the monkeys, and operates one monkey at time. Having the logic down in the
// monkeys feels a bit slower, and in some ways, more cumbersome, But it's a bit easier
// to handle the case where monkies might route to themselves. It also allows the monkeys
// to operate without knlowdge of the larger monkey network.
//
// Some places to improve efficiency:
//   - Instead of pop/push one item at a time, we could have the monkey run it's
//     entire stack. The trick with that approach is that if a monkey routes to themself
//     we'd need to be a bit more clever with the slices and the terminal condition
func (mn *monkeyNetwork) RunRound() error {
	for _, monkey := range mn.monkeys {
		for {
			item := monkey.Pop()
			if item == nil {
				break
			}

			if item.NextHop == -1 {
				return fmt.Errorf("item missing routing")
			}

			if item.NextHop >= len(mn.monkeys) {
				return fmt.Errorf("invalid NextHop")
			}

			// Since our worry level can get _very_ high, we need to reduce the value
			// to something we can calculate. big.Int is one approach, but I've opted to
			// whomp things by the common factor. After all -- the worry level only matters
			// for how the monkeys route, and the the common factor is a no-op. (And I'm
			// I'm slightly concerned we'll overflow big.Int, or be non-performant or something)
			//
			// Not sure this works when we square things. Hrm.....
			if mn.commonWorryFactor != 1 && item.Worry > uint64(mn.commonWorryFactor) {
				remain := item.Worry % uint64(mn.commonWorryFactor)

				//fmt.Printf("worry modulo: %d %% %d = %d\n", item.Worry, mn.commonWorryFactor, remain)
				item.Worry = remain
			}

			mn.monkeys[item.NextHop].Push(item)
		}
	}
	return nil
}

func NewMonkeyNetwork() *monkeyNetwork {
	return &monkeyNetwork{
		round:             0,
		monkeys:           []*Monkey{},
		commonWorryFactor: 1,
	}
}

func (mn *monkeyNetwork) AddMonkey(m *Monkey) {
	mn.monkeys = append(mn.monkeys, m)
}

func (mn *monkeyNetwork) Monkey(n int) *Monkey {
	return mn.monkeys[n]
}

func (mn *monkeyNetwork) PrintInfo() {
	for _, m := range mn.monkeys {
		fmt.Printf("%s", m)
		// fmt.Printf("inspected items %d times.\n" m.InspectCount)
	}
}

func (mn *monkeyNetwork) MonkeyBusiness() int {
	mosts := make([]int, len(mn.monkeys))
	for i, m := range mn.monkeys {
		mosts[i] = m.InspectCount
	}

	sort.Ints(mosts)

	last := len(mosts) - 1
	return mosts[last] * mosts[last-1]

}
