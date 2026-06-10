package core

import "strings"

// ExitKeys lists the keys that close the application.
// Renderers map these names to their platform-specific key codes.
var ExitKeys = []string{
	"esc",
	"enter",
	"space",
	"q",
	"backspace",
}

// IsExitKey reports whether a Bubble Tea key string should quit the app.
func IsExitKey(key string) bool {
	switch strings.ToLower(key) {
	case "esc", "enter", " ", "space", "q", "backspace":
		return true
	default:
		return false
	}
}