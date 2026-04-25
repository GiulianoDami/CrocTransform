package visualizer

import (
	"image/color"
	"os"
	"path/filepath"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// FossilData represents the structure of fossil data for visualization
type FossilData struct {
	SpecimenName string
	LimbLengths  []float64
	JointAngles  []float64
	GrowthData   []float64
}

// GeneratePlot creates a visualization plot from fossil data
func GeneratePlot(data FossilData, outputPath string) error {
	// Create a new page with charts
	page := components.NewPage()

	// Create line chart for limb lengths
	lineChart := charts.NewLine()
	lineChart.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Limb Length Analysis",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Length (mm)",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Age Stage",
		}),
	)

	lineChart.SetXAxis([]string{"Juvenile", "Subadult", "Adult"}).
		AddSeries("Limb Lengths", generateLineData(data.LimbLengths)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{
			ShowSymbol: opts.Bool(true),
		}))

	// Create bar chart for joint angles
	barChart := charts.NewBar()
	barChart.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Joint Angle Analysis",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Angle (degrees)",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Joint Type",
		}),
	)

	barChart.SetXAxis([]string{"Hip", "Knee", "Ankle", "Shoulder", "Elbow", "Wrist"}).
		AddSeries("Joint Angles", generateBarData(data.JointAngles)).
		SetSeriesOptions(charts.WithBarChartOpts(opts.BarChart{
			ShowBackground: opts.Bool(true),
		}))

	// Add charts to page
	page.AddCharts(lineChart, barChart)

	// Ensure output directory exists
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	// Save the chart to file
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return page.Render(f)
}

// Helper function to generate line chart data
func generateLineData(values []float64) []opts.LineData {
	var data []opts.LineData
	for i, v := range values {
		data = append(data, opts.LineData{
			Value: v,
			ItemStyle: &opts.ItemStyle{
				Color: color.RGBA{R: 255, G: 99, B: 71, A: 255},
			},
		})
	}
	return data
}

// Helper function to generate bar chart data
func generateBarData(values []float64) []opts.BarData {
	var data []opts.BarData
	for i, v := range values {
		data = append(data, opts.BarData{
			Value: v,
			ItemStyle: &opts.ItemStyle{
				Color: color.RGBA{R: 50, G: 205, B: 50, A: 255},
			},
		})
	}
	return data
}