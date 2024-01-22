package game

import (
	"jdemagiok-usermode/geometry"
	"jdemagiok-usermode/kernel"
	"jdemagiok-usermode/offset"
	"log"
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
	gameWorld.GameState = d.Read(gameWorld.Pointer + offset.GameStateOffset)
	gameWorld.PersistanceLevel.ActorArray = getArray(d, gameWorld)
	return gameWorld
}

func GetPersisntanceLevel(d *kernel.Driver, world SWorld) SPersistanceLevel {
	persistanceLevel := SPersistanceLevel{}
	persistanceLevel.Pointer = d.Read(world.Pointer + offset.PersistentLevelOffset)
	if persistanceLevel.Pointer == 0 {
		log.Fatal("se acabo instancia")
	}
	return persistanceLevel
}

func GetGameInstance(d *kernel.Driver, world SWorld) SGameInstance {
	gameInstance := SGameInstance{}
	gameInstance.Pointer = d.Read(world.Pointer + offset.OwningGameInstanceOffset)
	if gameInstance.Pointer == 0 {
		log.Fatal("se acabo instancia")
	}
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
	playerController.MinimalViewInfo = d.ReadvmMinimalView(playerController.PlayerCameraManager + offset.CacheCamOffset)
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
	pawn.Health = getHealth(d, pawn.Pointer)
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

func getHealth(d *kernel.Driver, pawnPointer uintptr) float32 {
	damageHandler := d.Read(pawnPointer + offset.DamageHandlerOffset)
	return d.ReadvmFloat(damageHandler + offset.CurrentHealthOffset)
}

func getUWorld(d *kernel.Driver) uintptr {
	uworld := uintptr(d.Readvm(d.Guardedregion+offset.World, 8))
	if uworld == 0 {
		log.Fatal("se acabo instancia")
	}
	var uworldOffset uintptr

	if uworld > 0x10000000000 {
		uworldOffset = uworld - 0x10000000000
	} else {
		uworldOffset = uworld - 0x8000000000
	}

	return d.Guardedregion + uworldOffset
}

func getArray(d *kernel.Driver, world SWorld) []SActor {
	array := d.ReadvmArray(world.GameState + offset.SpawnedCharacter)
	var actors []SActor
	myTeamid := world.GameInstance.LocalPlayer.PlayerController.Pawn.TeamID
	for i := 0; i < int(array.Count); i++ {
		actorPointerPawn := array.ReadAtIndex2(i, d)
		if actorPointerPawn == 0 {
			continue
		}
		pawnPtr := d.Read(actorPointerPawn + 0x918)
		if pawnPtr == 0 {
			continue
		}
		pawnTeamId := getTeamID(d, pawnPtr)
		if myTeamid == pawnTeamId {
			continue
		}
		pawn := SEnemyPawn{}
		pawn.Pointer = pawnPtr
		actors = append(actors, SActor{pawn})
	}
	return actors
}
