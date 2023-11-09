package dto

type ConfiguracionDTO struct {
	PrecioComun float64
	PrecioPausa float64
	TarifaExtra float64
}

type MonopatinDTO struct {
	Kilometros float64
	Parada     *string
}

type CuentaDTO struct {
	Id_cuenta  int64
	Habilitada bool
}

type CobroDTO struct {
	Id_cuenta int64
	Saldo     float64
}

type ReporteMonopatin struct {
	Id_monopatin int64
	Tiempo       float64
}

type ParadaDTO struct {
	Id_parada int64
}
