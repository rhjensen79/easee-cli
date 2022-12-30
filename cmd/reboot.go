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
	Short: "Reboot's charger",
	Long:  `Reboot's charger`,
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
