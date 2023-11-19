package terminalControl

import (
	"fmt"
	"strconv"
)

const Escape = "\033"
const Reset = Escape + "[0m"
const Red = Escape + "[31m"
const Green = Escape + "[32m"
const Yellow = Escape + "[33m"
const Blue = Escape + "[34m"
const Purple = Escape + "[35m"
const Cyan = Escape + "[36m"
const Gray = Escape + "[37m"
const White = Escape + "[97m"

func Colorize(str, color string) string {
	return color + str + Reset
}

func MoveCursorUp(step int) {
	fmt.Print(Escape + "[" + strconv.Itoa(step) + "A")
}

func MoveCursorDown(step int) {
	fmt.Print(Escape + "[" + strconv.Itoa(step) + "B")
}

func MoveCursorRight(step int) {
	fmt.Print(Escape + "[" + strconv.Itoa(step) + "C")
}

func MoveCursorLeft(step int) {
	fmt.Print(Escape + "[" + strconv.Itoa(step) + "D")
}

func ClearLine() {
	fmt.Print(Escape + "[K")
}

func MoveCursorToBeginning() {
	fmt.Print("\r")
}
