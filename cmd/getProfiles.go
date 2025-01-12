/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// getProfilesCmd represents the getProfiles command
var getProfilesCmd = &cobra.Command{
	Use:     "getProfiles",
	Short:   "Klaviyo Profile 데이터를 Bigquery에  저장",
	Long:    `Klaviyo API 를 통해서 Profile 데이터를 조회하여 BigQuery에 저장하는 명령어`,
	PreRunE: newApp.SetConfig,
	RunE:    newApp.RunE,
}

func init() {
	rootCmd.AddCommand(getProfilesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getProfilesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getProfilesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
