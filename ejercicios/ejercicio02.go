package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

)



func IngresoNumeroParaMultiplicar() string {
	scanner := bufio.NewScanner(os.Stdin)

	var num1 int
	var err error
	var texto string
	
		for  {
			fmt.Println("Ingrese un número")
			if scanner.Scan() {
				num1, err = strconv.Atoi(scanner.Text())
				if err != nil{
					continue
				} else {
					break
				}
			}	
		}

		for i := 1; i <= 10; i++ {
			texto += fmt.Sprintf("%d x %d = %d \n", num1, i,i*num1)
		}
		fmt.Println(texto)
		return texto
}
	
