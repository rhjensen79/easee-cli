/*
Copyright © 2022 Robert Jensen https://github.com/rhjensen79
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Create the config file Easee-cli",
	Long:  `Create the config file Easee-cli`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
		set_config()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

}

func set_config() {
	var answer string
	var username string
	var password string
	var chargerid string

	fmt.Print("Running this, will overwrite all existing data.\n\nPlease confirm y/n : ")

	fmt.Scanln(&answer)
	if answer == "y" {
		fmt.Println("Press enter to skip a question")
		fmt.Printf("Current Username 	: %v\n", viper.GetString("username"))
		fmt.Print("New Username 		: ")
		fmt.Scanln(&username)
		if username != "" {
			viper.Set("username", username)
		}
		fmt.Print("Password 		: ")
		fmt.Scanln(&password)
		if password != "" {
			viper.Set("password", password)
		}
		fmt.Printf("Current ChargerID 	: %v\n", viper.GetString("chargerid"))
		fmt.Print("New ChargerID 		: ")
		fmt.Scanln(&chargerid)
		if chargerid != "" {
			viper.Set("chargerid", chargerid)
		}
		viper.WriteConfig()
		log.Println("New config saved")
		log.Println("Run easee-cli login to get a token")
	} else {
		os.Exit(0)
		log.Println("Exiting")
	}

}
