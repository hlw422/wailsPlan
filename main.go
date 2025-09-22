package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	plan := NewPlan()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "计划管理",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 173, G: 216, B: 230, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			plan,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
