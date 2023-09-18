package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
)

var db Db

type Db struct {
	users []string
}

func initDb() Db {
	return Db{users: make([]string, 0)}
}

func (db *Db) addUser(mail string) error {
	db.users = append(db.users, mail)
	return nil
}

func main() {
	router := gin.Default()
	db = initDb()

	router.LoadHTMLGlob("templates/*")
	newUserForm := template.HTML(`
		`)

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", map[string]interface{}{"newUserForm": newUserForm, "userList": db.users})
	})

	router.POST("/user", handlePostUser)
	log.Fatal(router.Run(":8080"))
}

type NewUser struct {
	User string `form:"user"`
}

func handlePostUser(ctx *gin.Context) {
	var newUser NewUser
	ctx.Bind(&newUser)
	db.addUser(newUser.User)
	ctx.HTML(http.StatusOK, "UserList.tmpl", gin.H{"userList": db.users})
}
