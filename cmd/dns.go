// Copyright © 2019 Antonio Di Marco
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
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
)

const (
	DNSAPI = "https://api.dynu.com/v2/dns/"
)

var f_domain string
var f_id string
var f_ipv4 string
var f_all bool

func httpCall(url string, httpVerb string, jsonStr string) {
	req, _ := http.NewRequest(httpVerb, url, bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", GetToken()))
	req.Header.Set("accept", "application/json")

	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		fmt.Println("Client error:", err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Reading body error:", err)
	}

	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, respBody, "", "\t")
	if error != nil {
		fmt.Println(string(respBody))
		fmt.Println("JSON parse error:", error)
		return
	}

	fmt.Println(string(prettyJSON.Bytes()))
}

// dnsCmd represents the dns command
var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "Work with dynu dynamic dns",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new dns record and IP",
	Run: func(cmd *cobra.Command, args []string) {
		if f_ipv4 == "" {
			cmd.Help()
			return
		}

		body := fmt.Sprintf(`{"name":"%s", "ipv4Address": "%s"}`, f_domain, f_ipv4)
		httpCall(DNSAPI, "POST", body)
	},
}

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing dns record and IP",
	Run: func(cmd *cobra.Command, args []string) {
		if f_domain == "" || f_ipv4 == "" || f_id == "" {
			cmd.Help()
			return
		}

		jsonStr := fmt.Sprintf(`{"name":"%s", "ipv4Address": "%s"}`, f_domain, f_ipv4)
		infoURL := fmt.Sprint(DNSAPI, f_id)
		httpCall(infoURL, "POST", jsonStr)
	},
}

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a dns record and its IPs",
	Run: func(cmd *cobra.Command, args []string) {
		if f_id == "" {
			cmd.Help()
			return
		}

		deleteURL := fmt.Sprint(DNSAPI, f_id)
		httpCall(deleteURL, "DELETE", "")
	},
}

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Info about dns records and IPs",
	Run: func(cmd *cobra.Command, args []string) {
		var infoURL string
		if f_domain != "" {
			infoURL = fmt.Sprint(DNSAPI, f_domain)
		} else if f_all {
			infoURL = DNSAPI
		} else {
			cmd.Help()
			return
		}
		httpCall(infoURL, "GET", "")
	},
}

func init() {
	RootCmd.AddCommand(dnsCmd)

	dnsCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&f_domain, "domain", "d", "", "Dns record to add")
	addCmd.Flags().StringVarP(&f_ipv4, "ip", "", "", "IPV4 to add as resolution for the domain provided")

	dnsCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&f_id, "id", "i", "", "Dns id to update")
	updateCmd.Flags().StringVarP(&f_ipv4, "ip", "", "", "IPV4 to add as resolution for the id of the domain provided")
	updateCmd.Flags().StringVarP(&f_domain, "domain", "d", "", "Dns record to update")

	dnsCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&f_id, "id", "i", "", "Record ID to delete")

	dnsCmd.AddCommand(infoCmd)
	infoCmd.Flags().StringVarP(&f_domain, "domain", "d", "", "Domain to get DNS info for")
	infoCmd.Flags().BoolVarP(&f_all, "all", "a", false, "Get DNS info on all records")
}

