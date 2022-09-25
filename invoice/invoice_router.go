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
		i.POST("/", func(ctx *gin.Context) {
			var x invoice.Invoice
			ctx.Bind(&x)
			err := x.Save()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to generate new invoice #%d: %v", x.InvoiceNumber, err.Error())
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
				log.Printf("ERROR: Failed to delete invoice #%d: %v", x.InvoiceNumber, err.Error())
			}

			log.Printf("Deleted invoice invoice #%d.\n", x.InvoiceNumber)
			ctx.JSON(http.StatusOK, nil)
		})
	}

	transport := route.Group("/transport")
	{
		transport.POST("/", func(ctx *gin.Context) {
			var x invoice.Transport
			ctx.Bind(&x)
			err := x.Save()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to add transport vehicle \"%s\": %v", x.VehicleNum, err.Error())
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
				log.Printf("ERROR: Failed to delete transport vehicle \"%s\": %v", x.VehicleNum, err.Error())
			}

			log.Printf("Deleted transport vehicle: \"%s\"\n", x.VehicleNum)
			ctx.JSON(http.StatusOK, nil)
		})
	}

	transporter := route.Group("/transporter")
	{
		transporter.POST("/", func(ctx *gin.Context) {
			var x invoice.Transporter
			ctx.Bind(&x)
			err := x.Save()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Printf("ERROR: Failed to add transporter \"%s\": %v", x.Name, err.Error())
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
				log.Printf("ERROR: Failed to delete transporter \"%s\": %v", x.Name, err.Error())
			}

			log.Printf("Deleted transporter: \"%s\"\n", x.Name)
			ctx.JSON(http.StatusOK, nil)
		})
	}
}
