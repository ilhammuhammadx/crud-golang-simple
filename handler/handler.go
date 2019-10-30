package handler

import (
	"crud-golang-simple/model"
	"database/sql"	
	"fmt"
	// "html/template"
	// "log"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/kataras/go-sessions"
	// "os"
)

type user struct {
	ID int
	Username string
	FirstName string
	LastName string
	Password string
}

var username, password, host, nameDB, defaultDB string
var db *sql.DB
var err error

func init() {
	username = "root"
	password = ""
	host = "localhost"
	nameDB = "golang_crud"
	defaultDB = "mysql"
}

func QueryUser(username string) user {
	var users = user{}
	err = db.QueryRow(
		`SELECT id,
		username,
		first_name,
		last_name,
		password FROM users WHERE username=?`, username).
		
		Scan(
			&users.ID,
			&users.Username,
			&users.FirstName,
			&users.LastName,
			&users.Password,
		)
		
		return users
}

func checkErr(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {
		fmt.Println(r.Host + r.URL.Path)

		http.Redirect(w, r, r.Host + r.URL.Path, 301)
		return false
	}
	return true
}

func Register(w http.ResponseWriter, r *http.Request) {
	db, err = model.ConnectDB(username, password, host, nameDB)

	if err != nil {
		return
	}
	defer db.Close()

	if r.Method != "POST" {
		http.ServeFile(w,r, "views/register.html")
		return
	}

	username := r.FormValue("email")
	first_name := r.FormValue("first_name")	
	last_name := r.FormValue("last_name")	
	password := r.FormValue("password")

	users := QueryUser(username)

	if (user{}) == users {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if len(hashedPassword) != 0 && checkErr(w, r, err) {
			stmt, err := db.Prepare("INSERT INTO users SET username=?, password=?, first_name=?, last_name=?")
			if err == nil {
				_, err := stmt.Exec(&username, &hashedPassword, &first_name, &last_name)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
		}
	} else {
		http.Redirect(w, r, "/register", 302)
	}
}