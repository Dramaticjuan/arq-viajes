package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Dramaticjuan/arq3-viajes/internal/core/domain"
	"github.com/Dramaticjuan/arq3-viajes/internal/ports/in"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type ViajeHandler struct {
	svc in.ViajeService
}

func NewViajeHandler(svc in.ViajeService) *ViajeHandler {
	return &ViajeHandler{
		svc: svc,
	}
}

func (h *ViajeHandler) EmpezarViaje(w http.ResponseWriter, r *http.Request) {
	var viaje domain.Viaje
	err := json.NewDecoder(r.Body).Decode(&viaje)
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), 400)
		return
	}
	viaje_respuesta, err2 := h.svc.EmpezarViaje(viaje)
	if err2 != nil {
		// TODO: armar handler error
		http.Error(w, err2.Error(), 500)
		return
	}
	// TODO: armar respuesta, chi render
	render.JSON(w, r, viaje_respuesta)
}

func (h *ViajeHandler) TerminarViaje(w http.ResponseWriter, r *http.Request) {
	id_param := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(id_param, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), 400)
	}

	err = h.svc.TerminarViaje(id)
	if err != nil {
		http.Error(w, "Internal error: "+err.Error(), 500)
	}

	render.PlainText(w, r, "Viaje terminado")
}
func (h *ViajeHandler) PausarViaje(w http.ResponseWriter, r *http.Request) {
	id_param := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(id_param, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), 400)
	}

	err = h.svc.PausarViaje(id)
	if err != nil {
		http.Error(w, "Internal error: "+err.Error(), 500)
	}

	render.PlainText(w, r, "Viaje pausado")
}
func (h *ViajeHandler) ReanudarViaje(w http.ResponseWriter, r *http.Request) {
	id_param := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(id_param, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), 400)
	}

	err = h.svc.ReanudarViaje(id)
	if err != nil {
		http.Error(w, "Internal error: "+err.Error(), 500)
	}

	render.PlainText(w, r, "Viaje reanudado")
}

func (h *ViajeHandler) ReportConPausas(w http.ResponseWriter, r *http.Request) {
	id_monopatin_param := chi.URLParam(r, "id")
	id_monopatin, err := strconv.ParseInt(id_monopatin_param, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), 400)
	}

	report, err := h.svc.ReportConPausas(id_monopatin)
	if err != nil {
		http.Error(w, "Internal error: "+err.Error(), 500)
	}

	render.JSON(w, r, report)
}

func (h *ViajeHandler) ReportSinPausas(w http.ResponseWriter, r *http.Request) {
	id_monopatin_param := chi.URLParam(r, "id")
	id_monopatin, err := strconv.ParseInt(id_monopatin_param, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), 400)
	}

	report, err2 := h.svc.ReportSinPausas(id_monopatin)
	if err2 != nil {
		http.Error(w, "Internal error: "+err2.Error(), 500)
	}

	render.JSON(w, r, report)
}
