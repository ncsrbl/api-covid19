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

	//Em An�lise
	var r = []string{"Positivo"}
	SearchByResult(r, dataset)
	r = []string{"Em An�lise", ""}
	var w = []string{"Fortaleza"}
	SearchByMunicipios(w, SearchByObito(dataset))
	fmt.Println("\n\n")

	SearchByRegiaoDeSaude("JAGUARIBE", dataset)
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

		_, err = reader.ReadAll()
		if err != nil {
			return nil, err
		}
		return GetCSV(fileURL)
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

			if strings.EqualFold(strings.ToUpper(row[32]), strings.ToUpper(resultadoExame[i])) && (row[25] == "CE" || row[25] == "") {
				results = append(results, row)
			}
		}
	}
	fmt.Printf("Found %d results\n", len(results))
	return results
}

//SearchByObito retorna todos os casos de obito
func SearchByObito(dataset [][]string) (results [][]string) {
	fmt.Printf("Searching for obitos\n\t")
	var o = "true"
	for _, row := range dataset {
		if strings.EqualFold(strings.ToUpper(row[30]), strings.ToUpper(o)) {
			results = append(results, row)
		}
	}
	fmt.Printf("Found %d results\n", len(results))
	return results
}

//SearchByTodasRegioesDeSaude pesquisa pela regiao de saude
func SearchByTodasRegioesDeSaude(dataset [][]string) (results [][]string) {
	fmt.Printf("Seaching for all regions\n")
	r := []string{"Positivo"}
	SearchByResult(r, dataset)
	r = []string{"Negativo"}
	SearchByResult(r, dataset)
	r = []string{"Em An�lise", ""}
	SearchByResult(r, dataset)
	SearchByObito(dataset)
	return results
}

