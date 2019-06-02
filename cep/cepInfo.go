package cep

type CepInfo struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"complemnto"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Unidade     string `json:"unidade"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
}
