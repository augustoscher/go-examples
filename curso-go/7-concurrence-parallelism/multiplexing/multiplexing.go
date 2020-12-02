package main

// Multiplexing é uma tecnica onde pegamos dados de duas ou mais fontes
// e escrevemos no mesmo canal.
// Ou seja, temos dois channels gerando dados e juntamos em um único.

//Canal origin é somente leitura
func send(origin <-chan string, destination chan string) {
	for {
		destination <- <-origin //escreve dados do canal de origem no canal de destino
	}
}

func main() {

}
