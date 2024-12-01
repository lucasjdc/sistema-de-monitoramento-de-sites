package main

import (
	"fmt"
	"os"
	"net/http"
	"time"
	"bufio"
)

const monitoramentos = 3
const delay = 5

func main() {

    exibeIntroducao()
    for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
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
	fmt.Print("Opção: ")	
}

func leComando () int {
    var comandoLido int
    fmt.Scan(&comandoLido)   

    return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// sites := []string{"https://www.alura.com.br", "https://www.casadocodigo.com.br", "https://www.ead.senac.br/"}
	sites := leSitesDoArquivo()
		
	for i:= 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i+1, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
	}
	
	fmt.Println("")
	
}

func testaSite(site string) {
	resp, err := http.Get(site)
	
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")		
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}

    defer resp.Body.Close()
}

func leSitesDoArquivo() []string {
	
	var sites []string
	
	arquivo, err := os.Open("sites.txt")
	
	
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	
	leitor := bufio.NewReader(arquivo)
	linha, err := leitor.ReadString('\n')
	
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}	
	
	fmt.Println(linha)	
	
	return sites
}
