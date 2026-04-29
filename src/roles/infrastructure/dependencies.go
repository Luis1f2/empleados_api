package infrastructure

import (
	"github.com/jackc/pgx/v5/pgxpool"

	rolesApp "github.com/Luis1f2/empleados_api/src/roles/application"
	rolesController "github.com/Luis1f2/empleados_api/src/roles/infrastructure/controller"
	rolesRepo "github.com/Luis1f2/empleados_api/src/roles/infrastructure/repository"
)

type RolesDependencies struct {
	RolController *rolesController.RolController
}

func InitRolesDependencies(db *pgxpool.Pool) *RolesDependencies {
	rolRepository := rolesRepo.NewRolRepositoryPostgres(db)

	createRole := rolesApp.NewCreateRole(rolRepository)
	getRoleByID := rolesApp.NewGetRoleByID(rolRepository)
	listRoles := rolesApp.NewListRoles(rolRepository)
	updateRole := rolesApp.NewUpdateRole(rolRepository)
	deleteRole := rolesApp.NewDeleteRole(rolRepository)

	rolController := rolesController.NewRolController(
		createRole,
		getRoleByID,
		listRoles,
		updateRole,
		deleteRole,
	)

	return &RolesDependencies{
		RolController: rolController,
	}
}