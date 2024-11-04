package scenes

import (
	"simulador/models"
	veneno "simulador/poison"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
)

type Escenario struct {
	ventana   fyne.Window
	contenido *fyne.Container
}

func (e *Escenario) RenderizarVehiculo(parking *models.Estacionamiento) {
	for imagenVehiculo := range parking.AgregarVehiculo {
		e.contenido.Add(imagenVehiculo)
		e.ventana.Canvas().Refresh(e.contenido)
	}
}

func (e *Escenario) IniciarSimulacion() {
	parking := models.NuevoEstacionamiento(20)
	go veneno.GenerarVehiculos(100, parking)
	go e.RenderizarVehiculo(parking)
}

func (e *Escenario) inicializarContenido() {
	imagenFondo := canvas.NewImageFromURI(storage.NewFileURI("./assets/estacionamiento.jpg"))
	imagenFondo.Resize(fyne.NewSize(800, 600))
	imagenFondo.Move(fyne.NewPos(0, 0))

	e.contenido = container.NewWithoutLayout(imagenFondo)
	e.ventana.SetContent(e.contenido)
}

func NuevaEscena(ventana fyne.Window) *Escenario {
	escena := &Escenario{ventana: ventana}
	escena.Renderizar()
	return escena
}

func (e *Escenario) Renderizar() {
	e.inicializarContenido()
	e.IniciarSimulacion()
}
