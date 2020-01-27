package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/adm-github/godynu/config"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

const (
	URL = "https://api.dynu.com/v2/oauth2/token"
)

type Token struct {
	Access_Token string // `json:"accessToken"`
}

var versionCmd = &cobra.Command{
	Use:   "token",
	Short: "get the user API token from dynu",
	Run: func(cmd *cobra.Command, args []string) {
		access_Token := GetToken()
		fmt.Printf("Access token: %s\n", access_Token)
	},
}

func GetToken() string {
	v := config.InitializeConfig()
	user := fmt.Sprint(v.Get("clientID"))
	password := fmt.Sprint(v.Get("secret"))
	req, err := http.NewRequest("GET", URL, nil)
	req.SetBasicAuth(user, password)
	req.Header.Set("accept", "application/json")
	req.Header.Set("Accept-Language", "en_US")

	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		fmt.Println("Client error: ", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Reading body error: ", err)
	}

	var token = new(Token)
	err = json.Unmarshal(body, &token)
	if err != nil {
		fmt.Println("Unmarshal error", err)
	}
	return token.Access_Token
}
