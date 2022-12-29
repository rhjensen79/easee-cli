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

// pauseCmd represents the pause command
var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pause Charging",
	Long:  `Pause Charging`,
	Run: func(cmd *cobra.Command, args []string) {
		PauseCharging(viper.GetString("token"), viper.GetString("chargerid"))
	},
}

func init() {
	rootCmd.AddCommand(pauseCmd)

}

func PauseCharging(accesstoken string, chargerid string) {

	url := "https://api.easee.cloud/api/chargers/" + chargerid + "/commands/pause_charging"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("Authorization", accesstoken)

	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode == 200 {
		log.Println("Pause requested")
	}
	if res.StatusCode == 202 {
		log.Println("Pause requested")
	} else {
		log.Fatal("Error : ", res.StatusCode)
	}

	defer res.Body.Close()

}
