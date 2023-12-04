package main

import (
	"html/template"
	"log"
	"net/http"
)

// Metodos "template" que sirven para obtener información de una carpeta
var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	//localhost
	//Solicitud que sirve para acceder a la función de inicio
	http.HandleFunc("/", Inicio)
	log.Println("Servidor corriendo...")
	http.ListenAndServe(":8080", nil)
}

func Inicio(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola")
	//se crea un objeto de esa variable anterior usando ExecuteTemplate para mostrar el contenido de HTML, con el uso de template
	plantillas.ExecuteTemplate(w, "inicio", nil)
}
