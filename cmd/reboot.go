/*
Copyright © 2022 Robert Jensen https://github.com/rhjensen79
*/
package cmd

import (
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rebootCmd represents the reboot command
var rebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		RebootCharger(viper.GetString("token"), viper.GetString("chargerid"))
	},
}

func init() {
	chargerCmd.AddCommand(rebootCmd)

}

func RebootCharger(accesstoken string, chargerid string) {

	url := "https://api.easee.cloud/api/chargers/" + chargerid + "/commands/reboot"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("Authorization", accesstoken)

	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode == 200 {
		log.Println("Charger rebooting")
	}
	if res.StatusCode == 202 {
		log.Println("Charger rebooting")
	} else {
		log.Fatal("Error : ", res.StatusCode)
	}

	defer res.Body.Close()

}
