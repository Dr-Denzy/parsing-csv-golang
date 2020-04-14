package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	rows := readFromCSV("data.csv")

	rows = computeResults(rows)

	writeToCSV("data-output.csv", rows)
}

func readFromCSV(filename string) [][]string {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", filename, err.Error())
	}

	defer file.Close()

	input := csv.NewReader(file)
	input.Comma = ','
	rows, err := input.ReadAll()

	if err != nil {
		log.Fatalln("CSV file cannot be read: ", err.Error())

	}

	return rows

}

func computeResults(rows [][]string) [][]string {

	for i := range rows {
		if i == 0 {
			rows[0] = append(rows[0], "BatAvg", "HomeRunsPerHit", "OnBasePctg", "SlgPctg", "OnBaseSluggingPctg")
			continue
		}

		player := rows[i][0]

		atbats, err := strconv.ParseFloat(rows[i][5], 64)
		if err != nil {
			log.Fatalf("Cannot retrieve data -> %s: %s\n", player, err)
		}

		hits, err := strconv.ParseFloat(rows[i][7], 64)
		if err != nil {
			log.Fatalf("Cannot retrieve data -> %s: %s\n", player, err)
		}

		twobase, err := strconv.ParseFloat(rows[i][8], 64)
		if err != nil {
			log.Fatalf("Cannot retrieve data -> %s: %s\n", player, err)
		}

		threebase, err := strconv.ParseFloat(rows[i][9], 64)
		if err != nil {
			log.Fatalf("Cannot retrieve data -> %s: %s\n", player, err)
		}

		homeruns, err := strconv.ParseFloat(rows[i][10], 64)
		if err != nil {
			log.Fatalf("Cannot retrieve data -> %s: %s\n", player, err)
		}

		bbwalks, err := strconv.ParseFloat(rows[i][12], 64)
		if err != nil {
			log.Fatalf("Cannot retrieve data -> %s: %s\n", player, err)
		}

		// 	Calculate each players Batting Average (BatAvg).

		batAvg := hits / atbats

		rows[i] = append(rows[i], fmt.Sprintf("%.3f", batAvg))

		// Calculate each players Home Runs per Hit (HomeRunsPerHit).
		homeRunsPerHit := homeruns / hits

		rows[i] = append(rows[i], fmt.Sprintf("%.3f", homeRunsPerHit))

		// Calculate each players On Base Percentage (OnBasePctg).
		onBasePctg := (hits + bbwalks) / (atbats + bbwalks)

		rows[i] = append(rows[i], fmt.Sprintf("%.3f", onBasePctg))

		// Calculate each players Slugging Percentage (SlgPctg).
		slgPctg := ((hits) + (2 * twobase) + (3 * threebase) + (4 * homeruns)) / atbats

		rows[i] = append(rows[i], fmt.Sprintf("%.3f", slgPctg))

		// Calculate each players On – Base plus Slugging Percentage (OnBaseSluggingPctg).
		onBaseSluggingPctg := onBasePctg + slgPctg

		rows[i] = append(rows[i], fmt.Sprintf("%.3f", onBaseSluggingPctg))

	}

	return rows
}

func writeToCSV(filename string, rows [][]string) {

	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Cannot open '%s' : %s\n", filename, err.Error())
	}

	defer func() {
		e := file.Close()
		if e != nil {
			log.Fatalf("Cannot close '%s': %s\n", filename, e.Error())
		}
	}()

	w := csv.NewWriter(file)
	err = w.WriteAll(rows)

}
