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

package item

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/MikunoNaka/OpenBills-server/database"
	"github.com/MikunoNaka/OpenBills-server/brand"
	"go.mongodb.org/mongo-driver/mongo"
)

var brands *mongo.Collection = database.DB.Collection("Brands")

// Add item to db
func saveItem(i Item) (primitive.ObjectID, error) {
	res, err := db.InsertOne(context.TODO(), i)
	return res.InsertedID.(primitive.ObjectID), err
}

// Delete item from DB
func deleteItem(id primitive.ObjectID) error {
	_, err := db.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

// modify item in DB
func modifyItem(id primitive.ObjectID, ni Item) error {
	_, err := db.UpdateOne(context.TODO(), bson.D{{"_id", id}}, bson.D{{"$set", ni}})
	return err
}

/* GetItems queries the database and
 * returns items based on the given filter
 * if filter is nil every item is returned
 */
func getItems(filter bson.M) ([]Item, error) {
	var items []Item

	cursor, err := db.Find(context.TODO(), filter)
	if err != nil {
		return items, err
	}

	err = cursor.All(context.TODO(), &items)
	if err != nil {
	    return items, err
	}

	for id, i := range items {
    // continue if item doesn't have a brand
    if (i.Brand.Id == primitive.ObjectID{}) {
      continue
    }

		var b brand.Brand

		err := brands.FindOne(context.TODO(), bson.M{"_id": i.Brand.Id}).Decode(&b)
		if err != nil {
			return items, err
		}
		items[id].Brand = b
	}

	return items, err
}
