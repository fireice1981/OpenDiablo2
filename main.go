package main

import (
	"log"

	"github.com/essial/OpenDiablo2/Core"

	"github.com/essial/OpenDiablo2/MPQ"
	"github.com/hajimehoshi/ebiten"
)

var d2Engine *Core.Engine

func main() {
	log.Println("OpenDiablo2 - Open source Diablo 2 engine")
	MPQ.InitializeCryptoBuffer()
	d2Engine = Core.CreateEngine()
	ebiten.SetCursorVisible(false)
	ebiten.SetFullscreen(d2Engine.Settings.FullScreen)
	ebiten.SetRunnableInBackground(d2Engine.Settings.RunInBackground)
	ebiten.SetVsyncEnabled(d2Engine.Settings.VsyncEnabled)
	ebiten.SetMaxTPS(d2Engine.Settings.TicksPerSecond)
	if err := ebiten.Run(update, 800, 600, d2Engine.Settings.Scale, "OpenDiablo 2"); err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	d2Engine.Update()
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	d2Engine.Draw(screen)
	return nil
}
