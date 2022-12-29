/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getchargeridCmd represents the getchargerid command
var getchargeridCmd = &cobra.Command{
	Use:   "getchargerid",
	Short: "Get ChargerID of your charger",
	Long:  `Get ChargerID of your charger`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("chargerid", GetChargerId(viper.GetString("token")))
	},
}

func init() {
	configCmd.AddCommand(getchargeridCmd)

}

func GetChargerId(accesstoken string) (chargerid string) {
	type ResponseData struct {
		ID string `json:"id"`
	}

	url := "https://api.easee.cloud/api/chargers"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", accesstoken)

	res, _ := http.DefaultClient.Do(req)

	if res.StatusCode == 200 {
		log.Println("ChargerID Token recieved")
	} else {
		log.Fatal("Error Getting ChargerID code : ", res.StatusCode)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	st := string(body)
	st = strings.Trim(st, "[]")

	var responseData ResponseData
	json.Unmarshal([]byte(st), &responseData)

	return responseData.ID

}
