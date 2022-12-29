/*
Copyright © 2022 Robert Jensen https://github.com/rhjensen79
*/
package cmd

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Get a token valid for 24 hours",
	Long: `Use login, to get a new token, valid for 24 hours.
The token will be saves in the config file, and used for all api calls`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("token", Auth(viper.GetString("username"), viper.GetString("password")))
		viper.WriteConfig()

	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func Auth(username string, password string) (accesstoken string) {
	type ResponseData struct {
		AccessToken string `json:"accessToken"`
	}

	url := "https://api.easee.cloud/api/accounts/login"

	payload := strings.NewReader("{\"password\":\"" + password + "\",\"userName\":\"" + username + "\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/*+json")

	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode == 200 {
		log.Println("Access Token recieved")
	} else {
		log.Fatal("Error Getting Access Token code : ", res.StatusCode)
	}

	defer res.Body.Close()

	var responseData ResponseData
	json.NewDecoder(res.Body).Decode(&responseData)
	//accesstoken = responseData.AccessToken

	return "Bearer " + responseData.AccessToken

}
