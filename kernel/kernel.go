package kernel

import (
	"fmt"
	"jdemagiok-usermode/usermode"
	"syscall"
	"unsafe"
)

const (
	PROCESSID_REQUEST   = iota
	KERNEL_DRIVER_NAME  = "\\\\.\\linkjdemagiok82"
	FILE_DEVICE_UNKNOWN = 0x00000022
	METHOD_BUFFERED     = 0
	FILE_ANY_ACCESS     = 0
)

type KPROCESSID_REQUEST struct {
	processName string
}

type _KERNEL_MODULE_REQUEST struct {
	pid        int
	moduleName string
	getSize    bool
}

type KERNEL_READ_REQUEST struct {
	srcPid     int
	srcAddress uintptr
	pBuffer    *uintptr
	size       int
}

type KERNEL_READ_FLOAT_REQUEST struct {
	srcPid     int
	srcAddress uintptr
	pBuffer    *float32
	size       int
}

type _KERNEL_READ_GUARDED_REGION struct {
	srcPid  int
	pBuffer *uintptr
}

type KMOUSE_REQUEST struct {
	x, y        int32
	buttonFlags uint8
}

type Driver struct {
	handle        syscall.Handle
	processID     int
	Guardedregion uintptr
}

func CTL_CODE(deviceType, function, method, access uint32) uint32 {
	return (deviceType << 16) | (access << 14) | (function << 2) | method
}

func NewDriver() *Driver {
	name, _ := syscall.UTF16PtrFromString(KERNEL_DRIVER_NAME)
	handle, err := syscall.CreateFile(
		name,
		syscall.GENERIC_WRITE|syscall.GENERIC_READ,
		syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE,
		nil,
		syscall.OPEN_EXISTING,
		0,
		0,
	)
	if err != nil {
		panic(fmt.Sprintf("Failed to create file: %v", err))
	}
	processID := usermode.GetProcessID("VALORANT-Win64-Shipping.exe")
	fmt.Println("Communication with driver created")
	fmt.Println("Handle:", handle)
	fmt.Println("Proccess ID:", processID)
	d := &Driver{
		handle:    handle,
		processID: processID,
	}
	d.Guardedregion = d.ReadGuardedRegion()
	fmt.Printf("Guarded region: %x\n", d.Guardedregion)
	return d
}

func (d *Driver) Close() {
	syscall.CloseHandle(d.handle)
}

func (d *Driver) GetProcessID(processName string) int {
	request := KPROCESSID_REQUEST{processName: processName}

	var bytesReturned uint32
	err := syscall.DeviceIoControl(
		d.handle,
		CTL_CODE(FILE_DEVICE_UNKNOWN, 0x555, METHOD_BUFFERED, FILE_ANY_ACCESS),
		(*byte)(unsafe.Pointer(&request)),
		uint32(unsafe.Sizeof(request)),
		nil,
		0,
		&bytesReturned,
		nil,
	)
	if err != nil {
		fmt.Println("Error getting process ID:", err)
		return 0
	}

	return int(bytesReturned)
}

func (d *Driver) GetModuleBase(pid int, moduleName string, getSize bool) uintptr {
	request := _KERNEL_MODULE_REQUEST{
		pid:        pid,
		moduleName: moduleName,
		getSize:    getSize,
	}

	var bytesReturned uint32
	err := syscall.DeviceIoControl(
		d.handle,
		CTL_CODE(FILE_DEVICE_UNKNOWN, 0x777, METHOD_BUFFERED, FILE_ANY_ACCESS),
		(*byte)(unsafe.Pointer(&request)),
		uint32(unsafe.Sizeof(request)),
		nil,
		0,
		&bytesReturned,
		nil,
	)
	if err != nil {
		fmt.Println("Error getting module base:", err)
		return 0
	}

	return uintptr(bytesReturned)
}

func (d *Driver) Move(x, y int32, buttonFlags uint8) {
	request := KMOUSE_REQUEST{
		x:           x,
		y:           y,
		buttonFlags: buttonFlags,
	}

	err := syscall.DeviceIoControl(
		d.handle,
		CTL_CODE(FILE_DEVICE_UNKNOWN, 0x666, METHOD_BUFFERED, FILE_ANY_ACCESS),
		(*byte)(unsafe.Pointer(&request)),
		uint32(unsafe.Sizeof(request)),
		nil,
		0,
		nil,
		nil,
	)
	if err != nil {
		fmt.Println("Error moving mouse:", err)
	}
}

func (d *Driver) ReadGuardedRegion() uintptr {
	var buffer uintptr
	request := _KERNEL_READ_GUARDED_REGION{
		srcPid:  d.processID,
		pBuffer: &buffer,
	}

	err := syscall.DeviceIoControl(
		d.handle,
		CTL_CODE(FILE_DEVICE_UNKNOWN, 0x444, METHOD_BUFFERED, FILE_ANY_ACCESS),
		(*byte)(unsafe.Pointer(&request)),
		uint32(unsafe.Sizeof(request)),
		nil,
		0,
		nil,
		nil,
	)

	if err != nil {
		fmt.Println("Error reading memory:", err)
		return 0
	}

	return *request.pBuffer
}

func (d *Driver) Readvm(address uintptr, size int) uintptr {
	var buffer uintptr
	request := KERNEL_READ_REQUEST{
		srcPid:     d.processID,
		srcAddress: address,
		pBuffer:    &buffer,
		size:       8,
	}

	err := syscall.DeviceIoControl(
		d.handle,
		CTL_CODE(FILE_DEVICE_UNKNOWN, 0x888, METHOD_BUFFERED, FILE_ANY_ACCESS),
		(*byte)(unsafe.Pointer(&request)),
		uint32(unsafe.Sizeof(request)),
		nil,
		0,
		nil,
		nil,
	)
	if err != nil {
		fmt.Println("Error reading memory:", err)
		return 0
	}
	return *request.pBuffer
}

func (d *Driver) ReadvmFloat(address uintptr) float32 {
	var buffer float32
	request := KERNEL_READ_FLOAT_REQUEST{
		srcPid:     d.processID,
		srcAddress: address,
		pBuffer:    &buffer,
		size:       4,
	}

	err := syscall.DeviceIoControl(
		d.handle,
		CTL_CODE(FILE_DEVICE_UNKNOWN, 0x888, METHOD_BUFFERED, FILE_ANY_ACCESS),
		(*byte)(unsafe.Pointer(&request)),
		uint32(unsafe.Sizeof(request)),
		nil,
		0,
		nil,
		nil,
	)
	if err != nil {
		fmt.Println("Error reading memory:", err)
		return 0
	}
	return *request.pBuffer
}

func IsGuarded(pointer uintptr) bool {
	filter := uintptr(0xFFFFFFF000000000)
	result := pointer & filter
	return result == 0x8000000000 || result == 0x10000000000
}

func WardedTo(guardedRegion uintptr, pointer uintptr) uintptr {
	var offset uintptr

	if pointer > 0x10000000000 {
		offset = pointer - 0x10000000000
	} else {
		offset = pointer - 0x8000000000
	}

	return guardedRegion + offset
}
