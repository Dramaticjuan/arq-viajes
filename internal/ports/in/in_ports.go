package in

import (
	"github.com/Dramaticjuan/arq3-viajes/internal/core/domain"
)

type ViajeService interface {
	EmpezarViaje(viaje domain.Viaje) (*domain.Viaje, error)
	TerminarViaje(id_viaje int) error
	PausarViaje(id_viaje int) error
	ReanudarViaje(id_viaje int) (*domain.Viaje, error)
}
