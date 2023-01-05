/*
Реализовать паттерн «адаптер» на любом примере.
*/

package main

import (
	"fmt"
	"time"
)

// ---------- random generator ----------

type Lcg struct {
	val uint32
}

func (l *Lcg) GetUInt32() uint32 {
	l.val = (1103515245*l.val + 12345) % 2147483648
	return l.val
}

// --------- wanted interface ---------

type Randomizer interface {
	Byte() uint8
}

// ---- adapt to wanted interface ----

type LcgRandomAdapter struct {
	gen Lcg
}

func (l *LcgRandomAdapter) Byte() uint8 {
	return uint8(l.gen.GetUInt32())
}

// -------------------------------------

func main() {
	seed := uint32(time.Now().UnixNano())

	var r Randomizer = &LcgRandomAdapter{
		gen: Lcg{seed},
	}

	fmt.Println(r.Byte())
}
