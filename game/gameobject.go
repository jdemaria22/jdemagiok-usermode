package game

import (
	"jdemagiok-usermode/geometry"
)

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
	MinimalViewInfo     geometry.FMinimalViewInfo
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

type SEnemyPawn struct {
	Pointer                   uintptr
	TeamID                    int
	UniqueID                  int
	FNameID                   int
	RelativeLocation          geometry.FVector
	RelativeDistance          float32
	BIsDormant                bool
	Health                    float32
	SkeletalMesh              uintptr
	RelativeLocationProjected geometry.FVector
	HeadBone                  geometry.FVector
	HeadBoneProjected         geometry.FVector
	RootBone                  geometry.FVector
	RootBoneProjected         geometry.FVector
	RootBoneProjected2        geometry.FVector
	Distance                  float32
	BoxHeight                 float32
	BoxWidth                  float32
}
type SActor struct {
	Pawn SEnemyPawn
}
