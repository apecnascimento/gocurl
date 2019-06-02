package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	cepInfo "github.com/apecnascimento/gocurl/cep"
)

// CepInfo reponse body from api.
// type CepInfo struct {
// 	Cep         string `json:"cep"`
// 	Logradouro  string `json:"logradouro"`
// 	Complemento string `json:"complemento"`
// 	Bairro      string `json:"complemnto"`
// 	Localidade  string `json:"localidade"`
// 	Uf          string `json:"uf"`
// 	Unidade     string `json:"unidade"`
// 	Ibge        string `json:"ibge"`
// 	Gia         string `json:"gia"`
// }

func main() {
	// APIURL := "https://viacep.com.br/ws/61620450/json/"
	var apiurl string
	flag.StringVar(&apiurl, "url", "", "Api url to consume")
	flag.Parse()
	response, err := http.Get(apiurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	var cepInfo cepInfo.CepInfo

	// body, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(body))
	if err := json.NewDecoder(response.Body).Decode(&cepInfo); err != nil {
		panic(err)
	}
	fmt.Println("CEP: ", cepInfo.Cep)
	fmt.Println("Rua: ", cepInfo.Localidade)

}
