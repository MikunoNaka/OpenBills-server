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

package brand

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/MikunoNaka/OpenBills-server/database"
)

var items *mongo.Collection = database.DB.Collection("Items")

// Add brand to db
func saveBrand(b Brand) (primitive.ObjectID, error) {
	res, err := db.InsertOne(context.TODO(), b)
	return res.InsertedID.(primitive.ObjectID), err
}

// Delete brand from DB
func deleteBrand(id primitive.ObjectID) error {
	// delete brand
	_, err := db.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}

	// delete items associated with this brand
	_, err = items.DeleteMany(context.TODO(), bson.M{"Brand._id": id})
	return err
}

// modify brand in DB
func modifyBrand(id primitive.ObjectID, nb Brand) error {
	_, err := db.UpdateOne(context.TODO(), bson.D{{"_id", id}}, bson.D{{"$set", nb}})
	return err
}

/* GetBrands queries the database and
 * returns brands based on the given filter
 * if filter is nil every brand is returned
 */
func getBrands(filter bson.M) ([]Brand, error) {
	var brands []Brand

	cursor, err := db.Find(context.TODO(), filter)
	if err != nil {
		return brands, err
	}

	err = cursor.All(context.TODO(), &brands)
	return brands, err
}
