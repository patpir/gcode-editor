package pattern

import (
	"fmt"

	"github.com/patpir/svgshapes"
	"github.com/patpir/gcode-pattern-detect/analysis"
)

// Plot adds graphs to the target SVG document.
//
// The allPlots parameter gives a list of available plots,
// but only the plots listed in the display array will be plotted.
// For this purpose, the Id of each PlotInfo in display will be looked up
// in the allPlots map.
// If a PlotInfo's Id is not found in allPlots, it is silently ignored.
func Plot(target *svgshapes.Document, allPlots map[string]*analysis.Source, display []PlotInfo) {
	legendGroup := target.Group()
	graphGroup := target.Group()

	graphCount := 0

	for _, graph := range display {
		if graph.Scaling == 0.0 {
			graph.Scaling = 1.0
		}

		source, ok := allPlots[graph.Id]
		if !ok {
			continue
		}

		graphLine := graphGroup.Polyline()
		graphLine.Stroke = graph.Color
		graphLine.Fill = "none"

		keys := source.Keys()
		for _, key := range keys {
			for _, value := range source.ValuesAt(key) {
				graphLine.Points = append(graphLine.Points, svgshapes.Point{
					X: key,
					Y: value * graph.Scaling,
				})
			}
		}

		legendItem := legendGroup.Group()
		legendItem.Circle(20, float64(graphCount * 40 + 20), 10).Fill = graph.Color
		legendItem.Text(40, float64(graphCount * 40 + 20), graph.Name)

		graphCount += 1
	}

	graphGroup.Style = fmt.Sprintf("transform: translate(0, %d);", graphCount * 40)
}
