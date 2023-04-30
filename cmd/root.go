package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	HEALTH_URL string `json:"health_url"`
	API_URL    string `json:"api_url"`
	APIKEY     string `json:"api_key"`
	APISECRET  string `json:"api_secret"`
}

var (
	config      Config
	cfgFile     string
	userLicense string
	Verbose     bool
	HEALTH_URL  string
	// build       bool
	// author string
)

func init() {
	cobra.OnInitialize(initConfig)
	// Load config file
	data, err := os.ReadFile("sre-config.json")
	if err != nil {
		panic(err)
	}

	// Unmarshal config data into struct
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "Jayanta Paul", "Jayanta Paul")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	// viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	// viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "Jayanta Paul <jayantapaul.jp18@gmail.com>")
	viper.SetDefault("license", "apache")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sre",
	Short: "SRE Tools",
	Long:  `Tools for SRE Team`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// home, err := os.UserHomeDir()
		// cobra.CheckErr(err)
		// viper.WriteConfigAs(home + "/sre-config.yaml")
		// Use config values
		fmt.Printf("###################################### \n")
		fmt.Printf("API_URL: %s\n", config.API_URL)
		fmt.Printf("API key: %s\n", config.APIKEY)
		fmt.Printf("API secret: %s\n", config.APISECRET)
		fmt.Printf("###################################### \n")
		// Command Conditions
		if len(args) == 0 {
			// log.Fatal("Please provide input command! type -h for help")
			color.Cyan("Please provide input command! type -h for help")
		}
		if args[0] == "apm" || args[0] == "APM" {
			fmt.Println("REST API Call to APM")
			Apm()
		} else if args[0] == "du" || args[0] == "DU" {
			fmt.Println("REST API Call to DU")
		} else if args[0] == "health" || args[0] == "Health" {
			Health()
		}
		// } else {
		// 	color.Yellow("sorry not valid command! here some example command")
		// 	color.Cyan("sre health" + " / " + "sre apm" + " / " + "sre du" + " / " + "sre -h")
		// }
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		// color.Red("I'm crashing !")
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

//	func init() {
//		// rootCmd.Flags().StringVarP(&dotEnvPath, "envfilepath", "e", "", "Path to your environment file. if left empty, the program will only use system environment variables.")
//	}
func PrettyString(str []byte) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, str, "", "  "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func Health() {
	fmt.Printf("Inside HEALTH_URL: %s\n", config.HEALTH_URL)
	url := config.HEALTH_URL
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

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
}

func Apm() {
	fmt.Printf("Inside Apm - HEALTH_URL: %s\n", config.HEALTH_URL)
	url := config.HEALTH_URL
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

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
}
