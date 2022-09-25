package invoice

import (
	"github.com/gin-gonic/gin"
	"github.com/MikunoNaka/OpenBills-lib/invoice"
	"log"
	"net/http"
)

func Routes(route *gin.Engine) {
	i := route.Group("/invoice")
	{
		i.GET("/", func(ctx *gin.Context) {
			// TODO: add functionality to filter results
			invoices, err := invoice.GetInvoices(nil)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to read invoices from DB: %v\n", err.Error())
			}

			ctx.JSON(http.StatusOK, invoices)
		})

		i.POST("/", func(ctx *gin.Context) {
			var x invoice.Invoice
			ctx.Bind(&x)
			err := x.Save()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to generate new invoice #%d: %v\n", x.InvoiceNumber, err.Error())
			}

			log.Printf("Generated new invoice #%d.\n", x.InvoiceNumber)
			ctx.JSON(http.StatusOK, nil)
		})

		i.DELETE("/", func(ctx *gin.Context) {
			var x invoice.Invoice
			ctx.Bind(&x)
			err := x.Delete()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to delete invoice #%d: %v\n", x.InvoiceNumber, err.Error())
			}

			log.Printf("Deleted invoice invoice #%d.\n", x.InvoiceNumber)
			ctx.JSON(http.StatusOK, nil)
		})
	}

	transport := route.Group("/transport")
	{
		transport.GET("/", func(ctx *gin.Context) {
			// TODO: add functionality to filter results
			transports, err := invoice.GetTransports(nil)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to read transport vehicles from DB: %v\n", err.Error())
			}

			ctx.JSON(http.StatusOK, transports)
		})

		transport.POST("/", func(ctx *gin.Context) {
			var x invoice.Transport
			ctx.Bind(&x)
			err := x.Save()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to add transport vehicle \"%s\": %v\n", x.VehicleNum, err.Error())
			}

			log.Printf("Added new transport vehicle to database: \"%s\"\n", x.VehicleNum)
			ctx.JSON(http.StatusOK, nil)
		})

		transport.DELETE("/", func(ctx *gin.Context) {
			var x invoice.Transport
			ctx.Bind(&x)
			err := x.Delete()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to delete transport vehicle \"%s\": %v\n", x.VehicleNum, err.Error())
			}

			log.Printf("Deleted transport vehicle: \"%s\"\n", x.VehicleNum)
			ctx.JSON(http.StatusOK, nil)
		})
	}

	transporter := route.Group("/transporter")
	{
		transporter.GET("/", func(ctx *gin.Context) {
			// TODO: add functionality to filter results
			transporters, err := invoice.GetTransporters(nil)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to read transporters from DB: %v\n", err.Error())
			}

			ctx.JSON(http.StatusOK, transporters)
		})

		transporter.POST("/", func(ctx *gin.Context) {
			var x invoice.Transporter
			ctx.Bind(&x)
			err := x.Save()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to add transporter \"%s\": %v\n", x.Name, err.Error())
			}

			log.Printf("Added new transporter to database: \"%s\"\n", x.Name)
			ctx.JSON(http.StatusOK, nil)
		})

		transporter.DELETE("/", func(ctx *gin.Context) {
			var x invoice.Transporter
			ctx.Bind(&x)
			err := x.Delete()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to delete transporter \"%s\": %v\n", x.Name, err.Error())
			}

			log.Printf("Deleted transporter: \"%s\"\n", x.Name)
			ctx.JSON(http.StatusOK, nil)
		})
	}
}
