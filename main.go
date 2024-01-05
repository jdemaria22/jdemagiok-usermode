package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"elgordoloren/guncontrol"
	"elgordoloren/mouse"

	"github.com/kbinani/screenshot"
	"github.com/micmonay/keybd_event"
)

func main() {
	counter := &guncontrol.KeyPressCounter{Reset: make(chan struct{})}
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	go func() {
		for {
			select {
			case <-counter.Reset:
				counter.Count = 0
			}
		}
	}()

	for {
		img, err := screenshot.Capture((1920/2)-100, (1080/2)-100, 200, 200)
		if err != nil {
			log.Fatal(err)
		}

		recoil := guncontrol.ProcessImage(img, color.RGBA{R: 250, G: 100, B: 250, A: 255}, 10000)

		if recoil != nil && counter.Count < 2 {
			originalX := 910 + recoil.X
			originalY := 490 + recoil.Y

			finalX := originalX - (1920 / 2)
			finalY := originalY - (1080 / 2)

			mouse.MoveTo(int32(finalX), int32(finalY+5))

			kb, err := keybd_event.NewKeyBonding()
			if err != nil {
				fmt.Println(err)
				return
			}
			kb.SetKeys(keybd_event.VK_K)
			err = kb.Launching()
			if err != nil {
				fmt.Println(err)
				return
			}
			kb.SetKeys(keybd_event.VK_K)
			if err != nil {
				fmt.Println(err)
				return
			}
			err = kb.Launching()
			if err != nil {
				fmt.Println(err)
				return
			}
			counter.Count++
		}
		if counter.Count >= 2 {
			time.Sleep(300 * time.Millisecond)
			counter.Reset <- struct{}{}
		}

		lowerLimit := 30
		upperLimit := 144

		randomNum := rng.Intn(upperLimit-lowerLimit+1) + lowerLimit
		duration := time.Duration(randomNum) * time.Millisecond
		time.Sleep(duration)
	}
}
