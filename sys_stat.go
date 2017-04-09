// Copyright 2017 The Virtual Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package virtual

func init() {
	registerBuiltins(map[int]Opcode{
		dict.SID("fchmod"): fchmod,
		dict.SID("fstat"):  fstat,
		dict.SID("mkdir"):  mkdir,
		dict.SID("stat"):   stat,
	})
}