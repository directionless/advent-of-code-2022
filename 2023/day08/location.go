package day08

import (
	"bytes"
	"fmt"
)

type locationType struct {
	Name [3]byte
	R    [3]byte
	L    [3]byte
	zzz  bool
}

// AAA = (BBB, CCC)
// BBB = (DDD, EEE)
func locationFromLine(line []byte) (locationType, error) {
	loc := locationType{}
	fields := bytes.Split(line, []byte{' '})
	if len(fields) != 4 {
		return loc, fmt.Errorf(`line "%s" had wrong number of fields`, line)
	}

	loc.Name = byteArrTo3byte(fields[0])
	loc.L = byteArrTo3byte(bytes.TrimLeft(fields[2], "(,"))
	loc.R = byteArrTo3byte(bytes.TrimRight(fields[3], ")"))
	loc.zzz = loc.Name == [3]byte{'Z', 'Z', 'Z'}

	//fmt.Printf("%s -> %s ;;; %q\n", line, loc.DebugString(), fields)
	return loc, nil
}

func (loc locationType) ZZZ() bool {
	return loc.zzz
}

func (loc locationType) DebugString() string {
	return fmt.Sprintf("%s(L:%s R:%s)", loc.Name, loc.L, loc.R)
}

func byteArrTo3byte(in []byte) [3]byte {
	out := [3]byte{}
	for i := 0; i < 3; i++ {
		out[i] = in[i]
	}
	return out
}
