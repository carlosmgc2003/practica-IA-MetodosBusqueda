package busqueda_sin_dominio

type Nodo struct {
	Dato  string  //El contenido del nodo, una letra
	peso  int     //Peso del nodo, no utilizado todavia
	Hijos []*Nodo //Vector de punteros a nodos hijos.
}
