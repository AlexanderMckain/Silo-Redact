//go:build windows

package main

import (
	"runtime"
	"syscall"
	"unsafe"
)

const (
	ghostCFUnicodeText    = 13
	ghostGMEMMoveable     = 0x0002
	ghostMaxClipboardUTF16 = 1 << 20
)

var (
	ghostUser32   = syscall.NewLazyDLL("user32.dll")
	ghostKernel32 = syscall.NewLazyDLL("kernel32.dll")

	ghostOpenClipboard    = ghostUser32.NewProc("OpenClipboard")
	ghostCloseClipboard   = ghostUser32.NewProc("CloseClipboard")
	ghostEmptyClipboard   = ghostUser32.NewProc("EmptyClipboard")
	ghostGetClipboardData = ghostUser32.NewProc("GetClipboardData")
	ghostSetClipboardData = ghostUser32.NewProc("SetClipboardData")

	ghostGlobalAlloc   = ghostKernel32.NewProc("GlobalAlloc")
	ghostGlobalFree    = ghostKernel32.NewProc("GlobalFree")
	ghostGlobalSize    = ghostKernel32.NewProc("GlobalSize")
	ghostGlobalLock    = ghostKernel32.NewProc("GlobalLock")
	ghostGlobalUnlock  = ghostKernel32.NewProc("GlobalUnlock")
	ghostRtlMoveMemory = ghostKernel32.NewProc("RtlMoveMemory")
)

// ghostPollReadClipboard reads CF_UNICODETEXT with a single OpenClipboard attempt (no retry loop).
func ghostPollReadClipboard() (string, bool) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	ok, _, _ := ghostOpenClipboard.Call(0)
	if ok == 0 {
		return "", false
	}
	defer func() { _, _, _ = ghostCloseClipboard.Call() }()

	h, _, _ := ghostGetClipboardData.Call(ghostCFUnicodeText)
	if h == 0 {
		return "", false
	}

	l, _, _ := ghostGlobalLock.Call(h)
	if l == 0 {
		return "", false
	}
	defer func() { _, _, _ = ghostGlobalUnlock.Call(h) }()

	sz, _, _ := ghostGlobalSize.Call(h)
	if sz == 0 {
		return "", false
	}
	buf := make([]uint16, ghostMaxClipboardUTF16)
	maxBytes := uintptr(len(buf) * int(unsafe.Sizeof(uint16(0))))
	nBytes := uintptr(sz)
	if nBytes > maxBytes {
		nBytes = maxBytes
	}
	dest := uintptr(unsafe.Pointer(&buf[0]))
	_, _, _ = ghostRtlMoveMemory.Call(dest, l, nBytes)

	nLimit := int(nBytes / 2)
	n := 0
	for n < nLimit && buf[n] != 0 {
		n++
	}
	return syscall.UTF16ToString(buf[:n]), true
}

// ghostPollWriteClipboard replaces clipboard text with a single OpenClipboard attempt.
func ghostPollWriteClipboard(text string) bool {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	ok, _, _ := ghostOpenClipboard.Call(0)
	if ok == 0 {
		return false
	}
	defer func() { _, _, _ = ghostCloseClipboard.Call() }()

	_, _, _ = ghostEmptyClipboard.Call()

	utf16, err := syscall.UTF16FromString(text)
	if err != nil || len(utf16) == 0 {
		return false
	}
	nBytes := len(utf16) * int(unsafe.Sizeof(uint16(0)))
	h, _, _ := ghostGlobalAlloc.Call(ghostGMEMMoveable, uintptr(nBytes))
	if h == 0 {
		return false
	}

	p, _, _ := ghostGlobalLock.Call(h)
	if p == 0 {
		_, _, _ = ghostGlobalFree.Call(h)
		return false
	}
	src := uintptr(unsafe.Pointer(&utf16[0]))
	_, _, _ = ghostRtlMoveMemory.Call(p, src, uintptr(nBytes))
	_, _, _ = ghostGlobalUnlock.Call(h)

	r, _, _ := ghostSetClipboardData.Call(ghostCFUnicodeText, h)
	if r == 0 {
		_, _, _ = ghostGlobalFree.Call(h)
		return false
	}
	return true
}
