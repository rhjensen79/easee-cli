/*
Copyright © 2022 Robert Jensen https://github.com/rhjensen79
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var (
	// Used for flags.
	Username string

	rootCmd = &cobra.Command{
		Use:   "easee-cli",
		Short: "CLI for Easee charger using easee.cloud",
		Long:  `A CLI written in Go, for controlling a Easee charger, that is connected to easee.cloud`,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	ReadConfig()

}

func initConfig() {

}

func ReadConfig() {
	viper.SetConfigName("config")           // name of config file (without extension)
	viper.SetConfigType("yaml")             // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.easee-cli") // call multiple times to add many search paths
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			viper.SafeWriteConfig()
			log.Fatalln("No config file found - Creating one - Run easee-cli config to add data to it")

		} else {
			// Config file was found but another error was produced
			log.Fatalln("Error")
		}
	}
	viper.ReadInConfig()
}
