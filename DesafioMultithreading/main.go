package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

/*type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}*/

/*type BrasilAPI struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}*/

func main() {
	c1 := make(chan *http.Response)
	c2 := make(chan *http.Response)
	cep := "01153000"
	client := http.Client{}

	go func() {
		first, err := client.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			log.Print("Erro na requisição")
		} else {
			c2 <- first
		}
	}()

	go func() {
		first, err := client.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
		if err != nil {
			log.Print("Erro na requisição")
		} else {
			c1 <- first
		}
	}()

	select {
	case first := <-c1:
		defer first.Body.Close()
		body, err := io.ReadAll(first.Body)
		if err != nil {
			log.Print("Erro ao ler resposta")
		} else {
			fmt.Print("\n\nResposta da BrasilAPI:\n\n")
			fmt.Println(string(body))
		}
	case first := <-c2:
		defer first.Body.Close()
		body, err := io.ReadAll(first.Body)
		if err != nil {
			log.Print("Erro ao ler resposta")
		} else {
			fmt.Print("\n\nResposta da ViaCEP:\n\n")
			fmt.Println(string(body))
		}

	case <-time.After(time.Second):
		println("timeout")
	}
}
