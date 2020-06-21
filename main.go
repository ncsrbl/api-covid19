package main

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"
)

func main() {

	url := "https://indicadores.integrasus.saude.ce.gov.br/api/casos-coronavirus/export-csv"

	data, err := GetCSV(url)

	if err != nil {
		panic(err)
	}

	Search(data, "Positivo")
	//Search(data, "Morada Nova", "Positivo")

	for col, _ := range data {
		if col == 0 {
			//continue
		}
		if col == 10 {
			break
		}
		//fmt.Println(row)
	}

}

func GetCSV(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)
	reader.Comma = ','

	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func Search(data [][]string, filtros ...string) []string {
	var bairros []string
	numFilters := len(filtros)
	fmt.Printf("Seaching for %s\n", filtros)
	n := 0
	for _, row := range data {
		n = 0
		for i := range row {
			for j := range filtros {
				if i == 32 {
					//fmt.Printf("%s %s %t\n", row[i], filtros[j], row[i] == strings.ToUpper(filtros[j]))
				}
				if strings.ToUpper(row[i]) == strings.ToUpper(filtros[j]) {
					//fmt.Printf("%s %s => ", row[0], filtros[j])
					n += 1
				}
			}
		}
		if n == numFilters {
			//fmt.Printf("\t%s %s %s\n", row[0], row[32], row[33])
			bairros = append(bairros, row[32])
		}

	}
	fmt.Printf("Found %d results...\n", len(bairros))
	return bairros
}
