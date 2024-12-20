package game

import (
	"image/color"
	"jdemagiok-usermode/geometry"
	"jdemagiok-usermode/kernel"
	"jdemagiok-usermode/offset"
	"math"
	"unsafe"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (game *SGame) Draw(d *kernel.Driver, screen *ebiten.Image) {
	actorArray := game.World.PersistanceLevel.ActorArray
	for _, actorArray := range actorArray {
		enemyPawn := getEnemyPawn(d, actorArray.Pawn.Pointer)
		if enemyPawn.Health > 0 && enemyPawn.BIsDormant {
			enemyPawn = addAditionalInfoToEnemyPawn(d, enemyPawn, game)
			drawBox(screen, enemyPawn.RelativeLocationProjected.X-enemyPawn.BoxWidth/2, enemyPawn.RelativeLocationProjected.Y-enemyPawn.BoxHeight/2, enemyPawn.BoxHeight, enemyPawn.BoxWidth, int(enemyPawn.Health))
		}
	}
}

func drawBox(screen *ebiten.Image, x, y, height, width float32, health int) {
	vector.StrokeLine(screen, x, y, x+width, y, 2, color.RGBA{255, 0, 0, 255}, false)
	vector.StrokeLine(screen, x, y+height, x+width, y+height, 2, color.RGBA{255, 0, 0, 255}, false)
	vector.StrokeLine(screen, x, y, x, y+height, 2, color.RGBA{255, 0, 0, 0}, false)
	vector.StrokeLine(screen, x+width, y, x+width, y+height, 2, color.RGBA{255, 0, 0, 255}, false)

	percentage := float32(health) / 100.0
	lineHeight := height * percentage
	lineX := x + width
	lineY := y + height - lineHeight
	vector.StrokeLine(screen, lineX, lineY, lineX, y+height, 2, color.RGBA{0, 255, 0, 255}, false)
}

func getEnemyPawn(d *kernel.Driver, pawnPointer uintptr) SEnemyPawn {
	enemyPawn := SEnemyPawn{}
	enemyPawn.Pointer = pawnPointer
	enemyPawn.UniqueID = d.ReadvmInt(enemyPawn.Pointer + offset.ActorIDOffset)
	enemyPawn.FNameID = d.ReadvmInt(enemyPawn.Pointer + offset.FnameIDOffset)
	enemyPawn.RelativeLocation = getRelativePosition(d, enemyPawn.Pointer)
	enemyPawn.BIsDormant = d.ReadvmBool(enemyPawn.Pointer + offset.DormantOffset)
	enemyPawn.Health = getHealth(d, enemyPawn.Pointer)
	return enemyPawn
}

func getEntityBone(d *kernel.Driver, enemyPawn SEnemyPawn, index uintptr) geometry.FVector {
	array := d.Read(enemyPawn.SkeletalMesh + offset.BoneArrayOffset)
	if array == 0 {
		array = d.Read(enemyPawn.SkeletalMesh + offset.BoneArrayCacheOffset)
	}
	var transform geometry.FTransform
	val := unsafe.Sizeof(transform)
	bone := d.ReadvmFTransform(array + (index * val))
	componentToWorld := d.ReadvmFTransform(enemyPawn.SkeletalMesh + offset.ComponentToWorldOffset)
	matrix := geometry.MatrixMultiplication(bone.ToMatrixWithScale(), componentToWorld.ToMatrixWithScale())
	result := geometry.FVector{
		X: matrix.V_41,
		Y: matrix.V_42,
		Z: matrix.V_43,
	}
	return result
}

func addAditionalInfoToEnemyPawn(d *kernel.Driver, enemyPawn SEnemyPawn, game *SGame) SEnemyPawn {
	enemyPawn.SkeletalMesh = d.Read(enemyPawn.Pointer + offset.CurrentMeshOffset)
	enemyPawn.RelativeLocationProjected = geometry.ProjectWorldToScreen(enemyPawn.RelativeLocation, game.World.GameInstance.LocalPlayer.PlayerController.MinimalViewInfo)
	enemyPawn.RelativePosition = enemyPawn.RelativeLocation.Subtract(game.World.GameInstance.LocalPlayer.PlayerController.MinimalViewInfo.Location)
	enemyPawn.RelativeDistance = enemyPawn.RelativePosition.Length() / 10000 * 2
	enemyPawn.HeadBone = getEntityBone(d, enemyPawn, 8)
	enemyPawn.HeadBoneProjected = geometry.ProjectWorldToScreen(enemyPawn.HeadBone, game.World.GameInstance.LocalPlayer.PlayerController.MinimalViewInfo)
	enemyPawn.RootBone = getEntityBone(d, enemyPawn, 0)
	enemyPawn.RootBoneProjected = geometry.ProjectWorldToScreen(enemyPawn.RootBone, game.World.GameInstance.LocalPlayer.PlayerController.MinimalViewInfo)
	rootBone2 := geometry.FVector{X: enemyPawn.RootBone.X, Y: enemyPawn.RootBone.Y, Z: enemyPawn.RootBone.Z - 15}
	enemyPawn.RootBoneProjected2 = geometry.ProjectWorldToScreen(rootBone2, game.World.GameInstance.LocalPlayer.PlayerController.MinimalViewInfo)
	enemyPawn.Distance = game.World.GameInstance.LocalPlayer.PlayerController.Pawn.RelativeLocation.Distance(enemyPawn.RelativePosition)
	enemyPawn.BoxHeight = float32(math.Abs(float64(enemyPawn.HeadBoneProjected.Y) - float64(enemyPawn.RootBoneProjected.Y)))
	enemyPawn.BoxWidth = enemyPawn.BoxHeight * 0.40
	return enemyPawn
}
