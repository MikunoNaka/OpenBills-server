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
	"go.mongodb.org/mongo-driver/bson/primitive"
  "go.mongodb.org/mongo-driver/mongo"
	"github.com/MikunoNaka/OpenBills-server/database"
)

// initialise a database connection for this package
// not sure if I should do this but I am...
var db *mongo.Collection = database.DB.Collection("Clients")

/* each invoice has a client
 * you should be able to:
 * - add, modify, delete a client
 * - add client to invoices
 * - get all invoices associated with client, etc
 */

/* each contact has one name
 * but multiple contact addresses
 * it is assumed that the first one
 * has the highest priority
 */

type Contact struct {
  Name    string   `bson:"Name" json:"Name"`
  Phones  []string `bson:"Phones" json:"Phones"`
  Emails  []string `bson:"Emails" json:"Emails"`
  Website string   `bson:"Website" json:"Website"`
}

type Address struct {
	/* "Text" means the actual address lines.
	 * If address is 123, xyz colony, myCity, myState the Text
	 * will be 123, xyz colony, and
	 * State and City will be myCity and myState
	 *
	 * A multiline string is expected.
	 */
  Text       string `bson:"Text" json:"Text"`
  City       string `bson:"City" json:"City"`
  State      string `bson:"State" json:"State"`
  PostalCode string `bson:"PostalCode" json:"PostalCode"`
  Country    string `bson:"Country" json:"Country"`
}

type Client struct {
  Id primitive.ObjectID `bson:"_id,omitempty" json:"Id"`
  Name    string        `bson:"Name" json:"Name"`
  Contact Contact       `bson:"Contact" json:"Contact"`
  GSTIN   string        `bson:"GSTIN" json:"GSTIN"`
	/* if shipping address is empty it means that
	 * the billing address is also shipping address
	 */
  BillingAddress    Address   `bson:"BillingAddress" json:"BillingAddress"`
  ShippingAddresses []Address `bson:"ShippingAddresses,omitempty" json:"ShippingAddresses"`
}
