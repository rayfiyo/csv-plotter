//  汎用ステップティッカー

package plot_writer

import (
	"fmt"
	"math"
	"sort"

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

	// min / max 専用目盛り
	ensure := func(val float64) {
		const eps = 1e-9
		for _, tk := range ticks {
			if math.Abs(tk.Value-val) < eps {
				return // 既にある
			}
		}
		ticks = append(ticks, plot.Tick{
			Value: val,
			Label: fmt.Sprintf("%.0f", val), // 常にラベル付き
		})
	}
	ensure(min)
	ensure(max)

	// 並び順を保証
	sort.Slice(ticks, func(i, j int) bool { return ticks[i].Value < ticks[j].Value })

	return ticks
}
