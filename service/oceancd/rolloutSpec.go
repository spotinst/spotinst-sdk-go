package oceancd

import "time"

type RolloutSpec struct {
	FailurePolicy   *FailurePolicy       `json:"failurePolicy,omitempty"`
	Name            *string              `json:"name,omitempty"`
	SpotDeployment  *SpotDeployment      `json:"spotDeployment,omitempty"`
	SpotDeployments []*SpotDeployment    `json:"spotDeployments,omitempty"`
	Strategy        *RolloutSpecStrategy `json:"strategy,omitempty"`
	Traffic         *Traffic             `json:"traffic,omitempty"`

	// Read-only fields.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type FailurePolicy struct {
	Action *string `json:"action,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type SpotDeployment struct {
	ClusterId *string `json:"clusterId,omitempty"`
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type RolloutSpecStrategy struct {
	Args []*RolloutSpecArgs `json:"args,omitempty"`
	Name *string            `json:"name,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type RolloutSpecArgs struct {
	Name      *string               `json:"name,omitempty"`
	Value     *string               `json:"value,omitempty"`
	ValueFrom *RolloutSpecValueFrom `json:"valueFrom,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type RolloutSpecValueFrom struct {
	FieldRef *FieldRef `json:"fieldRef,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type FieldRef struct {
	FieldPath *string `json:"fieldPath,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Traffic struct {
	Alb           *Alb        `json:"alb,omitempty"`
	Ambassador    *Ambassador `json:"ambassador,omitempty"`
	CanaryService *string     `json:"canaryService,omitempty"`
	Istio         *Istio      `json:"istio,omitempty"`
	Nginx         *Nginx      `json:"nginx,omitempty"`
	PingPong      *PingPong   `json:"pingPong,omitempty"`
	Smi           *Smi        `json:"smi,omitempty"`
	StableService *string     `json:"stableService,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Alb struct {
	AnnotationPrefix *string           `json:"annotationPrefix,omitempty"`
	Ingress          *string           `json:"ingress,omitempty"`
	RootService      *string           `json:"rootService,omitempty"`
	ServicePort      *int              `json:"servicePort,omitempty"`
	StickinessConfig *StickinessConfig `json:"stickinessConfig,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type StickinessConfig struct {
	DurationSeconds *int  `json:"durationSeconds,omitempty"`
	Enabled         *bool `json:"enabled,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Ambassador struct {
	Mappings []string `json:"mappings,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Istio struct {
	DestinationRule *DestinationRule   `json:"destinationRule,omitempty"`
	VirtualServices []*VirtualServices `json:"virtualServices,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type DestinationRule struct {
	CanarySubsetName *string `json:"canarySubsetName,omitempty"`
	Name             *string `json:"name,omitempty"`
	StableSubsetName *string `json:"stableSubsetName,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type VirtualServices struct {
	Name      *string      `json:"name,omitempty"`
	Routes    []string     `json:"routes,omitempty"`
	TlsRoutes []*TlsRoutes `json:"tlsRoutes,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type TlsRoutes struct {
	Port     *int     `json:"port,omitempty"`
	SniHosts []string `json:"sniHosts,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Nginx struct {
	AdditionalIngressAnnotations *AdditionalIngressAnnotations `json:"additionalIngressAnnotations,omitempty"`
	AnnotationPrefix             *string                       `json:"annotationPrefix,omitempty"`
	StableIngress                *string                       `json:"stableIngress"`

	forceSendFields []string
	nullFields      []string
}

type AdditionalIngressAnnotations struct {
	CanaryByHeader *string `json:"canary-by-header,omitempty"`
	Key1           *string `json:"key1,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type PingPong struct {
	PingService *string `json:"pingService,omitempty"`
	PongService *string `json:"pongService,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Smi struct {
	RootService      *string `json:"rootService,omitempty"`
	TrafficSplitName *string `json:"trafficSplitName,omitempty"`

	forceSendFields []string
	nullFields      []string
}
