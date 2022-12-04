/* OpenBills-server - Server for libre billing software OpenBills-web
 * Copyright (C) 2022  Vidhu Kant Sharma <vidhukant@vidhukant.xyz>

 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.

 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.

 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package item

import (
	"github.com/gin-gonic/gin"
	"github.com/MikunoNaka/OpenBills-server/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func Routes(route *gin.Engine) {
	i := route.Group("/item")
	i.Use(auth.Authorize())
	{
		// TODO: add functionality to filter results
		// /all returns all the saved items
		i.GET("/all", func(ctx *gin.Context) {
			items, err := getItems(nil)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to read items from DB: %v\n", err.Error())
				return
			}

			ctx.JSON(http.StatusOK, items)
		})

		i.POST("/new", func(ctx *gin.Context) {
			var i Item
			ctx.BindJSON(&i)
			_, err := saveItem(i)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to add new item %v to DB: %v\n", i, err.Error())
				return
			}

			log.Printf("Successfully saved new item to DB: %v", i)
			ctx.JSON(http.StatusOK, nil)
		})

		i.PUT("/:itemId", func(ctx *gin.Context) {
			id := ctx.Param("itemId")
			objectId, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to modify item, Error parsing ID: %v\n", err.Error())
				return
			}

			var i Item
			ctx.BindJSON(&i)
			err = modifyItem(objectId, i)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to modify item %v: %v\n", objectId, err.Error())
				return
			}

			log.Printf("Modified item %v to %v.\n", objectId, i)
			ctx.JSON(http.StatusOK, nil)
		})

		i.DELETE("/:itemId", func(ctx *gin.Context) {
			id := ctx.Param("itemId")
			objectId, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to delete item, Error parsing ID: %v\n", err.Error())
				return
			}

			err = deleteItem(objectId)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to delete item %v: %v\n", objectId, err.Error())
				return
			}

			log.Printf("Deleted item %v from database.\n", objectId )
			ctx.JSON(http.StatusOK, nil)
		})
	}
}
