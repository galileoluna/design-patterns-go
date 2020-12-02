package singleton


type singleton struct {
	contador int
}

var Instancia *singleton

func GetInstancia() *singleton {
	if Instancia == nil {
		Instancia = new(singleton)
	}

	return Instancia
}

func (s *singleton) AgregarUno() int {
	s.contador++
	return s.contador
}

func (s *singleton) GetContador() int {
	return s.contador
}