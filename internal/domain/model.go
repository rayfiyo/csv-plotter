// 不変データ構造と共通メソッド

package domain

// Point は 2 次元上の 1 データ点（１つの観測点）を表すドメインオブジェクト
// X, Y ともに plot 軸でそのまま使えるよう float64 で
type Point struct {
	X, Y float64
}

// DataSet は散布図のデータ列
type DataSet []Point

// Bounds はグラフの境界（値の範囲）を返す
func (d DataSet) Bounds() (minX, maxX, minY, maxY float64) {
	// 事実上のエラーハンドリング
	if len(d) == 0 {
		return
	}

	// 基準の値（初期値に近い意味）
	minX, maxX = d[0].X, d[0].X
	minY, maxY = d[0].Y, d[0].Y

	// 値の更新
	for _, p := range d {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	return
}
