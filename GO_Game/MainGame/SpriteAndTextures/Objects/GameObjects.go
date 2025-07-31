package Objects

import (
	"Game/MainGame/SpriteAndTextures"
	"Game/MainGame/Utility"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Interface
type Moveable interface {
	CheckMovement()
}

type Animated interface {
	Animate()
}

type Updated interface {
	Update()
}

type Drawn interface {
	Draw()
}

// Objects
type GameObject struct {
	DrawObject *SpriteAndTextures.DrawableImage
	Data       *SpriteAndTextures.MetaData
}

func ContructGameObject(DrawObject *SpriteAndTextures.DrawableImage, Data *SpriteAndTextures.MetaData) GameObject {
	return GameObject{DrawObject, Data}
}

func (object *GameObject) DefaultDraw(screen *ebiten.Image) {
	screen.DrawImage(object.DrawObject.CurrentImage, object.DrawObject.DrawOptions)
}

type MoveableGameObject struct {
	GameObject
	Moveable
	Animated
	Updated
}

type AnimatedOnlyGameObject struct {
	GameObject
	Animated
	Updated
}

type StaticGameObject struct {
	GameObject
	Drawn
}

type Page struct {
	Name             string
	Map              []byte
	background_image *ebiten.Image
	MainPlayer       *Player

	AnimatedObjects map[string]*AnimatedOnlyGameObject
	StaticObjects   map[string]*StaticGameObject
	Enemies         map[string]*MoveableGameObject
}

func CreatePage(Name string, path string) (Page, []bool) {
	var result Page
	var result2 []bool = make([]bool, 4)

	result.Name = Name
	mappp, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	result.Map = mappp

	result.AnimatedObjects = make(map[string]*AnimatedOnlyGameObject)
	result.StaticObjects = make(map[string]*StaticGameObject)
	result.Enemies = make(map[string]*MoveableGameObject)

	result.MainPlayer = nil

	width, height := ebiten.WindowSize()
	parentImage := ebiten.NewImage(width, height)

	tileSize := 50
	tilesPerRow := width / tileSize

	var OffsetX float32 = float32((width % tileSize) / 2)
	var OffsetY float32 = float32((height % tileSize) / 2)
	var ByteOffset int = 0

	for i := 0; i < 4; i++ {
		result2[i] = result.Map[i] == 1
		ByteOffset++
	}

	for i := ByteOffset; i < len(result.Map); i++ {
		var x float64 = float64(((i - ByteOffset) % tilesPerRow) * tileSize)
		var y float64 = float64(((i - ByteOffset) / tilesPerRow) * tileSize)

		var Surface *ebiten.Image
		var err error
		switch result.Map[i] {
		case 0:
			Surface, _, err = ebitenutil.NewImageFromFile("Assets/TileSet/Ground.png")
		case 1:
			Surface, _, err = ebitenutil.NewImageFromFile("Assets/TileSet/Corner1.png")
		case 2:
			Surface, _, err = ebitenutil.NewImageFromFile("Assets/TileSet/Corner2.png")
		case 3:
			Surface, _, err = ebitenutil.NewImageFromFile("Assets/TileSet/Corner3.png")
		case 4:
			Surface, _, err = ebitenutil.NewImageFromFile("Assets/TileSet/Corner4.png")
		case 5:
			Surface, _, err = ebitenutil.NewImageFromFile("Assets/TileSet/Wall1.png")
		case 6:
			Surface, _, err = ebitenutil.NewImageFromFile("Assets/TileSet/Wall2.png")
		case 7:
			Surface, _, err = ebitenutil.NewImageFromFile("Assets/TileSet/Wall3.png")
		case 8:
			Surface, _, err = ebitenutil.NewImageFromFile("Assets/TileSet/Wall4.png")
		default:
			Surface, _, err = ebitenutil.NewImageFromFile("Assets/TileSet/Ground.png")
		}

		if err != nil {
			log.Fatalln(err)
		}
		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(OffsetX)+x, float64(OffsetY)+y)
		parentImage.DrawImage(Surface, &op)
	}
	result.background_image = parentImage

	return result, result2
}

