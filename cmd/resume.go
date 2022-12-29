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

// resumeCmd represents the resume command
var resumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ResumeCharging(viper.GetString("token"), viper.GetString("chargerid"))
	},
}

func init() {
	rootCmd.AddCommand(resumeCmd)

}

func ResumeCharging(accesstoken string, chargerid string) {

	url := "https://api.easee.cloud/api/chargers/" + chargerid + "/commands/resume_charging"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("Authorization", accesstoken)

	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode == 200 {
		log.Println("Resume requested")
	}
	if res.StatusCode == 202 {
		log.Println("Resume requested")
	} else {
		log.Fatal("Error : ", res.StatusCode)
	}

	defer res.Body.Close()

}
