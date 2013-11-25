package icalendar

import (
	"strings"
	"testing"
	"time"
)

func MustEqual(t *testing.T, got, want string) {
	if want != got {
		t.Errorf("want <<%v>>, got <<%v>>", want, got)
	}
}

func Join(lines []string) string {
	return strings.Join(lines, CrLf) + CrLf
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

func TestDates(t *testing.T) {
	td := time.Date(2008, time.April, 1, 0, 0, 0, 0, time.UTC)
	vd := VDate(td)
	MustEqual(t, vd.String(), "20080401")

	uo := VUtcOffset(-(5*3600+30*60))
	MustEqual(t, uo.String(), "-0530")

	uo = VUtcOffset((8*3600))
	MustEqual(t, uo.String(), "0800")
}

func TestDtstart(t *testing.T) {
	dt := time.Date(1970, 11, 1, 2, 0, 0, 0, time.UTC)
	ds := Property{ name: "DTSTART", value: VDateTime(dt) }
	ds.AddParameter("VALUE", VString("DATE-TIME"))
	MustEqual(t, ds.String(), "DTSTART;VALUE=DATE-TIME:19701101T020000")
}

