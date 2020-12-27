package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	analizarListaLineal(x)

	a := []int{7, 8, 9, 10, 11}
	analizarListaLineal(a)

	arregloNoLi := []int{1, 2, 4, 8, 16, 32}
	analizarListaLineal(arregloNoLi)
}

func analizarListaLineal(a []int) bool {
	y := restarNumeros(a)
	// fmt.Println(y)
	z := comprobarIgualdad(y)
	// fmt.Println(z)
	lineal := comprobarVerdad(z)
	r := comprobarLineal(lineal, a)
	return r
}

// obtiene la diferencia de un numero - el anterior
// y lo agrega a una lista
func restarNumeros(sucesion []int) []int {
	x := []int{}
	for i := 1; i < len(sucesion); i++ {
		// fmt.Println(sucesion[i])
		resto := sucesion[i] - sucesion[i-1]
		x = append(x, resto)
	}
	// fmt.Println(x)
	return x
}

// si un numero es igual al anterior agregara un true a la lista
// en caso contrario agregara un false y terminara la ejecucion
func comprobarIgualdad(sucesion []int) []bool {

	x := []bool{}
	if len(sucesion) > 1 {
		for i := 1; i < len(sucesion); i++ {
			if sucesion[i] == sucesion[i-1] {
				x = append(x, true)
				// fmt.Println(true)
			} else {
				x = append(x, false)
				// fmt.Println("falso")
				return x
			}
		}
		// fmt.Println(x)
	} else {
		if sucesion[0] == 1 {
			x = append(x, true)
		} else {
			// fmt.Println("No es igual")
			x = append(x, false)
		}
	}

	return x
}

// Si toda la lista tiene true regresara un trua
// si el ultimo termino es false regresara false
func comprobarVerdad(a []bool) bool {
	s := a[len(a)-1]

	return s
}

// Divide un numero con el numero anterior
// y regresa una lista con los cocientes
func obtenerDivision(arreglo []int) []int {
	divisiones := []int{}

	for i := 1; i < len(arreglo); i++ {
		division := arreglo[i] / arreglo[i-1]
		divisiones = append(divisiones, division)
	}
	// fmt.Println(divisiones)
	return divisiones
}

// regresa un true si el arreglo es lineal y lo imprime
func comprobarLineal(a bool, b []int) bool {
	var r bool
	if a {
		fmt.Printf("el arreglo %v es lineal\n", b)
		f := obtenerFuncionLineal(b)
		fmt.Println(f)
		r = true
	} else {
		fmt.Printf("el arreglo %v no es lineal\n", b)
		r = false
	}
	return r
}

func obtenerFuncionLineal(a []int) string {
	var r int
	r = a[0] - 1
	// fmt.Println(r)
	f := strconv.Itoa(r)
	// fmt.Println(f)
	funcion := "la funcion de esta lista lineal es an = n + " + f
	// fmt.Println(funcion)
	return funcion
}
