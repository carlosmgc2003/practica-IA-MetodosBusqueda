package busqueda_sin_dominio

import "container/list"

func Primero_profundidad(raiz *Nodo, buscado string) []string {
	solucion := make([]string, 0)
	pila := list.New()
	pila.PushFront(raiz)
	for pila.Len() > 0 {
		visitado := pila.Front()
		var hijo interface{}
		if len(visitado.Value.(*Nodo).Hijos) > 0 {
			hijo = visitado.Value.(*Nodo).Hijos[0]
		}
		if visitado.Value.(*Nodo).Dato == buscado {
			solucion = append(solucion, visitado.Value.(*Nodo).Dato)
			break
		} else {
			if hijo == nil {
				pila.Remove(visitado)
				solucion = solucion[:len(solucion)-1]
			} else {
				solucion = append(solucion, visitado.Value.(*Nodo).Dato)
				pila.PushFront(hijo)
				visitado.Value.(*Nodo).Hijos = visitado.Value.(*Nodo).Hijos[1:]
			}
		}
	}
	return solucion
}
