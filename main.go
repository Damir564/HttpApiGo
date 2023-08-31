package main

import (
	"log"

	"github.com/Damir564/HttpApiGo/config"
	_ "github.com/Damir564/HttpApiGo/docs"
	"github.com/Damir564/HttpApiGo/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	return r
}

// @title HTTP API GO
// @version 1.0
// @description Users and Segments

// @contact.name   Damir Nizamutdinov
// @contact.email  ddamir.nizamutdinov@yandex.ru

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host	localhost:8080
// @BasePath /
func main() {
	// os.Setenv("TZ", "Europe/Moscow")
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables %s", err.Error())
	}
	r := setupRouter()
	config.Connect()
	routes.Route(r)
	r.Run(":8080")
}
