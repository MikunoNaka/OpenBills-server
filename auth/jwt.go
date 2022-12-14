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
	"github.com/MikunoNaka/OpenBills-server/user"
	"github.com/MikunoNaka/OpenBills-server/util"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"context"
	"errors"
	"time"
)

var (
    errUserNotFound error = errors.New("user does not exist")
)

var accessSecret []byte
var refreshSecret []byte
func init() {
	conf := util.GetConfig().Crypto
	accessSecret = []byte(conf.AccessTokenSecret)
	refreshSecret = []byte(conf.RefreshTokenSecret)
}

func newAccessToken(userId string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims {
	    Issuer: userId,
		ExpiresAt: time.Now().Add(time.Second * 15).Unix(),
	})

	token, err := claims.SignedString(accessSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

/*
 * the refresh token has a long lifespan and is stored in
 * the database in case it needs to be revoked.
 *
 * this can be stored as an HTTP only cookie and will be used
 * when creating a new access token
 *
 * I'm using a different secret key for refresh tokens
 * for enhanced security
 */
func newRefreshToken(userId string) (string, int64, error) {
	// convert id from string to ObjectID
	id, _ := primitive.ObjectIDFromHex(userId)

	// check if user exists
	var u user.User
	if err := db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&u); err != nil {
		return "", 0, errUserNotFound
	}

	// generate refresh token
	expiresAt := time.Now().Add(time.Hour * 12).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims {
	    Issuer: userId,
		ExpiresAt: expiresAt,
	})
	token, err := claims.SignedString(refreshSecret)
	if err != nil {
		return "", expiresAt, err
	}

	// store refresh token in db with unique session name for ease in identification
	sessionName := time.Now().Format("01-02-2006.15:04:05") + "-" + u.UserName
	u.Sessions = append(u.Sessions, user.Session{Name: sessionName, Token: token})
	db.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.D{{"$set", u}})

	return token, expiresAt, nil
}
