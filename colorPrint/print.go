package colorPrint

import (
	"bufio"
	"fmt"

	"github.com/mattn/go-colorable"
)

var (
	Black   string = "\u001b[30m"
	Red     string = "\u001b[31m"
	Green   string = "\u001b[32m"
	Yellow  string = "\u001b[33m"
	Blue    string = "\u001b[34m"
	Magenta string = "\u001b[35m"
	Cyan    string = "\u001b[36m"
	White   string = "\u001b[37m"
	Default string = "\u001b[0m"
)

func colorChoice(color string) string {
	switch color {
	case "black":
		return Black
	case "red":
		return Red
	case "green":
		return Green
	case "yellow":
		return Yellow
	case "bule":
		return Blue
	case "magenta":
		return Magenta
	case "cyan":
		return Cyan
	case "white":
		return White
	default:
		return White
	}
}

func Out(text string, color string) {
	stdOut := bufio.NewWriter(colorable.NewColorableStdout())
	fmt.Fprint(stdOut, colorChoice(color)+text+Default+"\u001b[0m\n")
	stdOut.Flush()
}

// package colorPrint

// import (
// 	"github.com/fatih/color"
// )

// func Out(text string) {
// 	color.Cyan(text)
// }
