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

package client

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* TODO: Handle errors properly
 * Send an API error response instead of log.Fatal
 */

// Add client to db
func saveClient(c Client) (primitive.ObjectID, error) {
	res, err := db.InsertOne(context.TODO(), c)
	return res.InsertedID.(primitive.ObjectID), err
}

// Delete client from DB
func deleteClient(id primitive.ObjectID) error {
	_, err := db.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

// modify client in DB
func modifyClient(id primitive.ObjectID, nc Client) error {
	_, err := db.UpdateOne(context.TODO(), bson.D{{"_id", id}}, bson.D{{"$set", nc}})
	return err
}

/* GetClients queries the database and
 * returns clients based on the given filter
 * if filter is nil every client is returned
 */
func getClients(filter bson.M) ([]Client, error) {
	var clients []Client

	cursor, err := db.Find(context.TODO(), filter)
	if err != nil {
		return clients, err
	}

	err = cursor.All(context.TODO(), &clients)
	return clients, err
}
