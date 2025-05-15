package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fl4vis/sonar/app"
	"github.com/jroimartin/gocui"
)


func main() {
	// Get config directory
	configDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Error getting user config dir:", err)
		return
	}

	sonarDir := filepath.Join(configDir, "sonar")
	filePath := filepath.Join(sonarDir, "directories.txt")

	configDirFile := app.NewDirFile(sonarDir, filePath)

	layout := app.NewLayout(configDirFile)
	
	err = configDirFile.EnsureConfigExistence()
	if err != nil {
		return
	}

	// Arguments
	args := os.Args[1:]

	if len(args) > 0 && len(args) == 1 {
		var dir string

		switch args[0] {
		case "+":
			dir, _ = os.Getwd()
			configDirFile.AppenToFile(dir)
			fmt.Println("Succesfully appended file")
			break
		case "-":
			fmt.Println(os.Getwd())
			break
		}


		os.Exit(0)

	} else if len(args) > 1 {
		fmt.Println("Only one argument required")
		os.Exit(0)
	}

	// GUI
	var selection app.Selection
	k := app.NewKeyBindingHandler(&selection)

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err.Error())
	}

	defer g.Close()

	g.SetManagerFunc(layout.Layout)

	g.Update(func(g *gocui.Gui) error {
		_, _ = g.SetCurrentView("sonar")
		return nil
	})

	// Exit (Ctrl+C, q)
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, k.Quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", 'q', gocui.ModNone, k.Quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("sonar", 'j', gocui.ModNone, k.CursorDown); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("sonar", 'k', gocui.ModNone, k.CursorUp); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("sonar", gocui.KeyEnter, gocui.ModNone, k.HandleEnter); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatal(err.Error())
	}

}

