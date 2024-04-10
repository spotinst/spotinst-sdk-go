package oceancd

import (
	"time"
)

type VerificationProvider struct {
	CloudWatch *CloudWatch `json:"cloudWatch,omitempty"`
	ClusterIDs []string    `json:"clusterIds,omitempty"`
	DataDog    *DataDog    `json:"datadog,omitempty"`
	Jenkins    *Jenkins    `json:"jenkins,omitempty"`
	Name       *string     `json:"name,omitempty"`
	NewRelic   *NewRelic   `json:"newRelic,omitempty"`
	Prometheus *Prometheus `json:"prometheus,omitempty"`

	// Read-only fields.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type CloudWatch struct {
	IAmArn *string `json:"iamArn,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type DataDog struct {
	Address *string `json:"address,omitempty"`
	ApiKey  *string `json:"apiKey,omitempty"`
	AppKey  *string `json:"appKey,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Jenkins struct {
	ApiToken *string `json:"apiToken,omitempty"`
	BaseUrl  *string `json:"baseUrl,omitempty"`
	UserName *string `json:"username,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type NewRelic struct {
	AccountId        *string `json:"accountId,omitempty"`
	BaseUrlNerdGraph *string `json:"baseUrlNerdGraph,omitempty"`
	BaseUrlRest      *string `json:"baseUrlRest,omitempty"`
	PersonalApiKey   *string `json:"personalApiKey,omitempty"`
	Region           *string `json:"region,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Prometheus struct {
	Address *string `json:"address,omitempty"`

	forceSendFields []string
	nullFields      []string
}
