package main

import (
	"html/template"
	"log"
	"net/http"

	db "github.com/JoseAyala97/Empleados/Conexion"
	models "github.com/JoseAyala97/Empleados/Consultas"
)

// Metodos "template" que sirven para obtener información de una carpeta
var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	//localhost
	//Solicitud que sirve para acceder a la función de inicio
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/crear", Crear)
	http.HandleFunc("/insertar", Insertar)
	http.HandleFunc("/borrar", Borrar)
	http.HandleFunc("/editar", Editar)
	http.HandleFunc("/actualizar", Actualizar)

	log.Println("Servidor corriendo...")
	http.ListenAndServe(":8080", nil)
}

// Función para dirigirnos al inicio
func Inicio(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola")
	//se crea un objeto de esa variable anterior usando ExecuteTemplate para mostrar el contenido de HTML, con el uso de template

	//#Líneas de código para inserción directa a base de datos
	// conexionEstablecida := db.ConexionBD()
	// //Función prepare realiza una inserción directamente a la base de datos
	// insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO empleados(nombre, correo) VALUES ('lucas','lucas@gmail.com')")

	// if err != nil {
	// 	panic(err.Error())
	// }
	// //Método Exec ejecuta sentencia de inserción
	// insertarRegistros.Exec()

	//#Líneas de código para hacer consulta directa a base de datos
	conexionEstablecida := db.ConexionBD()
	//Función prepare realiza una inserción directamente a la base de datos
	registros, err := conexionEstablecida.Query("SELECT * FROM empleados")

	if err != nil {
		panic(err.Error())
	}

	empleado := models.Empleado{}
	arregloEmpleado := []models.Empleado{}
	for registros.Next() {
		var id int
		var nombre, correo string
		err = registros.Scan(&id, &nombre, &correo)
		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Correo = correo

		arregloEmpleado = append(arregloEmpleado, empleado)
	}

	plantillas.ExecuteTemplate(w, "inicio", arregloEmpleado)
}

// Función para crear
func Crear(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crear", nil)
}

// Función insertar, recepción de datos con el metodo post
func Insertar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")

		conexionEstablecida := db.ConexionBD()
		//Función prepare realiza una inserción directamente a la base de datos
		insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO empleados(nombre, correo) VALUES ($1,$2)")

		if err != nil {
			panic(err.Error())
		}
		insertarRegistros.Exec(nombre, correo)
		http.Redirect(w, r, "/", 301)
	}
}

func Borrar(w http.ResponseWriter, r *http.Request) {
	idEmpleado := r.URL.Query().Get("id")

	conexionEstablecida := db.ConexionBD()
	borrarRegistros, err := conexionEstablecida.Prepare("DELETE FROM empleados where id = $1")

	if err != nil {
		panic(err.Error())
	}
	borrarRegistros.Exec(idEmpleado)
	http.Redirect(w, r, "/", 301)
}

func Editar(w http.ResponseWriter, r *http.Request) {
	idEmpleado := r.URL.Query().Get("id")

	conexionEstablecida := db.ConexionBD()
	editarRegistro, err := conexionEstablecida.Query("SELECT * FROM empleados WHERE id=$1", idEmpleado)

	empleado := models.Empleado{}
	for editarRegistro.Next() {
		var id int
		var nombre, correo string
		err = editarRegistro.Scan(&id, &nombre, &correo)
		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Correo = correo
	}
	// Renderiza la página de edición y pasa el objeto empleado
	plantillas.ExecuteTemplate(w, "editar", empleado)
}

func Actualizar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")

		conexionEstablecida := db.ConexionBD()
		//Función prepare realiza una inserción directamente a la base de datos
		modificarRegistros, err := conexionEstablecida.Prepare("UPDATE empleados SET nombre=$1,correo=$2 WHERE id=$3")

		if err != nil {
			panic(err.Error())
		}
		modificarRegistros.Exec(nombre, correo, id)
		http.Redirect(w, r, "/", 301)
	}
}
