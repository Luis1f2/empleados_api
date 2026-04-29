package application

import domainRepository "github.com/Luis1f2/empleados_api/src/roles/domain/repository"

type DeleteRole struct {
	repo domainRepository.RolRepository
}

func NewDeleteRole(repo domainRepository.RolRepository) *DeleteRole {
	return &DeleteRole{repo: repo}
}

func (d *DeleteRole) Execute(id int) error {
	return d.repo.Delete(id)
}