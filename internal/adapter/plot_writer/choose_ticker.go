// 範囲→ステップを決めるヘルパ

package plot_writer

import (
	"math"

	"gonum.org/v1/plot"
)

func chooseTicker(rng float64) plot.Ticker {
	switch {
	case rng <= 10:
		return StepTicker{Minor: 1, Major: 5}
	case rng <= 100:
		return StepTicker{Minor: 5, Major: 10}
	case rng <= 2000:
		return StepTicker{Minor: 50, Major: 100}
	default:
		// “いい感じ”の値を自動で算出 (Renard 系)
		mag := math.Pow(10, math.Floor(math.Log10(rng)))
		base := rng / mag
		var major float64
		switch {
		case base <= 2.5:
			major = 0.5 * mag
		case base <= 5:
			major = 1 * mag
		case base <= 10:
			major = 2 * mag
		default:
			major = 5 * mag
		}
		return StepTicker{Minor: major / 2, Major: major}
	}
}
