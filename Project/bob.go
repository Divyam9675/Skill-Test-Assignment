package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Player struct {
	Id          int    `json:id`
	Title       string `json:name`
	Description string `json:role`
	Matches     string `json:matches`
	Age         string `json:age`
}

func dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(ctx *gin.Context) {
		//render only file, must full name with extension
		db := dbConn()
		selDB, err := db.Query("SELECT * FROM root ORDER BY id ASC")
		if err != nil {
			panic(err.Error())
		}
		player := Player{}
		res := []Player{}
		for selDB.Next() {
			var id int
			var matches, age string
			var name, role string
			err = selDB.Scan(&id, &name, &role, &matches, &age)
			if err != nil {
				panic(err.Error())
			}
			player.Id = id
			player.Title = name
			player.Description = role
			player.Matches = matches
			player.Age = age
			res = append(res, player)
		}
		//var a = "hello words"
		ctx.HTML(http.StatusOK, "page.html", gin.H{"title": "Home Page!!", "a": res})
	})

	r.GET("/add", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "add.html", gin.H{"title": "Add player!!"})
	})

	r.POST("/insert", func(ctx *gin.Context) {
		//render only file, must full name with extension
		var name, role string
		var matches, age string

		name = ctx.Request.FormValue("title")
		role = ctx.Request.FormValue("description")
		matches = ctx.Request.FormValue("date")
		age = ctx.Request.FormValue("priority")

		db := dbConn()
		insForm, err := db.Prepare("INSERT INTO root (Title, Description, Date, Priority) VALUES(?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, role, matches, age)
		// ctx.HTML(http.StatusOK, "updated.html", gin.H{"title": "Player"})

		selDB, err := db.Query("SELECT * FROM root ORDER BY id ASC")
		if err != nil {
			panic(err.Error())
		}
		player := Player{}
		res := []Player{}
		for selDB.Next() {
			var id int
			var matches, age string
			var name, role string
			err = selDB.Scan(&id, &name, &role, &matches, &age)
			if err != nil {
				panic(err.Error())
			}
			player.Id = id
			player.Title = name
			player.Description = role
			player.Matches = matches
			player.Age = age
			res = append(res, player)
		}
		//var a = "hello words"
		ctx.HTML(http.StatusOK, "page.html", gin.H{"title": "Home Page!!", "a": res})
	})

	r.Run(":8080")

}
