package app

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gmazoyer/peeringdb"
)

type App struct {
	Asns   [][]string
	PdbApi *peeringdb.API
}

// todo: build sqlite database
func ReadAsnsCsv() [][]string {
	f, err := os.Open("ipinfo_lite.csv")

	if err != nil {
		log.Fatal("couldn't read csv", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("couldn't parse csv", err)
	}

	return records
}

func NewApp() *App {
	return &App{
		Asns:   ReadAsnsCsv(),
		PdbApi: peeringdb.NewAPI(),
	}
}
