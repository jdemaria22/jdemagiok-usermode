package game

import (
	"jdemagiok-usermode/kernel"
	"jdemagiok-usermode/offset"
)

func GetUWorld(d *kernel.Driver) uintptr {
	uworld := d.Readvm(d.Guardedregion+offset.WORLD, 8)

	var uworldOffset uintptr

	if uworld > 0x10000000000 {
		uworldOffset = uworld - 0x10000000000
	} else {
		uworldOffset = uworld - 0x8000000000
	}

	return d.Guardedregion + uworldOffset
}

func GetGameInstance(d *kernel.Driver, uworld uintptr) uintptr {
	gameinstance := d.Readvm(uworld+offset.GAMEINSTANCE, 8)

	if kernel.IsGuarded(gameinstance) {
		return kernel.WardedTo(d.Guardedregion, gameinstance)
	}

	return gameinstance
}

func GetULocalPlayerArray(d *kernel.Driver, gameinstance uintptr) uintptr {
	uLocalPlayerArray := d.Readvm(gameinstance+offset.LOCALPLAYERARRAY, 8)

	if kernel.IsGuarded(uLocalPlayerArray) {
		return kernel.WardedTo(d.Guardedregion, uLocalPlayerArray)
	}

	return uLocalPlayerArray
}

func GetULocalPlayer(d *kernel.Driver, uLocalPlayerArray uintptr) uintptr {
	uLocalPlayer := d.Readvm(uLocalPlayerArray, 8)

	if kernel.IsGuarded(uLocalPlayer) {
		return kernel.WardedTo(d.Guardedregion, uLocalPlayer)
	}

	return uLocalPlayer
}

func GetAPlayerControllerPtr(d *kernel.Driver, uLocalPlayer uintptr) uintptr {
	aPlayerControllerPtr := d.Readvm(uLocalPlayer+offset.PLAYERCONTROLLERPTR, 8)

	if kernel.IsGuarded(aPlayerControllerPtr) {
		return kernel.WardedTo(d.Guardedregion, aPlayerControllerPtr)
	}

	return aPlayerControllerPtr
}

func GetAPawn(d *kernel.Driver, aPlayerControllerPtr uintptr) uintptr {
	aPawn := d.Readvm(aPlayerControllerPtr+offset.PAWN, 8)

	if kernel.IsGuarded(aPawn) {
		return kernel.WardedTo(d.Guardedregion, aPawn)
	}

	return aPawn
}

func GetDamageHandler(d *kernel.Driver, aPawn uintptr) uintptr {
	damageHandler := d.Readvm(aPawn+offset.DAMAGEHANDLER, 8)

	if kernel.IsGuarded(damageHandler) {
		return kernel.WardedTo(d.Guardedregion, damageHandler)
	}

	return damageHandler
}