func (p *Page) Update() {
	if p.MainPlayer != nil {
		p.MainPlayer.Update()
	}
}

func (p *Page) DrawOn(screen *ebiten.Image) {
	screen.DrawImage(p.background_image, &ebiten.DrawImageOptions{})
	if p.MainPlayer != nil {
		screen.DrawImage(p.MainPlayer.DrawObject.CurrentImage, p.MainPlayer.DrawObject.DrawOptions)
	}
}

type Teleporter struct {
	GameObject
	TransporterName string
	RoomName        string
}

func CreateTeleporter(RoomName string, TransporterName string, X, Y float64, Width, Height, Orientation int) Teleporter {
	// Create gradient image
	img := ebiten.NewImage(Width, Height)

	switch Orientation {
	case 0: // top to bottom
		for y := 0; y < Height; y++ {
			alpha := uint8(float64(y) / float64(Height) * 255)
			col := color.RGBA{0, 0, 0, alpha}
			for x := 0; x < Width; x++ {
				img.Set(x, y, col)
			}
		}
	case 1: // left to right
		for x := 0; x < Width; x++ {
			alpha := uint8(float64(x) / float64(Width) * 255)
			col := color.RGBA{0, 0, 0, alpha}
			for y := 0; y < Height; y++ {
				img.Set(x, y, col)
			}
		}
	case 2: // bottom to top
		for y := 0; y < Height; y++ {
			alpha := uint8(float64(Height-y-1) / float64(Height) * 255)
			col := color.RGBA{0, 0, 0, alpha}
			for x := 0; x < Width; x++ {
				img.Set(x, y, col)
			}
		}
	case 3: // right to left
		for x := 0; x < Width; x++ {
			alpha := uint8(float64(Width-x-1) / float64(Width) * 255)
			col := color.RGBA{0, 0, 0, alpha}
			for y := 0; y < Height; y++ {
				img.Set(x, y, col)
			}
		}
	default:
		// fallback to top to bottom
		for y := 0; y < Height; y++ {
			alpha := uint8(float64(y) / float64(Height) * 255)
			col := color.RGBA{0, 0, 0, alpha}
			for x := 0; x < Width; x++ {
				img.Set(x, y, col)
			}
		}
	}

	// Create DrawableImage
	drawable := &SpriteAndTextures.DrawableImage{
		CurrentImage: img,
		OtherImages:  make(map[string]*ebiten.Image),
		DrawOptions:  &ebiten.DrawImageOptions{},
	}

	// Set DrawOptions position
	// Translate to position
	drawable.DrawOptions.GeoM.Translate(X, Y)

	// Create MetaData
	metadata := SpriteAndTextures.CreateStaticMetaData(
		TransporterName, // name
		100,             // health
		100,             // hitpoints
		0,               // hit_damage
		[]float64{X, Y}, // position
		[]float64{float64(Width), float64(Height)}, // scale
		"", // image_path (can be added if needed)
	)

	// Assemble Teleporter
	return Teleporter{
		GameObject: GameObject{
			DrawObject: drawable,
			Data:       &metadata,
		},
		TransporterName: TransporterName,
		RoomName:        RoomName,
	}
}

func CreateTeleporterGridAligned(RoomName string, TransporterName string, GridX, GridY float64, Width, Height, Orientation int) Teleporter {
	Position := Utility.GridToMap([]float64{GridX, GridY})
	return CreateTeleporter(RoomName, TransporterName, Position[0], Position[1], Width, Height, Orientation)
}

type Globe struct {
	Pages []Page
	Gates map[string][]Teleporter
}

func CreateGlobe() Globe {
	return Globe{make([]Page, 0), make(map[string][]Teleporter)}
}

func (g *Globe) AddPage(page Page, Tps ...Teleporter) {
	g.Pages = append(g.Pages, page)
	g.Gates[page.Name] = make([]Teleporter, 0)
	g.Gates[page.Name] = append(g.Gates[page.Name], Tps...)
}

func (g *Globe) DrawTeleporters(page Page, screen *ebiten.Image) {
	for _, TP := range g.Gates[page.Name] {
		TP.DefaultDraw(screen)
	}
}
