package repo

import (
	"github.com/user/GoNeo/api/tratamento"
	"github.com/user/GoNeo/api/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

var dbConn *gorm.DB
var err error

type Compra struct {
    gorm.Model
    Cpf_cnpj_comprador      string
    Flg_private             string
    Flg_incompleto          string
    Dt_ultima_compra        string
    Vl_ticket_medio         string
    Vl_ticket_ult_compra    string
    Cnpj_loja_freq          string
    Cnpj_loja_ultima        string
}

//InsereDados metodo que insere os dados dos arquivo na base de dados
func InsereDados(mov []model.Compra) {

    connectionParams := "user=docker password=docker sslmode=disable host=db"
    dbConn, err = gorm.Open("postgres", connectionParams)

    if dbConn != nil {
        fmt.Println("Conectado!");
    } else {
        fmt.Println("Não Conectado!");
    }

    // Se tabela não existe então cria
    if !dbConn.HasTable(&Compra{}) {
    	fmt.Println("Criando Tabela!");
        dbConn.CreateTable(&Compra{})
    }

    for i := 0; i < len(mov); i++ {

    	if (i % 1000 == 0){
    		// var compra Compra
      		fmt.Println("Quantidade de Registros: ", i);
      		// fmt.Println(dbConn.First(&compra, i))
    	}

        ticket := Compra{
                Cpf_cnpj_comprador: tratamento.RemoveCaracteres(mov[i].Cpf_cnpj_comprador),Flg_private: mov[i].Flg_private, Flg_incompleto: mov[i].Flg_incompleto, Dt_ultima_compra: mov[i].Dt_ultima_compra, Vl_ticket_ult_compra: mov[i].Vl_ticket_ult_compra, Cnpj_loja_freq: tratamento.RemoveCaracteres(mov[i].Cnpj_loja_freq), Cnpj_loja_ultima: tratamento.RemoveCaracteres(mov[i].Cnpj_loja_ultima)}

        dbConn.Create(&ticket)
    }

}
