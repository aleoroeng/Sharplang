package controller

import (
	"fmt"
	"log"
	"sharplang/src/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var usersService *service.Users = service.NewUsers()

func (conn *DBConn) GetAllUsers(c *gin.Context) {
	fmt.Println(c.Param("id"))

	tx, err := conn.DB.Begin()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(usersService.GetAllUsers(tx))
}

func (conn *DBConn) GetUserById(c *gin.Context) {
	id, strToIntErr := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	tx, err := conn.DB.Begin()

	if err != nil || strToIntErr != nil {
		log.Fatal(err)
	}

	usersService.GetUsersById(tx, id)

}
