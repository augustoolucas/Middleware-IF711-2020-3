package shared

import (
	"fmt"
)

//Request is
type Request struct {
	Op string
	P1 int
	P2 int
}

type Reply struct {
	Result []interface{}
}

//CALCULATOR_PORT define a porta da aplicação
const CALCULATOR_PORT = 3300

//SAMPLE_SIZE quantidade de amostras
const SAMPLE_SIZE = 10000

//CLIENTS define quantidade de clientes
const CLIENTS = 5

//ChecaErro usado para checar erros de conexão tcp
func ChecaErro(err error, mensagem string) {
	if err != nil {
		fmt.Println(mensagem, err)
		return
	}
	//...
}
