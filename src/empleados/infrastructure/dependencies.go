package infrastructure

import (
	"github.com/jackc/pgx/v5/pgxpool"

	empleadosApp "github.com/Luis1f2/empleados_api/src/empleados/application"
	empleadosController "github.com/Luis1f2/empleados_api/src/empleados/infrastructure/controller"
	empleadosRepo "github.com/Luis1f2/empleados_api/src/empleados/infrastructure/repository"
)

type EmpleadosDependencies struct {
	EmpleadoController *empleadosController.EmpleadoController
}

func InitEmpleadosDependencies(db *pgxpool.Pool) *EmpleadosDependencies {
	empleadoRepository := empleadosRepo.NewEmpleadoRepositoryPostgres(db)

	createEmpleado := empleadosApp.NewCreateEmpleado(empleadoRepository)
	getEmpleadoByID := empleadosApp.NewGetEmpleadoByID(empleadoRepository)
	listEmpleados := empleadosApp.NewListEmpleados(empleadoRepository)
	updateEmpleado := empleadosApp.NewUpdateEmpleado(empleadoRepository)
	deleteEmpleado := empleadosApp.NewDeleteEmpleado(empleadoRepository)

	empleadoController := empleadosController.NewEmpleadoController(
		createEmpleado,
		getEmpleadoByID,
		listEmpleados,
		updateEmpleado,
		deleteEmpleado,
	)

	return &EmpleadosDependencies{
		EmpleadoController: empleadoController,
	}
}