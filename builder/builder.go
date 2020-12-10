package builder

type BuildProcess interface {
	SetRuedas() BuildProcess
	SetAsientos() BuildProcess
	SetEstructura() BuildProcess
	GetVehiculo() VehiculoProduto
}

//Director
type Director struct {
	builder BuildProcess
}

func (f *Director) Construct() {
	f.builder.SetAsientos().SetEstructura().SetRuedas()
}

func (f *Director) SetBuilder(b BuildProcess) {
	f.builder = b
}

// Estructura del VehiculoProducto
type VehiculoProduto struct {
	Ruedas     int
	Asientos   int
	Estructura string
}

//Constructor de tipo Auto
type AutoBuilder struct {
	v VehiculoProduto
}

func (c *AutoBuilder) SetRuedas() BuildProcess {
	c.v.Ruedas = 4
	return c
}

func (c *AutoBuilder) SetAsientos() BuildProcess {
	c.v.Asientos = 5
	return c
}

func (c *AutoBuilder) SetEstructura() BuildProcess {
	c.v.Estructura = "Auto"
	return c
}

func (c *AutoBuilder) GetVehiculo() VehiculoProduto {
	return c.v
}

//Constructor de tipo Moto
type MotoBuilder struct {
	v VehiculoProduto
}

func (b *MotoBuilder) SetRuedas() BuildProcess {
	b.v.Ruedas = 2
	return b
}

func (b *MotoBuilder) SetAsientos() BuildProcess {
	b.v.Asientos = 2
	return b
}

func (b *MotoBuilder) SetEstructura() BuildProcess {
	b.v.Estructura = "Motocicleta"
	return b
}

func (b *MotoBuilder) GetVehiculo() VehiculoProduto {
	return b.v
}

//Constructor de tipo Colectivo
type ColectivoBuilder struct {
	v VehiculoProduto
}

func (b *ColectivoBuilder) SetRuedas() BuildProcess {
	b.v.Ruedas = 8
	return b
}

func (b *ColectivoBuilder) SetAsientos() BuildProcess {
	b.v.Asientos = 30
	return b
}

func (b *ColectivoBuilder) SetEstructura() BuildProcess {
	b.v.Estructura = "Colectivo"
	return b
}

func (b *ColectivoBuilder) GetVehiculo() VehiculoProduto {
	return b.v
}
