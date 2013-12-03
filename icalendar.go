/*
Package icalendar implements a simple library for generating icalendar text (RFC 2445).
It is loosely based on https://github.com/collective/icalendar.
*/
package icalendar

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

const CrLf = "\r\n"

// Name holds the names of components, properties, and parameters
type Name string
func (n *Name) String() string { return strings.ToUpper(string(*n)) }

// Value holds the values of properties and parameters
type Value interface {
	String() string
}

// VString is a Value for strings
type VString string
func (s VString) String() string { return string(s) }

// VStringf generates a VString from format and args
func VStringf(format string, args ...interface{}) VString {
	return VString(fmt.Sprintf(format, args...))
}

type VDate time.Time
const datelayout = "20060102"
func (d VDate) String() string { return time.Time(d).Format(datelayout) }

type VDateTime time.Time
const datetimelayout = "20060102T150405"
func (d VDateTime) String() string { return time.Time(d).Format(datetimelayout) }

type VInt int
func (i VInt) String() string { return strconv.Itoa(int(i)) }

type VList []Value
func (vl VList) String() string {
	var vls []string
	for _, v := range vl {
		vls = append(vls, v.String())
	}
	return strings.Join(vls, ",")
}

type namevalue struct {
	name Name
	value Value
}

// VEnumList implements enumerated property values
type VEnumList []namevalue

func (vel VEnumList) String() string {
	var nvs []string
	for _, v := range vel {
		nvs = append(nvs, v.name.String() + "=" + v.value.String())
	}
	return strings.Join(nvs, ";")
}

func (vel *VEnumList) AddValue(name Name, value Value) {
	*vel = append(*vel, namevalue{ name, value })
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func hms(i int) (int, int, int) {
	hours := i / 3600
	minutes := (i % 3600) / 60
	seconds := i % 60
	return hours, minutes, seconds
}

type VUtcOffset int				// seconds east of UTC

func (u VUtcOffset) String() string {
	ua := abs(int(u))
	hours, minutes, seconds := hms(ua)
	if seconds != 0 {
		log.Printf("seconds (%v) not zero in VUtcOffset value (%v)", seconds, int(u))
	}
	if int(u) >= 0 {
		return fmt.Sprintf("%2.2d%2.2d", hours, minutes)
	} else {
		return fmt.Sprintf("-%2.2d%2.2d", hours, minutes)
	}
}

type Parameter struct {
	name Name
	value Value
}
func (param Parameter) String() string {
	return param.name.String() + "=" + param.value.String()
}

type Property struct {
	name Name
	value Value
	parameters []Parameter
}

func (prop Property) String() string {
	s := prop.name.String()
	for _, param := range prop.parameters {
		s += (";" + param.String())
	}
	if prop.value != nil {
		s += (":" + prop.value.String())
	}
	return s
}

func (prop *Property) AddParameter(name Name, value Value) {
	prop.parameters = append(prop.parameters, Parameter{name, value})
}

func NewProperty(name Name, value Value) Property {
	return Property{ name: name, value: value }
}

type Component struct {
	name Name
	properties []Property
	components []Component
}

func (c *Component) String() string {
	s := "BEGIN:" + c.name.String() + CrLf
	for _, prop := range c.properties {
		s += prop.String() + CrLf
	}
	for _, subc := range c.components {
		s += subc.String()
	}
	s += "END:" + c.name.String() + CrLf
	return s
}

func (c *Component) SetName(name string) {
	c.name = Name(name)
}

func (c *Component) Add(name Name, value Value) {
	c.properties = append(c.properties, Property{name: name, value: value})
}

func (c *Component) Set(name Name, value Value) {
	c.Add(name, value)			// TODO: remove duplicates
}

func (c *Component) AddProperty(prop Property) {
	c.properties = append(c.properties, prop)
}

func (c *Component) AddComponent(subc *Component) {
	c.components = append(c.components, *subc)
}
