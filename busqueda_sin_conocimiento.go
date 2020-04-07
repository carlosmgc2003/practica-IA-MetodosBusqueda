package main

import "container/list"

func primero_amplitud(raiz *Nodo, buscado string) []string {
	solucion := make([]string, 0)
	cola := list.New()
	cola.PushBack(raiz)
	for cola.Len() > 0 {
		visitado := cola.Front()
		if visitado.Value.(*Nodo).dato != buscado {
			for _, hijo := range visitado.Value.(*Nodo).hijos {
				cola.PushBack(hijo)
			}
		} else {
			break
		}
		solucion = append(solucion, visitado.Value.(*Nodo).dato)
		cola.Remove(visitado)
	}
	return solucion
}
