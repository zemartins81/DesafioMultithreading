package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetCepBrasilApi(cep string) (string, error) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	cl := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	resp, err := cl.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func getCepViaCep(cep string) (string, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	cl := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	resp, err := cl.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	var cep string
	fmt.Println("Digite o CEP: ")
	fmt.Scanln(&cep)

	go func() {
		resp, err := GetCepBrasilApi(cep)
		if err != nil {
			close(c1)
		}
		c1 <- resp
	}()

	go func() {
		resp, err := getCepViaCep(cep)
		if err != nil {
			close(c2)
		}
		c2 <- resp
	}()

	select {
	case res1 := <-c1:
		fmt.Println(res1)
		fmt.Println("https://brasilapi.com.br")
	case res2 := <-c2:
		fmt.Println(res2)
		fmt.Println("https://viacep.com.br")
	case <-time.After(time.Second):
		fmt.Println("Request Timeout")
	}
}
