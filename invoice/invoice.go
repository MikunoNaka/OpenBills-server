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

package invoice

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/MikunoNaka/OpenBills-server/client"
	"github.com/MikunoNaka/OpenBills-server/item"
	"github.com/MikunoNaka/OpenBills-server/database"
	"time"
)

// initialise a database connection for this package
// not sure if I should do this but I am...
var db *mongo.Database = database.DB

/* you should be able to:
 * - add, modify, delete an invoice
 * - add client to invoice
 * - add items to invoice
 */

/* Transporter details can be stored in
 * the DB. That is decided by the frontend.
 * You can optionally store Transporter
 * and Transport details which are often used
 */
type Transporter struct {
  Id primitive.ObjectID `bson:"_id,omitempty" json:"Id"`
  Name          string  `bson:"Name" json:"Name"`
  GSTIN         string  `bson:"GSTIN" json:"GSTIN"`
	// Issued ID for the transporter if any
  TransporterId string  `bson:"TransporterId,omitempty" json:"TransporterId"`
}

// transport vehicle details
type Transport struct {
  Id primitive.ObjectID   `bson:"_id,omitempty" json:"Id"`
  Transporter Transporter `bson:"Transporter,omitempty" json:"Transporter"`
  VehicleNum  string      `bson:"VehicleNum" json:"VehicleNum"`
  Note        string      `bson:"Note" json:"Note"`
  TransportMethod string  `bson:"TransportMethod" json:"TransportMethod"`
}

/* The *legendary* Invoice struct
 * Each Recipient, Item in invoice, Address
 * every detail that can change in the future is
 * saved in the database so even if values change
 * the invoice will have the old details
 *
 * The _id of the items/recipients will also be stored
 * so user can look at the new values of those fields
 * if needed. This system is better because if
 * item is deleted from the Db it won't mess
 * up the invoice collection
 *
 * Things like IGST, CGST, Discount, Quantity, etc
 * should be calculated on runtime.
 *
 * usually an invoice would store the currency
 * for payment. OpenBills does NOT support
 * international billing. The Db will hold the config
 * for the default currency, etc.
 */
// TODO: add place of supply
type Invoice struct {
  Id primitive.ObjectID         `bson:"_id,omitempty" json:"Id"` // not the same as invoice number
  InvoiceNumber   int           `bson:"InvoiceNumber" json:"InvoiceNumber"`
  CreatedAt       time.Time     `bson:"CreatedAt" json:"CreatedAt"`
  LastUpdated     time.Time     `bson:"LastUpdated,omitempty" json:"LastUpdated"`
  Recipient       client.Client `bson:"Recipient" json:"Recipient"`
  Paid            bool          `bson:"Paid" json:"Paid"`
  TransactionId   string        `bson:"TransactionId" json:"TransactionId"`
  Transport       Transport     `bson:"Transport" json:"Transport"`
  // user can apply a discount on the whole invoice
  // TODO: float64 isn't the best for this
  DiscountPercentage float64    `bson:"DiscountPercentage" json:"DiscountPercentage"`
  /* client may have multiple shipping
   * addresses but invoice only has one.
   * Empty ShippingAddress means shipping
   * address same as billing address
   */
  BillingAddress  client.Address     `bson:"BillingAddress" json:"BillingAddress"`
  ShippingAddress client.Address     `bson:"ShippingAddress,omitempty" json:"ShippingAddress"`
  Items           []item.InvoiceItem `bson:"Items" json:"Items"`
  // user can attach notes to the invoice
  // frontend decides if recipient sees this or not
  Note            string             `bson:"Note" json:"Note"`

  /* Invoices can be drafts
   * I personally like this functionality
   * because we can constantly save the
   * invoice to the DB as a draft
   * and if OpenBills crashes or is disconnected
   * we still have the progress
   */
  Draft bool `bson:"Draft" json:"Draft"`
}
