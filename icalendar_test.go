package icalendar

import (
	"testing"
)

func MustEqual(t *testing.T, got, want string) {
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestVevent(t *testing.T) {
	var e Vevent
	e.Add(Field{"SUMMARY", "foo"})
	MustEqual(t, e.String(), "BEGIN:VEVENT\nSUMMARY:foo\nEND:VEVENT\n")
}