//SearchByRegiaoDeSaude busca pela regiao de saude
func SearchByRegiaoDeSaude(regiao string, dataset [][]string) (results [][]string) {
	var municipios []string
	if strings.ToUpper(regiao) == "CARIRI" {
		municipios = []string{"BARBALHA", "IGUATU", "JUAZEIRO DO NORTE", "CRATO", "QUIXELO", "FARIAS BRITO", "BREJO SANTO", "CARIUS", "CARIRIACU",
			"MOMBACA", "JUCAS", "MISSAO VELHA", "ASSARE", "PIQUET CARNEIRO", "ICO", "ACOPIARA", "CATARINA", "NOVA OLINDA", "SANTANA DO CARIRI", "VARZEA ALEGRE", "CEDRO",
			"MILAGRES", "OROS", "BAIXIO", "PORTEIRAS", "LAVRAS DA MAGABEIRA", "JARDIM", "BARRO", "IPAUMIRIM", "AURORA", "UMARI", "ABAIARA", "CAMPOS SALES", "ARARIPE", "JATI",
			"GRANJEIRO", "ANTONINA DO NORTE", "TARRAFAS", "PENAFORTE", "SALITRE", "ALTANEIRA", "DEPUTADO IRAPUAN PINHEIRO", "SABOEIRO", "PONTENGI"}
	}
	if strings.ToUpper(regiao) == "SOBRAL" {
		municipios = []string{"SOBRAL", "CAMOCIM", "ACARAU", "ITAREMA", "TIANGUA", "BELA CRUZ", "MASSAPE", "CRATEUS", "VICOSA DO CEARA", "SANTA QUITERIA", "UBAJARA",
			"CRUZ", "GRANJA", "BARROQUINHA", "COREAU", "CHAVAL", "CARIRE", "MORRINHOS", "GROAIRAS", "ALCANTARAS", "MERUOCA", "JIJOCA DE JERICOACOARA", "MORAUJO", "URUOCA",
			"SENADOR SA", "SANTANA DO ACARAU", "VARJOTA", "SAO BENEDITO", "IPUEIRAS", "IBIAPINA", "GUARACIABA DO NORTE", "MUCAMBO", "NOVA RUSSAS", "IPU", "MARCO", "IRAUCUBA",
			"GRACA", "CARNAUBAL", "FRECHEIRINHA", "TAMBORIL", "FORQUILHA", "MONSENHOR TABOSA", "CATUNDA", "IPAPORANGA",
			"HIDROLANDIA", "RERIUTABA", "INDEPENDENCIA", "PACUJA", "QUITERIANOPOLIS", "MARTINOPOLE", "NOVO ORIENTE", "CROATA", "PORANGA", "PIRES FERREIRA", "ARARENDA"}
	}
	if strings.ToUpper(regiao) == "FORTALEZA" {
		municipios = []string{"FORTALEZA", "CAUCAIA", "MARACANAU", "ITAPIPOCA", "SAO GONCALO DO AMARANTE", "MARANGUAPE", "EUSEBIO", "PACATUBA", "CASCAVEL", "HORIZONTE",
			"REDENCAO", "PACAJUS", "AQUIRAZ", "AMONTADA", "URUBURETAMA", "ITAPAJE", "TRAIRI", "ITAITINGA", "PENTECOSTE", "ACARAPE", "ARACOIABA", "PARACURU", "BATURITÉ", "PARAIPABA",
			"CAPISTRANO", "BEBERIBE", "CHOROZINHO", "PINDORETAMA", "TEJUCUOCA", "MIRAIMA", "OCARA", "GUAIÚBA", "TURURU", "BARREIRA", "UMIRIM", "PALMACIA", "PACOTI", "ITAPIUNA",
			"APUIARES", "SAO LUÍS DO CURU", "GENERAL SAMPAIO", "ARATUBA", "MULUNGU", "GUARAMIRANGA"}
	}
	if strings.ToUpper(regiao) == "SERTÃO CENTRAL" {
		municipios = []string{"QUIXADA", "CANINDE", "QUIXERAMOBIM", "ITATIRA", "TAUA", "IBICUITINGA", "BANABUIU", "CHORÓ",
			"CARIDADE", "MADALENA", "PARAMBU", "BOA VIAGEM", "SOLONOPOLE", "MILHA", "PARAMOTI", "PEDRA BRANCA", "SENADOR POMPEU", "ARNEIROZ", "IBARETAMA", "AIUABA"}
	}
	if strings.ToUpper(regiao) == "LITORAL LESTE" || strings.ToUpper(regiao) == "JAGUARIBE" {
		municipios = []string{"MORADA NOVA", "RUSSAS", "ARACATI", "LIMOEIRO DO NORTE", "TABULEIRO DO NORTE", "JAGUARUANA", "ICAPUI", "JAGUARIBE", "QUIXERE", "JAGUARIBARA",
			"ITAICABA", "SAO JOAO DO JAGUARIBE", "JAGUARETAMA", "ALTO SANTO", "FORTIM", "ERERE", "PALHANO", "IRACEMA", "PEREIRO", "POTIRETAMA"}
	}
	r := []string{"Positivo"}
	positivos := SearchByResult(r, dataset)

	for i := range municipios {
		SearchByMunicipio(municipios[i], positivos)
	}

	return SearchByMunicipios(municipios, positivos)
}

func SearchByAreaDescentralizada(regiao string, dataset [][]string) (results [][]string) {

	return results
}

//SearchByMunicipios busca pelo resultado do exame
func SearchByMunicipios(municipios []string, dataset [][]string) (results [][]string) {
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

//SearchByMunicipio busca por um municipio em especifico
func SearchByMunicipio(municipio string, dataset [][]string) (results [][]string) {
	fmt.Printf("Searching for %s\n\t", municipio)
	for _, row := range dataset {
		if strings.ToUpper(row[29]) == strings.ToUpper(municipio) {
			results = append(results, row)
		}

	}
	//fmt.Printf("%s\n", results)
	fmt.Printf("Found %d results\n", len(results))

	return results
}
