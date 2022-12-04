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

package client

import (
  "github.com/MikunoNaka/OpenBills-server/auth"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Routes(route *gin.Engine) {
	c := route.Group("/client")
 	c.Use(auth.Authorize())
	{
		c.GET("/all", func(ctx *gin.Context) {
			// TODO: add functionality to filter results
			clients, err := getClients(nil)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to read clients from DB: %v\n", err.Error())
				return
			}

			ctx.JSON(http.StatusOK, clients)
		})

		c.POST("/new", func(ctx *gin.Context) {
			var c Client
			ctx.BindJSON(&c)
			_, err := saveClient(c)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to add new client %v to DB: %v\n", c, err.Error())
				return
			}

			log.Printf("Successfully saved new client to DB: %v", c)
			ctx.JSON(http.StatusOK, nil)
		})

		c.PUT("/:clientId", func(ctx *gin.Context) {
			id := ctx.Param("clientId")
			objectId, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to modify client, Error parsing ID: %v\n", err.Error())
				return
			}

			var c Client
			ctx.BindJSON(&c)
			err = modifyClient(objectId, c)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to modify client %v: %v\n", objectId, err.Error())
				return
			}

			log.Printf("Modified client %v to %v.\n", objectId, c)
			ctx.JSON(http.StatusOK, nil)
		})

		c.DELETE("/:clientId", func(ctx *gin.Context) {
			id := ctx.Param("clientId")
			objectId, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to delete client, Error parsing ID: %v\n", err.Error())
				return
			}

			err = deleteClient(objectId)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to delete client %v: %v\n", objectId, err.Error())
				return
			}

			log.Printf("Deleted client %v from database.\n", objectId )
			ctx.JSON(http.StatusOK, nil)
		})
	}
}
