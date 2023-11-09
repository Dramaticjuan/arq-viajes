package services

import "github.com/Dramaticjuan/arq3-viajes/internal/dto"

type UsuarioServiceImpl struct {
	url string
}

func NewUsuarioServiceImpl(url string) *UsuarioServiceImpl {
	return &UsuarioServiceImpl{
		url: url,
	}
}
func (usi *UsuarioServiceImpl) GetCuentaUsuario(id_usuario int64) (*dto.CuentaDTO, error) {
	// TODO
	return nil, nil
}
func (usi *UsuarioServiceImpl) CobrarViaje(id_cuenta int64, monto float64) (*dto.CobroDTO, error) {
	// TODO
	return nil, nil
}
