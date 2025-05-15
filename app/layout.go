package app

import (
	"github.com/jroimartin/gocui"
)

type Layout struct {
	ConfigDirFile *DirFile
}

func NewLayout(configDirFile *DirFile) *Layout {
	return &Layout{
		ConfigDirFile: configDirFile,
	}
}


func (l Layout) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	v, err := g.SetView("sonar", 0, 0, maxX-1, maxY-1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	v.Title = "Sonar"
	v.Highlight = true
	v.SelFgColor = gocui.ColorGreen
	v.Editable = false
	v.Autoscroll = true

	v.Clear()


	l.ConfigDirFile.ReadConfigFile(v)

	// fmt.Fprintf(v, "/home/fl4vis/Documents/Programming/sonar\n/home/fl4vis\n/var/www")

	return nil
}
