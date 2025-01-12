package metrics

type MetricAttributes struct {
	Name        string `json:"name"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
	Integration any    `json:"integration"`
}
