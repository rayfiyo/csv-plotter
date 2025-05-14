// CSVを読み Domain に変換する、読み込みの責務
// I/O に依存する実装なので internal/adapter に配置

package csv_reader

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/rayfiyo/csv-plotter/internal/domain"
)

// Reader は CSV → DataSet へ変換する
type Reader struct{}

// New は Reader のコンストラクタ
func New() *Reader { return &Reader{} }

// Read は CSVを読み Domain に変換
func (Reader) Read(path string) (domain.DataSet, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comment = '#'           // コメント行を無視
	r.TrimLeadingSpace = true // フィールドの頭にある空白を除外（1, 2 をうまく処理）
	recs, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	data := make(domain.DataSet, len(recs))
	for i, rec := range recs {
		if len(rec) < 2 {
			return nil, fmt.Errorf("line %d: want 2 columns", i+1)
		}
		x, err := strconv.ParseFloat(rec[0], 64)
		if err != nil {
			return nil, fmt.Errorf("line %d: %w", i+1, err)
		}
		y, err := strconv.ParseFloat(rec[1], 64)
		if err != nil {
			return nil, fmt.Errorf("line %d: %w", i+1, err)
		}
		data[i] = domain.Point{X: x, Y: y}
	}
	return data, nil
}
