package repo

import "database/sql"

type PostgresClient struct {
	*sql.DB
}

func NewPostgresClient(source string) *PostgresClient {
	db, err := sql.Open("postgres", source)

	if err != nil {
		panic(err)
	}

	db.Exec(`CREATE TABLE IF NOT EXISTS viaje(
		id BIGSERIAL primary key,
		id_cuenta        bigint NOT NULL,
		id_usuario       bigint NOT NULL,
		id_monopatin     bigint NOT NULL,
		comienzo         timestamp NOT NULL,
		fin              timestamp,
		km_inicio        float NOT NULL,
		km_fin           float,
		id_parada_inicio bigint NOT NULL,
		id_parada_fin    bigint,
		precio_final     float
	)`)

	db.Exec(`CREATE TABLE IF NOT EXISTS pausa(
		id BIGSERIAL primary key,
		comienzo         timestamp NOT NULL,
		fin              timestamp,
		id_viaje bigint  references viaje (id) not null
	);`)
	return &PostgresClient{db}
}
