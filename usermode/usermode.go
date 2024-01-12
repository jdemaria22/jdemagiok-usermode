package usermode

import (
	"syscall"
	"unsafe"
)

var (
	modKernel32                  = syscall.NewLazyDLL("kernel32.dll")
	procCreateToolhelp32Snapshot = modKernel32.NewProc("CreateToolhelp32Snapshot")
	procProcess32First           = modKernel32.NewProc("Process32FirstW")
	procProcess32Next            = modKernel32.NewProc("Process32NextW")
)

const (
	TH32CS_SNAPPROCESS = 0x00000002
)

type PROCESSENTRY32 struct {
	dwSize              uint32
	cntUsage            uint32
	th32ProcessID       uint32
	th32DefaultHeapID   uintptr
	th32ModuleID        uint32
	cntThreads          uint32
	th32ParentProcessID uint32
	pcPriClassBase      int32
	dwFlags             uint32
	szExeFile           [syscall.MAX_PATH]uint16
}

func GetProcessID(processName string) int {
	snapshot, _, _ := procCreateToolhelp32Snapshot.Call(uintptr(TH32CS_SNAPPROCESS), 0)
	defer syscall.CloseHandle(syscall.Handle(snapshot))

	var entry PROCESSENTRY32
	entry.dwSize = uint32(unsafe.Sizeof(entry))

	proc32First := uintptr(procProcess32First.Addr())
	ret, _, _ := syscall.SyscallN(proc32First, 2, snapshot, uintptr(unsafe.Pointer(&entry)), 0)

	if ret == 0 {
		return 0
	}
	var ID uint32 = 0
	for {
		if syscall.UTF16ToString(entry.szExeFile[:]) == processName {
			ID = entry.th32ProcessID
			return int(ID)
		}
		ret, _, _ = syscall.SyscallN(procProcess32Next.Addr(), 2, snapshot, uintptr(unsafe.Pointer(&entry)), 0)
		if ret == 0 {
			break
		}
	}
	panic("Failed to get proccessID")
}
