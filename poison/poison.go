package veneno

import (
	"math/rand"
	"simulador/models"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

func ObtenerImagenVehiculo() *canvas.Image {
	imagenVehiculo := canvas.NewImageFromURI(storage.NewFileURI("./assets/coche.png"))
	imagenVehiculo.Resize(fyne.NewSize(60, 70))
	imagenVehiculo.Move(fyne.NewPos(410, 0))
	return imagenVehiculo
}

func GenerarVehiculos(cantidad int, estacionamiento *models.Estacionamiento) {
	estacionamiento.EspaciosDisponibles <- true

	for i := 0; i < cantidad; i++ {
		imagenVehiculo := ObtenerImagenVehiculo()
		vehiculoNuevo := models.NuevoVehiculo(estacionamiento, imagenVehiculo)
		vehiculoNuevo.ID = i + 1
		estacionamiento.AgregarVehiculo <- imagenVehiculo
		time.Sleep(300 * time.Millisecond)

		go vehiculoNuevo.MoverVehiculo()
		time.Sleep(time.Duration(rand.ExpFloat64() * float64(time.Second)))
	}
}
