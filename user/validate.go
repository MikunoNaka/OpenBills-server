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
	"context"
	"strings"
	"net/mail"
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"errors"
)

var (
	errUsernameTaken = errors.New("username is taken")
	errUsernameTooShort = errors.New("username is too short")
	errUsernameInvalid = errors.New("invalid username")
	errEmailTaken = errors.New("email is taken")
	errEmailInvalid = errors.New("email is invalid")
	errPasswordTooShort = errors.New("password is too short")
)

func isUsernameTaken(username string) error {
	var x User
	err := db.FindOne(context.TODO(), bson.M{"UserName": username}).Decode(&x)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		} else {
			return err
		}
	} else {
		return errUsernameTaken
	}
	return nil
}

func isEmailTaken(email string) error {
	var x User
	err := db.FindOne(context.TODO(), bson.M{"Email": email}).Decode(&x)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		} else {
			return err
		}
	} else {
		return errEmailTaken
	}
	return nil
}

func validateUsername(username string) error {
	username = strings.Trim(username, " ")

	if len(username) < 2 {
		return errUsernameTooShort
	}

	if strings.Contains(username, " ") {
		return errUsernameInvalid
	}

	return isUsernameTaken(username)
}

func validateEmail(email string) error {
	email = strings.Trim(email, " ")

	// verify if string is valid email
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errEmailInvalid
	}

	return isEmailTaken(email)
}

func validatePassword(password string) error {
	// TODO: load password length from config
	if len(password) < 12 {
		return errPasswordTooShort
	}

	return nil
}

func validateMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var u User
		ctx.BindJSON(&u)

		// validate username
		isUsernameValid := validateUsername(u.UserName)
		switch isUsernameValid {
		case nil:
			break;
		case errUsernameTooShort, errUsernameInvalid:
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": isUsernameValid.Error()})
		case errUsernameTaken:
			ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": isUsernameValid.Error()})
		default:
		    log.Printf("Error while creating new user '%s': %s", u.UserName, isUsernameValid.Error())
		    ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error (cannot create user)"})
		}

		// validate email
		isEmailValid := validateEmail(u.Email)
		switch isEmailValid {
		case nil:
			break;
		case errEmailInvalid:
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": isEmailValid.Error()})
		case errEmailTaken:
			ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": isEmailValid.Error()})
		default:
		    log.Printf("Error while creating new user '%s': %s", u.UserName, isEmailValid.Error())
		    ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error (cannot create user)"})
		}

		// validate password
		isPasswordValid := validatePassword(u.Password)
		switch isPasswordValid {
		case nil:
			break;
		case errPasswordTooShort:
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": isPasswordValid.Error()})
		default:
		    ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error (cannot create user)"})
		    log.Printf("Error while creating new user '%s': %s", u.UserName, isPasswordValid.Error())
		}

		ctx.Set("user", u)
	}
}
