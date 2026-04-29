package repository

import (
	"context"
	"errors"

	"github.com/Luis1f2/empleados_api/src/roles/domain/entities"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RolRepositoryPostgres struct {
	db *pgxpool.Pool
}

func NewRolRepositoryPostgres(db *pgxpool.Pool) *RolRepositoryPostgres {
	return &RolRepositoryPostgres{db: db}
}

func (r *RolRepositoryPostgres) Create(nombre string) (*entities.Rol, error) {
	var rol entities.Rol

	err := r.db.QueryRow(
		context.Background(),
		`INSERT INTO rol (nombre, eliminado)
		 VALUES ($1, false)
		 RETURNING id, nombre, eliminado`,
		nombre,
	).Scan(&rol.ID, &rol.Nombre, &rol.Eliminado)

	if err != nil {
		return nil, err
	}

	return &rol, nil
}

func (r *RolRepositoryPostgres) GetByID(id int) (*entities.Rol, error) {
	var rol entities.Rol

	err := r.db.QueryRow(
		context.Background(),
		`SELECT id, nombre, eliminado
		 FROM rol
		 WHERE id = $1 AND eliminado = false`,
		id,
	).Scan(&rol.ID, &rol.Nombre, &rol.Eliminado)

	if err != nil {
		return nil, err
	}

	return &rol, nil
}

func (r *RolRepositoryPostgres) GetByName(nombre string) (*entities.Rol, error) {
	var rol entities.Rol

	err := r.db.QueryRow(
		context.Background(),
		`SELECT id, nombre, eliminado
		 FROM rol
		 WHERE nombre = $1 AND eliminado = false`,
		nombre,
	).Scan(&rol.ID, &rol.Nombre, &rol.Eliminado)

	if err != nil {
		return nil, err
	}

	return &rol, nil
}

func (r *RolRepositoryPostgres) List() ([]entities.Rol, error) {
	rows, err := r.db.Query(
		context.Background(),
		`SELECT id, nombre, eliminado
		 FROM rol
		 WHERE eliminado = false
		 ORDER BY id ASC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []entities.Rol

	for rows.Next() {
		var rol entities.Rol

		if err := rows.Scan(&rol.ID, &rol.Nombre, &rol.Eliminado); err != nil {
			return nil, err
		}

		roles = append(roles, rol)
	}

	return roles, nil
}

func (r *RolRepositoryPostgres) Update(id int, nombre string) (*entities.Rol, error) {
	var rol entities.Rol

	err := r.db.QueryRow(
		context.Background(),
		`UPDATE rol
		 SET nombre = $1
		 WHERE id = $2 AND eliminado = false
		 RETURNING id, nombre, eliminado`,
		nombre,
		id,
	).Scan(&rol.ID, &rol.Nombre, &rol.Eliminado)

	if err != nil {
		return nil, err
	}

	return &rol, nil
}

func (r *RolRepositoryPostgres) Delete(id int) error {
	var nombre string

	err := r.db.QueryRow(
		context.Background(),
		`SELECT nombre FROM rol WHERE id = $1 AND eliminado = false`,
		id,
	).Scan(&nombre)
	if err != nil {
		return err
	}

	if nombre == "admin" {
		return errors.New("no se puede eliminar el rol admin")
	}

	_, err = r.db.Exec(
		context.Background(),
		`UPDATE rol
		    SET eliminado = true
		    WHERE id = $1 AND eliminado = false`,
		id,
	)

	return err
}