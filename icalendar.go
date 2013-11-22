package icalendar

import (
	"strings"
)

var crlf = "\r\n"

type Value struct {
	vString string
}

type Parameter struct {
	name string
	value string
}

type Property struct {
	name string
	value string
	parameters []Parameter
}

type Properties []Property

type Component struct {
	name string
	properties Properties
	components []Component
}

func (properties Properties) String() string {
	return "dummy"
}

func (c *Component) String() string {
	s := "BEGIN:" + c.name + crlf
	for _, prop := range c.properties {
		s += strings.ToUpper(prop.name) + ":" + prop.value + crlf
	}
	for _, subc := range c.components {
		s += subc.String()
	}
	s += "END:" + c.name + crlf
	return s
}

func (c *Component) Add(name, value string) {
	c.properties = append(c.properties, Property{name: name, value: value})
}

func (c *Component) Set(name, value string) {
	c.Add(name, value)			// TODO: remove duplicates
}

func (c *Component) AddComponent(subc *Component) {
	c.components = append(c.components, *subc)
}
