package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
)

type Country struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	Currencies map[string]struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
}

func fetchCountries(url string) ([]Country, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var countries []Country
	if err := json.NewDecoder(resp.Body).Decode(&countries); err != nil {
		return nil, err
	}

	return countries, nil
}

func sortCountries(countries []Country, sortBy string, ascending bool) {
	sort.Slice(countries, func(i, j int) bool {
		if sortBy == "name" {
			if ascending {
				return countries[i].Name.Common < countries[j].Name.Common
			}
			return countries[i].Name.Common > countries[j].Name.Common
		}
		if len(countries[i].Currencies) > 0 && len(countries[j].Currencies) > 0 {
			var currencyI, currencyJ string
			for _, v := range countries[i].Currencies {
				currencyI = v.Name
				break
			}
			for _, v := range countries[j].Currencies {
				currencyJ = v.Name
				break
			}
			if ascending {
				return currencyI < currencyJ
			}
			return currencyI > currencyJ
		}
		return false
	})
}

func printCountries(countries []Country) {
	midpoint := len(countries) / 2
	if len(countries)%2 != 0 { // Make odd n of countries work
		midpoint++
	}

	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Country", "Currencies", "Country", "Currencies"}) // Double the columns for side-by-side display

	table.SetBorder(true)
	table.SetColumnSeparator("|")
	table.SetRowSeparator("-")
	table.SetRowLine(true)
	table.SetHeaderLine(true)

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor, tablewriter.BgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor, tablewriter.BgWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor, tablewriter.BgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor, tablewriter.BgWhiteColor})

	table.SetColumnColor(tablewriter.Colors{tablewriter.Normal},
		tablewriter.Colors{tablewriter.FgHiGreenColor},
		tablewriter.Colors{tablewriter.Normal},
		tablewriter.Colors{tablewriter.FgHiGreenColor})

	for i := 0; i < midpoint; i++ {
		var row []string
		row = append(row, countries[i].Name.Common, formatCurrencies(countries[i].Currencies))

		if i+midpoint < len(countries) {
			row = append(row, countries[i+midpoint].Name.Common, formatCurrencies(countries[i+midpoint].Currencies))
		} else {
			row = append(row, "", "")
		}
		table.Append(row)
	}

	table.Render()
}

func formatCurrencies(currencies map[string]struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}) string {
	var result []string
	for _, c := range currencies {
		result = append(result, fmt.Sprintf("%s (%s)", c.Name, c.Symbol))
	}
	return strings.Join(result, "\n")
}

func createCommand(region string) *cli.Command {
	return &cli.Command{
		Name:  region,
		Usage: fmt.Sprintf("Fetch and sort countries in %s by name in ascending or descending order", region),
		Action: func(c *cli.Context) error {
			var url string
			if region == "world" {
				url = "https://restcountries.com/v3.1/all"
			} else {
				url = fmt.Sprintf("https://restcountries.com/v3.1/region/%s", region)
			}

			fmt.Println("Fetching data from", url)
			countries, err := fetchCountries(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error fetching data: %v\n", err)
				return err
			}

			// Initially sort by country name in as default behaviour ascending order
			sortCountries(countries, "name", true)
			printCountries(countries)

			for {
				fmt.Println("Press 'a' for ascending, 'd' for descending order, and 'x' to exit:")
				var input string
				fmt.Scanln(&input)
				switch input {
				case "a":
					sortCountries(countries, "name", true)
					printCountries(countries)
				case "d":
					sortCountries(countries, "name", false)
					printCountries(countries)
				case "x":
					return nil
				default:
					fmt.Println("Invalid input. Please press 'a', 'd', or 'x'.")
				}
			}
		},
	}
}

func fetchAndSortCountriesCommands(regions []string) []*cli.Command {
	commands := []*cli.Command{}

	for _, region := range regions {
		commands = append(commands, createCommand(region))
	}

	return commands
}

func main() {
	app := &cli.App{
		Name:     "Country Info CLI",
		Usage:    "Fetch and display countries information",
		Commands: fetchAndSortCountriesCommands([]string{"europe", "world"}),
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Application error: %v\n", err)
		os.Exit(1)
	}
}
