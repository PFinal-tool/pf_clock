package main

import (
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	wr "github.com/wailsapp/wails/v2/pkg/runtime"
	"runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	app := NewApp()
	AppMenu := menu.NewMenu()
	ColckMenu := AppMenu.AddSubmenu("File")
	ColckMenu.AddText("退出", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		wr.Quit(app.ctx)
	})

	if runtime.GOOS == "darwin" {
		AppMenu.Append(menu.EditMenu()) // on macos platform, we should append EditMenu to enable Cmd+C,Cmd+V,Cmd+Z... shortcut
	}

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "clock",
		Width:  450,
		Height: 175,
		Menu:   AppMenu, // reference the menu above
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
			About: &mac.AboutInfo{
				Title:   "PFClock",
				Message: "PFinalClub Clock",
				Icon:    icon,
			},
			OnFileOpen: nil,
			OnUrlOpen:  nil,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
