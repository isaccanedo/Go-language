package main

import (
	"math"

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

const points = 100

func main() {
	plot := NewPlot(2)
	defer plot.Close()

	var pointsX [points]float64
	for i := 0; i < points; i++ {
		pointsX[i] = float64(i) * 2.0 * math.Pi / points
	}
	function := math.Sin

	plot.AddFunc2d("sin t", "lines", pointsX[:], function)

	plot.SetTitle("Plot #10")
	plot.SetXLabel("t")
	plot.SetYLabel("sin t")

	plot.SetXrange(0, int(math.Round(2.0*math.Pi)))
	plot.SetYrange(-1, 1)
	plot.Save("glot10.png")
}
