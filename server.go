package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"html/template"
)

var db Db

type Db struct {
	users []string
};

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
	if err := db.addUser("test@example.invalid"); err != nil {
		log.Fatal("You suck")
	}
	
	router.LoadHTMLGlob("templates/*")
	btn := template.HTML(`
	<button hx-post="/user"">
		Click Me
	</button>`)
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{ "msg": btn})
	})

	router.POST("/user", handlePostUser)
	log.Fatal(router.Run(":8080"))
}
func handlePostUser(ctx *gin.Context) {
	log.Print("yay")
}
