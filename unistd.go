// Copyright 2017 The Virtual Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package virtual

import (
	"syscall"
)

func init() {
	registerBuiltins(map[int]Opcode{
		dict.SID("access"):    access,
		dict.SID("close"):     close_,
		dict.SID("fchown"):    fchown,
		dict.SID("fsync"):     fsync,
		dict.SID("ftruncate"): ftruncate,
		dict.SID("getcwd"):    getcwd,
		dict.SID("geteuid"):   geteuid,
		dict.SID("getpid"):    getpid,
		dict.SID("lseek"):     lseek,
		dict.SID("read"):      read,
		dict.SID("readlink"):  readlink,
		dict.SID("rmdir"):     rmdir,
		dict.SID("sleep"):     sleep,
		dict.SID("sysconf"):   sysconf,
		dict.SID("unlink"):    unlink,
		dict.SID("write"):     write,
	})
}

// int close(int fd);
func (c *cpu) close() {
	fd := readI32(c.sp)
	r, _, err := syscall.Syscall(syscall.SYS_CLOSE, uintptr(fd), 0, 0)
	if err != 0 {
		c.setErrno(err)
	}
	writeI32(c.rp, int32(r))
}

// int fsync(int fildes);
func (c *cpu) fsync() {
	fildes := readI32(c.sp)
	r, _, err := syscall.Syscall(syscall.SYS_FSYNC, uintptr(fildes), 0, 0)
	if err != 0 {
		c.setErrno(err)
	}
	writeI32(c.rp, int32(r))
}

// char *getcwd(char *buf, size_t size);
func (c *cpu) getcwd() {
	sp, size := popLong(c.sp)
	buf := readPtr(sp)
	r, _, err := syscall.Syscall(syscall.SYS_GETCWD, buf, uintptr(size), 0)
	if err != 0 {
		c.setErrno(err)
	}
	writePtr(c.rp, r)
}

// uid_t geteuid(void);
func (c *cpu) geteuid() {
	r, _, _ := syscall.RawSyscall(syscall.SYS_GETEUID, 0, 0, 0)
	writeU32(c.rp, uint32(r))
}

// pid_t getpid(void);
func (c *cpu) getpid() {
	r, _, _ := syscall.RawSyscall(syscall.SYS_GETPID, 0, 0, 0)
	writeI32(c.rp, int32(r))
}

// off_t lseek(int fildes, off_t offset, int whence);
func (c *cpu) lseek() {
	sp, whence := popI32(c.sp)
	sp, offset := popLong(sp)
	fildes := readI32(sp)
	r, _, err := syscall.Syscall(syscall.SYS_LSEEK, uintptr(fildes), uintptr(offset), uintptr(whence))
	if err != 0 {
		c.setErrno(err)
	}
	writeLong(c.rp, int64(r))
}

// ssize_t read(int fd, void *buf, size_t count);
func (c *cpu) read() {
	sp, count := popLong(c.sp)
	sp, buf := popPtr(sp)
	fd := readI32(sp)
	r, _, err := syscall.Syscall(syscall.SYS_READ, uintptr(fd), buf, uintptr(count))
	if err != 0 {
		c.thread.setErrno(err)
	}
	writeLong(c.rp, int64(r))
}

// int unlink(const char *path);
func (c *cpu) unlink() {
	path := readPtr(c.sp)
	r, _, err := syscall.Syscall(syscall.SYS_UNLINK, path, 0, 0)
	if err != 0 {
		c.setErrno(err)
	}
	writeI32(c.rp, int32(r))
}

// ssize_t write(int fd, const void *buf, size_t count);
func (c *cpu) write() {
	sp, count := popLong(c.sp)
	sp, buf := popPtr(sp)
	fd := readI32(sp)
	r, _, err := syscall.Syscall(syscall.SYS_WRITE, uintptr(fd), buf, uintptr(count))
	if err != 0 {
		c.thread.setErrno(err)
	}
	writeLong(c.rp, int64(r))
}
