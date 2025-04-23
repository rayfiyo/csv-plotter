package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// threeTicks は最小・中央値・最大値だけラベルを付ける Ticker
type threeTicks struct{}

func (threeTicks) Ticks(min, max float64) []plot.Tick {
	mid := (min + max) / 2
	return []plot.Tick{
		{Value: min, Label: fmt.Sprintf("%.0f", min)},
		{Value: mid, Label: fmt.Sprintf("%.0f", mid)},
		{Value: max, Label: fmt.Sprintf("%.0f", max)},
	}
}

func main() {
	// 実行時引数の処理
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s <data.csv>", os.Args[0])
	}
	path := os.Args[1]

	// ファイルの開閉
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// CSV の読み取り
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// グラフ領域の上限下限
	pts := make(plotter.XYs, len(records))
	minX, maxX := 1e9, -1e9
	minY, maxY := 1e9, -1e9

	// バリデーション
	for i, rec := range records {
		if len(rec) < 2 {
			log.Fatalf("line %d: need 2 columns", i+1)
		}
		x, err := strconv.ParseFloat(rec[0], 64)
		if err != nil {
			log.Fatalf("line %d: %v", i+1, err)
		}
		y, err := strconv.ParseFloat(rec[1], 64)
		if err != nil {
			log.Fatalf("line %d: %v", i+1, err)
		}

		pts[i].X, pts[i].Y = x, y
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}

	// プロットの作成と調整
	p := plot.New()
	p.X.Label.Text = ""
	p.Y.Label.Text = ""

	p.X.Min, p.X.Max = minX, maxX
	p.Y.Min, p.Y.Max = minY, maxY

	p.X.Tick.Marker = threeTicks{}
	p.Y.Tick.Marker = threeTicks{}

	p.Add(plotter.NewGrid()) // 薄い格子

	// プロット
	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	p.Add(s)

	// 軸のスタイル調整

	// 画像を保存
	filename := filepath.Base(path[:len(os.Args[1])-len(filepath.Ext(os.Args[1]))])
	if err := p.Save(6*vg.Inch, 6*vg.Inch, filename+".png"); err != nil {
		log.Fatalf("save pdf: %v", err)
	}
}
