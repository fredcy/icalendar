package icalendar

import (
	"strings"
	"testing"
)

func MustEqual(t *testing.T, got, want string) {
	if want != got {
		t.Errorf("want <<%v>>, got <<%v>>", want, got)
	}
}

func Join(lines []string) string {
	return strings.Join(lines, crlf) + crlf
}

func TestVevent(t *testing.T) {
	e := Component{ name: Name("VEVENT") }
	e.Add("SUMMARY", VString("foo"))
	MustEqual(t, e.String(),
		Join([]string{"BEGIN:VEVENT", "SUMMARY:foo", "END:VEVENT"}))
}

func TestCalendar(t *testing.T) {
	c := Component{ name: "VCALENDAR" }
	c.Add("version", VString("2.0"))
	MustEqual(t, c.String(),
		Join([]string{"BEGIN:VCALENDAR", "VERSION:2.0", "END:VCALENDAR"}))
}

func TestSubComponents(t *testing.T) {
	timezone := Component{ name: "VTIMEZONE" }
	timezone.Set("tzid", VString("America/Chicago"))
	MustEqual(t, timezone.String(),
		Join([]string{"BEGIN:VTIMEZONE", "TZID:America/Chicago", "END:VTIMEZONE"}))

	daylight := Component{ name: "DAYLIGHT" }
	daylight.Add("tzname", VString("CDT"))
	timezone.AddComponent(&daylight)
	MustEqual(t, timezone.String(),
		Join([]string{"BEGIN:VTIMEZONE", "TZID:America/Chicago", "BEGIN:DAYLIGHT",
			"TZNAME:CDT", "END:DAYLIGHT", "END:VTIMEZONE"}))
}
