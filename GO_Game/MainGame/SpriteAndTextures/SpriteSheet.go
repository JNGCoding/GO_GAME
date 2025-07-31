package SpriteAndTextures

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SpriteSheet struct {
	Path        string
	SpriteSize  []int
	AllTextures []*ebiten.Image
}

func CreateSpriteSheet(Path string, SpriteSize []int, SpriteNumRow, SpriteNumColumn int) SpriteSheet {
	var result SpriteSheet

	result.Path = Path
	result.SpriteSize = SpriteSize
	result.AllTextures = make([]*ebiten.Image, 0)

	parentImage, _, err := ebitenutil.NewImageFromFile(Path)
	if err != nil {
		log.Fatalln(err)
	}

	for row := 0; row < SpriteNumRow; row++ {
		for col := 0; col < SpriteNumColumn; col++ {
			x := col * SpriteSize[0]
			y := row * SpriteSize[1]

			rect := image.Rect(x, y, x+SpriteSize[0], y+SpriteSize[1])
			texture := parentImage.SubImage(rect).(*ebiten.Image)
			result.AllTextures = append(result.AllTextures, texture)
		}
	}

	return result
}
