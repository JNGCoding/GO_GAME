package game

import (
	"Game/MainGame/SpriteAndTextures"
	"Game/MainGame/SpriteAndTextures/Objects"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func OverlapPercentage(x1, y1 float64, img1 *ebiten.Image, x2, y2 float64, img2 *ebiten.Image) float64 {
	w1, h1 := img1.Bounds().Dx(), img1.Bounds().Dy()
	w2, h2 := img2.Bounds().Dx(), img2.Bounds().Dy()

	// Determine overlap region
	left := int(math.Max(x1, x2))
	right := int(math.Min(x1+float64(w1), x2+float64(w2)))
	top := int(math.Max(y1, y2))
	bottom := int(math.Min(y1+float64(h1), y2+float64(h2)))

	if right <= left || bottom <= top {
		return 0.0
	}

	// Count overlapping pixels with non-zero alpha
	overlapCount := 0
	totalOverlapArea := (right - left) * (bottom - top)

	for y := top; y < bottom; y++ {
		for x := left; x < right; x++ {
			ax := x - int(x1)
			ay := y - int(y1)
			bx := x - int(x2)
			by := y - int(y2)

			if ax < 0 || ay < 0 || bx < 0 || by < 0 || ax >= w1 || ay >= h1 || bx >= w2 || by >= h2 {
				continue
			}

			c1 := img1.At(ax, ay).(color.RGBA)
			c2 := img2.At(bx, by).(color.RGBA)

			if c1.A > 0 && c2.A > 0 {
				overlapCount++
			}
		}
	}

	return float64(overlapCount) / float64(totalOverlapArea) * 100.0
}

func BoundingBoxCollision(m1 *SpriteAndTextures.MetaData, img1 *SpriteAndTextures.DrawableImage, m2 *SpriteAndTextures.MetaData, img2 *SpriteAndTextures.DrawableImage) bool {
	w1, h1 := img1.CurrentImage.Bounds().Dx(), img1.CurrentImage.Bounds().Dy()
	w2, h2 := img2.CurrentImage.Bounds().Dx(), img2.CurrentImage.Bounds().Dy()
	x1, y1, x2, y2 := m1.GetPosition()[0], m1.GetPosition()[1], m2.GetPosition()[0], m2.GetPosition()[1]

	return x1 < x2+float64(w2) &&
		x1+float64(w1) > x2 &&
		y1 < y2+float64(h2) &&
		y1+float64(h1) > y2
}

func PlayerPresentinGridSquare(player Objects.Player, GridX, GridY float64) {

}
