package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

// exemplo de chamada
// go run main.go 01001000
func main() {
	for _, cep := range os.Args[1:] {
		resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
		defer resp.Body.Close()

		res, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}

		var cep ViaCep
		err = json.Unmarshal(res, &cep)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}

		err = json.NewEncoder(os.Stdout).Encode(cep)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}

		file, err := os.Create("cidade.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", cep.Cep, cep.Localidade, cep.Uf))
	}
}
