//  汎用ステップティッカー

package plot_writer

import (
	"fmt"
	"math"

	"gonum.org/v1/plot"
)

type StepTicker struct {
	Minor, Major float64
}

func (t StepTicker) Ticks(min, max float64) []plot.Tick {
	// 端をキレイに揃える
	start := math.Floor(min/t.Minor) * t.Minor
	end := math.Ceil(max/t.Minor) * t.Minor

	ticks := make([]plot.Tick, 0, int((end-start)/t.Minor)+3)
	for v := start; v <= end+1e-9; v += t.Minor { // ε で丸め誤差対策
		lbl := ""
		if math.Mod(math.Abs(v), t.Major) < 1e-9 { // メジャー位置だけラベル
			lbl = fmt.Sprintf("%.0f", v)
		}
		ticks = append(ticks, plot.Tick{Value: v, Label: lbl})
	}

	return ticks
}
