package repository

import "github.com/Luis1f2/empleados_api/src/empleados/domain/entities"

type EmpleadoRepository interface {
	Create(empleado *entities.Empleado) (*entities.Empleado, error)
	GetByID(id int) (*entities.Empleado, error)
	GetByUser(user string) (*entities.Empleado, error)
	ExistsByUser(user string) (bool, error)
	List() ([]entities.Empleado, error)
	Update(id int, empleado *entities.Empleado) (*entities.Empleado, error)
	Delete(id int) error
}