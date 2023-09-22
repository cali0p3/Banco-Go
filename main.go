package main

import (
	"BANCO-GO/contas"
	"fmt"
)

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDaRoberta := contas.ContaPoupanca{}

	fmt.Println(contaDaRoberta)
}
