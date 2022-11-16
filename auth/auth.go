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
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/MikunoNaka/OpenBills-server/database"
	"github.com/MikunoNaka/OpenBills-server/user"
	"net/http"
    //"golang.org/x/crypto/bcrypt"
)

var db *mongo.Collection = database.DB.Collection("Users")

func checkPassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var u user.User
	    ctx.BindJSON(&u)

		filter := bson.M{
			"UserName": u.UserName,
			"$or": bson.M{"Email": u.Email},
		}

		err := db.FindOne(context.TODO(), filter).Decode(&u)
		if err != nil {
			panic(err)
		}
		fmt.Println(u)
	}
}

func Routes(route *gin.Engine) {
	u := route.Group("/auth")
	{
		u.POST("/login", func(ctx *gin.Context) {
			checkPassword()(ctx)
			ctx.HTML(http.StatusOK, "<h1>Hello World</h1>", nil)
		})
	}
}
