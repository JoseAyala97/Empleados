package conection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConexionBD() (conexion *sql.DB) {
	Usuario := "postgres"
	Contraseña := "Jose1997*"
	Nombre := "crud"

	// Formato de la cadena de conexión para PostgreSQL
	cadenaConexion := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", Usuario, Contraseña, Nombre)

	// Abre la conexión a la base de datos
	conexion, err := sql.Open("postgres", cadenaConexion)
	if err != nil {
		panic(err.Error())
	}

	// Intenta ping para verificar la conexión
	err = conexion.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Conexión exitosa a la base de datos")
	return conexion
}
