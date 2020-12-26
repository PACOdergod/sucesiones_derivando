package main

import (
	"fmt"
)

func main() {
	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// fmt.Println(x)

	y := _restarNumeros(x)
	fmt.Println(y)
	z := comprobarIgualdad(y)
	fmt.Println(z)
	lineal := comprobarLineal(z)
	fmt.Println(lineal)
	if lineal {
		fmt.Println("es lineal")
	} else {
		fmt.Println("no es lineal")
	}

}

func _restarNumeros(sucesion []int) []int {
	x := []int{}
	for i := 1; i < len(sucesion); i++ {
		// fmt.Println(sucesion[i])
		resto := sucesion[i] - sucesion[i-1]
		x = append(x, resto)
	}
	// fmt.Println(x)
	return x
}

func comprobarIgualdad(sucesion []int) []bool {
	/// si la sucesion en n es igual a n-1 regresara true
	/// hata terminar o encontrar unos numeros que no sean iguales
	/// si 2 numeros no son iguales agreara un false y regresara la lista
	/// terminando el ciclo, asi no tiene que recorrer todo el arreglo

	x := []bool{}
	if len(sucesion) > 1 {
		for i := 1; i < len(sucesion); i++ {
			if sucesion[i] == sucesion[i-1] {
				x = append(x, true)
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

func comprobarLineal(a []bool) bool {
	s := a[len(a)-1]
	return s
}
