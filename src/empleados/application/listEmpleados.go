package application

import (
	"github.com/Luis1f2/empleados_api/src/empleados/domain/entities"
	domainRepository "github.com/Luis1f2/empleados_api/src/empleados/domain/repository"
)

type ListEmpleados struct {
	repo domainRepository.EmpleadoRepository
}

func NewListEmpleados(repo domainRepository.EmpleadoRepository) *ListEmpleados {
	return &ListEmpleados{repo: repo}
}

func (l *ListEmpleados) Execute() ([]entities.Empleado, error) {
	return l.repo.List()
}