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

package util

import (
	"github.com/spf13/viper"
	"os"
	"log"
)

type Config struct {
	Crypto struct {
		AccessTokenSecret string `mapstructure:"access_token_secret"`
		RefreshTokenSecret string `mapstructure:"refresh_token_secret"`
	} `mapstructure:"cryptography"`
}

func init() {
	viper.SetConfigName("openbills")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// config file not found
			log.Println("Cannot find openbills config file in pwd or /etc")
		} else {
			// config file found but has errors
			log.Printf("Error while reading config file: %v\n", err)
			log.Println("Cannot start OpenBills Server")
		}
		os.Exit(1)
	}
}

func GetConfig() Config {
	var conf Config
	viper.Unmarshal(&conf)
	return conf
}
