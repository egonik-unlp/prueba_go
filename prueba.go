package prueba_go

import (
	"errors"
	"testing"
)

type Persona struct {
	Nombre string
	Edad   int
}

func (p Persona) CalcularAñoNacimiento() (int, error) {
	año := 2023 - p.Edad
	if año < 1 {
		return 0, errors.New("No naciste en el futuro")

	} else {
		return año, nil
	}

}

func (p *Persona) Agear() {
	p.Edad++
}

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
