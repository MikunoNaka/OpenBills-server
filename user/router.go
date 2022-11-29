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

package user

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)


func Routes(route *gin.Engine) {
	u := route.Group("/user")
	{
		u.POST("/new", validateMiddleware(), func(ctx *gin.Context) {
			u := ctx.MustGet("user").(User)
			// TODO: maybe add an invite code for some instances

			_, err := saveUser(u)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not login"})
				log.Printf("ERROR: Failed to add new user %v to DB: %v\n", u, err.Error())
				return
			}

			log.Printf("Successfully saved new user to DB: %s", u.UserName)
			ctx.JSON(http.StatusOK, nil)
		})

    u.PUT("/:userId", func(ctx *gin.Context) {
      id := ctx.Param("userId")
      objectId, err := primitive.ObjectIDFromHex(id)
      if err != nil {
          ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
          log.Printf("ERROR: Failed to modify user, Error parsing ID: %v\n", err.Error())
          return
      }

      var u User
      ctx.BindJSON(&u)
      err = modifyUser(objectId, u)
      if err != nil {
          ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          log.Printf("ERROR: Failed to modify user %v: %v\n", objectId, err.Error())
          return
      }

      log.Printf("Modified user %v to %v.\n", objectId, u)
      ctx.JSON(http.StatusOK, nil)
    })

		u.DELETE("/:userId", func(ctx *gin.Context) {
			id := ctx.Param("userId")
			objectId, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to delete user, Error parsing ID: %v\n", err.Error())
				return
			}

			err = deleteUser(objectId)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to delete user %v: %v\n", objectId, err.Error())
				return
			}

			log.Printf("Deleted user %v from database.\n", objectId )
			ctx.JSON(http.StatusOK, nil)
		})
	}
}
