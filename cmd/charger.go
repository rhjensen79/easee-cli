/*
Copyright © 2022 Robert Jensen https://github.com/rhjensen79
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// chargerCmd represents the charger command
var chargerCmd = &cobra.Command{
	Use:   "charger",
	Short: "The commands assosiated with the charger",
	Long:  `The commands assosiated with the charger`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(chargerCmd)

}
