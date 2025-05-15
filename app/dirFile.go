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

	result, _ := d.checkExistingDir(path)

	if result == false {
		newPath := path + "\n"

		_, err = f.WriteString(newPath)
		fmt.Println("Successfully appended dir")
		os.Exit(0)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Dir already exists")
		os.Exit(1)
	}
}

func (d DirFile) RemoveDirFromFile(path string) {

	// Read all lines
	file, err := os.Open(d.filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
		return
	}
	defer file.Close()

	result, lineToDelete := d.checkExistingDir(path)

	if result == true {

		var lines []string
		scanner := bufio.NewScanner(file)
		currentLine := 1

		for scanner.Scan() {
			if currentLine != lineToDelete {
				lines = append(lines, scanner.Text())
			}
			currentLine++
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}

		// Write lines back to file
		outputFile, err := os.Create(d.filePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			os.Exit(1)
		}
		defer outputFile.Close()

		writer := bufio.NewWriter(outputFile)
		for _, line := range lines {
			fmt.Fprintln(writer, line)
		}
		writer.Flush()

		fmt.Println("Dir deleted successfully")
		os.Exit(0)
	} else {
		fmt.Println("Dir doesn't exist")
		os.Exit(1)
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

func (d DirFile) checkExistingDir(path string) (bool, int) {

	file, err := os.Open(d.filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	defer file.Close()

	line := 1

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() == path {
			return true, line
		}
		line++
	}

	return false, line
}
