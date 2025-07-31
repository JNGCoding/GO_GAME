package Utility

import (
	Texture "Game/MainGame/SpriteAndTextures"
)

type BoundingRectangle struct {
	X, Y, Width, Height float64
}

func ForceConfineSprite(d *Texture.DrawableImage, ds *Texture.MetaData, boundRect BoundingRectangle) {
	PlayerPosition := ds.GetPosition()
	PlayerAbsoluteSize := d.CurrentImage.Bounds().Size()
	if PlayerPosition[0] > boundRect.Width-float64(PlayerAbsoluteSize.X) {
		PlayerPosition[0] = boundRect.Width - float64(PlayerAbsoluteSize.X)
	} else if PlayerPosition[0] < boundRect.X {
		PlayerPosition[0] = boundRect.X
	}

	if PlayerPosition[1] > boundRect.Height-float64(PlayerAbsoluteSize.Y) {
		PlayerPosition[1] = boundRect.Height - float64(PlayerAbsoluteSize.Y)
	} else if PlayerPosition[1] < boundRect.Y {
		PlayerPosition[1] = boundRect.Y
	}

	ds.MovePosition(PlayerPosition)
}

func MapToGrid(Position []float64) []float64 {
	if len(Position) < 2 {
		return []float64{0, 0}
	}
	tileSize := 50
	gridX := float64(int(Position[0]) / tileSize)
	gridY := float64(int(Position[1]) / tileSize)
	return []float64{gridX, gridY}
}

func GridToMap(Position []float64) []float64 {
	if len(Position) < 2 {
		return []float64{0, 0}
	}
	tileSize := 50
	mapX := Position[0] * float64(tileSize)
	mapY := Position[1] * float64(tileSize)
	return []float64{mapX + 20, mapY + 15}
}
