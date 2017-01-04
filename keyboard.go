package main

import "github.com/nsf/termbox-go"

type keyboardEventType int

const (
	MOVE_RIGHT keyboardEventType = 1 + iota
	MOVE_LEFT
	MOVE_UP
	MOVE_DOWN
	RETRY
	END
)

type keyboardEvent struct {
	eventType keyboardEventType
	key       termbox.Key
}

func keyToDirection(k keyboardEventType) direction {
	switch k {
	case MOVE_LEFT:
		return LEFT
	case MOVE_DOWN:
		return DOWN
	case MOVE_RIGHT:
		return RIGHT
	case MOVE_UP:
		return UP
	default:
		return 0
	}
}

func listenToKeyboard(evChan chan keyboardEvent) {
	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowLeft:
				evChan <- keyboardEvent{eventType: MOVE_LEFT, key: ev.Key}
			case termbox.KeyArrowDown:
				evChan <- keyboardEvent{eventType: MOVE_DOWN, key: ev.Key}
			case termbox.KeyArrowRight:
				evChan <- keyboardEvent{eventType: MOVE_RIGHT, key: ev.Key}
			case termbox.KeyArrowUp:
				evChan <- keyboardEvent{eventType: MOVE_UP, key: ev.Key}
			case termbox.KeyEsc:
				evChan <- keyboardEvent{eventType: END, key: ev.Key}
			default:
				switch ev.Ch {
				case 'h', 'H':
					evChan <- keyboardEvent{eventType: MOVE_LEFT, key: ev.Key}
				case 'j', 'J':
					evChan <- keyboardEvent{eventType: MOVE_DOWN, key: ev.Key}
				case 'k', 'K':
					evChan <- keyboardEvent{eventType: MOVE_UP, key: ev.Key}
				case 'l', 'L':
					evChan <- keyboardEvent{eventType: MOVE_RIGHT, key: ev.Key}
				case 'r', 'R':
					evChan <- keyboardEvent{eventType: RETRY, key: ev.Key}
				case 'q', 'Q':
					evChan <- keyboardEvent{eventType: END, key: ev.Key}
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
