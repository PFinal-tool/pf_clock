package main

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed config.json
var defaultDicts embed.FS

// App struct
type App struct {
	Config Config
	ctx    context.Context
}

type ClockConfig struct {
	ClockBgColor   string `json:"clock_bg_color"`
	ClockTextColor string `json:"clock_text_color"`
}

type TextConfig struct {
	BgColor     string `json:"text_bg_color"`
	TextColor   string `json:"text_color"`
	TextContent string `json:"text_content"`
}

type Config struct {
	ClockConfig
	TextConfig
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.initConfig()
	a.ctx = ctx
}

func (a *App) initConfig() {
	configPath := "config.json"
	config, err := a.loadDictsFromEmbedFS(configPath)
	if err != nil {
		fmt.Println(err)
	}
	a.Config = config
}

func (a *App) loadDictsFromEmbedFS(fileName string) (Config, error) {
	fmt.Println(fileName)
	var config Config
	data, err := defaultDicts.ReadFile(fileName)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	fmt.Println(config)
	return config, err
}

func (a *App) GetAppConfig() Config {
	fmt.Println(a.Config)
	return a.Config
}

func (a *App) SetAppConfig(config Config) {
	a.Config = config
}

func (a *App) Setting() string {
	runtime.WindowSetSize(a.ctx, 600, 600)
	// runtime.WindowReload(a.ctx)
	configString := fmt.Sprintf("%+v", a.Config)
	fmt.Println(configString)
	return configString
}

func (a *App) Reload() {
	runtime.WindowSetSize(a.ctx, 450, 175)
	runtime.WindowReload(a.ctx)
}
// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
