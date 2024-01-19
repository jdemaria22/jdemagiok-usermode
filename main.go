package main

import (
	"fmt"
	"jdemagiok-usermode/game"
	"jdemagiok-usermode/kernel"
)

func main() {
	driver := kernel.NewDriver()
	defer driver.Close()
	world := game.GetUWorld(driver)
	fmt.Printf("World %x\n", world)
	gameinstance := game.GetGameInstance(driver, world)
	fmt.Printf("gameinstance %x\n", gameinstance)
	uLocalPlayerArray := game.GetULocalPlayerArray(driver, gameinstance)
	fmt.Printf("uLocalPlayerArray %x\n", uLocalPlayerArray)
	uLocalPlayer := game.GetULocalPlayer(driver, uLocalPlayerArray)
	fmt.Printf("uLocalPlayer %x\n", uLocalPlayer)
	aPlayerControllerPtr := game.GetAPlayerControllerPtr(driver, uLocalPlayer)
	fmt.Printf("aPlayerControllerPtr %x\n", aPlayerControllerPtr)
	aPawn := game.GetAPawn(driver, aPlayerControllerPtr)
	fmt.Printf("aPawn %x\n", aPawn)
	damageHandler := game.GetDamageHandler(driver, aPawn)
	fmt.Printf("damageHandler %x\n", damageHandler)
	health := game.GetHealth(driver, damageHandler)
	fmt.Printf("health %f\n", health)
}
