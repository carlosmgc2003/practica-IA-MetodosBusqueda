package busqueda_sin_dominio

import "container/list"

func Primero_amplitud(raiz *Nodo) []string {
	//Recorre la totalidad de grafo, nivel a nivel.
	//Slice de strings
	solucion := make([]string, 0)
	//Lista doblemente enlazada de la libreria standart
	cola := list.New()
	cola.PushBack(raiz)
	for cola.Len() > 0 {
		visitado := cola.Front()
		//Hay que castear cola.Front() para poder utilizarlo por que es de tipo interface{}
		for _, hijo := range visitado.Value.(*Nodo).Hijos {
			cola.PushBack(hijo)
		}
		solucion = append(solucion, visitado.Value.(*Nodo).Dato)
		cola.Remove(visitado)
	}
	return solucion
}
