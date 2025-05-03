# plotter

- A Go program to read a CSV file and plot it on a graph
- CSV ファイルを読み取ってグラフにプロットする Go のプログラム

# dir

```
csv-plotter/
├── changelog.config.js
├── cmd/                    エントリポイント
│   └── main.go
├── go.mod
├── go.sum
├── internal/
│   ├── adapter/            I/O の窓口（CSV 読み・画像書き等）
│   │   ├── csv_reader
│   │   │   └── reader.go
│   │   └── plot_writer
│   │       └── writer.go
│   ├── domain/             抽象モデル（データ点・設定値）
│   │   └── model.go
│   └── usecase/            何をするか を示すサービス
│       └── plot_service.go
├── LICENSE
├── pkg/                    ライブラリとして公開するときのAPI用
└── README.md
```
