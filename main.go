package main

import (
	bsd "./busqueda_sin_dominio"
	"fmt"
)

func main() {

	//Inicializacion del arbol de acuerdo a lo visto en el apunte de clase.
	raiz := &bsd.Nodo{Dato: "A"}
	nodo1 := &bsd.Nodo{Dato: "D"}
	nodo2 := &bsd.Nodo{Dato: "F"}
	nodo3 := &bsd.Nodo{Dato: "G"}
	nodo4 := &bsd.Nodo{Dato: "H"}
	nodo5 := &bsd.Nodo{Dato: "J"}
	nodo6 := &bsd.Nodo{Dato: "C"}
	nodo7 := &bsd.Nodo{Dato: "E"}
	nodo8 := &bsd.Nodo{Dato: "P"}
	nodo9 := &bsd.Nodo{Dato: "Q"}
	nodo10 := &bsd.Nodo{Dato: "K"}
	nodo11 := &bsd.Nodo{Dato: "Z"}
	nodo12 := &bsd.Nodo{Dato: "W"}
	nodo13 := &bsd.Nodo{Dato: "B"}

	//Inicializacion de las relaciones
	raiz.Hijos = append(raiz.Hijos, nodo1)
	raiz.Hijos = append(raiz.Hijos, nodo2)
	raiz.Hijos = append(raiz.Hijos, nodo3)
	nodo1.Hijos = append(nodo1.Hijos, nodo4)
	nodo1.Hijos = append(nodo1.Hijos, nodo5)
	nodo2.Hijos = append(nodo2.Hijos, nodo6)
	nodo2.Hijos = append(nodo2.Hijos, nodo7)
	nodo4.Hijos = append(nodo4.Hijos, nodo8)
	nodo4.Hijos = append(nodo4.Hijos, nodo9)
	nodo5.Hijos = append(nodo5.Hijos, nodo10)
	nodo6.Hijos = append(nodo6.Hijos, nodo11)
	nodo6.Hijos = append(nodo6.Hijos, nodo12)
	nodo10.Hijos = append(nodo10.Hijos, nodo13)

	//Llamado a la solucion por primero amplitud, devuelve los nodos visitados.
	fmt.Println(bsd.Primero_amplitud(raiz, "B"))
	fmt.Println(bsd.Primero_profundidad(raiz, "B"))
}
