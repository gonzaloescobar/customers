// Instancio Go y JSON
package main

import (
	"encoding/json"
	"fmt"
)

// Defino los tipos de Datos de mi aplicación.
type Customer struct {
	ID       int
	Apellido string
	Nombre   string
	Edad     string
	Telefono string
	Ciudad   string
	Pais     string
}

// Creamos una Instancia con las Características del Cliente.
func main() {
	customer := Customer{
		ID:       1,
		Apellido: "Escobar",
		Nombre:   "Gonza",
		Edad:     "30",
		Telefono: "1558899345",
		Ciudad:   "Caseros",
		Pais:     "Argentina",
	}

	// Creamos nuestro JSON a partir de los datos que obtenemos desde la instancia Creada con las Características del Cliente.
	create_json, _ := json.Marshal(customer)

	// Convertimos los datos(bytes) en una cadena e imprimimos el contenido.
	convert_to_string := string(create_json)
	fmt.Println(convert_to_string)

}
