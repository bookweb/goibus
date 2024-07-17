package goibus

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/godbus/dbus/v5"
)

type Component struct {
	XMLName       xml.Name                `xml:"component" dbus:"-"`
	Name          string                  `xml:"-"`
	Attachments   map[string]dbus.Variant `xml:"-"`
	ComponentName string                  `xml:"name"`
	Description   string                  `xml:"description"`
	Version       string                  `xml:"version"`
	License       string                  `xml:"license"`
	Author        string                  `xml:"author"`
	Homepage      string                  `xml:"homepage"`
	Exec          string                  `xml:"exec"`
	Textdomain    string                  `xml:"textdomain"`
	ObservedPaths []dbus.Variant          `xml:"-"`
	EngineList    []dbus.Variant          `xml:"-"`
	Engines       []*EngineDesc           `xml:"engines>engine" dbus:"-"`
}

func NewComponent(name string, desc string, version string, license string, author string, homepage string, exec string, textdomain string) *Component {
	c := &Component{}
	c.Name = "IBusComponent"
	c.ComponentName = name
	c.Description = desc
	c.Version = version
	c.License = license
	c.Author = author
	c.Homepage = homepage
	c.Exec = exec
	c.Textdomain = textdomain

	return c
}

func (c *Component) AddEngine(e *EngineDesc) {
	c.Engines = append(c.Engines, e)
	c.EngineList = append(c.EngineList, dbus.MakeVariant(*e))
}

func (c *Component) OutputXML(w io.Writer) {

	data, err := xml.MarshalIndent(c, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "%s%s\n\n%s\n", xml.Header, "<!-- Generated By github.com/sarim/goibus -->", string(data))
}
