/*
Copyright © 2022 Robert Jensen https://github.com/rhjensen79
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// chargerCmd represents the charger command
var chargerCmd = &cobra.Command{
	Use:   "charger",
	Short: "The commands assosiated with the charger",
	Long:  `The commands assosiated with the charger`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(chargerCmd)

}
