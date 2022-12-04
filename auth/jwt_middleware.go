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
	"net/http"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.Request.Header["Authorization"]
		if tokenHeader != nil {
			token, err := jwt.ParseWithClaims(tokenHeader[0], &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(accessSecret), nil
			})
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "access token expired"})
			} else {
			    ctx.Set("userId", token.Claims.(*jwt.StandardClaims).Issuer)
			    ctx.Next()
			}
		} else {
		    // invalid Authorization header
	        ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not logged in"})
		}

	}
}

func verifyRefreshToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		refreshToken, err := ctx.Cookie("refreshToken")
		if err == nil {
			token, err := jwt.ParseWithClaims(refreshToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(refreshSecret), nil
			})
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "refresh token expired"})
			} else {
			    ctx.Set("userId", token.Claims.(*jwt.StandardClaims).Issuer)
			    ctx.Next()
			}
		} else {
		    // invalid Authorization header
	        ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not logged in"})
		}
	}
}
