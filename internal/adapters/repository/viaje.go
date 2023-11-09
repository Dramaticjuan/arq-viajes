package repo

import (
	"github.com/Dramaticjuan/arq3-viajes/internal/core/domain"
)

type RepoViajeImpl struct {
	db *PostgresClient
}

func NewRepoViajeImpl(db *PostgresClient) *RepoViajeImpl {
	return &RepoViajeImpl{db: db}
}

const empezarViaje = `-- name: EmpezarViaje :one
INSERT INTO viaje (
Id_cuenta        ,
Id_usuario       ,
Id_monopatin     ,
Comienzo         ,
Km_inicio        ,
Id_parada_inicio 
) VALUES (
  ?, ?, ?, ?, ?, ?
)
RETURNING id, id_cuenta, id_usuario, id_monopatin, comienzo, fin, km_inicio, km_fin, id_parada_inicio, id_parada_fin, precio_final
`

func (rv *RepoViajeImpl) EmpezarViaje(d domain.Viaje) (*domain.Viaje, error) {
	row := rv.db.QueryRow(empezarViaje,
		d.Id_cuenta,
		d.Id_usuario,
		d.Id_monopatin,
		d.Comienzo,
		d.Km_inicio,
		d.Id_parada_inicio,
	)
	i := domain.Viaje{}
	err := row.Scan(
		&i.Id,
		&i.Id_cuenta,
		&i.Id_usuario,
		&i.Id_monopatin,
		&i.Comienzo,
		&i.Fin,
		&i.Km_inicio,
		&i.Km_fin,
		&i.Id_parada_inicio,
		&i.Id_parada_fin,
		&i.Precio_final,
	)
	if err != nil {
		return nil, err
	}
	return &i, err
}

const getViajeById = `-- name: GetViajeById :one
SELECT id, id_cuenta, id_usuario, id_monopatin, comienzo, fin, km_inicio, km_fin, id_parada_inicio, id_parada_fin, precio_final FROM viaje
WHERE id = ?
`

func (rv *RepoViajeImpl) GetViajeById(id_viaje int64) (*domain.Viaje, error) {
	row := rv.db.QueryRow(getViajeById, id_viaje)
	i := domain.Viaje{}
	err := row.Scan(
		&i.Id,
		&i.Id_cuenta,
		&i.Id_usuario,
		&i.Id_monopatin,
		&i.Comienzo,
		&i.Fin,
		&i.Km_inicio,
		&i.Km_fin,
		&i.Id_parada_inicio,
		&i.Id_parada_fin,
		&i.Precio_final,
	)
	if err != nil {
		return nil, err
	}
	i.Pausas, err = rv.GetPausasByViaje(i.Id)
	return &i, err
}

const terminarViaje = `-- name: TerminarViaje :one
UPDATE viaje
SET fin = ?, km_fin = ?, Id_parada_fin = ?
WHERE id= ?
RETURNING id, id_cuenta, id_usuario, id_monopatin, comienzo, fin, km_inicio, km_fin, id_parada_inicio, id_parada_fin, precio_final
`

func (rv *RepoViajeImpl) TerminarViaje(d domain.Viaje) (*domain.Viaje, error) {
	row := rv.db.QueryRow(terminarViaje, d.Fin, d.Km_fin, d.Id_parada_fin, d.Id)
	err := row.Scan(
		&d.Id,
		&d.Id_cuenta,
		&d.Id_usuario,
		&d.Id_monopatin,
		&d.Comienzo,
		&d.Fin,
		&d.Km_inicio,
		&d.Km_fin,
		&d.Id_parada_inicio,
		&d.Id_parada_fin,
		&d.Precio_final,
	)
	return &d, err
}

const guardarPrecio = `-- name: guardarPrecio :exec
UPDATE vijae
SET precio_final= ?
WHERE id= ?
`

func (rv *RepoViajeImpl) GuardarPrecio(id_viaje int64, precio float64) error {
	_, err := rv.db.Exec(guardarPrecio, precio, id_viaje)
	if err != nil {
		return err
	}
	return nil
}

const listViajesByMonopatin = `-- name: ListViajesById :many
SELECT id, id_cuenta, id_usuario, id_monopatin, comienzo, fin, km_inicio, km_fin, id_parada_inicio, id_parada_fin, precio_final FROM viaje
WHERE id_monopatin= ?
`

func (rv *RepoViajeImpl) ListViajesByMonopatin(id_monopatin int64) ([]*domain.Viaje, error) {
	rows, err := rv.db.Query(listViajesByMonopatin, id_monopatin)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*domain.Viaje{}
	for rows.Next() {
		i := domain.Viaje{}
		if err = rows.Scan(
			&i.Id,
			&i.Id_cuenta,
			&i.Id_usuario,
			&i.Id_monopatin,
			&i.Comienzo,
			&i.Fin,
			&i.Km_inicio,
			&i.Km_fin,
			&i.Id_parada_inicio,
			&i.Id_parada_fin,
			&i.Precio_final,
		); err != nil {
			return nil, err
		}
		i.Pausas, err = rv.GetPausasByViaje(i.Id)
		if err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPausasByViaje = `-- name: GetPausasByViaje :many
SELECT id, comienzo, fin, id_viaje
FROM pausa
WHERE id_viaje= ?
`

func (rv *RepoViajeImpl) GetPausasByViaje(id_viaje int64) ([]*domain.Pausa, error) {
	rows, err := rv.db.Query(getPausasByViaje, id_viaje)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []*domain.Pausa{}
	for rows.Next() {
		i := domain.Pausa{}
		if err := rows.Scan(
			&i.Id,
			&i.Comienzo,
			&i.Fin,
			&i.Id_viaje,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const empezarPausa = `-- name: EmpezarPausa :exec
INSERT INTO pausa (comienzo, id_viaje)
VALUES(NOW()::TIMESTAMP, ?)
`

func (rv *RepoViajeImpl) EmpezarPausa(id_viaje int64) error {
	_, err := rv.db.Exec(empezarPausa, id_viaje)
	if err != nil {
		return err
	}

	return nil
}

const ultimaPausaSinTerminar = `-- name: UltimaPausaSinTerminar :one
SELECT id, comienzo, fin, id_viaje
FROM pausa
WHERE id_viaje= ? AND fin IS NULL
`

func (rv *RepoViajeImpl) UltimaPausaSinTerminar(id_viaje int64) (*domain.Pausa, error) {
	row := rv.db.QueryRow(ultimaPausaSinTerminar, id_viaje)
	var i domain.Pausa
	err := row.Scan(
		&i.Id,
		&i.Comienzo,
		&i.Fin,
		&i.Id_viaje,
	)
	return &i, err
}

const terminarPausa = `-- name: TerminarPausa :exec
UPDATE pausa
SET fin= NOW()::TIMESTAMP
WHERE id= ?
`

func (rv *RepoViajeImpl) TerminarPausa(id_viaje int64) error {
	_, err := rv.db.Exec(terminarPausa, id_viaje)
	if err != nil {
		return err
	}

	return nil
}
