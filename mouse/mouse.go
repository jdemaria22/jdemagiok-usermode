package mouse

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	FILE_DEVICE_UNKNOWN = 0x00000022
	METHOD_BUFFERED     = 0
	FILE_ANY_ACCESS     = 0
)

var (
	DeviceName  = "\\\\.\\linkelgordolorenzaish"
	pointerName *uint16
)

type KMOUSE_REQUEST struct {
	X           int32
	Y           int32
	ButtonFlags uint8
}

func init() {
	pointerName, _ = syscall.UTF16PtrFromString(DeviceName)
}

func CTL_CODE(deviceType, function, method, access uint32) uint32 {
	return (deviceType << 16) | (access << 14) | (function << 2) | method
}

func MoveTo(x int32, y int32) {
	handle, err := syscall.CreateFile(
		pointerName,
		syscall.GENERIC_READ|syscall.GENERIC_WRITE,
		0,
		nil,
		syscall.OPEN_EXISTING,
		0,
		0,
	)
	if err != nil {
		fmt.Printf("Error al abrir el archivo. Código de error: %v\n", err)
		return
	}
	defer syscall.CloseHandle(handle)

	var bytesReturned uint32
	mouseRequest := KMOUSE_REQUEST{
		X:           x,
		Y:           y,
		ButtonFlags: 0,
	}

	size := unsafe.Sizeof(mouseRequest)
	mouseRequestBytes := *(*[unsafe.Sizeof(mouseRequest)]byte)(unsafe.Pointer(&mouseRequest))

	err = syscall.DeviceIoControl(
		handle,
		CTL_CODE(FILE_DEVICE_UNKNOWN, 0x666, METHOD_BUFFERED, FILE_ANY_ACCESS),
		&mouseRequestBytes[0],
		uint32(size),
		nil,
		0,
		&bytesReturned,
		nil,
	)
	if err != nil {
		fmt.Printf("Error al enviar la solicitud al controlador. Código de error: %v\n", err)
		return
	}

}
