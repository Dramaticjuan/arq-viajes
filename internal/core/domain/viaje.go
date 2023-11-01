package domain

import (
	"fmt"
	"time"

	"github.com/Dramaticjuan/arq3-viajes/internal/dto"
)

type Viaje struct {
	Id               int
	Id_cuenta        int
	Id_usuario       int
	Id_monopatin     int
	Comienzo         time.Time
	Id_parada_inicio int
	Pausas           []Pausa
	Fin              time.Time
	Id_parada_fin    string
	Precio_final     int
}

func (v *Viaje) CalcularTiempoSinPausas() float64 {
	return v.Comienzo.Sub(v.Fin).Seconds()
}

func (v *Viaje) CalcularTiempoConPausas() float64 {
	tiempo:= v.CalcularTiempoSinPausas()
	tiempo -= v.calcularTimpoEnPausa()
	return tiempo
}

func (v *Viaje) contarPausasExtensas() int{
    var contador int
	for _, p := range v.Pausas{
        if p.calcularDuracion() >= 15.0{
            contador++
        }
	}
    return contador
}

func (v *Viaje) calcularTimpoEnPausa() float64{
    var total float64
	for _, p := range v.Pausas {
		total +=p.calcularDuracion()
	}
    return total
}

func (v *Viaje) CalcularPrecio(configuracion *dto.ConfiguracionDTO) float64 {
    var total float64
    total += float64(configuracion.PrecioComun)* v.CalcularTiempoSinPausas()
    total += float64(configuracion.PrecioPausa)* v.calcularTimpoEnPausa()
    total += float64(configuracion.TarifaExtra* v.contarPausasExtensas())
    return total 
}
