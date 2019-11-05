package main

// Organization ...
type Organization struct {
	Founded     int
	isDissolved bool
	Dissolved   int
}

// Data ...
type Data struct {
	Organizations []Organization
}

func main() {
	svg := &Svg{Width: 640, Height: 480, Xmlns: "http://www.w3.org/2000/svg", XmlnsSvg: "http://www.w3.org/2000/svg"}
	svg.Groups = append(svg.Groups, Group{Title: "This is a title"})
	svg.Groups[0].Circles = append(svg.Groups[0].Circles, Circle{ID: "circ", Radius: 50, X: 320, Y: 240, Fill: "#00ffff", Stroke: "#000000", StrokeWidth: 7})

	svg.Groups[0].Lines = append(svg.Groups[0].Lines, Line{X1: 0, Y1: 0, X2: 640, Y2: 0, Stroke: "black"})
	svg.Groups[0].Lines = append(svg.Groups[0].Lines, Line{X1: 640, Y1: 0, X2: 640, Y2: 480, Stroke: "black"})
	svg.Groups[0].Lines = append(svg.Groups[0].Lines, Line{X1: 640, Y1: 480, X2: 0, Y2: 480, Stroke: "black"})
	svg.Groups[0].Lines = append(svg.Groups[0].Lines, Line{X1: 0, Y1: 480, X2: 0, Y2: 0, Stroke: "black"})

	svg.Groups[0].Lines = append(svg.Groups[0].Lines, Line{X1: 10, Y1: 470, X2: 630, Y2: 470, Stroke: "grey"})

	saveSvg(svg, "picture.xml")
}
