package models

import (
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Vehiculo struct {
	Estacionamiento  *Estacionamiento
	ID               int
	posicionEstacion int
	imagen           *canvas.Image
}

func NuevoVehiculo(est *Estacionamiento, img *canvas.Image) *Vehiculo {
	return &Vehiculo{
		Estacionamiento: est,
		imagen:          img,
	}
}

func (v *Vehiculo) MoverVehiculo() {
	v.Estacionamiento.EspaciosDisponibles <- true
	v.Estacionamiento.Mutex.Lock()

	// Buscar un espacio disponible
	var espacio *Espacio
	for i := 0; i < len(v.Estacionamiento.Espacios); i++ {
		if !v.Estacionamiento.Espacios[i].Ocupado {
			espacio = &v.Estacionamiento.Espacios[i]
			break
		}
	}

	if espacio == nil {
		// No hay espacios disponibles
		v.Estacionamiento.Mutex.Unlock()
		return
	}

	// Mover el vehículo a la derecha, fuera de la pantalla
	v.imagen.Move(fyne.NewPos(800, espacio.Y)) // Suponiendo que la ventana tiene 800px de ancho
	v.imagen.Refresh()

	// Animación de entrada
	for posX := 800; float32(posX) > espacio.X; posX -= 5 { // Convertir posX a float32
		v.imagen.Move(fyne.NewPos(float32(posX), espacio.Y))
		v.imagen.Refresh()
		time.Sleep(10 * time.Millisecond) // Controla la velocidad de la animación
	}
	// Finalmente, coloca el vehículo en el espacio
	v.imagen.Move(fyne.NewPos(espacio.X, espacio.Y))
	v.imagen.Refresh()

	// Ocupa el espacio
	v.posicionEstacion = espacio.ID // Asegúrate de tener un campo ID en Espacio
	espacio.Ocupado = true

	fmt.Println("El vehículo", v.ID, "entró")
	time.Sleep(300 * time.Millisecond)

	// Desbloquear Mutex
	v.Estacionamiento.Mutex.Unlock()

	tiempoEspera := rand.Intn(5-1+1) + 1
	time.Sleep(time.Duration(tiempoEspera) * time.Second)

	// Volver a bloquear el Mutex antes de salir
	v.Estacionamiento.Mutex.Lock()

	// Liberar el espacio de vehículos
	<-v.Estacionamiento.EspaciosDisponibles
	espacio.Ocupado = false
	
	// Mover el vehículo fuera de la pantalla
	for posX := espacio.X; posX < 800; posX += 5 {
		v.imagen.Move(fyne.NewPos(float32(posX), espacio.Y))
		v.imagen.Refresh()
		time.Sleep(10 * time.Millisecond) // Controla la velocidad de la animación
	}

	fmt.Println("El vehículo", v.ID, "salió ")
	time.Sleep(300 * time.Millisecond)

	// Desbloquear el Mutex
	v.Estacionamiento.Mutex.Unlock()
}