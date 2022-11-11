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
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// add invoice to db
func saveInvoice(i Invoice) (primitive.ObjectID, error) {
	res, err := db.Collection("Invoices").InsertOne(context.TODO(), i)
	return res.InsertedID.(primitive.ObjectID), err
}

// add transporter to db
func saveTransporter(t Transporter) (primitive.ObjectID, error) {
	res, err := db.Collection("Transporters").InsertOne(context.TODO(), t)
	return res.InsertedID.(primitive.ObjectID), err
}

// add transport vehicle to db
func saveTransport(t *Transport) (primitive.ObjectID, error) {
	res, err := db.Collection("Transports").InsertOne(context.TODO(), t)
	return res.InsertedID.(primitive.ObjectID), err
}

// Delete invoice from DB
func deleteInvoice(id primitive.ObjectID) error {
	_, err := db.Collection("Invoices").DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

// Delete transporter from DB
func deleteTransporter(id primitive.ObjectID) error {
	_, err := db.Collection("Transporters").DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

// Delete transport vehicle from DB
func deleteTransport(id primitive.ObjectID) error {
	_, err := db.Collection("Transports").DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

// modify invoice in DB
func modifyInvoice(id primitive.ObjectID, ni Invoice) error {
	_, err := db.Collection("Invoices").UpdateOne(context.TODO(), bson.D{{"_id", id}}, bson.D{{"$set", ni}})
	return err
}

// modify transporter in DB
func modifyTransporter(id primitive.ObjectID, nt Transporter) error {
	_, err := db.Collection("Transporters").UpdateOne(context.TODO(), bson.D{{"_id", id}}, bson.D{{"$set", nt}})
	return err
}

// modify transport in DB
func modifyTransport(id primitive.ObjectID, nt Transport) error {
	_, err := db.Collection("Transports").UpdateOne(context.TODO(), bson.D{{"_id", id}}, bson.D{{"$set", nt}})
	return err
}

/* GetInvoices queries the database and
 * returns invoices based on the given filter
 * if filter is nil every invoice is returned
 */
func getInvoices(filter bson.M) ([]Invoice, error) {
	var invoices []Invoice

	cursor, err := db.Collection("Invoices").Find(context.TODO(), filter)
	if err != nil {
		return invoices, err
	}

	err = cursor.All(context.TODO(), &invoices)
	return invoices, err
}

func getTransporters(filter bson.M) ([]Transporter, error) {
	var transporters []Transporter

	cursor, err := db.Collection("Transporters").Find(context.TODO(), filter)
	if err != nil {
		return transporters, err
	}

	err = cursor.All(context.TODO(), &transporters)
	return transporters, err
}

func getTransports(filter bson.M) ([]Transport, error) {
	var transports []Transport

	cursor, err := db.Collection("Transports").Find(context.TODO(), filter)
	if err != nil {
		return transports, err
	}

	err = cursor.All(context.TODO(), &transports)
	return transports, err
}

func getInvoiceByNumber(invoiceNumber int) (Invoice, error) {
	var invoice Invoice
	err := db.Collection("Invoices").FindOne(context.TODO(), bson.M{"InvoiceNumber": invoiceNumber}).Decode(&invoice)
	return invoice, err
}

func getInvoiceById(invoiceId primitive.ObjectID) (Invoice, error) {
	var invoice Invoice
	err := db.Collection("Invoices").FindOne(context.TODO(), bson.M{"_id": invoiceId}).Decode(&invoice)
	return invoice, err
}
