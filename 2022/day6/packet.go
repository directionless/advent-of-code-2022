package day6

import "fmt"

// packet is tracks the last N byte (which make up a packet) and checks them for uniquenss.
// It's a bit akin to a ring buffer, and perhaps ought have been a layer over one. But this
// was quick. There are, no doubt, optimizations to be had.
type packet struct {
	sz   int
	idx  int
	data []byte
}

func NewPacket(sz int) *packet {
	return &packet{
		sz:   sz,
		data: make([]byte, sz),
	}
}

func (p *packet) Push(b byte) {
	p.data[p.idx%p.sz] = b
	p.idx++
}

func (p *packet) String() string {
	// Because we're using a ring, converting to a string is going to be rotated.
	return fmt.Sprintf("(rotated)%s", p.data)
}

func (p *packet) Uniq() bool {
	if len(p.data) < p.sz {
		return false
	}

	// This implementation has an n+1 -- ever call traverses the whole array. There should be an
	// optimization where we persist uniq, and only clear it on a collision. Then we would need
	// to re-traverse. That should reduce the number of array traversals. Not really worth it
	// for advent though

	uniq := make(map[byte]bool)

	for _, b := range p.data {
		// Uninitialized value, means not populated, which means false for the purposes of Uniq
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
