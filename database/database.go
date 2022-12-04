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

package database

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

/* This creates a new client and sets
 * it to the value of MongoClient which
 * can be imported by other packages
 *
 * I am not at all sure if this is the best
 * actually no, even if this is a decent way to do it.
 * But yea this seems to work
 * (remember to close the connections!)
 */
var DB *mongo.Database
var client *mongo.Client

func init() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

	  var err error
    client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
    if err != nil {
        log.Fatal(err)
    }

	  log.Println("Successfully connected to MongoDB database.")
    DB = client.Database("OpenBillsDB")
}

func DisconnectDB() {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully closed connection with MongoDB.")
}
