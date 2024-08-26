package main

import (
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "clock",
		Width:  450,
		Height: 160,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless:        true,
		DisableResize:    true,
		AlwaysOnTop:      true,
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Mac: &mac.Options{
			TitleBar:             nil,
			Appearance:           "",
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			Preferences:          nil,
			DisableZoom:          false,
			About:                nil,
			OnFileOpen:           nil,
			OnUrlOpen:            nil,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
