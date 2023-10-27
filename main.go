package main

import (
	//"fmt"
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

// Funcion para la conexion de la base de datos
func BDConnection() (connection *sql.DB) {
	Driver := "mysql"
	User := "root"
	Password := ""
	Name := "system"

	connection, err := sql.Open(Driver, User+":"+Password+"@tcp(127.0.0.1)/"+Name)
	if err != nil {
		panic(err.Error())
	}
	return connection
}

var templates = template.Must(template.ParseGlob("templates/*"))

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/Create", Create)
	log.Println("Running server...")
	http.ListenAndServe(":3000", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	SuccessfulConnection := BDConnection()
	InsertRegister, err := SuccessfulConnection.Prepare("INSERT INTO employees(name, email) VALUES('Jesus Cortez', '200348@utags.edu.mx')")
	if err != nil {
		panic(err.Error())
	}
	InsertRegister.Exec()

	templates.ExecuteTemplate(w, "Home", nil)
}

func Create(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Create", nil)
}
