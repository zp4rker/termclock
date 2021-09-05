package main

import (
	"github.com/gdamore/tcell/v2"
	"log"
	"os"
	"time"
)

var screen tcell.Screen
var plain = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

var (
	currentColour int
	militaryTime, showSeconds bool
)

var (
	colours = [...]tcell.Style{
		tcell.StyleDefault.Background(tcell.ColorWhite),
		tcell.StyleDefault.Background(tcell.ColorRed),
		tcell.StyleDefault.Background(tcell.ColorGreen),
		tcell.StyleDefault.Background(tcell.ColorBlue),
		tcell.StyleDefault.Background(tcell.ColorYellow),
		tcell.StyleDefault.Background(tcell.ColorOrange),
		tcell.StyleDefault.Background(tcell.ColorPurple),
	}

	numbers = map[rune][5][3]int{
		'0': {
			{1, 1, 1}, // # # #
			{1, 0, 1}, // #   #
			{1, 0, 1}, // #   #
			{1, 0, 1}, // #   #
			{1, 1, 1}, // # # #
		},
		'1': {
			{1, 1, 0}, // # #
			{0, 1, 0}, //   #
			{0, 1, 0}, //   #
			{0, 1, 0}, //   #
			{1, 1, 1}, // # # #
		},
		'2': {
			{1, 1, 1}, // # # #
			{0, 0, 1}, //     #
			{1, 1, 1}, // # # #
			{1, 0, 0}, // #
			{1, 1, 1}, // # # #
		},
		'3': {
			{1, 1, 1}, // # # #
			{0, 0, 1}, //     #
			{1, 1, 1}, // # # #
			{0, 0, 1}, //     #
			{1, 1, 1}, // # # #
		},
		'4': {
			{1, 0, 1}, // #   #
			{1, 0, 1}, // #   #
			{1, 1, 1}, // # # #
			{0, 0, 1}, //     #
			{0, 0, 1}, //     #
		},
		'5': {
			{1, 1, 1}, // # # #
			{1, 0, 0}, // #
			{1, 1, 1}, // # # #
			{0, 0, 1}, //     #
			{1, 1, 1}, // # # #
		},
		'6': {
			{1, 1, 1}, // # # #
			{1, 0, 0}, // #
			{1, 1, 1}, // # # #
			{1, 0, 1}, // #   #
			{1, 1, 1}, // # # #
		},
		'7': {
			{1, 1, 1}, // # # #
			{0, 0, 1}, //     #
			{0, 0, 1}, //     #
			{0, 0, 1}, //     #
			{0, 0, 1}, //     #
		},
		'8': {
			{1, 1, 1}, // # # #
			{1, 0, 1}, // #   #
			{1, 1, 1}, // # # #
			{1, 0, 1}, // #   #
			{1, 1, 1}, // # # #
		},
		'9': {
			{1, 1, 1}, // # # #
			{1, 0, 1}, // #   #
			{1, 1, 1}, // # # #
			{0, 0, 1}, //     #
			{1, 1, 1}, // # # #
		},
		':': {
			{0, 0, 0}, //
			{0, 1, 0}, //   #
			{0, 0, 0}, //
			{0, 1, 0}, //   #
			{0, 0, 0}, //
		},
		'A': {
			{1, 1, 1}, // # # #
			{1, 0, 1}, // #   #
			{1, 1, 1}, // # # #
			{1, 0, 1}, // #   #
			{1, 0, 1}, // #   #
		},
		'P': {
			{1, 1, 1}, // # # #
			{1, 0, 1}, // #   #
			{1, 1, 1}, // # # #
			{1, 0, 0}, // #
			{1, 0, 0}, // #
		},
		'M': {
			{1, 1, 1}, // # # #
			{1, 2, 1}, // # | #
			{1, 2, 1}, // # | #
			{1, 0, 1}, // #   #
			{1, 0, 1}, // #   #
		},
	}
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

	defer func() {
		recover()
	}()

	go handleInput()

	for {
		screen.Show()
		drawClock()
	}
}

func drawClock() {
	clearScreen()

	now := time.Now()

	var timeStr string
	if militaryTime {
		if showSeconds {
			timeStr = now.Format("15:04:05")
		} else {
			timeStr = now.Format("15:04")
		}
	} else {
		if showSeconds {
			timeStr = now.Format("03:04:05  PM")
		} else {
			if now.Second() % 2 == 0 {
				timeStr = now.Format("03:04  PM")
			} else {
				timeStr = now.Format("03 04  PM")
			}
		}
	}

	for i, r := range []rune(timeStr) {
		offset := i * 12
		for x := 0; x < 3; x++ {
			for y := 0; y < 5; y++ {
				v := numbers[r][y][x]
				for px := 0; px < 3; px++ {
					switch {
					case v == 0:
						screen.SetContent(offset + x + px, y, ' ', nil, plain)
					case v == 1, v == 2 && px == 1:
						screen.SetContent(offset + x + px, y, ' ', nil, colours[currentColour])
					}
				}
			}
		}
	}
}

func handleInput() {
	for {
		event := screen.PollEvent()
		switch event := event.(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			switch {
			case event.Key() == tcell.KeyCtrlC, event.Key() == tcell.KeyEsc:
				quit()
			case event.Key() == tcell.KeyCtrlR:
				clearScreen()
				drawClock()
				screen.Show()
			case event.Rune() == 'c':
				if currentColour + 1 >= len(colours) {
					currentColour = 0
				} else {
					currentColour++
				}
			case event.Rune() == 't':
				militaryTime = !militaryTime
			case event.Rune() == 's':
				showSeconds = !showSeconds
			}
		}
	}
}

func clearScreen() {
	screen.SetStyle(plain)
	screen.Clear()
}

func quit() {
	clearScreen()
	screen.ShowCursor(0, 0)
	screen.Fini()
	os.Exit(0)
}