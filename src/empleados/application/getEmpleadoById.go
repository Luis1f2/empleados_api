package application

import (
	"github.com/Luis1f2/empleados_api/src/empleados/domain/entities"
	domainRepository "github.com/Luis1f2/empleados_api/src/empleados/domain/repository"
)

type GetEmpleadoByID struct {
	repo domainRepository.EmpleadoRepository
}

func NewGetEmpleadoByID(repo domainRepository.EmpleadoRepository) *GetEmpleadoByID {
	return &GetEmpleadoByID{repo: repo}
}

func (g *GetEmpleadoByID) Execute(id int) (*entities.Empleado, error) {
	return g.repo.GetByID(id)
}