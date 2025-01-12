package internal

import "fmt"

type Config struct {
	Key                string   `json:"KEY" yaml:"key"`
	Revision           string   `json:"REVISION" yaml:"revision"`
	ClientSecretFile   string   `json:"CLIENT_SECRET_FILE" yaml:"client_secret_file"`
	ServiceAccountFile string   `json:"SERVICE_ACCOUNT_FILE" yaml:"service_account_file"`
	Scopes             []string `json:"SCOPES" yaml:"scopes"`
	PropertyID         string   `json:"PROPERTY_ID" yaml:"property_id"`
	FetchToDate        string   `json:"FETCH_TO_DATE" yaml:"fetch_to_date"`
	ProjectId          string   `json:"PROJECT_ID" yaml:"project_id"`
	DatasetID          string   `json:"DATASET_ID" yaml:"dataset_id"`
	DatasetLocation    string   `json:"DATASET_LOCATION" yaml:"dataset_location"`
	TablePrefix        string   `json:"TABLE_PREFIX" yaml:"table_prefix"`
	PartitionBy        string   `json:"PARTITION_BY" yaml:"partition_by"`
	ClusterBy          string   `json:"CLUSTER_BY" yaml:"cluster_by"`
}

func (c Config) AllConfig() string {
	return fmt.Sprintf("%+v", c)
}
