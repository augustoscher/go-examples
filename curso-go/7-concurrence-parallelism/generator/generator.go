package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Google IO 2012 - Go Concurrency Patterns

// <-chan é um canal somente leitura

// Title recebe uma lista de URLs e retorna um canal somente leitura dos titulos das URLs
func title(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			result, _ := http.Get(url)
			html, _ := ioutil.ReadAll(result.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	// O retorno é executado antes das requests terminarem
	// Quem ler o canal irá receber os retornos a medida que as request terminarem
	return c
}

func main() {
	c1 := title("http://augustoscher.com", "https://www.google.com.br")
	c2 := title("https://amazon.com", "https://www.youtube.com")

	//Imprime o primeiro e segundo resultado de cada channel
	//Não necessáriamente serão retornados os titulos na ordem que chamamos, vai depender de quão rápido o site carrega.
	fmt.Println("First: ", <-c1, " | ", <-c2)
	fmt.Println("Second: ", <-c1, " | ", <-c2)
}
