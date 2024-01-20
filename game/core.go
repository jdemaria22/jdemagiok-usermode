package game

import (
	"jdemagiok-usermode/geometry"
	"jdemagiok-usermode/kernel"
	"jdemagiok-usermode/offset"
)

func GetGame(d *kernel.Driver) SGame {
	gameObject := SGame{}
	gameObject.World = GetWorld(d)
	return gameObject
}

func GetWorld(d *kernel.Driver) SWorld {
	gameWorld := SWorld{}
	gameWorld.Pointer = getUWorld(d)
	gameWorld.GameInstance = GetGameInstance(d, gameWorld)
	gameWorld.PersistanceLevel = GetPersisntanceLevel(d, gameWorld)
	return gameWorld
}

func GetPersisntanceLevel(d *kernel.Driver, world SWorld) SPersistanceLevel {
	persistanceLevel := SPersistanceLevel{}
	persistanceLevel.Pointer = d.Read(world.Pointer + offset.PersistentLevelOffset)
	persistanceLevel.ActorArray = getActorArray(d, world, persistanceLevel)
	return persistanceLevel
}

func GetGameInstance(d *kernel.Driver, world SWorld) SGameInstance {
	gameInstance := SGameInstance{}
	gameInstance.Pointer = d.Read(world.Pointer + offset.OwningGameInstanceOffset)
	gameInstance.LocalPlayerArray = d.Read(gameInstance.Pointer + offset.LocalPlayersOffset)
	gameInstance.LocalPlayer = GetLocalPlayer(d, gameInstance)
	return gameInstance
}

func GetLocalPlayer(d *kernel.Driver, gameInstance SGameInstance) SLocalPlayer {
	localPlayer := SLocalPlayer{}
	localPlayer.Pointer = d.Read(gameInstance.LocalPlayerArray)
	localPlayer.PlayerController = GetPlayerController(d, localPlayer)
	return localPlayer
}

func GetPlayerController(d *kernel.Driver, localPlayer SLocalPlayer) SPlayerController {
	playerController := SPlayerController{}
	playerController.Pointer = d.Read(localPlayer.Pointer + offset.PlayerControllerOffset)
	playerController.PlayerCameraManager = d.Read(playerController.Pointer + offset.PlayerCameraOffset)
	playerController.AHUD = d.Read(playerController.Pointer + offset.MyHUDOffset)
	playerController.Pawn = GetPawn(d, playerController)
	return playerController
}

func GetPawn(d *kernel.Driver, playerController SPlayerController) SPawn {
	pawn := SPawn{}
	pawn.Pointer = d.Read(playerController.Pointer + offset.AcknowledgedPawnOffset)
	pawn.TeamID = getTeamID(d, pawn.Pointer)
	pawn.UniqueID = d.ReadvmInt(pawn.Pointer + offset.ActorIDOffset)
	pawn.FNameID = d.ReadvmInt(pawn.Pointer + offset.FnameIDOffset)
	pawn.RelativeLocation = getRelativePosition(d, pawn.Pointer)
	pawn.BIsDormant = d.ReadvmBool(pawn.Pointer + offset.DormantOffset)
	pawn.Health = getHealth(d, pawn)
	return pawn
}

func getTeamID(d *kernel.Driver, aPawn uintptr) int {
	playerState := d.Read(aPawn + offset.PlayerStateOffset)
	teamComponent := d.Read(playerState + offset.TeamComponentOffset)
	return d.ReadvmInt(teamComponent + offset.TeamIDOffset)
}

func getRelativePosition(d *kernel.Driver, aPawn uintptr) geometry.FVector {
	rootComponent := d.Read(aPawn + offset.RootComponentOffset)
	return d.ReadvmVector(rootComponent + offset.RelativeLocationOffset)
}

func getHealth(d *kernel.Driver, pawn SPawn) float32 {
	damageHandler := d.Read(pawn.Pointer + offset.DamageHandlerOffset)
	return d.ReadvmFloat(damageHandler + offset.CurrentHealthOffset)
}

func getUWorld(d *kernel.Driver) uintptr {
	uworld := uintptr(d.Readvm(d.Guardedregion+offset.World, 8))

	var uworldOffset uintptr

	if uworld > 0x10000000000 {
		uworldOffset = uworld - 0x10000000000
	} else {
		uworldOffset = uworld - 0x8000000000
	}

	return d.Guardedregion + uworldOffset
}

func getActorArray(d *kernel.Driver, world SWorld, persistanceLevel SPersistanceLevel) []SActor {
	actorArray := d.ReadvmArray(persistanceLevel.Pointer + offset.ActorArrayOffset)
	var actors []SActor
	for i := 0; i < int(actorArray.Count); i++ {
		actorPointerPawn := actorArray.ReadAtIndex(i, d)
		if actorPointerPawn != world.GameInstance.LocalPlayer.PlayerController.Pawn.Pointer {
			id := d.ReadvmInt(actorPointerPawn + offset.ActorIDOffset)
			if world.GameInstance.LocalPlayer.PlayerController.Pawn.UniqueID == id {
				pawn := getPawn(d, actorPointerPawn)
				pawn.Pointer = actorPointerPawn
				actors = append(actors, SActor{pawn})
			}
		}
	}
	return actors
}

func getPawn(d *kernel.Driver, pawnPointer uintptr) SPawn {
	pawn := SPawn{}
	pawn.Pointer = pawnPointer
	pawn.TeamID = getTeamID(d, pawn.Pointer)
	pawn.UniqueID = d.ReadvmInt(pawn.Pointer + offset.ActorIDOffset)
	pawn.FNameID = d.ReadvmInt(pawn.Pointer + offset.FnameIDOffset)
	pawn.RelativeLocation = getRelativePosition(d, pawn.Pointer)
	pawn.BIsDormant = d.ReadvmBool(pawn.Pointer + offset.DormantOffset)
	pawn.Health = getHealth(d, pawn)
	return pawn
}
