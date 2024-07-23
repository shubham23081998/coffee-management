package main

import (
	"backend/config"
	"backend/routes"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	r := gin.New()
	r.Use(cors.Default())
	config.Connect()
	routes.UserRoute(r)


//	enableCors(&*http.ResponseWriter)
	r.Run("localhost:8002")
}
