// Copyright 2017 The Virtual Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package virtual

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"sync"

	"syscall"
	"unsafe"

	"github.com/cznic/internal/buffer"
	"github.com/cznic/mathutil"
)

func init() {
	registerBuiltins(map[int]Opcode{
		dict.SID("fclose"):  fclose,
		dict.SID("fgetc"):   fgetc,
		dict.SID("fgets"):   fgets,
		dict.SID("fopen"):   fopen,
		dict.SID("fprintf"): fprintf,
		dict.SID("fread"):   fread,
		dict.SID("fwrite"):  fwrite,
		dict.SID("printf"):  printf,
		dict.SID("sprintf"): sprintf,
	})
}

const eof = -1

var (
	files      = &fmap{m: map[uintptr]*os.File{}}
	nullReader = bytes.NewBuffer(nil)
)

type fmap struct {
	m  map[uintptr]*os.File
	mu sync.Mutex
}

func (m *fmap) add(f *os.File, u uintptr) {
	m.mu.Lock()
	m.m[u] = f
	m.mu.Unlock()
}

func (m *fmap) reader(u uintptr, c *cpu) io.Reader {
	m.mu.Lock()
	f := m.m[u]
	m.mu.Unlock()
	switch {
	case f == os.Stdin:
		return c.m.stdin
	case f == os.Stdout:
		return nullReader
	case f == os.Stderr:
		return nullReader
	}
	return f
}

func (m *fmap) writer(u uintptr, c *cpu) io.Writer {
	m.mu.Lock()
	f := m.m[u]
	m.mu.Unlock()
	switch {
	case f == os.Stdin:
		return ioutil.Discard
	case f == os.Stdout:
		return c.m.stdout
	case f == os.Stderr:
		return c.m.stderr
	}
	return f
}

func (m *fmap) extract(u uintptr) *os.File {
	m.mu.Lock()
	f := m.m[u]
	delete(m.m, u)
	m.mu.Unlock()
	return f
}

type file struct{ _ int32 }

// int fclose(FILE *stream);
func (c *cpu) fclose() {
	u := readPtr(c.rp - ptrStackSz)
	f := files.extract(readPtr(u))
	if f == nil {
		c.thread.errno = int32(syscall.EBADF)
		writeI32(c.rp, eof)
		return
	}

	c.m.free(u)
	if err := f.Close(); err != nil {
		c.thread.errno = int32(syscall.EIO)
		writeI32(c.rp, eof)
		return
	}

	writeI32(c.rp, 0)
}

// int fgetc(FILE *stream);
func (c *cpu) fgetc() {
	p := buffer.Get(1)
	if _, err := files.reader(readPtr(c.rp-ptrStackSz), c).Read(*p); err != nil {
		writeI32(c.rp, eof)
		buffer.Put(p)
		return
	}

	writeI32(c.rp, int32((*p)[0]))
	buffer.Put(p)
}

// char *fgets(char *s, int size, FILE *stream);
func (c *cpu) fgets() {
	ap := c.rp - ptrStackSz
	s := readPtr(ap)
	ap -= i32StackSz
	size := int(readI32(ap))
	f := files.reader(readPtr(ap-ptrStackSz), c)
	p := buffer.Get(1)
	b := *p
	w := memWriter(s)
	ok := false
	for i := size - 1; i > 0; i-- {
		_, err := f.Read(b)
		if err != nil {
			if !ok {
				writePtr(c.rp, 0)
				buffer.Put(p)
				return
			}

			break
		}

		ok = true
		w.WriteByte(b[0])
		if b[0] == '\n' {
			break
		}
	}
	w.WriteByte(0)
	writePtr(c.rp, s)
	buffer.Put(p)

}

// FILE *fopen(const char *path, const char *mode);
func (c *cpu) fopen() {
	path := GoString(readPtr(c.rp - ptrStackSz))
	var f *os.File
	var err error
	switch path {
	case os.Stderr.Name():
		f = os.Stderr
	case os.Stdin.Name():
		f = os.Stdin
	case os.Stdout.Name():
		f = os.Stdout
	default:
		mode := GoString(readPtr(c.rp - 2*ptrStackSz))
		switch mode {
		case "r":
			if f, err = os.OpenFile(path, os.O_RDONLY, 0666); err != nil {
				switch {
				case os.IsNotExist(err):
					c.thread.errno = int32(syscall.ENOENT)
				case os.IsPermission(err):
					c.thread.errno = int32(syscall.EPERM)
				default:
					c.thread.errno = int32(syscall.EACCES)
				}
				writePtr(c.rp, 0)
				return
			}
		case "w":
			if f, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666); err != nil {
				switch {
				case os.IsPermission(err):
					c.thread.errno = int32(syscall.EPERM)
				default:
					c.thread.errno = int32(syscall.EACCES)
				}
				writePtr(c.rp, 0)
				return
			}
		default:
			panic(mode)
		}
	}

	u := c.m.malloc(int(unsafe.Sizeof(file{})))
	files.add(f, u)
	writePtr(c.rp, u)
}

