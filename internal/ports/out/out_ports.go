package out

import (
	"github.com/Dramaticjuan/arq3-viajes/internal/core/domain"
	"github.com/Dramaticjuan/arq3-viajes/internal/dto"
)

type ViajeProvider interface {
	GetViajeById(id_viaje int64) (*domain.Viaje, error)
	EmpezarViaje(domain.Viaje) (*domain.Viaje, error)
	TerminarViaje(domain.Viaje) (*domain.Viaje, error)
	GuardarPrecio(id_viaje int64, precio float64) error

	EmpezarPausa(id_viaje int64) error
	UltimaPausaSinTerminar(id_viaje int64) (*domain.Pausa, error)
	TerminarPausa(id_viaje int64) error
	ListViajesByMonopatin(id_monopatin int64) ([]*domain.Viaje, error)
}
type ConfiguracionService interface {
	GetPreciosActuales() (*dto.ConfiguracionDTO, error)
}
type MonopatinService interface {
	GetMonopatin(id_monopatin int64) (*dto.MonopatinDTO, error)
	GetParadaCercana(id_monopatin int64) (*dto.ParadaDTO, error)
	UpdateParadaMonopatin(id_monopatin int64, id_parada int64) error
}
type UsuarioService interface {
	GetCuentaUsuario(id_usuario int64) (*dto.CuentaDTO, error)
	CobrarViaje(id_cuenta int64, monto float64) error
}
