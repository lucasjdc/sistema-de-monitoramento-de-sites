package main

import "fmt"

func main() {
	fmt.Println("")
	fmt.Println("*** Monitoramento ***")
	fmt.Println("")

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
	var comando int
	fmt.Scan(&comando)

	if comando == 1 {
		fmt.Println("O comando escolhido foi", comando)
	} else if comando == 2 {
		fmt.Println("O comando escolhido foi", comando)
	} else if comando == 0 {
		fmt.Println("Encerrando o programa.")

	} else {
		fmt.Println("Comando desconhecido.")
	}
}
