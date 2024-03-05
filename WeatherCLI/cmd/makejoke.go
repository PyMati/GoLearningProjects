/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

const JokeEndpoint = "https://official-joke-api.appspot.com/random_joke"

type Joke struct {
	Type      string `json:type`
	Setup     string `json:setup`
	Punchline string `json:punchline`
	Id        int    `json:id`
}

// getjokeCmd represents the getjoke command
var getjokeCmd = &cobra.Command{
	Use:   "makejoke",
	Short: "Prints joke from joke api.",
	Long:  `Makes an http request to joke api in order to get random joke and prints it to console.`,
	Run: func(cmd *cobra.Command, args []string) {
		resp := makeApiCall()
		data := readDataFromResponse(resp)

		var parsedJoke Joke
		if err := json.Unmarshal(data, &parsedJoke); err != nil {
			fmt.Errorf("An error occured while parsing joke to json.")
		}

		fmt.Println(parsedJoke.Setup)
		fmt.Println(parsedJoke.Punchline)
	},
}

func init() {
	rootCmd.AddCommand(getjokeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getjokeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getjokeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func makeApiCall() *http.Response {
	resp, err := http.Get(JokeEndpoint)
	if err != nil {
		fmt.Errorf("An error occured while fetching joke from an api.")
	}

	return resp
}

func readDataFromResponse(resp *http.Response) []byte {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("An error occured while reading data from body request.")
	}

	return body
}
