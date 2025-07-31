package main

/*
Just as a reminder, I hate GO. Not my cup of tea in the slighest.
I don't hate the errors, the weird syntax but The fkin structs and interfaces of GO.
Why the fk can't google just create a basic OOP syntax for its language.
Simple class keyword, Simple inheritance, interface and so on.
Fkin disappointment of a language.

Well that is just a personal review. I may have over-reacted a little bit since This is my first real project in GO and I was not ready for this brother. Have fun reviewing my code.

Name : GO
Developer : Google
Co-Developer : C ka developer
*/

// Imports
import (
	GAME "Game/MainGame"
	"Game/MainGame/SpriteAndTextures/Objects"
	"Game/MainGame/Utility"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/ebitick"
)

// Variables
var GameInstance *GAME.Game = &GAME.Game{}

// Functions
func ConfigureEbiten() {
	ebiten.SetWindowTitle("Avengers")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSizeLimits(320, 240, -1, -1)
	// ebiten.MaximizeWindow()
	ebiten.SetTPS(60)

	// This hides the cursor and stops it from going out of the screen
	// ebiten.SetCursorMode(ebiten.CursorModeCaptured)
}

func GenerateTeleporters(page Objects.Page, tele_name string, GateNum []bool) []Objects.Teleporter {
	var result []Objects.Teleporter = make([]Objects.Teleporter, 0)

	for Index, Value := range GateNum {
		if Value {
			var tp Objects.Teleporter
			switch Index {
			case 0:
				tp = Objects.CreateTeleporterGridAligned(page.Name, "UP:"+tele_name, 6, 0, 50, 50, 2)
			case 1:
				tp = Objects.CreateTeleporterGridAligned(page.Name, "DOWN:"+tele_name, 6, 8, 50, 50, 0)
			case 2:
				tp = Objects.CreateTeleporterGridAligned(page.Name, "LEFT:"+tele_name, 0, 4, 50, 50, 3)
			case 3:
				tp = Objects.CreateTeleporterGridAligned(page.Name, "RIGHT:"+tele_name, 11, 4, 50, 50, 1)
			}
			result = append(result, tp)
		}
	}

	return result
}

func InitializeAndLoadAssets(GameInstance *GAME.Game) {
	GameInstance.TimerSystem = ebitick.NewTimerSystem() // Init TimerSystem
	GameInstance.WorldMap = Objects.CreateGlobe()

	// Pages
	HomePage, GateDir := Objects.CreatePage("HomePage", "Assets/Pages/test.bin")
	HomePage.MainPlayer = Objects.CreatePlayer("Assets/ninja.png", []float64{100, 100}, []float64{1, 1}, Utility.BoundingRectangle{X: 40, Y: 35, Width: 640 - 40, Height: 480 - 35})
	HomePageTeleporters := GenerateTeleporters(HomePage, "H2", GateDir)
	GameInstance.WorldMap.AddPage(HomePage, HomePageTeleporters...)

	// MainPlayers
	GameInstance.CurrentPage = GameInstance.WorldMap.Pages[0]
}

func LoadTimers(GameInstance *GAME.Game) {
}

func LoadGame(GameInstance *GAME.Game) {
	if err := ebiten.RunGame(GameInstance); err != nil {
		log.Fatalln(err)
	}
}

// Entry Point
func main() {
	ConfigureEbiten()
	InitializeAndLoadAssets(GameInstance)
	LoadTimers(GameInstance)
	LoadGame(GameInstance)
}
