// Domain を Plot にして PNG 保存する、書き込みの責務
// I/O に依存する実装なので internal/adapter に配置

package plot_writer

import (
	"fmt"
	"path/filepath"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"github.com/rayfiyo/csv-plotter/internal/domain"
)

// Writer は DataSet → PNG を担当
type Writer struct{}

// New は Writer のコンストラクタ
func New() *Writer { return &Writer{} }

// threeTicks は最小・中央値・最大値のみ表示する Ticker
type threeTicks struct{}

func (threeTicks) Ticks(min, max float64) []plot.Tick {
	mid := (min + max) / 2
	return []plot.Tick{
		{Value: min, Label: fmt.Sprintf("%.0f", min)},
		{Value: mid, Label: fmt.Sprintf("%.0f", mid)},
		{Value: max, Label: fmt.Sprintf("%.0f", max)},
	}
}

// Write は baseName.png へ散布図を保存する
func (Writer) Write(baseName string, data domain.DataSet) error {
	pts := make(plotter.XYs, len(data))
	for i, p := range data {
		pts[i].X, pts[i].Y = p.X, p.Y
	}

	p := plot.New()
	p.X.Tick.Marker = threeTicks{}
	p.Y.Tick.Marker = threeTicks{}

	minX, maxX, minY, maxY := data.Bounds()
	p.X.Min, p.X.Max = minX, maxX
	p.Y.Min, p.Y.Max = minY, maxY

	p.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(pts)
	if err != nil {
		return err
	}
	p.Add(s)

	out := filepath.Join(".", baseName+".png")
	return p.Save(6*vg.Inch, 6*vg.Inch, out)
}
