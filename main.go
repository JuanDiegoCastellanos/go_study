package mains

import (
	"fmt"
)

func mains() {
	fmt.Println("■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■")
	fmt.Println("■■■■■ WELCOME TO WATER MANAGEMENT ■■■■■■■■■")
	fmt.Println("■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■")

	var x = int64(12)
	const PI = 3.1415
	fmt.Println("x = ", x)

	fmt.Println(suma(2, 3))

	fmt.Println("■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■")
	fmt.Println("■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■")

	for i := 0; i < 10; i += 1 {
		fmt.Println(i)
	}

	listaNumerosPares := []int{2, 4, 6, 8}

	for xz, par := range listaNumerosPares {
		if par%2 == 0 {
			fmt.Println("Es par: ", par, "-", xz)
		}
	}

	fmt.Println("■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■")

	nombresGatitos := []string{"Nube", "Darki", "Manchas"}

	for _, gato := range nombresGatitos {
		fmt.Println("Nombre de gato: ", gato)
	}

	modulo := 5 % 2

	switch modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

}

func suma(x, a int64) (int, string) {
	y := x + a
	return int(y), fmt.Sprintf("El valor es: %v", y)
}
