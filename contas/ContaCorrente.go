package contas

import "fmt"

type ContaCorrente struct {
	       string
	numeroAgencia int
	numeroConta   int
	         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.
	if podeSacar {
		c. -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return " insuficiente"
	}
}