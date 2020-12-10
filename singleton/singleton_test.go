package singleton

import "testing"

func TestGetInstancia(t *testing.T) {
	contador1 := GetInstancia()
	if contador1 == nil {
		//Test 1 falla
		t.Error("El contador1 nunca puede ser igual a nulo")
	}
	contadorEsperado := contador1

	contadorActual := contador1.AgregarUno()
	if contadorActual != 1 {
		t.Errorf("Luego de sumarle 1 al contadorActual este deberia ser 1 pero es %d\n", contadorActual)
	}

	contador2 := GetInstancia()
	if contador2 != contadorEsperado {
		//Test 2 falla
		t.Error("Las intancias deben ser diferentes")
	}

	contadorActual = contador2.AgregarUno()
	if contadorActual != 2 {
		t.Errorf("Despues de llamar a la funcion AgregarUno usando el contador2,el contadorActual 2 deberia ser 2 pero es%d\n", contadorActual)
	}
}
