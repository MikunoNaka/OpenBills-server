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

package auth

import (
	"github.com/gin-gonic/gin"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/MikunoNaka/OpenBills-server/database"
	"github.com/MikunoNaka/OpenBills-server/user"
	"net/http"
	"log"
    "golang.org/x/crypto/bcrypt"
)

var db *mongo.Collection = database.DB.Collection("Users")

func checkPassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var u user.User
	    ctx.BindJSON(&u)

		filter := bson.M{
			"$or": []bson.M{
			    // u.UserName in this case can be either username or email
				{"Email": u.UserName},
				{"UserName": u.UserName},
			},
		}

		// check if the user exists in DB
		var user user.User
		err := db.FindOne(context.TODO(), filter).Decode(&user)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "user does not exist"})
			} else {
				log.Printf("Error while reading user from DB to check password: %v", err.Error())
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			}
			ctx.Abort()
		}

		// compare hash and password
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
		if err != nil {
			if err == bcrypt.ErrMismatchedHashAndPassword {
			    ctx.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect password"})
			} else {
				log.Printf("Error while checking password: %v", err.Error())
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			}
			ctx.Abort()
		}

		// everything's fine!
		ctx.Set("user", user)
		ctx.Next()
	}
}

func Routes(route *gin.Engine) {
	r := route.Group("/auth")
	{
		r.POST("/login", checkPassword(), func(ctx *gin.Context) {
			user := ctx.MustGet("user").(user.User)
			ctx.JSON(http.StatusOK, user)
		})
	}
}
