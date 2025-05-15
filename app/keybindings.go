package app

import (
	"fmt"
	"os"

	"github.com/jroimartin/gocui"
)

type KeyBinding struct {
    selection *Selection
}

func NewKeyBindingHandler(selection *Selection) *KeyBinding {
    return &KeyBinding{
        selection: selection,
    }
}

func (k KeyBinding) Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func (k KeyBinding) CursorDown(g *gocui.Gui, v *gocui.View) error {

	lines := v.BufferLines()
	if len(lines) == 0 {
		return nil
	}

	_, cy := v.Cursor()

	maxLine := len(lines) - 1

	if cy < maxLine {
		_ = v.SetCursor(0, cy+1)
	} else {
		_ = v.SetCursor(0, cy-maxLine)
	}

	_, yPos := v.Cursor()
	k.selection.Text = lines[yPos]

	return nil
}

func (k KeyBinding) CursorUp(g *gocui.Gui, v *gocui.View) error {

	lines := v.BufferLines()
	if len(lines) == 0 {
		return nil
	}

	_, cy := v.Cursor()

	maxLine := len(lines) - 1

	if cy > 0 {
		_ = v.SetCursor(0, cy-1)
	} else {
		_ = v.SetCursor(0, cy+maxLine)
	}

	_, yPos := v.Cursor()
	k.selection.Text = lines[yPos]

	return nil
}

func (k KeyBinding) HandleEnter(g *gocui.Gui, v *gocui.View) error {
	g.Close()

	// Select first one if not arrow actions
	if len(k.selection.Text) == 0 {
		lines := v.BufferLines()

		if len(lines) > 0 {
			fmt.Println(lines[0])
		}
	} else {
		fmt.Println(k.selection.Text)
	}

	os.Exit(100)

	return nil

}
