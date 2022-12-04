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

package brand

import (
  "github.com/MikunoNaka/OpenBills-server/auth"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)


func Routes(route *gin.Engine) {
	b := route.Group("/brand")
 	b.Use(auth.Authorize())
	{
		b.GET("/all", func(ctx *gin.Context) {
			// TODO: add functionality to filter results
			brands, err := getBrands(nil)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to read brands from DB: %v\n", err.Error())
				return
			}

			ctx.JSON(http.StatusOK, brands)
		})

		b.POST("/new", func(ctx *gin.Context) {
			var b Brand
			ctx.BindJSON(&b)
			_, err := saveBrand(b)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to add new brand %v to DB: %v\n", b, err.Error())
				return
			}

			log.Printf("Successfully saved new brand to DB: %v", b)
			ctx.JSON(http.StatusOK, nil)
		})

		b.PUT("/:brandId", func(ctx *gin.Context) {
			id := ctx.Param("brandId")
			objectId, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to modify brand, Error parsing ID: %v\n", err.Error())
				return
			}

			var b Brand
			ctx.BindJSON(&b)
			err = modifyBrand(objectId, b)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to modify brand %v: %v\n", objectId, err.Error())
				return
			}

			log.Printf("Modified brand %v to %v.\n", objectId, b)
			ctx.JSON(http.StatusOK, nil)
		})

		b.DELETE("/:brandId", func(ctx *gin.Context) {
			id := ctx.Param("brandId")
			objectId, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to delete brand, Error parsing ID: %v\n", err.Error())
				return
			}

			err = deleteBrand(objectId)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to delete brand %v: %v\n", objectId, err.Error())
				return
			}

			log.Printf("Deleted brand %v from database.\n", objectId )
			ctx.JSON(http.StatusOK, nil)
		})
	}
}
