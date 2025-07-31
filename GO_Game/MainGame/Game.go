package game

import (
	"Game/MainGame/SpriteAndTextures/Objects"
	"fmt"
	"image/color"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/ebitick"
)

type Game struct {
	WorldMap            Objects.Globe
	CurrentPage         Objects.Page
	TimerSystem         *ebitick.TimerSystem
	ThreadSyncWaitGroup sync.WaitGroup
}

func (g *Game) Update() error {
	ebiten.SetWindowTitle(fmt.Sprintf("Avengers, FPS - %f | TPS - %f", ebiten.ActualFPS(), ebiten.ActualTPS()))

	g.ThreadSyncWaitGroup.Add(1)
	go func() {
		defer g.ThreadSyncWaitGroup.Done()
		for _, TP := range g.WorldMap.Gates[g.CurrentPage.Name] {
			if BoundingBoxCollision(g.CurrentPage.MainPlayer.Data, g.CurrentPage.MainPlayer.DrawObject, TP.Data, TP.DrawObject) {
				fmt.Printf("Player Collided with TP, Name - %s\n", TP.TransporterName)
			}
		}
	}()

	g.CurrentPage.Update()
	g.TimerSystem.Update()

	g.ThreadSyncWaitGroup.Wait()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	g.CurrentPage.DrawOn(screen)
	g.WorldMap.DrawTeleporters(g.CurrentPage, screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Player Pos - (%.2f, %.2f)", g.CurrentPage.MainPlayer.Data.GetPosition()[0], g.CurrentPage.MainPlayer.Data.GetPosition()[1]))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
