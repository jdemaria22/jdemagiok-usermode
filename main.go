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
	uLevel := game.GetGameULevel(driver, world)
	fmt.Printf("uLevel %x\n", uLevel)
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

	nameId := game.GetFNameId(driver, aPawn)
	fmt.Printf("nameId %d\n", nameId)
	teamId := game.GetTeamID(driver, aPawn)
	fmt.Printf("teamId %d\n", teamId)
	location := game.GetRelativePosition(driver, aPawn)
	fmt.Println("location \n", location)

	uniqueId := game.GetUniqueID(driver, aPawn)
	fmt.Printf("uniqueId %d\n", uniqueId)

	actorArray := game.GetActorArray(driver, uLevel)
	fmt.Println("actorArray \n", actorArray)

	for i := 0; i < int(actorArray.Count); i++ {
		pawns := actorArray.ReadAtIndex(i, driver)
		if pawns != aPawn {
			pawnId := game.GetUniqueID(driver, pawns)
			fmt.Printf("pawnId %d\n", pawnId)
		}
	}

}
