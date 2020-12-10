package builder

import (
	"testing"
)

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := Director{}

	autoBuilder := &AutoBuilder{}
	manufacturingComplex.SetBuilder(autoBuilder)
	manufacturingComplex.Construct()

	auto := autoBuilder.GetVehiculo()

	if auto.Ruedas != 4 {
		t.Errorf("Ruedas deben ser 4 pero son %d\n", auto.Ruedas)
	}

	if auto.Estructura != "Auto" {
		t.Errorf("Estructura debe ser  'auto' pero es %s\n", auto.Estructura)
	}

	if auto.Asientos != 5 {
		t.Errorf("Un auto debe tener  5 pero tiene %d\n", auto.Asientos)
	}

	motoBuilder := &MotoBuilder{}

	manufacturingComplex.SetBuilder(motoBuilder)
	manufacturingComplex.Construct()

	moto := motoBuilder.GetVehiculo()
	moto.Asientos = 1

	if moto.Ruedas != 2 {
		t.Errorf("Ruedas debe  tener 2 ruedas pero tiene %d\n", moto.Ruedas)
	}

	if moto.Estructura != "Motocicleta" {
		t.Errorf("Estructura de la moto  debe ser 'Motocicleta' pero es %s\n", moto.Estructura)
	}

	colectivoBuilder := &ColectivoBuilder{}

	manufacturingComplex.SetBuilder(colectivoBuilder)
	manufacturingComplex.Construct()

	colectivo := colectivoBuilder.GetVehiculo()

	if colectivo.Ruedas != 8 {
		t.Errorf("Un colectivo debe tener 8 ruedas ,pero tiene ...  %d\n", colectivo.Ruedas)
	}

	if colectivo.Estructura != "Colectivo" {
		t.Errorf("La estructura del colectivo debe ser  'colectivo' pero es... %s\n", colectivo.Estructura)
	}

	if colectivo.Asientos != 30 {
		t.Errorf("Un colectivo debe tener 30 asientos pero tiene %d\n", colectivo.Asientos)
	}
}
