package main

import (
	"github.com/gdamore/tcell/v2"
	"log"
	"os"
)

var screen tcell.Screen

var (
	currentColour int
	militaryTime, showSeconds bool
)

func main() {
	var err error
	screen, err = tcell.NewScreen()
	if err != nil {
		log.Fatalln("Unable to create screen!")
	}
	if err = screen.Init(); err != nil {
		log.Fatalln("Unable to initialise screen!")
	}
}

func handleEvent(event tcell.EventKey) {
	switch {
	case event.Key() == tcell.KeyCtrlC, event.Key() == tcell.KeyEsc:
		quit()
	case event.Key() == tcell.KeyCtrlR:
		clearScreen()
		// refresh
	case event.Rune() == 'c':
		// change colour
	case event.Rune() == 't':
		// toggle military time
	case event.Rune() == 's':
		// toggle seconds
	}
}

func clearScreen() {
	screen.SetStyle(tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset))
	screen.Clear()
}

func quit() {
	screen.Fini()
	os.Exit(0)
}