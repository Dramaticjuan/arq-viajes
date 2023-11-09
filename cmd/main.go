package main

import (
	"fmt"

	"github.com/Dramaticjuan/arq3-viajes/internal/adapters/handlers"
	"github.com/Dramaticjuan/arq3-viajes/internal/adapters/services"
	"github.com/go-chi/chi/v5"
)

func main() {
	msi := services.NewMonopatineServiceImpl("http://localhost:3000/monopatin/")
	d, _ := msi.GetMonopatin(1)
	if d.Parada == nil {
		println(d.Kilometros)
	}
	er := msi.UpdateParadaMonopatin(2, 1)
	if er != nil {
		fmt.Println(er.Error())
	}

}

func Routes(h *handlers.ViajeHandler) *chi.Mux {
	r := chi.NewMux()
	r.Route("/viaje", func(r chi.Router) {
		r.Post("/", h.EmpezarViaje)
		r.Put("/{id}", h.PausarViaje)
		r.Put("/{id}", h.ReanudarViaje)
		r.Put("/{id}", h.TerminarViaje)

		r.Route("/report", func(r chi.Router) {
			r.Get("/{id}", h.ReportConPausas)
			r.Get("/{id}", h.ReportSinPausas)
		})
	})
	return r
}
