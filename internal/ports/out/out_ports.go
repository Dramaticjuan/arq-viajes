package out

import (
	"github.com/Dramaticjuan/arq3-viajes/internal/core/domain"
	"github.com/Dramaticjuan/arq3-viajes/internal/dto"
)

type ViajeProvider interface {
	EmpezarViaje(domain.Viaje) (*domain.Viaje, error)
	GetViajeById(id_viaje int) (*domain.Viaje, error)
	TerminarViaje(id_viaje int) (*domain.Viaje, error)
	PausarViaje(id_viaje int) error
	UltimaPausa(id_viaje int) (*domain.Pausa, error)
	ReanudarViaje(id_viaje int) error
}
type ConfiguracionService interface {
	GetPreciosActuales() (*dto.ConfiguracionDTO, error)
}
type MonopatinService interface {
	GetMonopatin(id_monopatin int) (*dto.MonopatinDTO, error)
	UpdateParadaMonopatin(id_monopatin int, id_parada int) error
}
type UsuarioService interface {
	GetCuentaUsuario(id_usuario int) (*dto.CuentaDTO, error)
	CobrarViaje(id_cuenta int, monto float64) (*dto.CobroDTO, error)
}
