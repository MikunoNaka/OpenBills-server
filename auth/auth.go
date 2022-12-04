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
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/MikunoNaka/OpenBills-server/database"
	"github.com/MikunoNaka/OpenBills-server/user"
	"net/http"
	"log"
)

var db *mongo.Collection = database.DB.Collection("Users")

func Routes(route *gin.Engine) {
	r := route.Group("/auth")
	{
		r.POST("/login", checkPassword(), func(ctx *gin.Context) {
			user := ctx.MustGet("user").(user.User)

			accessToken, err := newAccessToken(user.Id.Hex())
			if err != nil {
				log.Printf("Error while generating new access token: %v", err)
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error (cannot login)"})
			}

			refreshToken, expiresAt, err := newRefreshToken(user.Id.Hex())
			if err != nil {
				log.Printf("Error while generating new refresh token: %v", err)
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error (cannot login)"})
			}

			ctx.SetCookie("refreshToken", refreshToken, int(expiresAt), "", "", true, true)
			ctx.JSON(http.StatusOK, gin.H{"accessToken": accessToken})
		})

		r.POST("/refresh", verifyRefreshToken(), func (ctx *gin.Context) {
			u := ctx.MustGet("user").(user.User)
			accessToken, err := newAccessToken(u.Id.Hex())
			if err != nil {
				log.Printf("Error while generating new access token: %v", err)
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error (cannot refresh session)"})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"accessToken": accessToken})
			}
		})
	}
}
