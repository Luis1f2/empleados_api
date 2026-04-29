package entities

type Rol struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Eliminado bool   `json:"eliminado"`
}