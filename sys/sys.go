package sys

import (
	"syscall"
	"unsafe"
)

var (
	moduser32               = syscall.NewLazyDLL("user32.dll")
	procFindWindow          = moduser32.NewProc("FindWindowW")
	procSetWindowLong       = moduser32.NewProc("SetWindowLongW")
	procSetForegroundWindow = moduser32.NewProc("SetForegroundWindow")
)

func FindWindow(lpClassName, lpWindowName *uint16) (syscall.Handle, error) {
	ret, _, err := procFindWindow.Call(uintptr(unsafe.Pointer(lpClassName)), uintptr(unsafe.Pointer(lpWindowName)))
	if ret == 0 {
		return 0, err
	}
	return syscall.Handle(ret), nil
}

func SetWindowLong(hWnd syscall.Handle, nIndex int, dwNewLong int32) (int32, error) {
	ret, _, err := procSetWindowLong.Call(uintptr(hWnd), uintptr(nIndex), uintptr(dwNewLong))
	if ret == 0 {
		return 0, err
	}
	return int32(ret), nil
}

func SetForegroundWindow(hWnd syscall.Handle) error {
	_, _, err := procSetForegroundWindow.Call(uintptr(hWnd))
	if err != nil {
		return err
	}
	return nil
}
