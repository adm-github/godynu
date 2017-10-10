package cmd

import (
	"fmt"
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/adm-github/godynu/config"
	"net/http"
	"io/ioutil"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

const (
		URL = "https://api.dynu.com/v1/oauth2/token"
)

type Token struct {
		//Scope string     // `json:"scope"`
		AccessToken string // `json:"accessToken"`
		//TokenType string // `json:"tokenType"`
		//Roles []string   // `json:"roles"`
		//ExpiresIn int    // `json:"expiresIn"`
}

var versionCmd = &cobra.Command{
	Use:   "token",
	Short: "get the user API token from dynu",
	Run: func(cmd *cobra.Command, args []string) {
		accessToken := GetToken()
		fmt.Printf("Access token: %s\n", accessToken)
	},
}

func GetToken() string {
		v := config.InitializeConfig()
		user := fmt.Sprint(v.Get("clientID"))
		password := fmt.Sprint(v.Get("secret"))

		req, err := http.NewRequest("GET", URL, nil)
		req.SetBasicAuth(user, password)
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Accept-Language", "en_US")

		cli := &http.Client{}
		resp, err := cli.Do(req)
		if err != nil {
				fmt.Println("Client error: ",err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
				fmt.Println("Reading body error: ",err)
		}

		var token = new(Token)
		err = json.Unmarshal(body, &token)
		if err != nil {
				fmt.Println("Unmarshal error", err)
		}
		return token.AccessToken
}
