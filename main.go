package main

import (
	"fmt"

	"github.com/puricalvo/godesde0/goroutines"
)

func main() {
	/* estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto) */
	/* if os := runtime.GOOS; os=="linux" || os=="OS X."{
		fmt.Println("Esto no es Windows, es ", os)
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os:= runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)

	}

	numero, texto := ejercicios.ConvertirStringaNumero("500")
	fmt.Println(numero)
	fmt.Println(texto)


	teclado.IngresoNumeros()
	*/

	//iteraciones.Iterar()
	
	//fmt.Println(ejercicios.IngresoNumeroParaMultiplicar()) 
	
	//files.GrabaTabla()

	//files.SumaTabla()

	//files.LeoArchivo()

	//funciones.Calculos()

	//funciones.LlamarClosure()

	//funciones.Exponencia(2)

	//arreglos_slices.MuestroArreglos()

	//arreglos_slices.MuestroSlice()

	// arreglos_slices.Capacidad()

	// mapas.Mostrarmapas()

	// users.AltaUsuario()

	/* Pedro := new(modelos.Hombre)
	e.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria) */

	// d.VemosDefer()

	// d.EjemploPanic()

	canal1 := make(chan bool)
	go goroutines.MiNombrelento("Puri Calvo", canal1)
	defer func(){ 
		<-canal1 
	}()
	fmt.Println("Estoy aquí")

	
	

}