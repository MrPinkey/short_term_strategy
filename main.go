package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"shortTermStrategy/pkg/router"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
