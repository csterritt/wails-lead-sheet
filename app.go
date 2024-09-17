package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"wails-lead-sheet/parser"
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
func (a *App) RetrieveFileContents(filePath string) (parser.ParsedContent, error) {
	prsr := parser.ParsedContent{}
	contents, err := os.ReadFile(filePath)
	if err != nil {
		runtime.LogPrintf(a.ctx, "Retrieve contents of %s contains caught %v\n", filePath, err)
		return prsr, err
	}

	err = prsr.ParseContent(string(contents))
	if err != nil {
		return prsr, err
	}

	return prsr, nil
}

// TransposeUpOneStep transposes the given content up one step
func (a *App) TransposeUpOneStep(content parser.ParsedContent) parser.ParsedContent {
	content.TransposeUpOneStep()

	return content
}

// TransposeDownOneStep transposes the given content down one step
func (a *App) TransposeDownOneStep(content parser.ParsedContent) parser.ParsedContent {
	content.TransposeDownOneStep()

	return content
}
