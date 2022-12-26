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

package main

import (
	_ "github.com/MikunoNaka/OpenBills-server/util"
	"github.com/MikunoNaka/OpenBills-server/brand"
	"github.com/MikunoNaka/OpenBills-server/client"
	"github.com/MikunoNaka/OpenBills-server/database"
	"github.com/MikunoNaka/OpenBills-server/invoice"
	"github.com/MikunoNaka/OpenBills-server/item"
	"github.com/MikunoNaka/OpenBills-server/user"
	"github.com/MikunoNaka/OpenBills-server/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	defer database.DisconnectDB()
	r := gin.New()

	item.Routes(r)
	brand.Routes(r)
	client.Routes(r)
	invoice.Routes(r)
	user.Routes(r)
	auth.Routes(r)

	// ping server and check if logged in
	r.POST("/ping", auth.Authorize(), func (ctx *gin.Context) {
		ctx.Status(200)
	})

	r.Run(":6969")
}
