package brand

import (
	"github.com/gin-gonic/gin"
	"github.com/MikunoNaka/OpenBills-lib/brand"
	"log"
	"net/http"
)


func Routes(route *gin.Engine) {
	b := route.Group("/brand")
	{
		b.POST("/", func(ctx *gin.Context) {
			var x brand.Brand
			ctx.Bind(&x)
			err := x.Save()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to add new brand \"%s\": %v", x.Name, err.Error())
			}

			log.Println("Added new brand to database: ", x.Name)
			ctx.JSON(http.StatusOK, nil)
		})

		b.DELETE("/", func(ctx *gin.Context) {
			var x brand.Brand
			ctx.Bind(&x)
			err := x.Delete()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to delete brand \"%s\": %v", x.Name, err.Error())
			}

			log.Println("Delete brand: ", x.Name)
			ctx.JSON(http.StatusOK, nil)
		})
	}
}
