package day9

type videoMaker struct {
	min point
	max point

	moves []map[int]point
}

func (vm *videoMaker) AddMove(pos map[int]*point) {
	// the head is going to move the most, so we only need to check the head
	// against the min/max positions
	head := pos[0]
	if head.Row < vm.min.Row {
		vm.min.Row = head.Row
	}
	if head.Row > vm.max.Row {
		vm.max.Row = head.Row
	}
	if head.Col < vm.min.Col {
		vm.min.Col = head.Col
	}
	if head.Col > vm.max.Col {
		vm.max.Col = head.Col
	}

	if vm.moves == nil {
		vm.moves = make([]map[int]point, 0)
	}

	// The move is a series of pointers, and we need to resolve the pointers,
	// lest they change out from under us.
	move := make(map[int]point, len(pos))
	for i, p := range pos {
		move[i] = point{p.Row, p.Col}
	}
	vm.moves = append(vm.moves, move)
}
