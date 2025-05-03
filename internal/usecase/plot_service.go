// reader.go → write.go の組織化
// このプログラム全体が、CSV を読み取って、プロットし、png に書き出すことを示す

package usecase

import (
	"path/filepath"
	"strings"

	"github.com/rayfiyo/csv-plotter/internal/domain"
)

// Reader と Writer のポート定義
type Reader interface {
	Read(path string) (domain.DataSet, error)
}
type Writer interface {
	Write(baseName string, data domain.DataSet) error
}

// PlotService はユースケースの集約
type PlotService struct {
	reader Reader
	writer Writer
}

// NewPlotService は DI コンストラクタ
func NewPlotService(r Reader, w Writer) *PlotService {
	return &PlotService{reader: r, writer: w}
}

// Execute は CSV を読み取って PNG を出力する
func (s *PlotService) Execute(csvPath string) error {
	data, err := s.reader.Read(csvPath)
	if err != nil {
		return err
	}

	base := strings.TrimSuffix(filepath.Base(csvPath), filepath.Ext(csvPath))
	return s.writer.Write(base, data)
}
