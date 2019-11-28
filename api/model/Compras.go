package model

import (
	"strings"
)

//Compra representa informações da compra
type Compra struct {
    Cpf_cnpj_comprador 		string `json:"cpf_cnpj_comprador"`
    Flg_private 		string `json:"flg_private"`
    Flg_incompleto 		string `json:"flg_incompleto"`
    Dt_ultima_compra 		string `json:"dt_ultima_compra"`
    Vl_ticket_medio 		string `json:"vl_ticket_medio"`
    Vl_ticket_ult_compra 	string `json:"vl_ticket_ult_compra"`
    Cnpj_loja_freq 		string `json:"cnpj_loja_freq"`
    Cnpj_loja_ultima 		string `json:"cnpj_loja_ultima"`
}

//GetFlgPrivate converte valor inteiro para boolean
func (i *Compra) GetFlgPrivate() bool {
    if i.Flg_private == "0" {
		return true
	}
	return false
}

//GetFlgIncompleto converte valor inteiro para boolean
func (i *Compra) GetFlgIncompleto() bool {
    if i.Flg_incompleto == "0" {
		return true
	}
	return false
}


//GetFlgIncompleto altera a virgula por ponto
func (i *Compra) GetVlTicketMedio() string  {
    return strings.Replace(i.Vl_ticket_medio, ",", ".", -1)
}

//GetVlTicketUltCompra altera a virgula por ponto
func (i *Compra) GetVlTicketUltCompra() string  {
    return strings.Replace(i.Vl_ticket_ult_compra, ",", ".", -1)
}
