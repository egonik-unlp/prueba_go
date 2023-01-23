package prueba_go

import (
	"errors"
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
