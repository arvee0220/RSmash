package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"

	"time"
)

func main() {
	fmt.Println("Listening for Alt+key combos. Press NumPad9 to exit.")

	events := hook.Start()
	defer hook.End()

	altPressed := false

	keyMap := map[uint16]string{
		16: "num7", // Q
		17: "num8", // W
		30: "num4", // A
		31: "num5", // S
		44: "num1", // Z
		45: "num2", // X
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
					robotgo.KeyToggle("alt", "up")
					robotgo.KeyTap(numpadKey)
					//fmt.Println("Keypad:", numpadKey)
					//fmt.Println("KeyCode:", ev.Keycode)
					robotgo.KeyToggle("alt", "down")

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
