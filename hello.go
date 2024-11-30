package main

import (
	"fmt"
	"os"
)

func main() {

    exibeIntroducao()
    exibeMenu()
	comando := leComando()

	switch comando {
	case 1:
        fmt.Println("Monitorando...")
    case 2:
        fmt.Println("Exibindo logs...")
    case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
    default:
        fmt.Println("Não conheço este comando")
        os.Exit(-1)
	}
}

func exibeIntroducao() {
    nome := "Lucas"
    var versao float32 = 1.1
    fmt.Println("Olá, sr.", nome)
    fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")	
}

func leComando () int {
    var comandoLido int
    fmt.Scan(&comandoLido)
    fmt.Println("O comando esolhido foi: ", comandoLido)

    return comandoLido
}
