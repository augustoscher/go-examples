package main

import (
	"fmt"

	"github.com/augustoscher/goarea/html"
)

// precisa atualizar o pachage:
// go get -u github.com/augustoscher/goarea

// Multiplexing é uma tecnica onde pegamos dados de duas ou mais fontes
// e escrevemos no mesmo canal.
// Ou seja, temos dois channels gerando dados e juntamos em um único.

//Canal origin é somente leitura
func send(origin <-chan string, destination chan string) {
	for {
		destination <- <-origin //escreve dados do canal de origem no canal de destino
	}
}

// Responsável por misturar mensagens em um canal.
// Recebe dois canais somente leitura
func multiplex(entry1, entry2 <-chan string) <-chan string {
	c := make(chan string)
	go send(entry1, c)
	go send(entry2, c)
	return c
}

func main() {
	c := multiplex(
		//cada chamada das funcs abaixo retornam um canal
		html.Title("http://augustoscher.com", "https://www.google.com.br"),
		html.Title("https://amazon.com", "https://www.youtube.com"),
	)

	fmt.Println("First: ", <-c, " | Second: ", <-c)
	fmt.Println("Third: ", <-c, " | Fourth: ", <-c)
}
