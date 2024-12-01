package main

import (
	"fmt"
	"os"
	"io"
	"net/http"
	"time"
	"bufio"
	"strings"
	"strconv"
	"io/ioutil"
)

const monitoramentos = 2
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
			imprimeLogs()
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
	
	verificaErro(err)
	
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}

    defer resp.Body.Close()
}

func leSitesDoArquivo() []string {
	
	var sites []string
	
	arquivo, err := os.Open("sites.txt")
	verificaErro(err)
	leitor := bufio.NewReader(arquivo)
	
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}	
	}
	arquivo.Close()
	
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	verificaErro(err)
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLogs() {
    arquivo, err := ioutil.ReadFile("log.txt")
    verificaErro(err)
    fmt.Println(string(arquivo))
}

func verificaErro(err error) {
    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
        return
    }
}
