package app

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jroimartin/gocui"
)

type DirFile struct {
	sonarDir string
	filePath string
}

func NewDirFile(sonarDir, filePath string) *DirFile {
	return &DirFile{
		sonarDir: sonarDir,
		filePath: filePath,
	}
}

func (d DirFile) EnsureConfigExistence() error {

	// Ensure the config directory exists
	if _, err := os.Stat(d.sonarDir); os.IsNotExist(err) {
		err := os.MkdirAll(d.sonarDir, 0755)
		if err != nil {
			fmt.Println("Error creating config directory:", err)
			return err
		}
	}

	// Ensure the file exists
	if _, err := os.Stat(d.filePath); os.IsNotExist(err) {
		f, err := os.Create(d.filePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return err
		}
		defer f.Close()
	}

	return nil
}

func (d DirFile) AppenToFile(path string) {
	f, err := os.OpenFile(d.filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file for appending:", err)
		os.Exit(1)
	}
	defer f.Close()

	newPath := path + "\n"

	_, err = f.WriteString(newPath)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func (d DirFile) RemoveDirFromFile(path string) {
	f, err := os.OpenFile(d.filePath, os.O_APPEND|os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file for appending:", err)
		os.Exit(1)
	}
	defer f.Close()

	newPath := path + "\n"

	_, err = f.WriteString(newPath)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func (d DirFile) ReadConfigFile(v *gocui.View) {
	file, err := os.Open(d.filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var prev string
	first := true

	for scanner.Scan() {
		if !first {
			fmt.Fprintln(v, prev)
		} else {
			first = false
		}
		prev = scanner.Text()
	}

	// After the loop, print the last line without newline
	if prev != "" {
		fmt.Fprint(v, prev)
	}
}
