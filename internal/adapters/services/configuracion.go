package services

import (
	"github.com/Dramaticjuan/arq3-viajes/internal/dto"
)

type ConfiguracionServiceImpl struct {
	url string
}

func NewConfiguracionServiceImpl(url string) *ConfiguracionServiceImpl {
	return &ConfiguracionServiceImpl{
		url: url,
	}
}

func (csi *ConfiguracionServiceImpl) GetPreciosActuales() (*dto.ConfiguracionDTO, error) {
	return nil, nil
}
