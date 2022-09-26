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
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/MikunoNaka/OpenBills-lib/brand"
	"log"
	"net/http"
)


func Routes(route *gin.Engine) {
	b := route.Group("/brand")
	{
		b.GET("/all", func(ctx *gin.Context) {
			// TODO: add functionality to filter results
			brands, err := brand.GetBrands(nil)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to read brands from DB: %v\n", err.Error())
				return
			}

			ctx.JSON(http.StatusOK, brands)
		})

		b.DELETE("/:brandId", func(ctx *gin.Context) {
			id := ctx.Param("brandId")
			objectId, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to delete brand, Error parsing ID: %v\n", err.Error())
				return
			}

			err = brand.DeleteBrand(objectId)
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
