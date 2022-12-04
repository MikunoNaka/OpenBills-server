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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/MikunoNaka/OpenBills-server/database"
    "golang.org/x/crypto/bcrypt"
)

var db *mongo.Collection = database.DB.Collection("Users")

// per-user config can be shared to DB
type Config struct {
}

type Session struct {
	Name string `bson:"Name" json:"Name"`
	Token string `bson:"Token" json:"Token"`
}

type User struct {
    Id   primitive.ObjectID `bson:"_id,omitempty" json:"Id"`
    UserName string         `bson:"UserName" json:"UserName"`
    Email    string         `bson:"Email" json:"Email"`
    Password string         `bson:"Password" json:"Password"`
	Config   Config         `bson:"Config" json:"Config"`
	Sessions []Session      `bson:"Sessions" json:"Sessions"`
	// some actions are only available when email is verified
	Verified bool           `bson:"Verified" json:"Verified"`
}

func (u *User) hashPassword() error {
	// TODO: password validation
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	return nil
}
