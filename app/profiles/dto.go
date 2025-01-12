package profiles

import "go-klaviyo-to-bigquery/internal"

type ProfileData struct {
	Type          string                 `json:"type"`
	Id            string                 `json:"id"`
	Attributes    ProfileAttributes      `json:"attributes"`
	Links         internal.Links         `json:"links"`
	Relationships internal.Relationships `json:"relationships"`
}

type ProfileAttributes struct {
	Email               string              `json:"email"`
	PhoneNumber         string              `json:"phone_number"`
	ExternalId          string              `json:"external_id"`
	FirstName           string              `json:"first_name"`
	LastName            string              `json:"last_name"`
	Organization        string              `json:"organization"`
	Locale              string              `json:"locale"`
	Title               string              `json:"title"`
	Image               string              `json:"image"`
	Created             string              `json:"created"`
	Updated             string              `json:"updated"`
	LastEventDate       string              `json:"last_event_date"`
	Location            Location            `json:"location"`
	Properties          map[string]string   `json:"properties"`
	Subscriptions       Subscriptions       `json:"subscriptions"`
	PredictiveAnalytics PredictiveAnalytics `json:"predictive_analytics"`
}

type Location struct {
	Address1  string `json:"address1"`
	Address2  string `json:"address2"`
	City      string `json:"city"`
	Country   string `json:"country"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Region    string `json:"region"`
	Zip       string `json:"zip"`
	Timezone  string `json:"timezone"`
	Ip        string `json:"ip"`
}

type Subscriptions struct {
	Email      MarketingSubscription `json:"email"`
	Sms        MarketingSubscription `json:"sms"`
	MobilePush MarketingSubscription `json:"mobile_push"`
}

type MarketingSubscription struct {
	Marketing MarketingDetails `json:"marketing"`
}

type MarketingDetails struct {
	CanReceiveMarketing bool              `json:"can_receive_marketing"`
	Consent             string            `json:"consent"`
	ConsentTimestamp    string            `json:"consent_timestamp"`
	Method              string            `json:"method"`
	MethodDetail        string            `json:"method_detail"`
	CustomMethodDetail  string            `json:"custom_method_detail"`
	DoubleOptin         string            `json:"double_optin"`
	Suppression         []Suppression     `json:"suppression"`
	ListSuppressions    []ListSuppression `json:"list_suppressions"`
}

type Suppression struct {
	Reason    string `json:"reason"`
	Timestamp string `json:"timestamp"`
}

type ListSuppression struct {
	ListId    string `json:"list_id"`
	Reason    string `json:"reason"`
	Timestamp string `json:"timestamp"`
}

type PredictiveAnalytics struct {
	HistoricClv              float64 `json:"historic_clv"`
	PredictedClv             float64 `json:"predicted_clv"`
	TotalClv                 float64 `json:"total_clv"`
	HistoricNumberOfOrders   int     `json:"historic_number_of_orders"`
	PredictedNumberOfOrders  float64 `json:"predicted_number_of_orders"`
	AverageDaysBetweenOrders int     `json:"average_days_between_orders"`
	AverageOrderValue        float64 `json:"average_order_value"`
	ChurnProbability         float64 `json:"churn_probability"`
	ExpectedDateOfNextOrder  string  `json:"expected_date_of_next_order"`
}
