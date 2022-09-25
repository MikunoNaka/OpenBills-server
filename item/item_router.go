package item

import (
	"github.com/gin-gonic/gin"
	"github.com/MikunoNaka/OpenBills-lib/item"
	"log"
	"net/http"
)


func Routes(route *gin.Engine) {
	i := route.Group("/item")
	{
		i.POST("/", func(ctx *gin.Context) {
			var x item.Item
			ctx.Bind(&x)
			err := x.Save()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}

			log.Println("Added new item to database: ", x.Name)
			ctx.JSON(http.StatusOK, nil)
		})
	}
}
