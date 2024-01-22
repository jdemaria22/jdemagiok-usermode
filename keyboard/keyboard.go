package keyboard

import (
	"jdemagiok-usermode/win"

	"github.com/micmonay/keybd_event"
)

func IsKeyDown(key int) bool {
	return win.GetAsyncKeyState(int32(key)) < 0
}

var keyStateMap = make(map[int]bool)

func IsKeyPressed(key int) bool {
	currentState := win.GetAsyncKeyState(int32(key)) < 0
	previousState, ok := keyStateMap[key]
	keyStateMap[key] = currentState
	return currentState && !previousState && ok
}
func InitKeyBonding() (keybd_event.KeyBonding, error) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		return kb, err
	}
	kb.SetKeys(keybd_event.VK_K)
	return kb, nil
}
func ListenForKeyboardEvents(ch chan bool) {
	for {
		// Escucha eventos de teclado y envía el código de la tecla al canal
		keyCode := IsKeyDown(win.VK_LSHIFT) // Obtén el código de tecla
		ch <- keyCode
	}
}
