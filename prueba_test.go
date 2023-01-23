package prueba_go

import "testing"

func testAgearAges(t *testing.T) {
	Nombre := "Yaul"
	Edad := 28
	NAge := 10 // Podria ser un rn mas adelante
	p := Persona{Nombre: Nombre, Edad: Edad}
	i := 0
	for i < NAge {
		p.Agear()
		i++
	}
	if p.Edad != Edad+i {
		t.Fatalf("No funciona el aging")
	}
}
