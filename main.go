package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
)

func contains(array []string, element string) bool {
	for _, e := range array {
		if e == element {
			return true
		}
	}
	return false
}

type filter func(string, string) bool

func filterIndex(args []string, element string, fn filter) int {
	for i, e := range args {
		if fn(e, element) {
			return i
		}
	}
	return -1
}

func catFile(file_name string, style string) {
	contents, err := os.ReadFile(file_name)
	if err != nil {
		log.Fatal(err)
	}

	markdown, err := glamour.Render(string(contents), style)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(markdown)
}

func catStdin(style string) {
	scanner := bufio.NewScanner(os.Stdin)
	var contents string
	for scanner.Scan() {
		contents += scanner.Text() + "\n"
	}

	markdown, err := glamour.Render(contents, style)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(markdown)
}

func main() {
	help_text := `Markdown cli renderer
A simple cli tool, written in Go, that outputs formattted markdown, using the Go library 'github.com/charmbracelet/glamour'.

Usage:
Input arguments work similarly to 'cat': list one or more file names to output the file contents.
If no file name is given, stdin is used as input.

Options:
  --help                Print this help text
  --style=dark          Set a glamour style. Possible options: ascii | dark | dracula | light | notty | pink (default: dark)
`
	if contains(os.Args, "--help") {
		fmt.Println(help_text)
		return
	}

	style := "dark"
	style_index := filterIndex(os.Args, "--style", strings.HasPrefix)
	if style_index >= 0 {
		style = os.Args[style_index]
		style = strings.TrimPrefix(style, "--style=")
	}

	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i++ {
			if i != style_index {
				catFile(os.Args[i], style)
			}
		}
	} else {
		catStdin(style)
	}
}
