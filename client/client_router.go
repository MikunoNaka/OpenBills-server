package client

import (
	"github.com/gin-gonic/gin"
	"github.com/MikunoNaka/OpenBills-lib/client"
	"log"
	"net/http"
)

func Routes(route *gin.Engine) {
	c := route.Group("/client")
	{
		c.POST("/", func(ctx *gin.Context) {
			var x client.Client
			ctx.Bind(&x)
			err := x.Save()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to add new client \"%s\": %v", x.Name, err.Error())
			}

			log.Println("Added new client to database: ", x.Name)
			ctx.JSON(http.StatusOK, nil)
		})

		c.DELETE("/", func(ctx *gin.Context) {
			var x client.Client
			ctx.Bind(&x)
			err := x.Delete()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to delete client \"%s\": %v", x.Name, err.Error())
			}

			log.Println("Deleted client: ", x.Name)
			ctx.JSON(http.StatusOK, nil)
		})
	}
}
