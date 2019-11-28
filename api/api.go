package main

import (
	"database"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"

	"io/ioutil"
	"bufio"
	"regexp"

	"github.com/user/GoDoRP/api/model"
	"github.com/user/GoDoRP/api/repo"
	"github.com/user/GoDoRP/api/tratamento"

)

func processaArquivo(caminhaArquivo string ) {

	file, err := os.Open("./Dados/" + caminhaArquivo)
    if err != nil {
		log.Fatal("[api] Ocorreu erro ao abrir o arquivo. Erro: ",err.Error())
		return
	}

	fmt.Println(file.Name())

	defer file.Close()
	
    scanner := bufio.NewScanner(file)
    r := regexp.MustCompile("[^\\s]+")

    list_movimento:=[]model.Compra{}

    for scanner.Scan() {
		results := r.FindAllString(scanner.Text(), -1)
		
		mov := model.Compra{}
		mov.Cpf_cnpj_comprador = results[0]
		mov.Flg_private = results[1]
		mov.Flg_incompleto = results[2]
		mov.Dt_ultima_compra = results[3]
		mov.Vl_ticket_medio = results[4]
		mov.Vl_ticket_ult_compra = results[5]
		mov.Cnpj_loja_freq = results[6]
		mov.Cnpj_loja_ultima = results[7]

		if tratamento.ValidateCpfCnpj(results[0]) && tratamento.ValidateCpfCnpj(results[7]) && tratamento.ValidateCpfCnpj(results[7]) {
				list_movimento=append(list_movimento,mov)    
		}
    }

    fmt.Println("CPF/CNPJ que estão válidos para inserir na Base de Dados: ", len(list_movimento))

	fmt.Println("Inserir dados no Banco de Dados!")
	repo.InsereDados(list_movimento)
	fmt.Println("Fim do processamento de Dados!")
}

func main() {
	defer database.DB.Close()

	// add router and routes
	router := httprouter.New()

	// add database
	_, err := database.Init()
	if err != nil {
		log.Println("Conexão com a base de dados falhou, abortando...")
		log.Fatal(err)
	}

	log.Println("Conexão com a Base de Dados realizada!")

	// print env
	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running api server in production mode")
	} else {
		log.Println("Running api server in dev mode")
	}

	dir := "./Dados/"

    files, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatal("[main] Ocorreu na pasta de diretório dos arquivos. Erro: ",err.Error())
    }

    for _, f := range files {
		fmt.Println(f.Name())
		processaArquivo(f.Name())
	}

	http.ListenAndServe(":8080", router)
}
