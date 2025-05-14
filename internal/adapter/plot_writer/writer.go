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

	// 余白を 5 % 追加
	dx := maxX - minX
	dy := maxY - minY
	const pad = 0.05 // 5 %
	if dx == 0 {
		dx = 1
	} // 全点同値でも 1px 分は空ける
	if dy == 0 {
		dy = 1
	}

	p.X.Min = minX - dx*pad
	p.X.Max = maxX + dx*pad
	p.Y.Min = minY - dy*pad
	p.Y.Max = maxY + dy*pad

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
