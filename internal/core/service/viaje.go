package service

import (
	"errors"

	"github.com/Dramaticjuan/arq3-viajes/internal/core/domain"
	"github.com/Dramaticjuan/arq3-viajes/internal/ports/in"
	"github.com/Dramaticjuan/arq3-viajes/internal/ports/out"
)

type ViajeServiceImpl struct {
	repo          out.ViajeProvider
	usuarios      out.UsuarioService
	monopatines   out.MonopatinService
	configuracion out.ConfiguracionService
}

func NewViajeServiceImpl(repo out.ViajeProvider, usuarios out.UsuarioService, monopatines out.MonopatinService, configuracion out.ConfiguracionService,) *ViajeServiceImpl {
	return &ViajeServiceImpl{
		repo:          repo,
		usuarios:      usuarios,
		monopatines:   monopatines,
		configuracion: configuracion,
	}
}

func (v *ViajeServiceImpl) EmpezarViaje(viaje domain.Viaje) (*domain.Viaje, error) {
    cuenta, err :=v.usuarios.GetCuentaUsuario(viaje.Id_usuario)
    if err != nil{
        return nil, err
    }
    viaje.Id_cuenta= cuenta.Id_cuenta
	return v.repo.EmpezarViaje(viaje)
}

func (v *ViajeServiceImpl) TerminarViaje(id_viaje int) (*domain.Viaje, error){
	viaje, err := v.repo.GetViajeById(id_viaje)
	if err != nil {
		return nil, err
	}
	if !viaje.Fin.IsZero() {
		return nil, errors.New("El viaje ya ha terminado")
	}
	pausa, errp := v.repo.UltimaPausa(id_viaje)
	if errp != nil {
		return nil, errp
	}
	if pausa.Fin.IsZero() {
		v.PausarViaje(id_viaje)
	}
    // TODO: chequear que monopatin esté en la parada y escribir los kilometros finales

    viaje, err = v.repo.TerminarViaje(id_viaje)
    if err != nil{
        return nil, err
    }
    precios_actuales, err := v.configuracion.GetPreciosActuales()
    if err != nil{
        return nil, err
    }
    precio:= viaje.CalcularPrecio(precios_actuales)

    v.usuarios.CobrarViaje(viaje.Id_cuenta, precio)
    return viaje, nil
}

func (v *ViajeServiceImpl) PausarViaje(id_viaje int) error {
	ult_pausa, err := v.repo.UltimaPausa(id_viaje)
	if err != nil {
		return err
	}
	if ult_pausa != nil && ult_pausa.Fin.IsZero() {
		return errors.New("Hay una pausa vigente")
	}
	return v.repo.PausarViaje(id_viaje)
}

func (v *ViajeServiceImpl) ReanudarViaje(id_viaje int) error {
	ult_pausa, err := v.repo.UltimaPausa(id_viaje)
	if err != nil {
		return err
	}
	if ult_pausa != nil && ult_pausa.Fin.IsZero() {
		return v.repo.ReanudarViaje(id_viaje)
	}
	return errors.New("No hay una pausa vigente")
}
