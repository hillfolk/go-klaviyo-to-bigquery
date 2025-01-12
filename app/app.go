package app

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/bigquery"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"go-klaviyo-to-bigquery/app/bq"
	"go-klaviyo-to-bigquery/app/client"
	"go-klaviyo-to-bigquery/app/events"
	"go-klaviyo-to-bigquery/app/metrics"
	"go-klaviyo-to-bigquery/app/profiles"
	"go-klaviyo-to-bigquery/internal"
)

func (a *App) SetConfig(cmd *cobra.Command, args []string) error {
	var err error

	cfgFile, err := cmd.Flags().GetString("config")

	if err != nil {
		return errors.Wrap(err, "failed to get config file")
	}

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.Getwd()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("json")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".local")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		a.cfg = &internal.Config{
			Key:                viper.GetString("KEY"),
			ClientSecretFile:   viper.GetString("CLIENT_SECRET_FILE"),
			ServiceAccountFile: viper.GetString("SERVICE_ACCOUNT_FILE"),
			Scopes:             viper.GetStringSlice("SCOPES"),
			PropertyID:         viper.GetString("PROPERTY_ID"),
			ProjectId:          viper.GetString("PROJECT_ID"),
			FetchToDate:        viper.GetString("FETCH_TO_DATE"),
			DatasetID:          viper.GetString("DATASET_ID"),
			TablePrefix:        viper.GetString("TABLE_PREFIX"),
		}
		fmt.Println(a.cfg.AllConfig())
	} else {
		return errors.Wrap(err, "failed to read config")
	}

	return nil
}

type App struct {
	key    string
	client *client.Client
	cfg    *internal.Config
}

func NewApp() *App {
	return &App{}
}

// Run runs the Ga4DataFetcher
func (a *App) RunE(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	fmt.Println("Running go-klaviyo-to-bigquery")
	newClient := client.NewClient(a.cfg.Key)
	newBqClient, err := bigquery.NewClient(ctx, a.cfg.ProjectId)
	if err != nil {
		return errors.Wrap(err, "failed to create BigQuery client")
	}

	newDataInsert := bq.NewBigQueryDataInsert(newBqClient)

	// 상위 Command는 Google Analytics Data API를 이용하여 데이터를 조회 하는 방식을 결정합니다.
	switch cmd.Use {
	case "getEvents":
		return events.NewHandler(newClient, newDataInsert, a.cfg).Handle(ctx)
	case "getProfiles":
		return profiles.NewHandler(newClient, newDataInsert, a.cfg).Handle(ctx)
	case "getMetrics":
		return metrics.NewHandler(newClient, newDataInsert, a.cfg).Handle(ctx)
	case "all":
		if err := events.NewHandler(newClient, newDataInsert, a.cfg).Handle(ctx); err != nil {
			return errors.Wrap(err, "failed to get events")
		}
		fmt.Println("Events done")
		if err := profiles.NewHandler(newClient, newDataInsert, a.cfg).Handle(ctx); err != nil {
			return errors.Wrap(err, "failed to get profiles")
		}
		fmt.Println("Profiles done")
		if err := metrics.NewHandler(newClient, newDataInsert, a.cfg).Handle(ctx); err != nil {
			return errors.Wrap(err, "failed to get metrics")
		}
		fmt.Println("Metrics done")
	default:
		errors.New("invalid command")
	}
	fmt.Println("Stop go-klaviyo-to-bigquery")
	return nil
}
