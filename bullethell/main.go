package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	//"fmt"
	"./gamesystem"
)
func main() {
	rl.InitWindow(sys.WinWidth, sys.WinHeight, "Test game")
	rl.SetTargetFPS(60)
	sys.Onload()
	for !rl.WindowShouldClose() {
		sys.Update()
		rl.BeginDrawing()
		sys.Draw()
		rl.ClearBackground(rl.RayWhite)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
