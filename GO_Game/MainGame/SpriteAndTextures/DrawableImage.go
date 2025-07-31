package SpriteAndTextures

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type DrawableImage struct {
	CurrentImage *ebiten.Image
	OtherImages  map[string]*ebiten.Image
	DrawOptions  *ebiten.DrawImageOptions
}

func ConstructImage(s MetaData) DrawableImage {
	var result DrawableImage
	Surface, _, err := ebitenutil.NewImageFromFile(s.current_image_path)
	if err != nil {
		log.Fatalln(err)
	}

	result.CurrentImage = Surface
	result.DrawOptions = &ebiten.DrawImageOptions{}

	result.DrawOptions.GeoM.Translate(s.position[0], s.position[1])
	result.DrawOptions.GeoM.Scale(s.size[0], s.size[1])

	return result
}

func (d *DrawableImage) LoadAnimationTexturesFromSpriteSheet(sheet SpriteSheet) {
	d.OtherImages = make(map[string]*ebiten.Image, len(sheet.AllTextures))
	for x := 0; x < len(sheet.AllTextures); x++ {
		d.OtherImages[fmt.Sprintf("Image%d", x)] = sheet.AllTextures[x]
	}
}

func (d *DrawableImage) Update(s MetaData) {
	d.DrawOptions.GeoM = ebiten.GeoM{}
	d.DrawOptions.GeoM.Translate(s.position[0], s.position[1])
	d.DrawOptions.GeoM.Scale(s.size[0], s.size[1])
}

func (d *DrawableImage) ChangeTexture(key string) {
	image, ok := d.OtherImages[key]
	if !ok {
		log.Fatalln("Animation Error : { MetaData doesn't exist. }")
	}

	d.CurrentImage = image
}

func (d *DrawableImage) DrawOn(screen *ebiten.Image) {
	screen.DrawImage(d.CurrentImage, d.DrawOptions)
}
