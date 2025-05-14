//  各設定を反映した書き込み（プロット）処理

package plot_writer

import (
	"path/filepath"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"github.com/rayfiyo/csv-plotter/internal/domain"
)

type Writer struct{}

func New() *Writer { return &Writer{} }

func (Writer) Write(baseName string, data domain.DataSet) error {
	pts := make(plotter.XYs, len(data))
	for i, p := range data {
		pts[i].X, pts[i].Y = p.X, p.Y
	}

	p := plot.New()

	// 軸範囲
	minX, maxX, minY, maxY := data.Bounds()
	p.X.Min, p.X.Max = minX, maxX
	p.Y.Min, p.Y.Max = minY, maxY

	// ティッカー選定
	p.X.Tick.Marker = chooseTicker(maxX - minX)
	p.Y.Tick.Marker = chooseTicker(maxY - minY)

	// p.Add(plotter.NewGrid()) // グリッド線

	// プロット
	s, err := plotter.NewScatter(pts)
	if err != nil {
		return err
	}
	p.Add(s)

	out := filepath.Join(".", baseName+".png")
	return p.Save(6*vg.Inch, 6*vg.Inch, out)
}
