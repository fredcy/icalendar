package icalendar

type Field struct {
	name string
	value string
}

type Component struct {
	name string
	properties map[string]string
}

func (c *Component) String() string {
	s := "BEGIN:" + c.name + "\n"
	for k, v := range c.properties {
		s += k + ":" + v + "\n"
	}
	s += "END:" + c.name + "\n"
	return s
}

func (c *Component) Add(name, value string) {
	if c.properties == nil {
		c.properties = make(map[string]string)
	}
	c.properties[name] = value
}

type Vevent struct {
	Component
}


