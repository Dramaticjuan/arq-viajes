package domain

import (
	"time"

	"github.com/Dramaticjuan/arq3-viajes/internal/dto"
	"gopkg.in/guregu/null.v4"
)

type Viaje struct {
	Id               int64
	Id_cuenta        int64
	Id_usuario       int64
	Id_monopatin     int64
	Comienzo         time.Time
	Fin              null.Time //opt
	Km_inicio        float64
	Km_fin           null.Float //opt
	Id_parada_inicio int64
	Id_parada_fin    null.Int //opt
	Pausas           []*Pausa
	Precio_final     null.Float //opt
}

func (v *Viaje) CalcularTiempoSinPausas() float64 {
	return v.Comienzo.Sub(v.Fin.Time).Seconds()
}

func (v *Viaje) CalcularTiempoConPausas() float64 {
	tiempo := v.CalcularTiempoSinPausas()
	tiempo -= v.calcularTimpoEnPausa()
	return tiempo
}

func (v *Viaje) contarPausasExtensas() int {
	var contador int
	for _, p := range v.Pausas {
		if p.calcularDuracion() >= 15.0 {
			contador++
		}
	}
	return contador
}

func (v *Viaje) calcularTimpoEnPausa() float64 {
	var total float64
	for _, p := range v.Pausas {
		total += p.calcularDuracion()
	}
	return total
}

func (v *Viaje) CalcularPrecio(configuracion *dto.ConfiguracionDTO) float64 {
	var total float64
	total += configuracion.PrecioComun * v.CalcularTiempoSinPausas()
	total += configuracion.PrecioPausa * v.calcularTimpoEnPausa()
	total += configuracion.TarifaExtra * float64(v.contarPausasExtensas())
	v.Precio_final.SetValid(total)
	return total
}
