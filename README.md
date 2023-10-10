<div align="center">

# Markdown cli renderer
A cli tool that renders markdown, based on [glamour](https://github.com/charmbracelet/glamour.git)

![last commit](https://img.shields.io/github/last-commit/loenard97/markdown-cli?&style=for-the-badge&logo=github&color=00ADD8)
![repo size](https://img.shields.io/github/repo-size/loenard97/markdown-cli?&style=for-the-badge&logo=github&color=00ADD8)

</div>


## Usage
A simple cli tool, written in Go, that outputs formattted markdown, using the Go library 'github.com/charmbracelet/glamour'.

Usage:
Input arguments work similarly to 'cat': list one or more file names to output the file contents.
If no file name is given, stdin is used as input.

Options:
-  --help                Print this help text
-  --style=dark          Set a glamour style. Possible options: ascii | dark | dracula | light | notty | pink (default: dark)
