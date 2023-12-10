package main

import (
	"html/template"
	"log"
	"net/http"

	db "github.com/JoseAyala97/Empleados/Conexion"
)

// Metodos "template" que sirven para obtener información de una carpeta
var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	//localhost
	//Solicitud que sirve para acceder a la función de inicio
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/crear", Crear)
	http.HandleFunc("/insertar", Insertar)

	log.Println("Servidor corriendo...")
	http.ListenAndServe(":8080", nil)
}

// Función para dirigirnos al inicio
func Inicio(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola")
	//se crea un objeto de esa variable anterior usando ExecuteTemplate para mostrar el contenido de HTML, con el uso de template

	// conexionEstablecida := db.ConexionBD()
	// //Función prepare realiza una inserción directamente a la base de datos
	// insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO empleados(nombre, correo) VALUES ('lucas','lucas@gmail.com')")

	// if err != nil {
	// 	panic(err.Error())
	// }
	// //Método Exec ejecuta sentencia de inserción
	// insertarRegistros.Exec()

	plantillas.ExecuteTemplate(w, "inicio", nil)
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
