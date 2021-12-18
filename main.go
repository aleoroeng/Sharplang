package main

import (
	"database/sql"
	"fmt"
	"sharplang/src/auth"
	"sharplang/src/controller"
	"sharplang/src/middleware"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func setupConn() *sql.DB {
	connString := "dbname=postgres user=postgres password=postgres port=5432 sslmode=disable"
	connector, errConnector := pq.NewConnector(connString)

	if errConnector != nil {
		fmt.Println(errConnector)
		return nil
	}
	fmt.Println("Connecting")
	return sql.OpenDB(connector)
}

func main() {
	//Connection to DB which will
	dbConn := controller.DBConn{DB: setupConn()}
	router := gin.Default()
	router.LoadHTMLGlob("files/templates/*")

	router.GET("/index", controller.GetIndex)
	router.GET("/product/:id", dbConn.GetProductById)
	router.GET("/products", dbConn.GetAllProducts)

	router.GET("/users", dbConn.GetAllUsers)
	router.GET("/user/:id", dbConn.GetUserById)
	router.GET("/auth", auth.Init)
	router.GET("/auth2", auth.GetToken)
	router.Use(middleware.SetupStaticFiles())

	router.Run("localhost:8080")
}
