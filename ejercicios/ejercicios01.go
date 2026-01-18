package ejercicios

import (
	"strconv"
)

func ConvertirStringaNumero(texto string) (int, string) {
	numero, err := strconv.Atoi(texto)
	if err != nil{
		return 0, "Hubo un error en el mensaje" + err.Error()
	}
	if numero > 100 {
		return numero, "Es mayor a 100"
	}else {
		return numero, "No es mayor a 100"
	}

	
}