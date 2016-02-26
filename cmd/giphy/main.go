package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/nickvanw/gifgo"
)

const (
	apiKeyFlag = "apikey"
)

func main() {
	app := cli.NewApp()
	app.Name = "giphy"
	app.Usage = "Get GIFs from GIPHY!"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   apiKeyFlag,
			Value:  "",
			Usage:  "API Key to use when calling GIPHY",
			EnvVar: "GIPHY_KEY",
		},
	}
	app.Commands = []cli.Command{
		cli.Command{
			Name:    "search",
			Aliases: []string{"s"},
			Usage:   "Search giphy",
			Action:  searchGIF,
		},
	}
	app.Run(os.Args)

}

func searchGIF(c *cli.Context) {
	client := newClient(c)
	query := strings.Join(c.Args(), " ")
	req := gifgo.SearchReq{Query: query}
	results, err := client.Search(req)
	if err != nil {
		fmt.Printf("Search Failed: %s\n", err)
		os.Exit(2)
	}
	fmt.Printf("%d Results from GIPHY\nListing the first %d\n",
		results.Pagination.TotalCount, len(results.Data))
	for k, v := range results.Data {
		fmt.Printf("%d: %s; rating: %s\n", k+1, v.URL, v.Rating)
	}
}

func newClient(c *cli.Context) *gifgo.Client {
	opts := []gifgo.OptFunc{}
	if apiKey := c.String(apiKeyFlag); apiKey != "" {
		opts = append(opts, gifgo.APIKey(apiKey))
	}
	client, err := gifgo.New(opts...)
	if err != nil {
		fmt.Printf("Unable to create client: %s\n", err)
		os.Exit(1)
	}
	return client
}
