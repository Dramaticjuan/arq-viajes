package in

import (
	"github.com/Dramaticjuan/arq3-viajes/internal/core/domain"
	"github.com/Dramaticjuan/arq3-viajes/internal/dto"
)

type ViajeService interface {
	EmpezarViaje(viaje domain.Viaje) (*domain.Viaje, error)
	TerminarViaje(id_viaje int64) error
	PausarViaje(id_viaje int64) error
	ReanudarViaje(id_viaje int64) error
	ReportConPausas(id_monopatin int64) (*dto.ReporteMonopatin, error)
	ReportSinPausas(id_monopatin int64) (*dto.ReporteMonopatin, error)
}
