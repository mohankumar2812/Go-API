package app

import (
	"fmt"
	"log"
	"os"

	"example/model"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/gin-contrib/cors"
)

var (
	r = gin.Default()
)

func StartApp() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	config := cors.DefaultConfig()

	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Authorization"}

	r.Use(cors.New(config))
	
	model.DBConnect()
	Router()
	port := os.Getenv("PORT")
	r.SetTrustedProxies([]string{"http://localhost:3000"})
	r.Run(port)

	fmt.Println("Port number: ", port)
	
}
