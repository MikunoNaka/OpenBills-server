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
	"github.com/gin-gonic/gin"
	"github.com/MikunoNaka/OpenBills-lib/invoice"
	"log"
	"net/http"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Routes(route *gin.Engine) {
	i := route.Group("/invoice")
	{
		i.GET("/all", func(ctx *gin.Context) {
			// TODO: add functionality to filter results
			invoices, err := invoice.GetInvoices(nil)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to read invoices from DB: %v\n", err.Error())
				return
			}

			ctx.JSON(http.StatusOK, invoices)
		})

		i.DELETE("/:invoiceId", func(ctx *gin.Context) {
			id := ctx.Param("invoiceId")
			objectId, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to delete invoice, Error parsing ID: %v\n", err.Error())
				return
			}

			err = invoice.DeleteInvoice(objectId)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to delete invoice %v: %v\n", objectId, err.Error())
				return
			}

			log.Printf("Deleted invoice %v from database.\n", objectId )
			ctx.JSON(http.StatusOK, nil)
		})
	}

	transport := route.Group("/transport")
	{
		transport.GET("/all", func(ctx *gin.Context) {
			// TODO: add functionality to filter results
			transports, err := invoice.GetTransports(nil)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to read transport vehicles from DB: %v\n", err.Error())
				return
			}

			ctx.JSON(http.StatusOK, transports)
		})

		transport.DELETE("/:transportId", func(ctx *gin.Context) {
			id := ctx.Param("transportId")
			objectId, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to delete transport vehicle, Error parsing ID: %v\n", err.Error())
				return
			}

			err = invoice.DeleteTransport(objectId)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to delete transport vehicle %v: %v\n", objectId, err.Error())
				return
			}

			log.Printf("Deleted transport vehicle %v from database.\n", objectId )
			ctx.JSON(http.StatusOK, nil)
		})
	}

	transporter := route.Group("/transporter")
	{
		transporter.GET("/all", func(ctx *gin.Context) {
			// TODO: add functionality to filter results
			transporters, err := invoice.GetTransporters(nil)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to read transporters from DB: %v\n", err.Error())
				return
			}

			ctx.JSON(http.StatusOK, transporters)
		})

		transporter.DELETE("/:transporterId", func(ctx *gin.Context) {
			id := ctx.Param("transporterId")
			objectId, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to delete transporter, Error parsing ID: %v\n", err.Error())
				return
			}

			err = invoice.DeleteTransporter(objectId)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to delete transporter %v: %v\n", objectId, err.Error())
				return
			}

			log.Printf("Deleted transporter %v from database.\n", objectId )
			ctx.JSON(http.StatusOK, nil)
		})
	}
}
