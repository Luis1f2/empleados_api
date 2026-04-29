package repository

import "github.com/Luis1f2/empleados_api/src/roles/domain/entities"

type RolRepository interface {
	Create(nombre string) (*entities.Rol, error)
	GetByID(id int) (*entities.Rol, error)
	GetByName(nombre string) (*entities.Rol, error)
	List() ([]entities.Rol, error)
	Update(id int, nombre string) (*entities.Rol, error)
	Delete(id int) error
}