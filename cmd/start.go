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

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Charging",
	Long:  `Start Charging`,
	Run: func(cmd *cobra.Command, args []string) {
		StartCharging(viper.GetString("token"), viper.GetString("chargerid"))
	},
}

func init() {
	chargerCmd.AddCommand(startCmd)

}

func StartCharging(accesstoken string, chargerid string) {

	url := "https://api.easee.cloud/api/chargers/" + chargerid + "/commands/start_charging"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("Authorization", accesstoken)

	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode == 200 {
		log.Println("Start requested")
	}
	if res.StatusCode == 202 {
		log.Println("Start requested")
	} else {
		log.Fatal("Error : ", res.StatusCode)
	}

	defer res.Body.Close()

}
