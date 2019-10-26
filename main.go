package main

import (
    "encoding/xml"
    "fmt"
    "os"
    "io"
)

func main() {
    type Circle struct {
        ID string `xml:"id,attr"`
        Radius int `xml:"r,attr"`
        X int `xml:"cx,attr"`
        Y int `xml:"cy,attr"`
        Fill string `xml:"fill,attr"`
        Stroke string `xml:"stroke,attr"`
        StrokeWidth int `xml:"stroke-width,attr"`
    }

    type Group struct {
        Title string `xml:"title"`
        Circles []Circle `xml:"circle"`
    }

    type Svg struct {
        XMLName xml.Name `xml:"svg"`
        Width int `xml:"width,attr"`
        Height int `xml:"height,attr"`
        Xmlns string `xml:"xmlns,attr"`
        XmlnsSvg string `xml:"xmlns:svg,attr"`
        Groups []Group `xml:"g"`
    }

    svg := &Svg{Width: 640, Height: 480, Xmlns: "http://www.w3.org/2000/svg", XmlnsSvg: "http://www.w3.org/2000/svg"}
    svg.Groups = append(svg.Groups, Group{Title: "This is a title"})
    svg.Groups[0].Circles = append(svg.Groups[0].Circles, Circle{ID: "circ", Radius: 50, X: 320, Y: 240, Fill: "#00ffff", Stroke: "#000000", StrokeWidth: 7})

    filename := "svg.xml"
    file, _ := os.Create(filename)

    xmlWriter := io.Writer(file)

    enc := xml.NewEncoder(xmlWriter)
    enc.Indent("  ", "    ")
    if err := enc.Encode(svg); err != nil {
            fmt.Printf("error: %v\n", err)
    }
}

