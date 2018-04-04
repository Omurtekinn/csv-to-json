package main

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
)

type Record struct {
	Plaka int64
	Bolge string
	Il    string
	Ilce  string
	// High, Low, Close
}

func main() {
	src, err := os.Open("illerilceler.csv")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create("il-ilce.json")
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	rows, err := csv.NewReader(src).ReadAll()
	if err != nil {
		panic(err)
	}

	records := make([]Record, 0, len(rows))
	for _, row := range rows {
		plaka, _ := strconv.ParseInt(row[0], 0, 64)
		bolge := row[1]
		il := row[2]
		ilce := row[3]

		records = append(records, Record{
			Plaka: plaka,
			Bolge: bolge,
			Il:    il,
			Ilce:  ilce,
		})
	}

	err = json.NewEncoder(dst).Encode(records)
	if err != nil {
		panic(err)
	}

}
