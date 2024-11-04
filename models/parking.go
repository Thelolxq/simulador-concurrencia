package models

import (
	"sync"

	"fyne.io/fyne/v2/canvas"
)

type Espacio struct {
	ID      int
	X       float32
	Y       float32
	Ocupado bool
}

type Estacionamiento struct {
	EspaciosDisponibles chan bool
	Espacios            []Espacio
	AgregarVehiculo     chan *canvas.Image
	Mutex               sync.Mutex
}

func NuevoEstacionamiento(numeroEspacios int) *Estacionamiento {
	coordenadas := []struct {
		x float32
		y float32
	}{
		{200, 100}, {325, 100}, {425, 100}, {550, 100},
		{200, 400}, {325, 400}, {425, 400}, {550, 400},
	}

	espacios := make([]Espacio, len(coordenadas))
	for i, coord := range coordenadas {
		espacios[i] = Espacio{
			X:       coord.x,
			Y:       coord.y,
			Ocupado: false,
		}
	}

	return &Estacionamiento{
		EspaciosDisponibles: make(chan bool, numeroEspacios+1),
		AgregarVehiculo:     make(chan *canvas.Image, 100),
		Espacios:            espacios,
	}
}
