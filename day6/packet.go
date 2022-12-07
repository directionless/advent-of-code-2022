package day6

import "fmt"

// packet is tracks the last N byte (which make up a packet) and checks them for uniquenss.
// It's a bit akin to a ring buffer, and perhaps ought have been a layer over one. But this
// was quick. There are, no doubt, optimizations to be had.
type packet struct {
	sz   int
	data []byte
}

func NewPacket(sz int) *packet {
	return &packet{
		sz:   sz,
		data: make([]byte, sz),
	}
}

func (p *packet) Push(b byte) {
	// Instead of the array shift, it would be interesting to track where in the index we are,
	// and overwrite the character directly.
	if len(p.data) >= p.sz {
		p.data = p.data[1:]
	}

	p.data = append(p.data, b)
}

func (p *packet) String() string {
	return fmt.Sprintf("(%d/%d)%s", p.sz, len(p.data), p.data)
}

func (p *packet) Uniq() bool {
	if len(p.data) < p.sz {
		return false
	}

	uniq := make(map[byte]bool)

	for _, b := range p.data {
		// Uninitialized value, means not populated, means false
		if b == 0 {
			return false
		}

		if uniq[b] {
			return false
		}
		uniq[b] = true
	}

	return true
}
