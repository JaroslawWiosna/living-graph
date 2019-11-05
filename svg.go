package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

// Circle ...
type Circle struct {
	ID          string `xml:"id,attr"`
	Radius      int    `xml:"r,attr"`
	X           int    `xml:"cx,attr"`
	Y           int    `xml:"cy,attr"`
	Fill        string `xml:"fill,attr"`
	Stroke      string `xml:"stroke,attr"`
	StrokeWidth int    `xml:"stroke-width,attr"`
}

// Line ...
type Line struct {
	X1     int    `xml:"x1,attr"`
	Y1     int    `xml:"y1,attr"`
	X2     int    `xml:"x2,attr"`
	Y2     int    `xml:"y2,attr"`
	Stroke string `xml:"stroke,attr"`
}

// Group is a substruct of Svg
type Group struct {
	Title   string   `xml:"title"`
	Circles []Circle `xml:"circle"`
	Lines   []Line   `xml:"line"`
}

// Svg is the stuct to be saved to a file.
type Svg struct {
	XMLName  xml.Name `xml:"svg"`
	Width    int      `xml:"width,attr"`
	Height   int      `xml:"height,attr"`
	Xmlns    string   `xml:"xmlns,attr"`
	XmlnsSvg string   `xml:"xmlns:svg,attr"`
	Groups   []Group  `xml:"g"`
}

func saveSvg(svg *Svg, filename string) error {
	file, _ := os.Create(filename)

	xmlWriter := io.Writer(file)

	enc := xml.NewEncoder(xmlWriter)
	enc.Indent("  ", "    ")
	if err := enc.Encode(svg); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return nil

}
