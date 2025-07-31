package Objects

import (
	"Game/MainGame/SpriteAndTextures"
	"Game/MainGame/Utility"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	MoveableGameObject
	bounding_rect Utility.BoundingRectangle
	Action        string
	dx, dy        float64
}

func CreatePlayer(spritesheet_path string, starting_pos, object_scale []float64, boundRect Utility.BoundingRectangle) *Player {
	PlayerData := SpriteAndTextures.CreateMetaData("Player", true, 10, 100, 1000, 50, starting_pos, object_scale, spritesheet_path, true)
	DrawObject := SpriteAndTextures.ConstructImage(PlayerData)

	PlayerSpriteSheet := SpriteAndTextures.CreateSpriteSheet(spritesheet_path, []int{32, 32}, 7, 4)
	DrawObject.LoadAnimationTexturesFromSpriteSheet(PlayerSpriteSheet)

	result := Player{
		MoveableGameObject: MoveableGameObject{
			GameObject: ContructGameObject(&DrawObject, &PlayerData),
		},
		bounding_rect: boundRect,
		Action:        "IDLE",
	}

	return &result
}

func (p *Player) CheckMovement() {
	p.dx = 0
	p.dy = 0

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.dx = float64(p.Data.GetSpeed())
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.dx = -float64(p.Data.GetSpeed())
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.dy = -float64(p.Data.GetSpeed())
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.dy = float64(p.Data.GetSpeed())
	}

	if p.dx == 0 && p.dy == 0 {
		p.Action = "IDLE"
	} else {
		p.Action = "NOT IDLE"
	}

	Utility.ForceConfineSprite(p.DrawObject, p.Data, p.bounding_rect)
}

func (p *Player) Animate() {
	if p.Action == "IDLE" {
		p.DrawObject.ChangeTexture("Image0")
	} else {
		p.DrawObject.ChangeTexture("Image1")
	}
}

func (p *Player) Update() {
	p.CheckMovement()
	p.Data.MoveBy(p.dx, p.dy)
	p.Animate()
	p.DrawObject.Update(*p.Data)
}
