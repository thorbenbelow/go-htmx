package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

type Db struct {
	pool *sql.DB
}

func initDb(path string) (*Db, error) {
	pool, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	if _, err := pool.Exec(`CREATE TABLE IF NOT EXISTS users (id integer not null primary key autoincrement, name text);`); err != nil {
		return nil, err
	}

	return &Db{pool}, nil
}

func (db *Db) addUser(mail string) error {
	stmt, err := db.pool.Prepare(`insert into users(name) values(?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(mail); err != nil {
		return err
	}
	return nil
}

func (db *Db) getUsers() ([]string, error) {
	rows, err := db.pool.Query(`SELECT name FROM users;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := make([]string, 0)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		log.Print(name)
		res = append(res, name)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	log.Print(res)
	return res, nil
}

func main() {
	router := gin.Default()
	db, err := initDb("./test.db")
	if err != nil {
		log.Fatal(err)
	}

	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(ctx *gin.Context) {
		users, err := db.getUsers()
		if err != nil {
			log.Print(err)
		}
		ctx.HTML(http.StatusOK, "index.tmpl", map[string]interface{}{"newUserForm": "", "userList": users})
	})

	router.POST("/user", func(ctx *gin.Context) {
		var newUser NewUser
		ctx.Bind(&newUser)
		db.addUser(newUser.User)
		users, err := db.getUsers()
		if err != nil {
			log.Print(err)
		}
		ctx.HTML(http.StatusOK, "UserList.tmpl", gin.H{"userList": users})
	})
	log.Fatal(router.Run(":8080"))
}

type NewUser struct {
	User string `form:"user"`
}
