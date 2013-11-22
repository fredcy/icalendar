package icalendar

import (
	"strings"
)

type Property []string

func (p *Property) String() string {
	return strings.Join(*p, ",")
}

type Properties map[string]Property

type Component struct {
	name string
	properties Properties
}

func (c *Component) String() string {
	s := "BEGIN:" + c.name + "\n"
	for k, v := range c.properties {
		s += k + ":" + v.String() + "\n"
	}
	s += "END:" + c.name + "\n"
	return s
}

func (c *Component) Add(name, value string) {
	if c.properties == nil {
		c.properties = make(Properties)
	}
	c.properties[name] = Property{value}
}

func (c *Component) Set(name, value string) {
	if c.properties == nil {
		c.properties = make(Properties)
	}
	c.properties[name] = Property{value}
}

type Vevent struct {
	Component
}


