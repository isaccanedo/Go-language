package main

import (
	"github.com/Arafatk/glot"
)

// Plot represents context for a graph/chart
type Plot struct {
	*glot.Plot
}

// Save method saves graph or chart into PNG file
func (plot *Plot) Save(filename string) {
	plot.Cmd("set terminal pngcairo")
	plot.Cmd("set output '" + filename + "'")
	plot.Cmd("replot")
}

// NewPlot function is a constructor for Plot data structure
func NewPlot(dimensions int) *Plot {
	plot, err := glot.NewPlot(dimensions, false, false)
	if err != nil {
		panic(err)
	}
	return &Plot{plot}

}

func main() {
	plot := NewPlot(2)
	defer plot.Close()

	plot.AddPointGroup("Measured data", "lines", []float32{1.2, 1.9, 3.9, 8.2, 9.5, 8.1, 4.1, 1.8, 1.5})
	plot.AddPointGroup("Expected data", "lines", []float32{1, 2, 4, 8, 9, 8, 4, 2, 1})

	plot.SetXrange(-2, 10)
	plot.SetYrange(0, 10)
	plot.Save("glot06.png")
}
