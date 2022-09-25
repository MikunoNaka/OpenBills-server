package main

import (
	"github.com/MikunoNaka/OpenBills-server/brand"
	"github.com/MikunoNaka/OpenBills-server/item"
	"github.com/MikunoNaka/OpenBills-server/client"
	"github.com/MikunoNaka/OpenBills-server/invoice"
	"github.com/MikunoNaka/OpenBills-lib/database"

	"github.com/gin-gonic/gin"
)

func main() {
	defer database.DisconnectDB()
	r := gin.New()

	item.Routes(r)
	brand.Routes(r)
	client.Routes(r)
	invoice.Routes(r)

	r.Run(":6969")
}
