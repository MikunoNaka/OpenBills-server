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
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Add user to db
func saveUser(u User) (primitive.ObjectID, error) {
	u.hashPassword()
	res, err := db.InsertOne(context.TODO(), u)
	return res.InsertedID.(primitive.ObjectID), err
}

// Delete user from DB
func deleteUser(id primitive.ObjectID) error {
	_, err := db.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

// modify user in DB
func modifyUser(id primitive.ObjectID, nu User) error {
	fmt.Println(nu.Password)
	_, err := db.UpdateOne(context.TODO(), bson.D{{"_id", id}}, bson.D{{"$set", nu}})
	return err
}
