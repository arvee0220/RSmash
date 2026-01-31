package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("Listening for Alt+key combos. Press NumPad9 to exit.")

	events := hook.Start()
	defer hook.End()

	altPressed := false

	keyMap := map[uint16]string{
		16: "7", // Q
		17: "8", // W
		30: "4", // A
		31: "5", // S
		44: "1", // Z
		45: "2", // X
	}

	for ev := range events {
		switch ev.Kind {
		case hook.KeyDown:
			if ev.Keycode == 56 {

				altPressed = true
				continue
			}

			if altPressed {
				if numpadKey, ok := keyMap[ev.Keycode]; ok {
					err := robotgo.KeyTap(numpadKey, "numpad")
					if err != nil {
						return
					}
				}
			}

			if ev.Keycode == 73 {
				time.Sleep(100 * time.Millisecond)
				fmt.Println("NumPad9 — exiting")
				return
			}

		case hook.KeyUp:
			if ev.Keycode == 56 {
				time.Sleep(100 * time.Millisecond)
				altPressed = false
			}
		}

	}
}
