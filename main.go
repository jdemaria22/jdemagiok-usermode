package main

import (
	"fmt"
	"jdemagiok-usermode/game"
	"jdemagiok-usermode/kernel"
	"jdemagiok-usermode/sys"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"syscall"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth       = 1920
	screenHeight      = 1080
	RANDOM            = 12354
	MAX_TPS           = 1000
	GWL_EXSTYLE       = -20
	WS_EX_COMPOSITED  = 0x02000000
	WS_EX_LAYERED     = 0x00080000
	WS_EX_TRANSPARENT = 0x00000020
	WS_EX_TOOLWINDOW  = 0x00000080
	WS_EX_TOPMOST     = 0x00000008
)

var count = 0
var NAME string
var gamec game.SGame
var driver *kernel.Driver

type GameE struct {
}

func main() {
	NAME = strconv.Itoa(rand.Intn(RANDOM))
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle(NAME)
	ebiten.SetWindowResizable(false)
	ebiten.SetWindowDecorated(false)
	ebiten.SetScreenTransparent(true)
	ebiten.SetWindowFloating(true)
	ebiten.SetInitFocused(true)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetMaxTPS(165)

	driver = kernel.NewDriver()
	defer driver.Close()

	err := ebiten.RunGame(NewGameE())
	if err != nil {
		log.Fatal(err)
	}
}

func NewGameE() *GameE {
	return nil
}

func (g *GameE) Update() error {
	if count == 0 {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			count++
			wnproc, err := sys.FindWindow(nil, syscall.StringToUTF16Ptr(NAME))
			if err != nil {
				fmt.Println("error in  FindWindow: ", err)
			}
			r2, err := sys.SetWindowLong(wnproc, GWL_EXSTYLE, WS_EX_COMPOSITED|WS_EX_LAYERED|WS_EX_TRANSPARENT|WS_EX_TOOLWINDOW|WS_EX_TOPMOST)
			if err != nil {
				fmt.Println("error in  setWindowLong: ", err, r2)
			}
			sys.SetForegroundWindow(wnproc)
		}()
	}
	gamec = game.GetGame(driver)
	return nil
}

func (g *GameE) Draw(screen *ebiten.Image) {
	gamec.Draw(driver, screen)
}

func (g *GameE) Layout(outsideWidth, outsideHeight int) (int, int) {
	s := ebiten.DeviceScaleFactor()
	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
}
