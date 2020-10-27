package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5435
	user     = "boloindya_prod"
	password = "boloadmin3011"
	dbname   = "boloindya"
)

func main() {
	router := gin.Default()

	router.GET("/user", getUser)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	router.Run(":8080")
}

func getUser(c *gin.Context) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	rows, err := db.Query("select id, name from forum_user_userprofile limit 1")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println("id=", id)
		fmt.Println("name=", name)
	}

	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
}
