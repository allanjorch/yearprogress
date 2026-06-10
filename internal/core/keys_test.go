package core

import "testing"

func TestIsExitKey(t *testing.T) {
	exit := []string{"esc", "enter", " ", "space", "q", "Q", "backspace"}
	for _, key := range exit {
		if !IsExitKey(key) {
			t.Fatalf("%q should be an exit key", key)
		}
	}

	stay := []string{"a", "tab", "ctrl+c", "up", "1"}
	for _, key := range stay {
		if IsExitKey(key) {
			t.Fatalf("%q should not be an exit key", key)
		}
	}
}

func TestExitKeys_placeholder(t *testing.T) {
	if len(ExitKeys) != 5 {
		t.Fatalf("got %d exit keys, want 5", len(ExitKeys))
	}
}