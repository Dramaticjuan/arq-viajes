package domain

import (
	"time"
)

type Pausa struct {
	Comienzo time.Time
	Fin      time.Time
}

func (p *Pausa) calcularDuracion() float64 {
	return p.Comienzo.Sub(p.Fin).Seconds()
}
