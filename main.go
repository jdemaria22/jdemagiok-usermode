package main

import (
	"fmt"
	"jdemagiok-usermode/game"
	"jdemagiok-usermode/kernel"
)

func main() {
	driver := kernel.NewDriver()
	defer driver.Close()
	gameObject := game.SGame{}
	gameWorld := game.SWorld{}
	gameInstance := game.SGameInstance{}
	persistanceLevel := game.SPersistanceLevel{}
	gameWorld.GameInstance = gameInstance
	gameWorld.PersistanceLevel = persistanceLevel

	gameWorld.Pointer = game.GetUWorld(driver)
	gameInstance.Pointer = game.GetGameInstance(driver, gameWorld.Pointer)
	gameInstance.LocalPlayerArray = game.GetULocalPlayerArray(driver, gameInstance.Pointer)

	localPlayer := game.SLocalPlayer{}
	localPlayer.Pointer = game.GetULocalPlayer(driver, gameInstance.LocalPlayerArray)
	playerController := game.SPlayerController{}
	playerController.Pointer = game.GetAPlayerControllerPtr(driver, localPlayer.Pointer)

	pawn := game.SPawn{}
	playerController.Pawn = pawn

	gameObject.World = gameWorld

	// world := game.GetUWorld(driver)
	// fmt.Printf("World %x\n", world)
	// gameinstance := game.GetGameInstance(driver, world)
	// fmt.Printf("gameinstance %x\n", gameinstance)
	// uLevel := game.GetGameULevel(driver, world)
	// fmt.Printf("uLevel %x\n", uLevel)
	// uLocalPlayerArray := game.GetULocalPlayerArray(driver, gameinstance)
	// fmt.Printf("uLocalPlayerArray %x\n", uLocalPlayerArray)

	// uLocalPlayer := game.GetULocalPlayer(driver, uLocalPlayerArray)
	// fmt.Printf("uLocalPlayer %x\n", uLocalPlayer)
	// aPlayerControllerPtr := game.GetAPlayerControllerPtr(driver, uLocalPlayer)
	// fmt.Printf("aPlayerControllerPtr %x\n", aPlayerControllerPtr)
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

	gameStateBase := game.GetGameStateBase(driver, world)
	fmt.Printf("gameStateBase %x\n", gameStateBase)

	arrayOfPLayer := game.GetULocalPlayerArray(driver, gameStateBase)
	fmt.Printf("arrayOfPLayer %x\n", arrayOfPLayer)

	uniqueId := game.GetUniqueID(driver, aPawn)
	fmt.Printf("uniqueId %d\n", uniqueId)

	actorArray := game.GetActorArray(driver, uLevel)
	fmt.Println("actorArray \n", actorArray)

	var actors []uintptr
	for i := 0; i < int(actorArray.Count); i++ {
		pawns := actorArray.ReadAtIndex(i, driver)
		if pawns != aPawn {
			id := game.GetUniqueID(driver, pawns)
			if uniqueId == id {
				actors = append(actors, pawns)
			}
		}
	}
	fmt.Printf("pawns %x\n", actors)
	for _, actor := range actors {
		actorrelativeLocation := game.GetRelativePosition(driver, actor)
		fmt.Println("location \n", actorrelativeLocation)
	}
}
