package application

import domainRepository "github.com/Luis1f2/empleados_api/src/empleados/domain/repository"

type DeleteEmpleado struct {
	repo domainRepository.EmpleadoRepository
}

func NewDeleteEmpleado(repo domainRepository.EmpleadoRepository) *DeleteEmpleado {
	return &DeleteEmpleado{repo: repo}
}

func (d *DeleteEmpleado) Execute(id int) error {
	return d.repo.Delete(id)
}