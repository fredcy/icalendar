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

func TestRrule(t *testing.T) {
	rr := Property{ name: "RRULE" }
	rr.AddParameter("FREQ", VString("YEARLY"))
	rr.AddParameter("BYDAY", VString("1SU"))
	rr.AddParameter("BYMONTH", VInt(11))
	MustEqual(t, rr.String(), "RRULE;FREQ=YEARLY;BYDAY=1SU;BYMONTH=11")

	// alternatively, set the parameters with a literal [so clumsy]
	r2 := Property{ name: "RRULE",
		parameters: []Parameter{
			{"FREQ", VString("YEARLY") },
			{"BYDAY", VString("2SU") },
			{"BYMONTH", VInt(3)},
		},
	}
	MustEqual(t, r2.String(), "RRULE;FREQ=YEARLY;BYDAY=2SU;BYMONTH=3")
}

func TestVList(t *testing.T) {
	p := Property{ name: "FOO" }
	p.AddParameter("BAR", VList{VInt(11), VInt(12)})
	MustEqual(t, p.String(), "FOO;BAR=11,12")

	v2 := VEnumList{}
	v2.AddValue("FREQ", VString("YEARLY"))
	v2.AddValue("BYMONTH", VInt(4))
	p2 := NewProperty("foo2", v2)
	MustEqual(t, p2.String(), "FOO2:FREQ=YEARLY;BYMONTH=4")
}

func TestName(t *testing.T) {
	c := Component{}
	c.SetName("blatz")
	MustEqual(t, c.name.String(), "BLATZ")
}

func TestProperty(t *testing.T) {
	p := NewProperty("thename", VString("thevalue"))
	p.AddParameter("P1", VString("foo"))
	MustEqual(t, p.name.String(), "THENAME")
	MustEqual(t, p.String(), "THENAME;P1=foo:thevalue")
}

func TestVStringf(t *testing.T) {
	v := VStringf("%s %d", "foo", 3)
	MustEqual(t, v.String(), "foo 3")
}
