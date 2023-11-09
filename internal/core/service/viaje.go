package service

import (
	"errors"
	"time"

	"github.com/Dramaticjuan/arq3-viajes/internal/core/domain"
	"github.com/Dramaticjuan/arq3-viajes/internal/dto"
	"github.com/Dramaticjuan/arq3-viajes/internal/ports/out"
)

type ViajeServiceImpl struct {
	repo          out.ViajeProvider
	usuarios      out.UsuarioService
	monopatines   out.MonopatinService
	configuracion out.ConfiguracionService
}

func NewViajeServiceImpl(repo out.ViajeProvider, usuarios out.UsuarioService, monopatines out.MonopatinService, configuracion out.ConfiguracionService) *ViajeServiceImpl {
	return &ViajeServiceImpl{
		repo:          repo,
		usuarios:      usuarios,
		monopatines:   monopatines,
		configuracion: configuracion,
	}
}

func (v *ViajeServiceImpl) EmpezarViaje(viaje domain.Viaje) (*domain.Viaje, error) {
	cuenta, err := v.usuarios.GetCuentaUsuario(viaje.Id_usuario)
	if err != nil {
		return nil, err
	}
	viaje.Id_cuenta = cuenta.Id_cuenta
	return v.repo.EmpezarViaje(viaje)
}

func (v *ViajeServiceImpl) TerminarViaje(id_viaje int64) (*domain.Viaje, error) {
	viaje, err := v.repo.GetViajeById(id_viaje)
	if err != nil {
		return nil, err
	}
	if !viaje.Fin.IsZero() {
		return nil, errors.New("el viaje ya ha terminado")
	}
	pausa, errp := v.repo.UltimaPausaSinTerminar(id_viaje)
	if errp != nil {
		return nil, errp
	}
	if pausa.Fin.IsZero() {
		v.PausarViaje(id_viaje)
	}

	parada, err := v.monopatines.GetParadaCercana(viaje.Id_monopatin)
	if err != nil {
		return nil, err
	}
	monopatin, err := v.monopatines.GetMonopatin(viaje.Id_monopatin)
	if err != nil {
		return nil, err
	}
	viaje.Fin.SetValid(time.Now())
	viaje.Id_parada_fin.SetValid(parada.Id_parada)
	viaje.Km_fin.SetValid(monopatin.Kilometros)
	viaje, err = v.repo.TerminarViaje(*viaje)
	if err != nil {
		return nil, err
	}
	v.monopatines.UpdateParadaMonopatin(viaje.Id_monopatin, viaje.Id_parada_fin.ValueOrZero())
	precios_actuales, err := v.configuracion.GetPreciosActuales()
	if err != nil {
		return nil, err
	}
	precio := viaje.CalcularPrecio(precios_actuales)
	v.repo.GuardarPrecio(id_viaje, precio)

	err = v.usuarios.CobrarViaje(viaje.Id_cuenta, precio)
	if err != nil {
		return nil, err
	}
	return viaje, nil
}

func (v *ViajeServiceImpl) PausarViaje(id_viaje int64) error {
	viaje_vigente, err := v.repo.GetViajeById(id_viaje)
	if err != nil {
		return err
	}
	if !viaje_vigente.Fin.ValueOrZero().IsZero() {
		return errors.New("el viaje ya terminó")
	}
	pausa_vigente, err := v.repo.UltimaPausaSinTerminar(id_viaje)
	if err != nil {
		return err
	}
	if pausa_vigente != nil {
		return errors.New("hay una pausa vigente")
	}
	return v.repo.EmpezarPausa(id_viaje)
}

func (v *ViajeServiceImpl) ReanudarViaje(id_viaje int64) error {
	pausa_vigente, err := v.repo.UltimaPausaSinTerminar(id_viaje)
	if err != nil {
		return err
	}
	if pausa_vigente != nil {
		return v.repo.TerminarPausa(id_viaje)
	}
	return errors.New("no hay una pausa vigente")
}

func (v *ViajeServiceImpl) ReportConPausas(id_monopatin int64) (*dto.ReporteMonopatin, error) {
	viajes, err := v.repo.ListViajesByMonopatin(id_monopatin)
	if err != nil {
		return nil, err
	}
	r := dto.ReporteMonopatin{Id_monopatin: id_monopatin}
	for i := 0; i < len(viajes); i++ {
		r.Tiempo -= viajes[i].CalcularTiempoConPausas()
	}
	return &r, nil
}
func (v *ViajeServiceImpl) ReportSinPausas(id_monopatin int64) (*dto.ReporteMonopatin, error) {
	viajes, err := v.repo.ListViajesByMonopatin(id_monopatin)
	if err != nil {
		return nil, err
	}
	r := dto.ReporteMonopatin{Id_monopatin: id_monopatin}
	for i := 0; i < len(viajes); i++ {
		r.Tiempo -= viajes[i].CalcularTiempoSinPausas()
	}
	return &r, nil
}
