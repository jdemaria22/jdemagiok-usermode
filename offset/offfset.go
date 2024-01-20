package offset

// const (
// 	GAMEINSTANCE        = 0x1A0
// 	LOCALPLAYERARRAY    = 0x40
// 	PLAYERCONTROLLERPTR = 0x38
// 	PAWN                = 0x468
// 	DAMAGEHANDLER       = 0x9F0
// 	HEALTH              = 0x1B0
// )

const (
	World                           uintptr = 0x60
	FnamePoolOffset                 uintptr = 0xA23E540
	PersistentLevelOffset           uintptr = 0x38
	OwningGameInstanceOffset        uintptr = 0x1A0
	GameStateOffset                 uintptr = 0x140
	LevelsOffset                    uintptr = 0x158
	LocalPlayersOffset              uintptr = 0x40
	ActorArrayOffset                uintptr = 0xA0
	ActorCountOffset                uintptr = 0xA8
	PlayerControllerOffset          uintptr = 0x38
	AcknowledgedPawnOffset          uintptr = 0x468
	PlayerCameraOffset              uintptr = 0x480
	ControlRotationOffset           uintptr = 0x448
	RootComponentOffset             uintptr = 0x238
	DamageHandlerOffset             uintptr = 0x9F0
	ActorIDOffset                   uintptr = 0x38
	FnameIDOffset                   uintptr = 0x18
	DormantOffset                   uintptr = 0x100
	PlayerStateOffset               uintptr = 0x3f8
	CurrentMeshOffset               uintptr = 0x438
	InventoryOffset                 uintptr = 0x990
	OutlineComponentOffset          uintptr = 0x11D0
	PortraitMapOffset               uintptr = 0x1188
	CharacterMapOffset              uintptr = 0x1180
	CurrentEquippableOffset         uintptr = 0x248
	LocalObserverOffset             uintptr = 0x530
	IsVisibleOffset                 uintptr = 0x501
	ComponentToWorldOffset          uintptr = 0x250
	BoneArrayOffset                 uintptr = 0x5D8
	BoneCountOffset                 uintptr = 0x5E0
	LastSubmitTimeOffset            uintptr = 0x380
	LastRenderTimeOffset            uintptr = 0x384
	OutlineModeOffset               uintptr = 0x330
	AttachChildrenOffset            uintptr = 0x110
	AttachChildrenCountOffset       uintptr = 0x118
	TeamComponentOffset             uintptr = 0x630
	TeamIDOffset                    uintptr = 0xf8
	CurrentHealthOffset             uintptr = 0x1B0
	MaxLifeOffset                   uintptr = 0x1B4
	RelativeLocationOffset          uintptr = 0x164
	RelativeRotationOffset          uintptr = 0x170
	DefaultFOVOffset                uintptr = 0x3EC
	CameraCacheOffset               uintptr = 0x4d0
	POVOffset                       uintptr = 0x10
	LocationOffset                  uintptr = 0x0
	RotationOffset                  uintptr = 0x0C
	CurrentFOVOffset                uintptr = 0x18
	EnemyOutlineOffset              uintptr = 0x2B1
	BoneArrayCacheOffset            uintptr = 0x5C8
	CurrentDefuseSectionOffset      uintptr = 0x524
	MagazineAmmoOffset              uintptr = 0xFF0
	AuthResourceAmountOffset        uintptr = 0x120
	MaxAmmoOffset                   uintptr = 0x140
	SpikeTimerOffset                uintptr = 0x4e4
	MyHUDOffset                     uintptr = 0x478
	HPOffset                        uintptr = 0x134
	MaxHPOffset                     uintptr = 0x138
	DamageTypeOffset                uintptr = 0x128
	DamageSectionsOffset            uintptr = 0x198
	CurrentEquippableVFXStateOffset uintptr = 0xcf8
	FresnelIntensityOffset          uintptr = 0x1c8
	CacheCamOffset                  uintptr = 0x1FB0
)
