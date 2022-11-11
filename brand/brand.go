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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/MikunoNaka/OpenBills-server/database"
)

// initialise a database connection for this package
// not sure if I should do this but I am...
var db *mongo.Collection = database.DB.Collection("Brands")

/* An item may or may not be
 * assigned to a brand
 *
 * brands can be used to group products
 * to perform certain actions
 */
type Brand struct {
  Id   primitive.ObjectID `bson:"_id,omitempty" json:"Id"`
  Name string             `bson:"Name" json:"Name"`
}
