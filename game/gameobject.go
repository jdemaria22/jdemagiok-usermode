package game

import "jdemagiok-usermode/geometry"

type SGame struct {
	World SWorld
}

type SWorld struct {
	Pointer          uintptr
	GameInstance     SGameInstance
	PersistanceLevel SPersistanceLevel
}

type SGameInstance struct {
	Pointer          uintptr
	LocalPlayerArray uintptr
	LocalPlayer      SLocalPlayer
}

type SLocalPlayer struct {
	Pointer          uintptr
	PlayerController SPlayerController
}

type SPlayerController struct {
	Pointer             uintptr
	PlayerCameraManager uintptr
	AHUD                uintptr
	Pawn                SPawn
}

type SPersistanceLevel struct {
	Pointer    uintptr
	ActorArray []SActor
}

type SPawn struct {
	Pointer          uintptr
	TeamID           int
	UniqueID         int
	FNameID          int
	RelativeLocation geometry.FVector
	BIsDormant       bool
	Health           float32
}

type SActor struct {
	Pawn SPawn
}
