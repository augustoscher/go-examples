package main

import (
	"fmt"
	"time"

	"github.com/augustoscher/goarea/html"
)

func fastest(url1, url2, url3 string) string {
	c1 := html.Title(url1)
	c2 := html.Title(url2)
	c3 := html.Title(url3)

	// Estrutura de controle específica para concorrência.
	// Retorna o primeiro retorno dos canais considerados
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perdera!"
		// default:
		// 	return "Sem resposta!"
	}
}

func main() {
	winner := fastest("http://augustoscher.com", "https://www.google.com.br", "https://amazon.com")
	fmt.Println(winner)
}
