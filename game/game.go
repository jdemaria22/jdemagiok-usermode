package game

import (
	"image/color"
	"jdemagiok-usermode/geometry"
	"jdemagiok-usermode/kernel"
	"jdemagiok-usermode/offset"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (game *SGame) Draw(d *kernel.Driver, screen *ebiten.Image) {
	actorArray := game.World.PersistanceLevel.ActorArray
	for _, actorArray := range actorArray {
		enemyPawn := getEnemyPawn(d, actorArray.Pawn.Pointer)
		if enemyPawn.Health > 0 || enemyPawn.BIsDormant {
			enemyPawn = addAditionalInfoToEnemyPawn(d, enemyPawn, game)
			vector.StrokeLine(screen, enemyPawn.RelativeLocationProjected.X, enemyPawn.RelativeLocationProjected.Y, enemyPawn.RelativeLocationProjected.X+50, enemyPawn.RelativeLocationProjected.Y, 1, color.RGBA{255, 0, 0, 255}, false)
			vector.StrokeLine(screen, enemyPawn.RelativeLocationProjected.X, enemyPawn.RelativeLocationProjected.Y+30, enemyPawn.RelativeLocationProjected.X+50, enemyPawn.RelativeLocationProjected.Y+30, 1, color.RGBA{255, 0, 0, 255}, false)
			vector.StrokeLine(screen, enemyPawn.RelativeLocationProjected.X, enemyPawn.RelativeLocationProjected.Y, enemyPawn.RelativeLocationProjected.X, enemyPawn.RelativeLocationProjected.Y+30, 1, color.RGBA{255, 0, 0, 0}, false)
			vector.StrokeLine(screen, enemyPawn.RelativeLocationProjected.X+50, enemyPawn.RelativeLocationProjected.Y, enemyPawn.RelativeLocationProjected.X+50, enemyPawn.RelativeLocationProjected.Y+30, 1, color.RGBA{255, 0, 0, 255}, false)
		}
	}
}

func (game *SGame) Loop(d *kernel.Driver) {
	actorArray := game.World.PersistanceLevel.ActorArray
	for _, actorArray := range actorArray {
		enemyPawn := getEnemyPawn(d, actorArray.Pawn.Pointer)
		if enemyPawn.Health > 0 || enemyPawn.BIsDormant {
			enemyPawn = addAditionalInfoToEnemyPawn(d, enemyPawn, game)
		}
	}
}

func getEnemyPawn(d *kernel.Driver, pawnPointer uintptr) SEnemyPawn {
	enemyPawn := SEnemyPawn{}
	enemyPawn.Pointer = pawnPointer
	enemyPawn.TeamID = getTeamID(d, enemyPawn.Pointer)
	enemyPawn.UniqueID = d.ReadvmInt(enemyPawn.Pointer + offset.ActorIDOffset)
	enemyPawn.FNameID = d.ReadvmInt(enemyPawn.Pointer + offset.FnameIDOffset)
	enemyPawn.RelativeLocation = getRelativePosition(d, enemyPawn.Pointer)
	enemyPawn.BIsDormant = d.ReadvmBool(enemyPawn.Pointer + offset.DormantOffset)
	enemyPawn.Health = getHealth(d, enemyPawn.Pointer)
	return enemyPawn
}

func addAditionalInfoToEnemyPawn(d *kernel.Driver, enemyPawn SEnemyPawn, game *SGame) SEnemyPawn {
	enemyPawn.SkeletalMesh = d.Read(enemyPawn.Pointer + offset.CurrentMeshOffset)
	enemyPawn.RelativeLocationProjected = geometry.ProjectWorldToScreen(enemyPawn.RelativeLocation, game.World.GameInstance.LocalPlayer.PlayerController.MinimalViewInfo)
	// fmt.Printf("Projected Screen Location: (%.2f, %.2f)\n", enemyPawn.RelativeLocationProjected.X, enemyPawn.RelativeLocationProjected.Y)
	return enemyPawn
}
