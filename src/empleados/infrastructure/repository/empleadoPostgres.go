package repository

import (
	"context"

	"github.com/Luis1f2/empleados_api/src/empleados/domain/entities"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmpleadoRepositoryPostgres struct {
	db *pgxpool.Pool
}

func NewEmpleadoRepositoryPostgres(db *pgxpool.Pool) *EmpleadoRepositoryPostgres {
	return &EmpleadoRepositoryPostgres{db: db}
}

func (r *EmpleadoRepositoryPostgres) Create(empleado *entities.Empleado) (*entities.Empleado, error) {
	var created entities.Empleado

	err := r.db.QueryRow(
		context.Background(),
		`INSERT INTO empleado (
			nombre,
			apellidos,
			mail,
			password,
			username,
			desactivado,
			rol_id,
			eliminado,
			created_at,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, false, $6, NULL, NOW(), NOW())
		RETURNING id, nombre, apellidos, mail, password, username, desactivado, rol_id, eliminado, created_at, updated_at`,
		empleado.Nombre,
		empleado.Apellidos,
		empleado.Mail,
		empleado.Password,
		empleado.User,
		empleado.RolID,
	).Scan(
		&created.ID,
		&created.Nombre,
		&created.Apellidos,
		&created.Mail,
		&created.Password,
		&created.User,
		&created.Desactivado,
		&created.RolID,
		&created.Eliminado,
		&created.CreatedAt,
		&created.UpdatedAt,
	)

	return &created, err
}

func (r *EmpleadoRepositoryPostgres) GetByID(id int) (*entities.Empleado, error) {
	var empleado entities.Empleado

	err := r.db.QueryRow(
		context.Background(),
		`SELECT id, nombre, apellidos, mail, password, username, desactivado, rol_id, eliminado, created_at, updated_at
		 FROM empleado
		 WHERE id = $1 AND eliminado IS NULL`,
		id,
	).Scan(
		&empleado.ID,
		&empleado.Nombre,
		&empleado.Apellidos,
		&empleado.Mail,
		&empleado.Password,
		&empleado.User,
		&empleado.Desactivado,
		&empleado.RolID,
		&empleado.Eliminado,
		&empleado.CreatedAt,
		&empleado.UpdatedAt,
	)

	return &empleado, err
}

func (r *EmpleadoRepositoryPostgres) GetByUser(user string) (*entities.Empleado, error) {
	var empleado entities.Empleado

	err := r.db.QueryRow(
		context.Background(),
		`SELECT id, nombre, apellidos, mail, password, username, desactivado, rol_id, eliminado, created_at, updated_at
		 FROM empleado
		 WHERE username = $1 AND eliminado IS NULL AND desactivado = false`,
		user,
	).Scan(
		&empleado.ID,
		&empleado.Nombre,
		&empleado.Apellidos,
		&empleado.Mail,
		&empleado.Password,
		&empleado.User,
		&empleado.Desactivado,
		&empleado.RolID,
		&empleado.Eliminado,
		&empleado.CreatedAt,
		&empleado.UpdatedAt,
	)

	return &empleado, err
}

func (r *EmpleadoRepositoryPostgres) ExistsByUser(user string) (bool, error) {
	var exists bool

	err := r.db.QueryRow(
		context.Background(),
		`SELECT EXISTS(
			SELECT 1 FROM empleado 
			WHERE username = $1 AND eliminado IS NULL
		)`,
		user,
	).Scan(&exists)

	return exists, err
}

func (r *EmpleadoRepositoryPostgres) List() ([]entities.Empleado, error) {
	rows, err := r.db.Query(
		context.Background(),
		`SELECT id, nombre, apellidos, mail, password, username, desactivado, rol_id, eliminado, created_at, updated_at
		 FROM empleado
		 WHERE eliminado IS NULL
		 ORDER BY id ASC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var empleados []entities.Empleado

	for rows.Next() {
		var empleado entities.Empleado

		err := rows.Scan(
			&empleado.ID,
			&empleado.Nombre,
			&empleado.Apellidos,
			&empleado.Mail,
			&empleado.Password,
			&empleado.User,
			&empleado.Desactivado,
			&empleado.RolID,
			&empleado.Eliminado,
			&empleado.CreatedAt,
			&empleado.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		empleados = append(empleados, empleado)
	}

	return empleados, nil
}

func (r *EmpleadoRepositoryPostgres) Update(id int, empleado *entities.Empleado) (*entities.Empleado, error) {
	var updated entities.Empleado

	err := r.db.QueryRow(
		context.Background(),
		`UPDATE empleado
		 SET nombre = $1,
		     apellidos = $2,
		     mail = $3,
		     username = $4,
		     password = $5,
		     rol_id = $6,
		     updated_at = NOW()
		 WHERE id = $7 AND eliminado IS NULL
		 RETURNING id, nombre, apellidos, mail, password, username, desactivado, rol_id, eliminado, created_at, updated_at`,
		empleado.Nombre,
		empleado.Apellidos,
		empleado.Mail,
		empleado.User,
		empleado.Password,
		empleado.RolID,
		id,
	).Scan(
		&updated.ID,
		&updated.Nombre,
		&updated.Apellidos,
		&updated.Mail,
		&updated.Password,
		&updated.User,
		&updated.Desactivado,
		&updated.RolID,
		&updated.Eliminado,
		&updated.CreatedAt,
		&updated.UpdatedAt,
	)

	return &updated, err
}

func (r *EmpleadoRepositoryPostgres) Delete(id int) error {
	_, err := r.db.Exec(
		context.Background(),
		`UPDATE empleado
		 SET eliminado = NOW(), updated_at = NOW()
		 WHERE id = $1 AND eliminado IS NULL`,
		id,
	)

	return err
}