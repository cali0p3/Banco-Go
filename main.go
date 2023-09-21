package main

import (
	"fmt"
	"github.com/roberta.longoni\Documents\Dev\Go\Banco-Go/contas"
)
func main() {
	contaDaSilvia := contas.ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Println(contaDaSilvia.saldo)

	fmt.Println(contaDaSilvia.Sacar(400))
	fmt.Println(contaDaSilvia.saldo)

}
