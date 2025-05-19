package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Linha struct {
	Atributo1 string
	Atributo2 int
	Atributo3 float64
	Atributo4 bool
}

const jsonFile = "arquivo.json"
const jsonFileEncoder = "arquivoEncoder.json"
const csvFile = "arquivo.csv"

func main() {
	fmt.Println("Criando arquivo com a função createJsonWithMarshal")
	createJsonWithMarshal()
	fmt.Println("Lendo arquivo com a função readJsonWithUnmarshal")
	contentJson1 := readJsonWithUnmarshal(jsonFile)
	for _, linha := range contentJson1 {
		fmt.Println("Atributo1:", linha.Atributo1)
		fmt.Println("Atributo2:", linha.Atributo2)
		fmt.Println("Atributo3:", linha.Atributo3)
		fmt.Println("Atributo4:", linha.Atributo4)
		fmt.Println()
	}

	fmt.Println()
	fmt.Println("Criando arquivo com a função createJsonWithEncoder")
	createJsonWithEncoder()
	fmt.Println("Lendo arquivo com a função readJsonWithDecoder")
	contentJson2 := readJsonWithDecoder(jsonFileEncoder)
	for _, linha := range contentJson2 {
		fmt.Println("Atributo1:", linha.Atributo1)
		fmt.Println("Atributo2:", linha.Atributo2)
		fmt.Println("Atributo3:", linha.Atributo3)
		fmt.Println("Atributo4:", linha.Atributo4)
		fmt.Println()
	}

	fmt.Println()
	fmt.Println("Criando arquivo com a função createCsv")
	createCsv()
	fmt.Println("Lendo arquivo com a função readCsv")
	contentCsv := readCsv()
	for _, linha := range contentCsv {
		fmt.Println("Atributo1:", linha.Atributo1)
		fmt.Println("Atributo2:", linha.Atributo2)
		fmt.Println("Atributo3:", linha.Atributo3)
		fmt.Println("Atributo4:", linha.Atributo4)
		fmt.Println()
	}

	deleteMultipleFiles(jsonFile, jsonFileEncoder, csvFile)
}

func createJsonWithMarshal() {
	file, err := os.Create(jsonFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	linhas := []Linha{
		{
			Atributo1: "Nome1",
			Atributo2: 10,
			Atributo3: 11.2277,
			Atributo4: true,
		},
		{
			Atributo1: "Nome2",
			Atributo2: 20,
			Atributo3: 22.4477,
			Atributo4: false,
		},
	}

	linhaJSON, err := json.MarshalIndent(linhas, "", "  ")
	if err != nil {
		panic(err)
	}

	_, err = file.Write(linhaJSON)
	if err != nil {
		panic(err)
	}
}

func readJsonWithUnmarshal(fileName string) []Linha {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	var linhas []Linha
	err = json.Unmarshal(data, &linhas)
	if err != nil {
		panic(err)
	}

	return linhas
}

func createJsonWithEncoder() {
	file, err := os.Create(jsonFileEncoder)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	linhas := []Linha{
		{
			Atributo1: "Nome1",
			Atributo2: 10,
			Atributo3: 11.2277,
			Atributo4: true,
		},
		{
			Atributo1: "Nome2",
			Atributo2: 20,
			Atributo3: 22.4477,
			Atributo4: false,
		},
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(linhas)
	if err != nil {
		panic(err)
	}

	// se for remover o /n na ultima linha, usar o seguinte
	// mas geralmente é melhor usar o json.MarshalIndent, pois salva em memoria
	// e depois escreve tudo de uma vez
	// var buf bytes.Buffer
	// enc := json.NewEncoder(&buf)
	// enc.SetIndent("", "  ")
	// err = enc.Encode(linhas)
	// if err != nil {
	// 	panic(err)
	// }
	// jsonSemQuebra := bytes.TrimRight(buf.Bytes(), "\n")
	// file.Write(jsonSemQuebra)
}

func readJsonWithDecoder(fileName string) []Linha {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	// Espera que o JSON seja um array
	_, err = decoder.Token() // Lê o '['
	if err != nil {
		panic(err)
	}

	linhas := make([]Linha, 0)
	for decoder.More() {
		var item Linha
		err := decoder.Decode(&item)
		if err != nil {
			panic(err)
		}
		linhas = append(linhas, item)
	}

	// Lê o ']' do array
	_, err = decoder.Token()
	if err != nil {
		panic(err)
	}

	return linhas
}

func createCsv() {
	file, err := os.Create(csvFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	linhas := []Linha{
		{
			Atributo1: "Nome1",
			Atributo2: 10,
			Atributo3: 11.2277,
			Atributo4: true,
		},
		{
			Atributo1: "Nome2",
			Atributo2: 20,
			Atributo3: 22.4477,
			Atributo4: false,
		},
	}

	writer := csv.NewWriter(file)
	defer writer.Flush() // força o buffer a ir pro disco

	// Cabeçalho
	err = writer.Write([]string{"Index", "Atributo1", "Atributo2", "Atributo3", "Atributo4"})
	if err != nil {
		panic(err)
	}

	for index, linha := range linhas {
		err = writer.Write([]string{
			strconv.Itoa(index),
			linha.Atributo1,
			strconv.Itoa(linha.Atributo2),
			strconv.FormatFloat(linha.Atributo3, 'f', 4, 64),
			strconv.FormatBool(linha.Atributo4),
		})
		if err != nil {
			panic(err)
		}
	}
}

func readCsv() []Linha {
	file, err := os.Open(csvFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Lê o cabeçalho e ignora
	if _, err := reader.Read(); err != nil {
		panic(err)
	}

	linhas := make([]Linha, 0)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		var item Linha
		item.Atributo1 = record[1]
		item.Atributo2, err = strconv.Atoi(record[2])
		if err != nil {
			panic(err)
		}
		item.Atributo3, err = strconv.ParseFloat(record[3], 64)
		if err != nil {
			panic(err)
		}
		item.Atributo4, err = strconv.ParseBool(record[4])
		if err != nil {
			panic(err)
		}
		linhas = append(linhas, item)
	}
	return linhas
}

func deleteMultipleFiles(names ...string) {
	for _, name := range names {
		err := os.Remove(name)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("Arquivo %s já não existe\n", name)
				continue
			}
			panic(fmt.Errorf("erro ao remover %s: %w", name, err))
		}
		fmt.Printf("Arquivo %s removido com sucesso\n", name)
	}
}
