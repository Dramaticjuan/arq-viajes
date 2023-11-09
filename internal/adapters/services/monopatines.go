package services

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Dramaticjuan/arq3-viajes/internal/dto"
)

type MonopatineServiceImpl struct {
	url string
}

func NewMonopatineServiceImpl(url string) *MonopatineServiceImpl {
	return &MonopatineServiceImpl{
		url: url,
	}
}

func (msi *MonopatineServiceImpl) GetMonopatin(id int64) (*dto.MonopatinDTO, error) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", msi.url+strconv.FormatInt(id, 10), nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var monopatin *dto.MonopatinDTO
	err = json.NewDecoder(res.Body).Decode(&monopatin)
	if err != nil {
		return nil, err
	}

	return monopatin, nil
}

func (msi *MonopatineServiceImpl) UpdateParadaMonopatin(id int64, id_parada int64) error {
	req, err := http.NewRequestWithContext(context.Background(), "PATCH", msi.url+strconv.FormatInt(id, 10)+"/parada/"+strconv.FormatInt(id_parada, 10), nil)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		data, err := ioutil.ReadAll(res.Body)
		reqBodyString := string(data)
		if err != nil {
			return err
		}
		return errors.New(reqBodyString)
	}
	return nil
}

func (msi *MonopatineServiceImpl) GetParadaCercana(id_monopatin int64) (*dto.ParadaDTO, error) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", msi.url+strconv.FormatInt(id_monopatin, 10)+"/parada", nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var parada *dto.ParadaDTO
	err = json.NewDecoder(res.Body).Decode(&parada)
	if err != nil {
		return nil, err
	}

	return parada, nil
}
