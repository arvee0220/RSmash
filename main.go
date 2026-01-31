package main

import (
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"time"
)

func main() {
	events := hook.Start()
	defer hook.End()

	altPressed := false

	for ev := range events {
		switch ev.Kind {
		case hook.KeyDown:
			if ev.Keycode == 56 {
				altPressed = true
				continue
			}

			if altPressed {
				switch ev.Keycode {
				case 16:
					robotgo.KeyTap("7", "numpad")
				case 30:
					robotgo.KeyTap("4", "numpad")
				case 44:
					robotgo.KeyTap("1", "numpad")
				case 17:
					robotgo.KeyTap("8", "numpad")
				case 31:
					robotgo.KeyTap("5", "numpad")
				case 45:
					robotgo.KeyTap("2", "numpad")
				}
			}

			if ev.Keycode == 67 {
				time.Sleep(100 * time.Millisecond)
				return
			}
		case hook.KeyUp:
			if ev.Keycode == 56 {
				altPressed = false
			}
		}
	}
}
