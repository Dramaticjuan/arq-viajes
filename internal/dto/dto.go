package dto

type ConfiguracionDTO struct {
	PrecioComun int
	PrecioPausa int
	TarifaExtra int
}

type MonopatinDTO struct {
	Kilometros float64
	Parada     string
}

type CuentaDTO struct {
	Id_cuenta  int
    Habilitada bool
}

type CobroDTO struct {
	id_cuenta int
	Saldo     int
}
