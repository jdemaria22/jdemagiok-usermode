package main

import (
	"fmt"
	"jdemagiok-usermode/game"
	"jdemagiok-usermode/kernel"
)

func main() {
	driver := kernel.NewDriver()
	defer driver.Close()
	gameObject := game.GetGame(driver)
	uniqueId := gameObject.World.GameInstance.LocalPlayer.PlayerController.Pawn.UniqueID
	fmt.Println("uniqueId \n", uniqueId)
	actorArray := gameObject.World.PersistanceLevel.ActorArray
	for _, actorArray := range actorArray {
		fmt.Println("RelativeLocation \n", actorArray.Pawn.RelativeLocation)
	}
}
