package main

import (
	"log"
	"os"

	"github.com/rayfiyo/csv-plotter/internal/adapter/csv_reader"
	"github.com/rayfiyo/csv-plotter/internal/adapter/plot_writer"
	"github.com/rayfiyo/csv-plotter/internal/usecase"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s <data.csv>", os.Args[0])
	}
	csvPath := os.Args[1]

	reader := csv_reader.New()
	writer := plot_writer.New()
	service := usecase.NewPlotService(reader, writer)

	if err := service.Execute(csvPath); err != nil {
		log.Fatal(err)
	}
}
