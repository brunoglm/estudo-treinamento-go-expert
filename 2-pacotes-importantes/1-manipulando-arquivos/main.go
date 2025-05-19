package main

import (
	"encoding/json"
	"fmt"
	"os"
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
	fmt.Println("Lendo arquivo com a função readJsonWithUnmarshal")

	deleteMultipleFiles(jsonFile, jsonFileEncoder)
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

// func testWithCSV() {
// 	//criando um arquivo CSV
// 	file, err := os.Create("arquivo.csv")
// 	if err != nil {
// 		panic(err)
// 	}

// 	linhas := []Linha{
// 		{
// 			Atributo1: "Nome1",
// 			Atributo2: 10,
// 			Atributo3: 11.2277,
// 			Atributo4: true,
// 		},
// 		{
// 			Atributo1: "Nome2",
// 			Atributo2: 20,
// 			Atributo3: 22.4477,
// 			Atributo4: false,
// 		},
// 	}

// 	//gravando conteudo no CSV
// 	for index, linha := range linhas {
// 		linhaCSV := fmt.Sprintf("%d,%s,%d,%.4f,%t\n", index, linha.Atributo1, linha.Atributo2, linha.Atributo3, linha.Atributo4)
// 		file.WriteString(linhaCSV)
// 	}
// 	file.Close()

// 	//lendo um arquivo csv
// 	file, err = os.Open("arquivo.csv")
// 	if err != nil {
// 		panic(err)
// 	}
// 	reader := bufio.NewReader(file)
// 	for {
// 		line, _, err := reader.ReadLine()
// 		if err != nil {
// 			break
// 		}
// 		a := strings.Split(string(line), ",")
// 		fmt.Println("Index linha", a[0])
// 		fmt.Println("posição 2", a[1])
// 		fmt.Println("posição 3", a[2])
// 		fmt.Println("posição 4", a[3])
// 		fmt.Println("posição 5", a[4])
// 	}
// 	file.Close()

// 	//deletando o arquivo csv
// 	err = os.Remove("arquivo.csv")
// 	if err != nil {
// 		panic(err)
// 	}
// }

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
