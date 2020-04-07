package main

import (
	"fmt"
)

type Nodo struct {
	dato  string
	peso  int
	hijos []*Nodo
}

func main() {
	raiz := &Nodo{dato: "A"}
	nodo1 := &Nodo{dato: "D"}
	nodo2 := &Nodo{dato: "F"}
	nodo3 := &Nodo{dato: "G"}
	nodo4 := &Nodo{dato: "H"}
	nodo5 := &Nodo{dato: "J"}
	nodo6 := &Nodo{dato: "C"}
	nodo7 := &Nodo{dato: "E"}
	nodo8 := &Nodo{dato: "P"}
	nodo9 := &Nodo{dato: "Q"}
	nodo10 := &Nodo{dato: "K"}
	nodo11 := &Nodo{dato: "Z"}
	nodo12 := &Nodo{dato: "W"}
	nodo13 := &Nodo{dato: "B"}

	raiz.hijos = append(raiz.hijos, nodo1)
	raiz.hijos = append(raiz.hijos, nodo2)
	raiz.hijos = append(raiz.hijos, nodo3)
	nodo1.hijos = append(nodo1.hijos, nodo4)
	nodo1.hijos = append(nodo1.hijos, nodo5)
	nodo2.hijos = append(nodo2.hijos, nodo6)
	nodo2.hijos = append(nodo2.hijos, nodo7)
	nodo4.hijos = append(nodo4.hijos, nodo8)
	nodo4.hijos = append(nodo4.hijos, nodo9)
	nodo5.hijos = append(nodo5.hijos, nodo10)
	nodo6.hijos = append(nodo6.hijos, nodo11)
	nodo6.hijos = append(nodo6.hijos, nodo12)
	nodo10.hijos = append(nodo10.hijos, nodo13)
	fmt.Println(primero_amplitud(raiz, "B"))

}
