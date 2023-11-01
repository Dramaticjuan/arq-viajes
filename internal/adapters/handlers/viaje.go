package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Dramaticjuan/arq3-viajes/internal/core/domain"
	"github.com/Dramaticjuan/arq3-viajes/internal/ports/in"
	"github.com/go-chi/chi/v5"
)

type ViajeHandler struct {
    svc in.ViajeService
}

func NewViajeHandler(svc in.ViajeService) *ViajeHandler {
    return &ViajeHandler{
        svc: svc,
    }
}

func (h *ViajeHandler) EmpezarViaje(w http.ResponseWriter, r *http.Request){
    var viaje domain.Viaje
    err := json.NewDecoder(r.Body).Decode(&viaje)
    if err != nil{
        // TODO: armar handler error
        http.Error(w, "Bad Request", 400)
        return
    }
    viaje_respuesta, err2 := h.svc.EmpezarViaje(viaje)
    if err2 != nil{
        // TODO: armar handler error
        http.Error(w, err2.Error(), 500)
        return
    }
    // TODO: armar respuesta, chi render
}


func (h *ViajeHandler) TerminarViaje(w http.ResponseWriter, r *http.Request){
    id := chi.URLParam(r, "id")

    h.svc.TerminarViaje(id)
}
func (h *ViajeHandler) PausarViaje(w http.ResponseWriter, r *http.Request){
    h.svc.PausarViaje()
}
func (h *ViajeHandler) ReanudarViaje(w http.ResponseWriter, r *http.Request){
    h.svc.ReanudarViaje()
}
