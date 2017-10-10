// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

const (
		PINGURL = "https://api.dynu.com/v1/ping"
)

type Ping struct {
		Type string     // `json:"type"`
		Message string  // `json:"message"`
}

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "ping API call",
	Run: func(cmd *cobra.Command, args []string) {
		token := GetToken()
		tokenString := fmt.Sprintf("Bearer %s", token)

		req, err := http.NewRequest("GET", PINGURL, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", tokenString)

		cli := &http.Client{}
		resp, err := cli.Do(req)
		if err != nil {
				fmt.Println("Client error: ",err)
		}

		if resp.StatusCode == 200 {
			fmt.Println("Success")
		} else {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
					fmt.Println("Reading body error: ",err)
			}

			var ping = new(Ping)
			err = json.Unmarshal(body, &ping)
			if err != nil {
					fmt.Println("Unmarshal error", err)
			}

			fmt.Println("Status code: ", resp.StatusCode)
			fmt.Println("Type: ",    ping.Type)
			fmt.Println("Message: ", ping.Message)
		}
	},
}

func init() {
	RootCmd.AddCommand(pingCmd)
}
