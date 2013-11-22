package icalendar

import (
	"strings"
	"time"
)

var crlf = "\r\n"

type Name string
func (n *Name) String() string { return strings.ToUpper(string(*n)) }

type Value interface {
	String() string
}

type VString string
func (s VString) String() string { return string(s) }

type VDate time.Time
var datelayout = "20060102"
func (d VDate) String() string { return time.Time(d).Format(datelayout) }

type Parameter struct {
	name Name
	value Value
}

type Property struct {
	name Name
	value Value
	parameters []Parameter
}

type Properties []Property

type Component struct {
	name Name
	properties Properties
	components []Component
}

func (properties Properties) String() string {
	return "dummy TODO"
}

func (c *Component) String() string {
	s := "BEGIN:" + c.name.String() + crlf
	for _, prop := range c.properties {
		s += prop.name.String() + ":" + prop.value.String() + crlf
	}
	for _, subc := range c.components {
		s += subc.String()
	}
	s += "END:" + c.name.String() + crlf
	return s
}

func (c *Component) Add(name Name, value Value) {
	c.properties = append(c.properties, Property{name: name, value: value})
}

func (c *Component) Set(name Name, value Value) {
	c.Add(name, value)			// TODO: remove duplicates
}

func (c *Component) AddComponent(subc *Component) {
	c.components = append(c.components, *subc)
}
