package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

var lastDirectory string

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// ChooseFile lets the user choose an input file
func (a *App) ChooseFile() string {
	if lastDirectory == "" {
		//lastDirectory = os.Getenv("HOME")
		lastDirectory = "/Users/chris/hacks/music/wails-tab/song-tabs"
	}

	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		DefaultDirectory:     lastDirectory,
		DefaultFilename:      "",
		Title:                "Choose Song File",
		CanCreateDirectories: false,
		Filters:              []runtime.FileFilter{},
	})
	if err != nil {
		return fmt.Sprintf("Error: Unable to choose Song File: %v", err)
	}

	lastDirectory = filepath.Dir(file)

	return file
}

// RetrieveFileContents retrieves the contents from the given file path
func (a *App) RetrieveFileContents(filePath string) (string, error) {
	contents, err := os.ReadFile(filePath)
	if err != nil {
		runtime.LogPrintf(a.ctx, "Retrieve contents of %s contains caught %v\n", filePath, err)
	}

	return string(contents), err
}
