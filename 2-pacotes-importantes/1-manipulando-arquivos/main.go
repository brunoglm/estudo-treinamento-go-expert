package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	//criação arquivo
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	//escrita arquivo
	tamanho, err := f.Write([]byte("Hello, World bra bre bri bro bru!"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)
	f.Close()

	//leitura arquivo
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("arquivo readFile: ", string(arquivo))

	//leitura de pouco em pouco abrindo o arquivo
	file, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println("N: ", n, string(buffer[:n]))
	}
	file.Close()

	//excluindo arquivo
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("arquivo excluido")

	testWithJSON()
	testWithCSV()
}

type Linha struct {
	Atributo1 string
	Atributo2 int
	Atributo3 float64
	Atributo4 bool
}

func testWithJSON() {
	//Criando arquivo JSON
	file, err := os.Create("arquivo.json")
	if err != nil {
		panic(err)
	}

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

	//gravando no JSON
	linhaJSON, err := json.Marshal(linhas)
	if err != nil {
		panic(err)
	}
	file.Write(linhaJSON)
	file.Close()

	//lendo o arquivo JSON
	file, err = os.Open("arquivo.json")
	if err != nil {
		panic(err)
	}

	// Cria um decodificador JSON
	// Se o arquivo for pequeno, pode ser usado um simples json Unmarshal
	decoder := json.NewDecoder(file)

	// Abre o arquivo JSON de saída
	outFile, err := os.Create("saida.json")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	// Cria um codificador JSON
	encoder := json.NewEncoder(outFile)
	// Escreve o início do array no arquivo de saída
	outFile.Write([]byte("[\n"))

	// Lê o início do array
	_, err = decoder.Token()
	if err != nil {
		panic(err)
	}

	first := true
	// Itera sobre os itens no array
	for decoder.More() {
		var item Linha
		err := decoder.Decode(&item)
		if err != nil {
			panic(err)
		}

		// Exibe os dados na tela
		fmt.Printf("Atributo1: %s\n", item.Atributo1)
		fmt.Printf("Atributo2: %d\n", item.Atributo2)
		fmt.Printf("Atributo3: %.4f\n", item.Atributo3)
		fmt.Printf("Atributo4: %t\n", item.Atributo4)

		if !first {
			outFile.Write([]byte(","))
		}
		first = false

		// Escreve o item no arquivo de saída
		err = encoder.Encode(&item)
		if err != nil {
			panic(err)
		}
	}

	// Escreve o final do array no arquivo de saída
	outFile.Write([]byte("]\n"))

	// Lê o final do array
	_, err = decoder.Token()
	if err != nil {
		panic(err)
	}
	file.Close()
	outFile.Close()

	//deletando o JSON
	deleteMultipleFiles("arquivo.json", "saida.json")
}

func testWithCSV() {
	//criando um arquivo CSV
	file, err := os.Create("arquivo.csv")
	if err != nil {
		panic(err)
	}

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

	//gravando conteudo no CSV
	for index, linha := range linhas {
		linhaCSV := fmt.Sprintf("%d,%s,%d,%.4f,%t\n", index, linha.Atributo1, linha.Atributo2, linha.Atributo3, linha.Atributo4)
		file.WriteString(linhaCSV)
	}
	file.Close()

	//lendo um arquivo csv
	file, err = os.Open("arquivo.csv")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		a := strings.Split(string(line), ",")
		fmt.Println("Index linha", a[0])
		fmt.Println("posição 2", a[1])
		fmt.Println("posição 3", a[2])
		fmt.Println("posição 4", a[3])
		fmt.Println("posição 5", a[4])
	}
	file.Close()

	//deletando o arquivo csv
	err = os.Remove("arquivo.csv")
	if err != nil {
		panic(err)
	}
}

func deleteMultipleFiles(names ...string) {
	for _, name := range names {
		err := os.Remove(name)
		if err != nil {
			panic(err)
		}
	}
}
