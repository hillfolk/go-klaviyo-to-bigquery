/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// getEventsCmd represents the getEvents command
var getEventsCmd = &cobra.Command{
	Use:     "getEvents",
	Short:   "Klaviyo 이벤트 데이터를 Bigquery에  저장",
	Long:    `Klaviyo API 를 통해서 이벤트 데이터를 조회하여 BigQuery에 저장하는 명령어`,
	PreRunE: newApp.SetConfig,
	RunE:    newApp.RunE,
}

func init() {
	rootCmd.AddCommand(getEventsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getEventsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//getEventsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
