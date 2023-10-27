package main

import (
	//"fmt"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	Id    int
	Name  string
	Email string
}

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
	http.HandleFunc("/Insert", Insert)
	http.HandleFunc("/Delete", Delete)
	http.HandleFunc("/Edit", Edit)
	http.HandleFunc("/Update", Update)

	log.Println("Running server...")
	http.ListenAndServe(":3000", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	SuccessfulConnection := BDConnection()
	Registers, err := SuccessfulConnection.Query("SELECT * FROM employees")
	if err != nil {
		panic(err.Error())
	}
	employee := Employee{}
	arrayEmployee := []Employee{}

	for Registers.Next() {
		var id int
		var name, email string
		err = Registers.Scan(&id, &name, &email)

		if err != nil {
			panic(err.Error())
		}
		employee.Id = id
		employee.Name = name
		employee.Email = email

		arrayEmployee = append(arrayEmployee, employee)
	}
	//fmt.Println(arrayEmployee)

	templates.ExecuteTemplate(w, "Home", arrayEmployee)
}

func Create(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Create", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		email := r.FormValue("email")

		SuccessfulConnection := BDConnection()
		InsertRegister, err := SuccessfulConnection.Prepare("INSERT INTO employees(name, email) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		InsertRegister.Exec(name, email)

		http.Redirect(w, r, "/", 301)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idEmployee := r.URL.Query().Get("id")
	fmt.Println(idEmployee)

	SuccessfulConnection := BDConnection()
	RegisterEdit, err := SuccessfulConnection.Query("SELECT * FROM employees WHERE id=?", idEmployee)

	employee := Employee{}
	for RegisterEdit.Next() {
		var id int
		var name, email string
		err = RegisterEdit.Scan(&id, &name, &email)

		if err != nil {
			panic(err.Error())
		}
		employee.Id = id
		employee.Name = name
		employee.Email = email
	}
	fmt.Println(employee)
	templates.ExecuteTemplate(w, "Edit", employee)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		email := r.FormValue("email")

		SuccessfulConnection := BDConnection()
		UpdateRegister, err := SuccessfulConnection.Prepare("UPDATE employees SET name=?, email=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		UpdateRegister.Exec(name, email, id)

		http.Redirect(w, r, "/", 301)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idEmployee := r.URL.Query().Get("id")
	fmt.Println(idEmployee)

	SuccessfulConnection := BDConnection()
	DeleteRegister, err := SuccessfulConnection.Prepare("DELETE FROM employees WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	DeleteRegister.Exec(idEmployee)

	http.Redirect(w, r, "/", 301)
}
