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

// updatefirmwareCmd represents the updatefirmware command
var updatefirmwareCmd = &cobra.Command{
	Use:   "updatefirmware",
	Short: "Updates the firmware of the charger, to latest version",
	Long:  `Updates the firmware of the charger, to latest version`,
	Run: func(cmd *cobra.Command, args []string) {
		UpdateFirmware(viper.GetString("token"), viper.GetString("chargerid"))
	},
}

func init() {
	chargerCmd.AddCommand(updatefirmwareCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updatefirmwareCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updatefirmwareCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func UpdateFirmware(accesstoken string, chargerid string) {

	url := "https://api.easee.cloud/api/chargers/" + chargerid + "/commands/update_firmware"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("Authorization", accesstoken)

	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode == 200 {
		log.Println("Update requested")
	}
	if res.StatusCode == 202 {
		log.Println("Update requested")
	} else {
		log.Fatal("Error : ", res.StatusCode)
	}

	defer res.Body.Close()

}
