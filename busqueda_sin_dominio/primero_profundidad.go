package busqueda_sin_dominio

import "container/list"

func Primero_profundidad(raiz *Nodo, buscado string) []string {
	//Algoritmo que recorre el grafo, rama a rama, haciendo backtracking cuando no encuentra la solucion, el problema
	//por ahora es que destruye el grafo.
	solucion := make([]string, 0)
	pila := list.New()
	pila.PushFront(raiz)
	for pila.Len() > 0 {
		visitado := pila.Front()
		var hijo interface{} //La variable hijo tiene que ser de tipo interface porque viene "Encapsulada" del contenedor
		if len(visitado.Value.(*Nodo).Hijos) > 0 {
			hijo = visitado.Value.(*Nodo).Hijos[0]
		}
		if visitado.Value.(*Nodo).Dato == buscado {
			solucion = append(solucion, visitado.Value.(*Nodo).Dato)
			break
		} else {
			if hijo == nil {
				pila.Remove(visitado)
				solucion = solucion[:len(solucion)-1] //Slice para eliminar la ultima solucion que no funcionó (backtrack)
			} else {
				solucion = append(solucion, visitado.Value.(*Nodo).Dato)
				pila.PushFront(hijo)
				visitado.Value.(*Nodo).Hijos = visitado.Value.(*Nodo).Hijos[1:] //Elimino al primero puntero hijo que no
				//me sirvio... esto destruye el arbol por ahora es la unica manera que lo hago.
			}
		}
	}
	return solucion
}
