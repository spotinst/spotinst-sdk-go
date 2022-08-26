package spark

import "time"

type Cluster struct {
	ID                  *string `json:"id,omitempty"`
	ControllerClusterID *string `json:"controllerClusterId,omitempty"`
	OceanClusterID      *string `json:"oceanClusterId,omitempty"`
	Region              *string `json:"region,omitempty"`
	Config              *Config `json:"config,omitempty"`
	State               *string `json:"state,omitempty"`

	// Read-only fields.
	K8sVersion            *string    `json:"k8sVersion,omitempty"`
	OperatorVersion       *string    `json:"operatorVersion,omitempty"`
	OperatorLastHeartbeat *time.Time `json:"operatorLastHeartbeat,omitempty"`
	CreatedAt             *time.Time `json:"createdAt,omitempty"`
	UpdatedAt             *time.Time `json:"updatedAt,omitempty"`
}

type Config struct {
	Ingress       *IngressConfig       `json:"ingress,omitempty"`
	Webhook       *WebhookConfig       `json:"webhook,omitempty"`
	Compute       *ComputeConfig       `json:"compute,omitempty"`
	LogCollection *LogCollectionConfig `json:"logCollection,omitempty"`
}

type LogCollectionConfig struct {
	CollectDriverLogs *bool `json:"collectDriverLogs,omitempty"`
}

type ComputeConfig struct {
	UseTaints  *bool `json:"useTaints,omitempty"`
	CreateVngs *bool `json:"createVngs,omitempty"`
}

type WebhookConfig struct {
	UseHostNetwork   *bool  `json:"useHostNetwork,omitempty"`
	HostNetworkPorts []*int `json:"hostNetworkPorts,omitempty"`
}

type IngressConfig struct {
	ServiceAnnotations map[string]string `json:"serviceAnnotations,omitempty"`
	DeployIngress      *bool             `json:"deployIngress,omitempty"`
}

type ListClustersInput struct {
	ControllerClusterID *string `json:"controllerClusterId,omitempty"`
	ClusterState        *string `json:"clusterState,omitempty"`
}

type ListClustersOutput struct {
	Clusters []*Cluster `json:"clusters,omitempty"`
}

type ReadClusterInput struct {
	ClusterID *string `json:"clusterId,omitempty"`
}

type ReadClusterOutput struct {
	Cluster *Cluster `json:"cluster,omitempty"`
}

type CreateClusterInput struct {
	Cluster *CreateClusterRequest `json:"cluster,omitempty"`
}

type CreateClusterRequest struct {
	OceanClusterID *string `json:"oceanClusterId,omitempty"`
	Config         *Config `json:"config,omitempty"`
}

type CreateClusterOutput struct {
	Cluster *Cluster `json:"cluster,omitempty"`
}

type UpdateClusterInput struct {
	ClusterID *string               `json:"-"`
	Cluster   *UpdateClusterRequest `json:"cluster,omitempty"`
}

type UpdateClusterRequest struct {
	Config *Config `json:"config,omitempty"`
}

type UpdateClusterOutput struct{}

type DeleteClusterInput struct {
	ClusterID *string `json:"clusterId,omitempty"`
}

type DeleteClusterOutput struct{}
