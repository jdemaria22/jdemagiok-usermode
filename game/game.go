package game

import (
	"jdemagiok-usermode/kernel"
	"jdemagiok-usermode/offset"
)

func GetUWorld(d *kernel.Driver) uintptr {
	uworld := uintptr(d.Readvm(d.Guardedregion+offset.World, 8))

	var uworldOffset uintptr

	if uworld > 0x10000000000 {
		uworldOffset = uworld - 0x10000000000
	} else {
		uworldOffset = uworld - 0x8000000000
	}

	return d.Guardedregion + uworldOffset
}

func GetGameInstance(d *kernel.Driver, uworld uintptr) uintptr {
	gameinstance := uintptr(d.Readvm(uworld+offset.OwningGameInstanceOffset, 8))

	if kernel.IsGuarded(gameinstance) {
		return kernel.WardedTo(d.Guardedregion, gameinstance)
	}

	return gameinstance
}

func GetGameULevel(d *kernel.Driver, uworld uintptr) uintptr {
	ulevel := uintptr(d.Readvm(uworld+offset.PersistentLevelOffset, 8))

	if kernel.IsGuarded(ulevel) {
		return kernel.WardedTo(d.Guardedregion, ulevel)
	}

	return ulevel
}

func GetULocalPlayerArray(d *kernel.Driver, gameinstance uintptr) uintptr {
	uLocalPlayerArray := uintptr(d.Readvm(gameinstance+offset.LocalPlayersOffset, 8))

	if kernel.IsGuarded(uLocalPlayerArray) {
		return kernel.WardedTo(d.Guardedregion, uLocalPlayerArray)
	}

	return uLocalPlayerArray
}

func GetULocalPlayer(d *kernel.Driver, uLocalPlayerArray uintptr) uintptr {
	uLocalPlayer := uintptr(d.Readvm(uLocalPlayerArray, 8))

	if kernel.IsGuarded(uLocalPlayer) {
		return kernel.WardedTo(d.Guardedregion, uLocalPlayer)
	}

	return uLocalPlayer
}

func GetAPlayerControllerPtr(d *kernel.Driver, uLocalPlayer uintptr) uintptr {
	aPlayerControllerPtr := uintptr(d.Readvm(uLocalPlayer+offset.PlayerControllerOffset, 8))

	if kernel.IsGuarded(aPlayerControllerPtr) {
		return kernel.WardedTo(d.Guardedregion, aPlayerControllerPtr)
	}

	return aPlayerControllerPtr
}

func GetAPawn(d *kernel.Driver, aPlayerControllerPtr uintptr) uintptr {
	aPawn := uintptr(d.Readvm(aPlayerControllerPtr+offset.AcknowledgedPawnOffset, 8))

	if kernel.IsGuarded(aPawn) {
		return kernel.WardedTo(d.Guardedregion, aPawn)
	}

	return aPawn
}

func GetDamageHandler(d *kernel.Driver, aPawn uintptr) uintptr {
	damageHandler := uintptr(d.Readvm(aPawn+offset.DamageHandlerOffset, 8))

	if kernel.IsGuarded(damageHandler) {
		return kernel.WardedTo(d.Guardedregion, damageHandler)
	}

	return damageHandler
}

func GetBIsDormant(d *kernel.Driver, aPawn uintptr) bool {
	return d.ReadvmBool(aPawn + offset.DormantOffset)
}

func GetFNameId(d *kernel.Driver, aPawn uintptr) int {
	return d.ReadvmInt(aPawn + offset.FnameIDOffset)
}

func GetUniqueID(d *kernel.Driver, aPawn uintptr) int {
	return d.ReadvmInt(aPawn + offset.FnameIDOffset)
}

func GetHealth(d *kernel.Driver, aPawn uintptr) float32 {
	return d.ReadvmFloat(aPawn + offset.CurrentHealthOffset)
}

func GetTeamID(d *kernel.Driver, aPawn uintptr) int {
	playerState := d.Readvm(aPawn+offset.PlayerStateOffset, 8)
	if kernel.IsGuarded(playerState) {
		playerState = kernel.WardedTo(d.Guardedregion, playerState)
	}
	teamComponent := d.Readvm(playerState+offset.TeamComponentOffset, 8)
	if kernel.IsGuarded(teamComponent) {
		teamComponent = kernel.WardedTo(d.Guardedregion, teamComponent)
	}
	return d.ReadvmInt(teamComponent + offset.TeamIDOffset)
}

func GetRelativePosition(d *kernel.Driver, aPawn uintptr) kernel.FVector {
	rootComponent := d.Readvm(aPawn+offset.RootComponentOffset, 8)
	if kernel.IsGuarded(rootComponent) {
		rootComponent = kernel.WardedTo(d.Guardedregion, rootComponent)
	}
	return d.ReadvmVector(rootComponent + offset.RelativeLocationOffset)
}

func GetActorArray(d *kernel.Driver, uLevel uintptr) kernel.TArrayDrink {
	return d.ReadvmArray(uLevel + offset.ActorArrayOffset)
}

func GetActorArrayBase(d *kernel.Driver, uLevel uintptr) uintptr {
	return d.Readvm(uLevel+offset.ActorArrayOffset, 8)
}

func GetActorArrayCount(d *kernel.Driver, uLevel uintptr) int {
	return d.ReadvmInt(uLevel + offset.ActorCountOffset)
}

func GetGameStateBase(d *kernel.Driver, uworld uintptr) uintptr {
	gameStateBase := d.Readvm(uworld+offset.GameStateOffset, 8)
	if kernel.IsGuarded(gameStateBase) {
		gameStateBase = kernel.WardedTo(d.Guardedregion, gameStateBase)
	}
	return gameStateBase
}
