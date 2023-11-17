package main

import (
	"fmt"

	"github.com/Godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(1234)
	fmt.Println(estado)
	fmt.Println(texto)
}
