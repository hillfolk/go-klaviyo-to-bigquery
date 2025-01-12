package internal

type ResponseData struct {
	Data     []Data     `json:"data"`
	Links    Links      `json:"links"`
	Included []Included `json:"included"`
}

type Data struct {
	Type          string        `json:"type"`
	Id            string        `json:"id"`
	Attributes    any           `json:"attributes"`
	Links         Links         `json:"links"`
	Relationships Relationships `json:"relationships"`
}

type Links struct {
	Self  string `json:"self"`
	First string `json:"first"`
	Last  string `json:"last"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
}

type Relationships struct {
	Profile      RelationshipData `json:"profile"`
	Metric       RelationshipData `json:"metric"`
	Attributions AttributionData  `json:"attributions"`
}

type RelationshipData struct {
	Data  RelationshipDetails `json:"data"`
	Links Links               `json:"links"`
}

type RelationshipDetails struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type AttributionData struct {
	Data  []RelationshipDetails `json:"data"`
	Links Links                 `json:"links"`
}

type Included struct {
	Type          string        `json:"type"`
	Id            string        `json:"id"`
	Attributes    any           `json:"attributes,omitempty"`
	Relationships Relationships `json:"relationships,omitempty"`
	Links         Links         `json:"links"`
}
