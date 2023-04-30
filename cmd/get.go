package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

type Response struct {
	Healthy     string `json:"isHealthy"`
	Application string `json:"application"`
}

func init() {
	rootCmd.AddCommand(duCmd)
}

var duCmd = &cobra.Command{
	Use:   "get",
	Short: "Project get",
	Long:  `Project details get will be in github wiki`,
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		if len(args) >= 0 {
			fmt.Println("from cmd/get : ", args[0])
			Get(query)
		} else {
			fmt.Println("never get!")
		}
	},
}

func Get(query string) {
	fmt.Printf("QUERY: %s\n", query)
	fmt.Printf("Inside GET - API_URL: %s\n", config.API_URL)
	fmt.Printf("\n")
	apiUrl := config.API_URL + "?tls=" + query
	fmt.Println(apiUrl)
	fmt.Printf("\n")
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, apiUrl, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyString(body))
	// fmt.Println((body))

	// Response using struct
	var data Response
	// var tdata []Response
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("\n")
	fmt.Println("Healthy: ", data.Healthy)

	// Write response to Table
	// table := tablewriter.NewWriter(os.Stdout)
	// table.SetHeader([]string{"Is Healthy", "Name", "Application"})
	// for _, d := range tdata {
	// 	table.Append([]string{
	// 		fmt.Println(d.isHealthy),
	// 		// d.application,
	// 	})
	// }
	// table.Render()
}
