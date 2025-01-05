package main

import "fmt"
import "os"
import "net/http"
//import "reflect"

func main() {
    /*
	fmt.Println("Olá Mundo")
	var nome string = "Douglas"
    var idade int = 24
    var versao float32 = 1.1
    fmt.Println("Olá sr.", nome, "sua idade é", idade)
    fmt.Println("Este programa está na versão", versao)
	
	var nome = "Douglas"
    var idade = 24
    var versao = 1.1
    fmt.Println("Olá sr.", nome, "sua idade é", idade)
    fmt.Println("Este programa está na versão", versao)
    fmt.Println("O tipo da variável idade é", reflect.TypeOf(versao))
	
	
	nome := "Douglas"
    idade := 24
    versao := 1.1
    fmt.Println("Olá sr.", nome, "sua idade é", idade)
    fmt.Println("Este programa está na versão", versao)
	

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")
    var comando int
	fmt.Printf("Digite:")
    fmt.Scan(&comando)
    fmt.Println("O comando escolhido foi", comando)  

	if comando == 1 {
        fmt.Println("Monitorando...")
    } else if comando == 2 {
        fmt.Println("Exibindo Logs...")
    } else if comando == 0 {
        fmt.Println("Saindo do programa...")
    } else {
        fmt.Println("Não conheço este comando")
    }

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
    default:
		fmt.Println("Não conheço este comando")
	}
	*/

	// TESTE DE RETORNO MULTIPLO
	nome, idade, _ := retornoMultiplo() // O _ indica que o terceiro retorno é inutil para mim.
	fmt.Println(nome, "tem", idade, "anos")  // Nâo precisa considerar espaço entre o "tem" e "idade". O Println já adiciona automaticamente

	exibeIntroducao()

	for {
		exibeMenu()
		comando := leEntrada()
		
		switch comando {
    	case 1:
    	    fmt.Println("Monitorando com array...")
			iniciarMonitoramentoComArray()
    	case 2:
    	    fmt.Println("Monitorando com slice...")
			iniciarMonitoramentoComSlice()
    	case 0:
    	    fmt.Println("Saindo do programa...")
    	    os.Exit(0)
    	default:
    	    fmt.Println("Não conheço este comando")
    	    os.Exit(-1)
    	}
	}
}

func leEntrada() int {
	var comandoLido int
	fmt.Printf("Digite aqui a opção de interesse:")
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi ", comandoLido)  
	return comandoLido
}

func exibeMenu() {
    fmt.Println("1- Iniciar Monitoramento com array")
    fmt.Println("2- Iniciar Monitoramento com slice")
    fmt.Println("0- Sair do Programa")
}

func exibeIntroducao() {
	mensagem := "Boas vindas ao treinamento GO"
	versao := 1.0
	fmt.Println(mensagem ," versão: ", versao)
}

func iniciarMonitoramentoComArray() {
	var sites [4]string
    sites[0] = "https://random-status-code.herokuapp.com/"
    sites[1] = "https://www.alura.com.br"
    sites[2] = "https://www.caelum.com.br"

	for i := 0; len(sites) > i && sites[i] != "" ;i++ {
	resp, _ := http.Get(sites[i])
		if resp.StatusCode == 200 {
    	    fmt.Println("Site:", sites[i], "foi carregado com sucesso!")
    	} else {
    	    fmt.Println("Site:", sites[i], "está com problemas. Status Code:", resp.StatusCode)
    	}
	}

}

func iniciarMonitoramentoComSlice() {
	sites := []string
	append(sites, "https://random-status-code.herokuapp.com/")
	append(sites, "https://www.alura.com.br")
	append(sites, "https://www.caelum.com.br")

	for i := 0; len(sites) > i ;i++ {
	resp, _ := http.Get(sites[i])
		if resp.StatusCode == 200 {
    	    fmt.Println("Site:", sites[i], "foi carregado com sucesso!")
    	} else {
    	    fmt.Println("Site:", sites[i], "está com problemas. Status Code:", resp.StatusCode)
    	}
	}

}

func retornoMultiplo() (string, int, string) {
    nome := "Douglas"
    idade := 24
	inutil := "Não tenho interesse em usar essa informação, mas tenho que mantê-la no código"
    return nome, idade, inutil
}