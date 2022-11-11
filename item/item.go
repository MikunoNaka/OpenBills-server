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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/MikunoNaka/OpenBills-server/database"
	"github.com/MikunoNaka/OpenBills-server/brand"
)

// initialise a database connection for this package
// not sure if I should do this but I am...
var db *mongo.Collection = database.DB.Collection("Items")

/* each invoice must contain at least one item
 * you should be able to:
 * - add, modify, delete an item
 * - add item to invoice
 */

/* An item is any product
 * or service that can be sold
 * Items may have a max and min quanity
 * and some default fields like GST and Unit Price
 * that don't need to be entered manually
 *
 * the front-end may or may not implement
 * the default fields
 *
 * Items can be assigned brands
 * and certain actions can be performed
 * on the products of a brand altogether
 */
type Item struct {
	Id primitive.ObjectID      `bson:"_id,omitempty" json:"Id"`
	Brand brand.Brand          `bson:"Brand,omitempty" json:"Brand"`
	UnitOfMeasure      string  `bson:"UnitOfMeasure" json:"UnitOfMeasure"`
	HasDecimalQuantity bool    `bson:"HasDecimalQuantity, json:"HasDecimalQuantity"`
    // just the defaults, can be overridden in an invoice
	Name               string  `bson:"Name" json:"Name"`
	Description        string  `bson:"Description" json:"Description"`
	HSN                string  `bson:"HSN" json:"HSN"`
	UnitPrice          float64 `bson:"UnitPrice" json:"UnitPrice"`
	// default tax percentage
	GSTPercentage      float64 `bson:"GSTPercentage" json:"GSTPercentage"`
	MaxQuantity        float64 `bson:"MaxQuantity" json:"MaxQuantity"`
	MinQuantity        float64 `bson:"MinQuantity" json:"MinQuantity"`
}

// Item but with extra fields an invoice might require
type InvoiceItem struct {
	Item
	/* Each product must have a quantity
	 * but it is upto the backend to enforce that
	 */
	// TODO: float64 isn't ideal, find a better way
	Quantity           float64 `bson:"Quantitiy" json:"Quantity"`
	DiscountPercentage float64 `bson:"DiscountPercentage,omitempty" json:"DiscountPercentage"`
}
