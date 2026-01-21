package ejer_interfaces

import (
	"fmt"

	"github.com/puricalvo/godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {

	hu.Respirar()
	fmt.Printf("Soy un/a %s,  estoy respirando, y estoy vivo ? %t \n", hu.Sexo(), hu.EstaVivo())
}



