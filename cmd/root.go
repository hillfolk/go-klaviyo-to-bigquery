/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"go-klaviyo-to-bigquery/app"
)

var (
	newApp   = app.NewApp()
	config   string
	revision string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "all",
	Short: "Klaviyo API 를 통해서 이벤트 Profile Metrics Event 조회하여 BigQuery에 저장합니다.",
	Long:  `Klaviyo API 를 통해서 이벤트 Profile Metrics Event 조회하여 BigQuery에 저장합니다.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	PreRunE: newApp.SetConfig,
	RunE:    newApp.RunE,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().StringVar(&config, "config", "", "config file (default is config.json)")
	rootCmd.PersistentFlags().StringVar(&revision, "revision", "", "Api revision")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
