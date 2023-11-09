package domain

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Pausa struct {
	Id       int64
	Comienzo time.Time
	Fin      null.Time
	Id_viaje int64
}

func (p *Pausa) calcularDuracion() float64 {
	return p.Comienzo.Sub(p.Fin.ValueOrZero()).Seconds()
}
