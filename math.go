// Copyright 2017 The Virtual Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package virtual

import (
	"math"
)

func init() {
	registerBuiltins(map[int]Opcode{
		dict.SID("__builtin_acos"):  acos,
		dict.SID("__builtin_asin"):  asin,
		dict.SID("__builtin_atan"):  atan,
		dict.SID("__builtin_ceil"):  ceil,
		dict.SID("__builtin_cos"):   cos,
		dict.SID("__builtin_cosh"):  cosh,
		dict.SID("__builtin_exp"):   exp,
		dict.SID("__builtin_fabs"):  fabs,
		dict.SID("__builtin_floor"): floor,
		dict.SID("__builtin_log"):   log,
		dict.SID("__builtin_log10"): log10,
		dict.SID("__builtin_pow"):   pow,
		dict.SID("__builtin_round"): round,
		dict.SID("__builtin_sin"):   sin,
		dict.SID("__builtin_sinh"):  sinh,
		dict.SID("__builtin_sqrt"):  sqrt,
		dict.SID("__builtin_tan"):   tan,
		dict.SID("__builtin_tanh"):  tanh,
	})
}

func (c *cpu) acos()  { writeF64(c.rp, math.Acos(readF64(c.sp))) }
func (c *cpu) asin()  { writeF64(c.rp, math.Asin(readF64(c.sp))) }
func (c *cpu) atan()  { writeF64(c.rp, math.Atan(readF64(c.sp))) }
func (c *cpu) ceil()  { writeF64(c.rp, math.Ceil(readF64(c.sp))) }
func (c *cpu) cos()   { writeF64(c.rp, math.Cos(readF64(c.sp))) }
func (c *cpu) cosh()  { writeF64(c.rp, math.Cosh(readF64(c.sp))) }
func (c *cpu) exp()   { writeF64(c.rp, math.Exp(readF64(c.sp))) }
func (c *cpu) fabs()  { writeF64(c.rp, math.Abs(readF64(c.sp))) }
func (c *cpu) floor() { writeF64(c.rp, math.Floor(readF64(c.sp))) }
func (c *cpu) log()   { writeF64(c.rp, math.Log(readF64(c.sp))) }
func (c *cpu) log10() { writeF64(c.rp, math.Log10(readF64(c.sp))) }
func (c *cpu) pow()   { writeF64(c.rp, math.Pow(readF64(c.sp+f64StackSz), readF64(c.sp))) }
func (c *cpu) sin()   { writeF64(c.rp, math.Sin(readF64(c.sp))) }
func (c *cpu) sinh()  { writeF64(c.rp, math.Sinh(readF64(c.sp))) }
func (c *cpu) sqrt()  { writeF64(c.rp, math.Sqrt(readF64(c.sp))) }
func (c *cpu) tan()   { writeF64(c.rp, math.Tan(readF64(c.sp))) }
func (c *cpu) tanh()  { writeF64(c.rp, math.Tanh(readF64(c.sp))) }

func (c *cpu) round() {
	v := readF64(c.sp)
	switch {
	case v < 0:
		v = math.Ceil(v - 0.5)
	case v > 0:
		v = math.Floor(v + 0.5)
	}
	writeF64(c.rp, v)
}
