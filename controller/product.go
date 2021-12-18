package controller

import (
	"database/sql"
	"fmt"
	"log"
	"sharplang/src/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DBConn struct {
	DB *sql.DB
}

var productService *service.Product = service.NewProduct()

func (conn *DBConn) GetAllProducts(c *gin.Context) {
	fmt.Println(c.Param("id"))

	tx, err := conn.DB.Begin()

	if err != nil {
		log.Fatal(err)
	}

	productService.GetAllProducts(tx)
}

func (conn *DBConn) GetProductById(c *gin.Context) {
	id, strToIntErr := strconv.Atoi(c.Param("id"))

	tx, err := conn.DB.Begin()

	if err != nil || strToIntErr != nil {
		log.Fatal(err)
	}

	productService.GetProductById(tx, id)

}
