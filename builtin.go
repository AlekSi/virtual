// Copyright 2017 The Virtual Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package virtual

import (
	"github.com/cznic/mathutil"
)

func init() {
	registerBuiltins(map[int]Opcode{
		dict.SID("__builtin_clrsb"):      clrsb,
		dict.SID("__builtin_clrsbl"):     clrsbl,
		dict.SID("__builtin_clrsbll"):    clrsbll,
		dict.SID("__builtin_clz"):        clz,
		dict.SID("__builtin_clzl"):       clzl,
		dict.SID("__builtin_clzll"):      clzll,
		dict.SID("__builtin_ctz"):        ctz,
		dict.SID("__builtin_ctzl"):       ctzl,
		dict.SID("__builtin_ctzll"):      ctzll,
		dict.SID("__builtin_parity"):     parity,
		dict.SID("__builtin_parityl"):    parityl,
		dict.SID("__builtin_parityll"):   parityll,
		dict.SID("__builtin_popcount"):   popcount,
		dict.SID("__builtin_popcountl"):  popcountl,
		dict.SID("__builtin_popcountll"): popcountll,
	})
}

// int __builtin_clrsb (int x);
func (c *cpu) clrsb() {
	x := readI32(c.rp - i32StackSz)
	i := int32(1)
	n := x >> 31 & 1
	for ; i < 32 && x>>(31-uint(i))&1 == n; i++ {
	}
	writeI32(c.rp, i-1)
}

// int __builtin_clrsbl (long x);
func (c *cpu) clrsbl() {
	x := readLong(c.rp - stackAlign)
	i := int32(1)
	n := x >> 63 & 1
	for ; i < 64 && x>>(63-uint(i))&1 == n; i++ {
	}
	writeI32(c.rp, i-1)
}

// int __builtin_clrsbll (long long x);
func (c *cpu) clrsbll() {
	x := readI64(c.rp - i64StackSz)
	i := int32(1)
	n := x >> 63 & 1
	for ; i < 64 && x>>(63-uint(i))&1 == n; i++ {
	}
	writeI32(c.rp, i-1)
}

// int __builtin_clz (unsigned x);
func (c *cpu) clz() {
	x := readU32(c.rp - i32StackSz)
	var i int32
	for ; i < 32 && x&(1<<uint(31-i)) == 0; i++ {
	}
	writeI32(c.rp, i)
}

// int __builtin_clzl (unsigned long x);
func (c *cpu) clzl() {
	x := uint64(readLong(c.rp - stackAlign))
	var i int32
	for ; i < 64 && x&(1<<uint(63-i)) == 0; i++ {
	}
	writeI32(c.rp, i)
}

// int __builtin_clzll (unsigned long long x);
func (c *cpu) clzll() {
	x := readU64(c.rp - i64StackSz)
	var i int32
	for ; i < 64 && x&(1<<uint(63-i)) == 0; i++ {
	}
	writeI32(c.rp, i)
}

// int __builtin_ctz (unsigned x);
func (c *cpu) ctz() {
	x := readU32(c.rp - i32StackSz)
	var i int32
	for ; i < 32 && x&(1<<uint(i)) == 0; i++ {
	}
	writeI32(c.rp, i)
}

// int __builtin_ctzl (unsigned long x);
func (c *cpu) ctzl() {
	x := uint64(readLong(c.rp - stackAlign))
	var i int32
	for ; i < 64 && x&(1<<uint(i)) == 0; i++ {
	}
	writeI32(c.rp, i)
}

// int __builtin_ctzll (unsigned long long x);
func (c *cpu) ctzll() {
	x := readU64(c.rp - i32StackSz)
	var i int32
	for ; i < 64 && x&(1<<uint(i)) == 0; i++ {
	}
	writeI32(c.rp, i)
}

// int __builtin_parity(unsigned x);
func (c *cpu) parity() { writeI32(c.rp, int32(mathutil.PopCountUint32(readU32(c.rp-i32StackSz)))&1) }

// int __builtin_parityl(unsigned long x);
func (c *cpu) parityl() {
	writeI32(c.rp, int32(mathutil.PopCountUint64(uint64(readLong(c.rp-stackAlign))))&1)
}

// int __builtin_parityll(unsigned long long x);
func (c *cpu) parityll() { writeI32(c.rp, int32(mathutil.PopCountUint64(readU64(c.rp-i64StackSz)))&1) }

// int __builtin_popcount(unsigned x);
func (c *cpu) popcount() { writeI32(c.rp, int32(mathutil.PopCountUint32(readU32(c.rp-i32StackSz)))) }

// int __builtin_popcountl(unsigned long x);
func (c *cpu) popcountl() {
	writeI32(c.rp, int32(mathutil.PopCountUint64(uint64(readLong(c.rp-stackAlign)))))
}

// int __builtin_popcountll(unsigned long long x);
func (c *cpu) popcountll() { writeI32(c.rp, int32(mathutil.PopCountUint64(readU64(c.rp-i64StackSz)))) }
