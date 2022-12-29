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

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop Charging",
	Long:  `Stop Charging`,
	Run: func(cmd *cobra.Command, args []string) {
		StopCharging(viper.GetString("token"), viper.GetString("chargerid"))
	},
}

func init() {
	chargerCmd.AddCommand(stopCmd)

}

func StopCharging(accesstoken string, chargerid string) {

	url := "https://api.easee.cloud/api/chargers/" + chargerid + "/commands/stop_charging"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("Authorization", accesstoken)

	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode == 200 {
		log.Println("Stop requested")
	}
	if res.StatusCode == 202 {
		log.Println("Stop requested")
	} else {
		log.Fatal("Error : ", res.StatusCode)
	}

	defer res.Body.Close()

}