// int fprintf(FILE * stream, const char *format, ...);
func (c *cpu) fprintf() {
	ap := c.rp - ptrStackSz
	stream := readPtr(ap)
	ap -= ptrStackSz
	writeI32(c.rp, goFprintf(files.writer(stream, c), readPtr(ap), ap))
}

// size_t fread(void *ptr, size_t size, size_t nmemb, FILE *stream);
func (c *cpu) fread() {
	ap := c.rp - ptrStackSz
	ptr := readPtr(ap)
	ap -= stackAlign
	size := readSize(ap)
	ap -= stackAlign
	nmemb := readSize(ap)
	ap -= ptrStackSz
	hi, lo := mathutil.MulUint128_64(size, nmemb)
	if hi != 0 || lo > math.MaxInt32 {
		c.thread.errno = int32(syscall.E2BIG)
		writeSize(c.rp, 0)
		return
	}

	n, err := files.reader(readPtr(ap), c).Read((*[math.MaxInt32]byte)(unsafe.Pointer(ptr))[:lo])
	if err != nil {
		c.thread.errno = int32(syscall.EIO)
	}
	writeSize(c.rp, uint64(n)/size)
}

// size_t fwrite(const void *ptr, size_t size, size_t nmemb, FILE *stream);
func (c *cpu) fwrite() {
	ap := c.rp - ptrStackSz
	ptr := readPtr(ap)
	ap -= stackAlign
	size := readSize(ap)
	ap -= stackAlign
	nmemb := readSize(ap)
	ap -= ptrStackSz
	hi, lo := mathutil.MulUint128_64(size, nmemb)
	if hi != 0 || lo > math.MaxInt32 {
		c.thread.errno = int32(syscall.E2BIG)
		writeSize(c.rp, 0)
		return
	}

	n, err := files.writer(readPtr(ap), c).Write((*[math.MaxInt32]byte)(unsafe.Pointer(ptr))[:lo])
	if err != nil {
		c.thread.errno = int32(syscall.EIO)
	}
	writeSize(c.rp, uint64(n)/size)
}

func goFprintf(w io.Writer, format, argp uintptr) int32 {
	var b buffer.Bytes
	written := 0
	for {
		ch := readI8(format)
		format++
		switch ch {
		case 0:
			_, err := b.WriteTo(w)
			b.Close()
			if err != nil {
				return -1
			}

			return int32(written)
		case '%':
			modifiers := ""
		more:
			ch := readI8(format)
			format++
			switch ch {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
				modifiers += string(ch)
				goto more
			case 'c':
				argp -= i32StackSz
				arg := readI32(argp)
				n, _ := fmt.Fprintf(&b, fmt.Sprintf("%%%sc", modifiers), arg)
				written += n
			case 'd', 'i':
				argp -= i32StackSz
				arg := readI32(argp)
				n, _ := fmt.Fprintf(&b, fmt.Sprintf("%%%sd", modifiers), arg)
				written += n
			case 'f':
				argp -= f64StackSz
				arg := readF64(argp)
				n, _ := fmt.Fprintf(&b, fmt.Sprintf("%%%sf", modifiers), arg)
				written += n
			case 's':
				argp -= ptrStackSz
				arg := readPtr(argp)
				if arg == 0 {
					break
				}

				var b2 buffer.Bytes
				for {
					c := readI8(arg)
					arg++
					if c == 0 {
						n, _ := fmt.Fprintf(&b, fmt.Sprintf("%%%ss", modifiers), b2.Bytes())
						b2.Close()
						written += n
						break
					}

					b2.WriteByte(byte(c))
				}
			default:
				panic(fmt.Errorf("TODO %q", "%"+string(ch)))
			}
		default:
			b.WriteByte(byte(ch))
			written++
			if ch == '\n' {
				if _, err := b.WriteTo(w); err != nil {
					b.Close()
					return -1
				}
				b.Reset()
			}
		}
	}
}

// int printf(const char *format, ...);
func (c *cpu) printf() {
	writeI32(c.rp, goFprintf(c.m.stdout, readPtr(c.rp-ptrStackSz), c.rp-ptrStackSz))
}

// int sprintf(char *str, const char *format, ...);
func (c *cpu) sprintf() {
	ap := c.rp - ptrStackSz
	w := memWriter(readPtr(ap))
	ap -= ptrStackSz
	writeI32(c.rp, goFprintf(&w, readPtr(ap), ap))
	writeI8(uintptr(w), 0)
}
