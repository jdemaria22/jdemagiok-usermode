package main

import (
	"fmt"
	"jdemagiok-usermode/kernel"
	"jdemagiok-usermode/usermode"
)

func main() {
	driver := kernel.NewDriver()
	defer driver.Close()

	processID := usermode.GetProcessID("VALORANT-Win64-Shipping.exe")
	fmt.Println("Process ID:", processID)

}
