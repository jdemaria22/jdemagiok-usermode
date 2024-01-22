package win

import (
	"syscall"
	"unsafe"
)

// DEFINED IN THE DWM API BUT NOT IMPLEMENTED BY MS:
// DwmAttachMilContent
// DwmDetachMilContent
// DwmEnableComposition
// DwmGetGraphicsStreamClient
// DwmGetGraphicsStreamTransformHint

var (
	moddwmapi = syscall.NewLazyDLL("dwmapi.dll")

	procDwmDefWindowProc          = moddwmapi.NewProc("DwmDefWindowProc")
	procDwmEnableBlurBehindWindow = moddwmapi.NewProc("DwmEnableBlurBehindWindow")
)

func DwmDefWindowProc(hWnd HWND, msg uint, wParam, lParam uintptr) (bool, uint) {
	var result uint
	ret, _, _ := procDwmDefWindowProc.Call(
		uintptr(hWnd),
		uintptr(msg),
		wParam,
		lParam,
		uintptr(unsafe.Pointer(&result)))
	return ret != 0, result
}

type DWM_BLURBEHIND struct {
	DwFlags                uint32
	fEnable                BOOL
	hRgnBlur               HRGN
	fTransitionOnMaximized BOOL
}

type MARGINS struct {
	CxLeftWidth, CxRightWidth, CyTopHeight, CyBottomHeight int32
}

func DwmEnableBlurBehindWindow(hWnd HWND) HRESULT {
	var data DWM_BLURBEHIND
	hrgn := CreateRectRgn(0, 0, -1, -1)
	data.DwFlags = DWM_BB_ENABLE | DWM_BB_BLURREGION
	var enable BOOL = TRUE
	data.fEnable = enable
	data.hRgnBlur = hrgn
	ret, _, _ := procDwmEnableBlurBehindWindow.Call(
		uintptr(hWnd),
		uintptr(unsafe.Pointer(&data)))
	return HRESULT(ret)
}

const (
	DWM_BB_ENABLE                = 0x00000001 //     A value for the fEnable member has been specified.
	DWM_BB_BLURREGION            = 0x00000002 //     A value for the hRgnBlur member has been specified.
	DWM_BB_TRANSITIONONMAXIMIZED = 0x00000004 //     A value for the fTransitionOnMaximized member has been specified.
)
