package singleton


type singleton struct {
	contador int
}
//definimos un puntero a una estructura de tipo Singleton como nil, 
//y lo asignamos a la variable  Instancia.
var Instancia *singleton
//Verificamos que si la instancia no se inicializa se crea una instancia en el espacio ya asignado
func GetInstancia() *singleton {
	if Instancia == nil {
		Instancia = new(singleton)
	}

	return Instancia
}
//Le agregamos 1 al valor del contador y nos retorna el valor del contador
func (s *singleton) AgregarUno() int {
	s.contador++
	return s.contador
}

func (s *singleton) GetContador() int {
	return s.contador
}