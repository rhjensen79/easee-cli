/*
Copyright © 2022 Robert Jensen https://github.com/rhjensen79
*/
package cmd

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setcablelockCmd represents the setcablelock command
var setcablelockCmd = &cobra.Command{
	Use:   "setcablelock",
	Short: "Set Cable Lock",
	Long: `
setcablelock true
Set's the cable to be locked.
	
setcablelock false
Set's the cable to be unlocked`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] != "true" {
			args[0] = "false"
		}

		SetCableLock(viper.GetString("token"), viper.GetString("chargerid"), args[0])
	},
}

func init() {
	chargerCmd.AddCommand(setcablelockCmd)

}

func SetCableLock(accesstoken string, chargerid string, value string) {
	url := "https://api.easee.cloud/api/chargers/" + chargerid + "/commands/lock_state"

	payload := strings.NewReader("{\"state\":" + value + "}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Authorization", accesstoken)

	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode == 200 {
		log.Printf("Cablelock set to %v requested", value)
	}
	if res.StatusCode == 202 {
		log.Printf("Cablelock set to %v requested", value)
	} else {
		log.Fatal("Error : ", res.StatusCode)
	}

	defer res.Body.Close()
}
