package icalendar

type Field struct {
	name string
	value string
}

type Vevent struct {
	components []Field
}

func (e *Vevent) String() string {
	s := "BEGIN:VEVENT\n"
	for _, c := range e.components {
		s += c.name + ":" + c.value + "\n"
	}
	s += "END:VEVENT\n"
	return s
}

func (e *Vevent) Add(f Field) {
	e.components = append(e.components, f)
}
