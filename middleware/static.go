package middleware

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

//Middleware for serving js and css files in respective directories
func SetupStaticFiles() (js gin.HandlerFunc, css gin.HandlerFunc) {
	js = static.Serve("/files/js", static.LocalFile("./files/js", true))
	css = static.Serve("/files/styles", static.LocalFile("./files/styles", true))
	return
}
