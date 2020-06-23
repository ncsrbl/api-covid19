package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	url := "https://indicadores.integrasus.saude.ce.gov.br/api/casos-coronavirus/export-csv"

	dataset, err := GetFileByURL("covid.csv", url)

	var r = []string{"Positivo", "Negativo"}
	SearchByResult(r, dataset)
	r = []string{"Positivo"}
	var w = []string{"Iguatu"}
	SearchByMunicipio(w, SearchByResult(r, dataset))
	if err != nil {
		panic(err)
	}

}

//GetCSV Pega o csv da url
func GetCSV(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)
	fmt.Printf("Data -> %s\n", resp.Header.Get("Date"))
	date := resp.Header.Get("Date")
	split := strings.Split(string(date), " ")
	for i := range split {
		fmt.Printf("%s -- ", split[i])
	}
	fmt.Printf("\n")
	reader.Comma = ','

	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

//GetFileByURL bla bla bla
func GetFileByURL(filepath string, fileURL string) ([][]string, error) {
	file, err := os.Open("covid.csv")

	if err != nil {
		resp, err := http.Get(fileURL)

		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		out, err := os.Create(filepath)

		if err != nil {
			return nil, err
		}
		defer out.Close()
		_, err = io.Copy(out, resp.Body)

		reader := csv.NewReader(resp.Body)
		reader.Comma = ','

		dataset, err := reader.ReadAll()
		if err != nil {
			return nil, err
		}
		return dataset, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','

	dataset, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return dataset, err
}

//Search faz a pesquisa de acordo com as strings passadas
func Search(resultadoExame []string, municipio []string, regiao string, area string) (res []string) {

	return res
}

//SearchByResult busca pelo resultado do exame
func SearchByResult(resultadoExame []string, dataset [][]string) (results [][]string) {
	fmt.Printf("Searching for %s\n\t", resultadoExame)
	for _, row := range dataset {
		for i := range resultadoExame {
			if row[32] == resultadoExame[i] {
				results = append(results, row)
			}
		}
	}
	fmt.Printf("Found %d results\n", len(results))
	return results
}

//SearchByMunicipio busca pelo resultado do exame
func SearchByMunicipio(municipios []string, dataset [][]string) (results [][]string) {
	fmt.Printf("Searching for %s\n\t", municipios)
	for _, row := range dataset {
		for i := range municipios {
			if strings.ToUpper(row[29]) == strings.ToUpper(municipios[i]) {
				results = append(results, row)
			}
		}
	}
	//fmt.Printf("%s\n", results)
	fmt.Printf("Found %d results\n", len(results))

	return results
}

// var regiaoCariri = string[] {"",}
