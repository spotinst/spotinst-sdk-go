package aws

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/jsonutil"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/uritemplates"
)

// A Product represents the type of an operating system.
type Product int

const (
	// ProductWindows represents the Windows product.
	ProductWindows Product = iota

	// ProductWindowsVPC represents the Windows (Amazon VPC) product.
	ProductWindowsVPC

	// ProductLinuxUnix represents the Linux/Unix product.
	ProductLinuxUnix

	// ProductLinuxUnixVPC represents the Linux/Unix (Amazon VPC) product.
	ProductLinuxUnixVPC

	// ProductSUSELinux represents the SUSE Linux product.
	ProductSUSELinux

	// ProductSUSELinuxVPC represents the SUSE Linux (Amazon VPC) product.
	ProductSUSELinuxVPC
)

var ProductName = map[Product]string{
	ProductWindows:      "Windows",
	ProductWindowsVPC:   "Windows (Amazon VPC)",
	ProductLinuxUnix:    "Linux/UNIX",
	ProductLinuxUnixVPC: "Linux/UNIX (Amazon VPC)",
	ProductSUSELinux:    "SUSE Linux",
	ProductSUSELinuxVPC: "SUSE Linux (Amazon VPC)",
}

var ProductValue = map[string]Product{
	"Windows":                 ProductWindows,
	"Windows (Amazon VPC)":    ProductWindowsVPC,
	"Linux/UNIX":              ProductLinuxUnix,
	"Linux/UNIX (Amazon VPC)": ProductLinuxUnixVPC,
	"SUSE Linux":              ProductSUSELinux,
	"SUSE Linux (Amazon VPC)": ProductSUSELinuxVPC,
}

func (p Product) String() string {
	return ProductName[p]
}

type Group struct {
	ID          *string      `json:"id,omitempty"`
	Name        *string      `json:"name,omitempty"`
	Description *string      `json:"description,omitempty"`
	Region      *string      `json:"region,omitempty"`
	Capacity    *Capacity    `json:"capacity,omitempty"`
	Compute     *Compute     `json:"compute,omitempty"`
	Strategy    *Strategy    `json:"strategy,omitempty"`
	Scaling     *Scaling     `json:"scaling,omitempty"`
	Scheduling  *Scheduling  `json:"scheduling,omitempty"`
	Integration *Integration `json:"thirdPartiesIntegration,omitempty"`

	// forceSendFields is a list of field names (e.g. "Keys") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	forceSendFields []string

	// nullFields is a list of field names (e.g. "Keys") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	nullFields []string
}

type Integration struct {
	EC2ContainerService *EC2ContainerServiceIntegration `json:"ecs,omitempty"`
	ElasticBeanstalk    *ElasticBeanstalkIntegration    `json:"elasticBeanstalk,omitempty"`
	CodeDeploy          *CodeDeployIntegration          `json:"codeDeploy,omitempty"`
	OpsWorks            *OpsWorksIntegration            `json:"opsWorks,omitempty"`
	Rancher             *RancherIntegration             `json:"rancher,omitempty"`
	Kubernetes          *KubernetesIntegration          `json:"kubernetes,omitempty"`
	Mesosphere          *MesosphereIntegration          `json:"mesosphere,omitempty"`
	Multai              *MultaiIntegration              `json:"mlbRuntime,omitempty"`
	Nomad               *NomadIntegration               `json:"nomad,omitempty"`
	Chef                *ChefIntegration                `json:"chef,omitempty"`
	Gitlab              *GitlabIntegration              `json:"gitlab,omitempty"`
	Route53             *Route53Integration             `json:"route53,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type AutoScale struct {
	IsEnabled    *bool              `json:"isEnabled,omitempty"`
	IsAutoConfig *bool              `json:"isAutoConfig,omitempty"`
	Cooldown     *int               `json:"cooldown,omitempty"`
	Headroom     *AutoScaleHeadroom `json:"headroom,omitempty"`
	Down         *AutoScaleDown     `json:"down,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type AutoScaleECS struct {
	AutoScale                                             // embedding
	Attributes                     []*AutoScaleAttributes `json:"attributes,omitempty"`
	ShouldScaleDownNonServiceTasks *bool                  `json:"shouldScaleDownNonServiceTasks,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type AutoScaleKubernetes struct {
	AutoScale                   // embedding
	Labels    []*AutoScaleLabel `json:"labels,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type AutoScaleNomad struct {
	AutoScale                          // embedding
	Constraints []*AutoScaleConstraint `json:"constraints,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type AutoScaleHeadroom struct {
	CPUPerUnit    *int `json:"cpuPerUnit,omitempty"`
	MemoryPerUnit *int `json:"memoryPerUnit,omitempty"`
	NumOfUnits    *int `json:"numOfUnits,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type AutoScaleDown struct {
	EvaluationPeriods *int `json:"evaluationPeriods,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type AutoScaleConstraint struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type AutoScaleLabel struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type AutoScaleAttributes struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ElasticBeanstalkIntegration struct {
	EnvironmentID         *string                `json:"environmentId,omitempty"`
	DeploymentPreferences *DeploymentPreferences `json:"deploymentPreferences,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type DeploymentPreferences struct {
	AutomaticRoll       *bool              `json:"automaticRoll,omitempty"`
	BatchSizePercentage *int               `json:"batchSizePercentage,omitempty"`
	GracePeriod         *int               `json:"gracePeriod,omitempty"`
	Strategy            *BeanstalkStrategy `json:"strategy,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type BeanstalkStrategy struct {
	Action               *string `json:"action,omitempty"`
	ShouldDrainInstances *bool   `json:"shouldDrainInstances,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type CodeDeployIntegration struct {
	DeploymentGroups           []*DeploymentGroup `json:"deploymentGroups,omitempty"`
	CleanUpOnFailure           *bool              `json:"cleanUpOnFailure,omitempty"`
	TerminateInstanceOnFailure *bool              `json:"terminateInstanceOnFailure,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type DeploymentGroup struct {
	ApplicationName     *string `json:"applicationName,omitempty"`
	DeploymentGroupName *string `json:"deploymentGroupName,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type OpsWorksIntegration struct {
	LayerID   *string `json:"layerId,omitempty"`
	StackType *string `json:"stackType,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type RancherIntegration struct {
	MasterHost *string `json:"masterHost,omitempty"`
	AccessKey  *string `json:"accessKey,omitempty"`
	SecretKey  *string `json:"secretKey,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type EC2ContainerServiceIntegration struct {
	ClusterName  *string       `json:"clusterName,omitempty"`
	AutoScaleECS *AutoScaleECS `json:"autoScale,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type KubernetesIntegration struct {
	IntegrationMode     *string              `json:"integrationMode,omitempty"`
	ClusterIdentifier   *string              `json:"clusterIdentifier,omitempty"`
	Server              *string              `json:"apiServer,omitempty"`
	Token               *string              `json:"token,omitempty"`
	AutoScaleKubernetes *AutoScaleKubernetes `json:"autoScale,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type MesosphereIntegration struct {
	Server *string `json:"apiServer,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type MultaiIntegration struct {
	DeploymentID *string `json:"deploymentId,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type NomadIntegration struct {
	MasterHost     *string         `json:"masterHost,omitempty"`
	MasterPort     *int            `json:"masterPort,omitempty"`
	ACLToken       *string         `json:"aclToken,omitempty"`
	AutoScaleNomad *AutoScaleNomad `json:"autoScale,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ChefIntegration struct {
	Server       *string `json:"chefServer,omitempty"`
	Organization *string `json:"organization,omitempty"`
	User         *string `json:"user,omitempty"`
	PEMKey       *string `json:"pemKey,omitempty"`
	Version      *string `json:"chefVersion,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Route53Integration struct {
	Domains []*Domain `json:"domains,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Domain struct {
	HostedZoneID *string      `json:"hostedZoneId,omitempty"`
	RecordSets   []*RecordSet `json:"recordSets,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type RecordSet struct {
	UsePublicIP *bool   `json:"usePublicIp,omitempty"`
	Name        *string `json:"name,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type GitlabIntegration struct {
	Runner *GitlabRunner `json:"runner,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type GitlabRunner struct {
	IsEnabled *bool `json:"isEnabled,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Scheduling struct {
	Tasks []*Task `json:"tasks,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Task struct {
	IsEnabled           *bool   `json:"isEnabled,omitempty"`
	Type                *string `json:"taskType,omitempty"`
	Frequency           *string `json:"frequency,omitempty"`
	CronExpression      *string `json:"cronExpression,omitempty"`
	StartTime           *string `json:"startTime,omitempty"`
	ScaleTargetCapacity *int    `json:"scaleTargetCapacity,omitempty"`
	ScaleMinCapacity    *int    `json:"scaleMinCapacity,omitempty"`
	ScaleMaxCapacity    *int    `json:"scaleMaxCapacity,omitempty"`
	BatchSizePercentage *int    `json:"batchSizePercentage,omitempty"`
	GracePeriod         *int    `json:"gracePeriod,omitempty"`
	TargetCapacity      *int    `json:"targetCapacity,omitempty"`
	MinCapacity         *int    `json:"minCapacity,omitempty"`
	MaxCapacity         *int    `json:"maxCapacity,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Scaling struct {
	Up     []*ScalingPolicy `json:"up,omitempty"`
	Down   []*ScalingPolicy `json:"down,omitempty"`
	Target []*ScalingPolicy `json:"target,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ScalingPolicy struct {
	PolicyName        *string      `json:"policyName,omitempty"`
	MetricName        *string      `json:"metricName,omitempty"`
	Namespace         *string      `json:"namespace,omitempty"`
	Source            *string      `json:"source,omitempty"`
	Statistic         *string      `json:"statistic,omitempty"`
	Unit              *string      `json:"unit,omitempty"`
	Threshold         *float64     `json:"threshold,omitempty"`
	Adjustment        *int         `json:"adjustment,omitempty"`
	MinTargetCapacity *int         `json:"minTargetCapacity,omitempty"`
	MaxTargetCapacity *int         `json:"maxTargetCapacity,omitempty"`
	EvaluationPeriods *int         `json:"evaluationPeriods,omitempty"`
	Period            *int         `json:"period,omitempty"`
	Cooldown          *int         `json:"cooldown,omitempty"`
	Operator          *string      `json:"operator,omitempty"`
	Dimensions        []*Dimension `json:"dimensions,omitempty"`
	Action            *Action      `json:"action,omitempty"`
	Target            *float64     `json:"target,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Action struct {
	Type              *string `json:"type,omitempty"`
	Adjustment        *string `json:"adjustment,omitempty"`
	MinTargetCapacity *string `json:"minTargetCapacity,omitempty"`
	MaxTargetCapacity *string `json:"maxTargetCapacity,omitempty"`
	Maximum           *string `json:"maximum,omitempty"`
	Minimum           *string `json:"minimum,omitempty"`
	Target            *string `json:"target,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Dimension struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Strategy struct {
	Risk                     *float64      `json:"risk,omitempty"`
	OnDemandCount            *int          `json:"onDemandCount,omitempty"`
	DrainingTimeout          *int          `json:"drainingTimeout,omitempty"`
	AvailabilityVsCost       *string       `json:"availabilityVsCost,omitempty"`
	LifetimePeriod           *string       `json:"lifetimePeriod,omitempty"`
	UtilizeReservedInstances *bool         `json:"utilizeReservedInstances,omitempty"`
	FallbackToOnDemand       *bool         `json:"fallbackToOd,omitempty"`
	SpinUpTime               *int          `json:"spinUpTime,omitempty"`
	Signals                  []*Signal     `json:"signals,omitempty"`
	Persistence              *Persistence  `json:"persistence,omitempty"`
	RevertToSpot             *RevertToSpot `json:"revertToSpot,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Persistence struct {
	ShouldPersistPrivateIP    *bool   `json:"shouldPersistPrivateIp,omitempty"`
	ShouldPersistBlockDevices *bool   `json:"shouldPersistBlockDevices,omitempty"`
	ShouldPersistRootDevice   *bool   `json:"shouldPersistRootDevice,omitempty"`
	BlockDevicesMode          *string `json:"blockDevicesMode,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type RevertToSpot struct {
	PerformAt   *string  `json:"performAt,omitempty"`
	TimeWindows []string `json:"timeWindows,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Signal struct {
	Name    *string `json:"name,omitempty"`
	Timeout *int    `json:"timeout,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Capacity struct {
	Minimum *int    `json:"minimum,omitempty"`
	Maximum *int    `json:"maximum,omitempty"`
	Target  *int    `json:"target,omitempty"`
	Unit    *string `json:"unit,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Compute struct {
	Product                    *string              `json:"product,omitempty"`
	InstanceTypes              *InstanceTypes       `json:"instanceTypes,omitempty"`
	LaunchSpecification        *LaunchSpecification `json:"launchSpecification,omitempty"`
	AvailabilityZones          []*AvailabilityZone  `json:"availabilityZones,omitempty"`
	PreferredAvailabilityZones []string             `json:"preferredAvailabilityZones,omitempty"`
	ElasticIPs                 []string             `json:"elasticIps,omitempty"`
	EBSVolumePool              []*EBSVolume         `json:"ebsVolumePool,omitempty"`
	PrivateIPs                 []string             `json:"privateIps,omitempty"`
	SubnetIDs                  []string             `json:"subnetIds,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type EBSVolume struct {
	DeviceName *string  `json:"deviceName,omitempty"`
	VolumeIDs  []string `json:"volumeIds,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type InstanceTypes struct {
	OnDemand      *string               `json:"ondemand,omitempty"`
	Spot          []string              `json:"spot,omitempty"`
	PreferredSpot []string              `json:"preferredSpot,omitempty"`
	Weights       []*InstanceTypeWeight `json:"weights,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type InstanceTypeWeight struct {
	InstanceType *string `json:"instanceType,omitempty"`
	Weight       *int    `json:"weightedCapacity,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type AvailabilityZone struct {
	Name               *string `json:"name,omitempty"`
	SubnetID           *string `json:"subnetId,omitempty"`
	PlacementGroupName *string `json:"placementGroupName,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type LaunchSpecification struct {
	LoadBalancerNames                             []string              `json:"loadBalancerNames,omitempty"`
	LoadBalancersConfig                           *LoadBalancersConfig  `json:"loadBalancersConfig,omitempty"`
	SecurityGroupIDs                              []string              `json:"securityGroupIds,omitempty"`
	HealthCheckType                               *string               `json:"healthCheckType,omitempty"`
	HealthCheckGracePeriod                        *int                  `json:"healthCheckGracePeriod,omitempty"`
	HealthCheckUnhealthyDurationBeforeReplacement *int                  `json:"healthCheckUnhealthyDurationBeforeReplacement,omitempty"`
	ImageID                                       *string               `json:"imageId,omitempty"`
	KeyPair                                       *string               `json:"keyPair,omitempty"`
	UserData                                      *string               `json:"userData,omitempty"`
	ShutdownScript                                *string               `json:"shutdownScript,omitempty"`
	Tenancy                                       *string               `json:"tenancy,omitempty"`
	Monitoring                                    *bool                 `json:"monitoring,omitempty"`
	EBSOptimized                                  *bool                 `json:"ebsOptimized,omitempty"`
	IAMInstanceProfile                            *IAMInstanceProfile   `json:"iamRole,omitempty"`
	BlockDeviceMappings                           []*BlockDeviceMapping `json:"blockDeviceMappings,omitempty"`
	NetworkInterfaces                             []*NetworkInterface   `json:"networkInterfaces,omitempty"`
	Tags                                          []*Tag                `json:"tags,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type LoadBalancersConfig struct {
	LoadBalancers []*LoadBalancer `json:"loadBalancers,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type LoadBalancer struct {
	Name          *string `json:"name,omitempty"`
	Arn           *string `json:"arn,omitempty"`
	Type          *string `json:"type,omitempty"`
	BalancerID    *string `json:"balancerId,omitempty"`
	TargetSetID   *string `json:"targetSetId,omitempty"`
	ZoneAwareness *bool   `json:"azAwareness,omitempty"`
	AutoWeight    *bool   `json:"autoWeight,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type NetworkInterface struct {
	ID                             *string  `json:"networkInterfaceId,omitempty"`
	Description                    *string  `json:"description,omitempty"`
	DeviceIndex                    *int     `json:"deviceIndex,omitempty"`
	SecondaryPrivateIPAddressCount *int     `json:"secondaryPrivateIpAddressCount,omitempty"`
	AssociatePublicIPAddress       *bool    `json:"associatePublicIpAddress,omitempty"`
	DeleteOnTermination            *bool    `json:"deleteOnTermination,omitempty"`
	SecurityGroupsIDs              []string `json:"groups,omitempty"`
	PrivateIPAddress               *string  `json:"privateIpAddress,omitempty"`
	SubnetID                       *string  `json:"subnetId,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type BlockDeviceMapping struct {
	DeviceName  *string `json:"deviceName,omitempty"`
	VirtualName *string `json:"virtualName,omitempty"`
	EBS         *EBS    `json:"ebs,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type EBS struct {
	DeleteOnTermination *bool   `json:"deleteOnTermination,omitempty"`
	Encrypted           *bool   `json:"encrypted,omitempty"`
	KmsKeyId            *string `json:"kmsKeyId,omitempty"`
	SnapshotID          *string `json:"snapshotId,omitempty"`
	VolumeType          *string `json:"volumeType,omitempty"`
	VolumeSize          *int    `json:"volumeSize,omitempty"`
	IOPS                *int    `json:"iops,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type IAMInstanceProfile struct {
	Name *string `json:"name,omitempty"`
	Arn  *string `json:"arn,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Instance struct {
	ID               *string    `json:"instanceId,omitempty"`
	SpotRequestID    *string    `json:"spotInstanceRequestId,omitempty"`
	InstanceType     *string    `json:"instanceType,omitempty"`
	Status           *string    `json:"status,omitempty"`
	Product          *string    `json:"product,omitempty"`
	AvailabilityZone *string    `json:"availabilityZone,omitempty"`
	PrivateIP        *string    `json:"privateIp,omitempty"`
	PublicIP         *string    `json:"publicIp,omitempty"`
	CreatedAt        *time.Time `json:"createdAt,omitempty"`
}

type RollStrategy struct {
	Action               *string `json:"action,omitempty"`
	ShouldDrainInstances *bool   `json:"shouldDrainInstances,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ListGroupsInput struct{}

type ListGroupsOutput struct {
	Groups []*Group `json:"groups,omitempty"`
}

type CreateGroupInput struct {
	Group *Group `json:"group,omitempty"`
}

type CreateGroupOutput struct {
	Group *Group `json:"group,omitempty"`
}

type ReadGroupInput struct {
	GroupID *string `json:"groupId,omitempty"`
}

type ReadGroupOutput struct {
	Group *Group `json:"group,omitempty"`
}

type UpdateGroupInput struct {
	Group                *Group `json:"group,omitempty"`
	ShouldResumeStateful *bool  `json:"-"`
}

type UpdateGroupOutput struct {
	Group *Group `json:"group,omitempty"`
}

type DeleteGroupInput struct {
	GroupID              *string               `json:"groupId,omitempty"`
	StatefulDeallocation *StatefulDeallocation `json:"statefulDeallocation,omitempty"`
}

type StatefulDeallocation struct {
	ShouldDeleteImages            *bool `json:"shouldDeleteImages,omitempty"`
	ShouldDeleteNetworkInterfaces *bool `json:"shouldDeleteNetworkInterfaces,omitempty"`
	ShouldDeleteVolumes           *bool `json:"shouldDeleteVolumes,omitempty"`
	ShouldDeleteSnapshots         *bool `json:"shouldDeleteSnapshots,omitempty"`
}

type DeleteGroupOutput struct{}

type StatusGroupInput struct {
	GroupID *string `json:"groupId,omitempty"`
}

type StatusGroupOutput struct {
	Instances []*Instance `json:"instances,omitempty"`
}

type DetachGroupInput struct {
	GroupID                       *string  `json:"groupId,omitempty"`
	InstanceIDs                   []string `json:"instancesToDetach,omitempty"`
	ShouldDecrementTargetCapacity *bool    `json:"shouldDecrementTargetCapacity,omitempty"`
	ShouldTerminateInstances      *bool    `json:"shouldTerminateInstances,omitempty"`
	DrainingTimeout               *int     `json:"drainingTimeout,omitempty"`
}

type DetachGroupOutput struct{}

type RollGroupInput struct {
	GroupID             *string       `json:"groupId,omitempty"`
	BatchSizePercentage *int          `json:"batchSizePercentage,omitempty"`
	GracePeriod         *int          `json:"gracePeriod,omitempty"`
	HealthCheckType     *string       `json:"healthCheckType,omitempty"`
	Strategy            *RollStrategy `json:"strategy,omitempty"`
}

type RollGroupOutput struct{}

type ImportBeanstalkInput struct {
	EnvironmentName *string `json:"environmentName,omitempty"`
	Region          *string `json:"region,omitempty"`
}

type ImportBeanstalkOutput struct {
	Group *Group `json:"group,omitempty"`
}

func groupFromJSON(in []byte) (*Group, error) {
	b := new(Group)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func groupsFromJSON(in []byte) ([]*Group, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Group, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := groupFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func groupsFromHttpResponse(resp *http.Response) ([]*Group, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return groupsFromJSON(body)
}

func instanceFromJSON(in []byte) (*Instance, error) {
	b := new(Instance)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func instancesFromJSON(in []byte) ([]*Instance, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Instance, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := instanceFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func instancesFromHttpResponse(resp *http.Response) ([]*Instance, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return instancesFromJSON(body)
}

func (s *ServiceOp) List(ctx context.Context, input *ListGroupsInput) (*ListGroupsOutput, error) {
	r := client.NewRequest(http.MethodGet, "/aws/ec2/group")
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := groupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListGroupsOutput{Groups: gs}, nil
}

func (s *ServiceOp) Create(ctx context.Context, input *CreateGroupInput) (*CreateGroupOutput, error) {
	r := client.NewRequest(http.MethodPost, "/aws/ec2/group")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := groupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateGroupOutput)
	if len(gs) > 0 {
		output.Group = gs[0]
	}

	return output, nil
}

func (s *ServiceOp) Read(ctx context.Context, input *ReadGroupInput) (*ReadGroupOutput, error) {
	path, err := uritemplates.Expand("/aws/ec2/group/{groupId}", uritemplates.Values{
		"groupId": spotinst.StringValue(input.GroupID),
	})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodGet, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := groupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadGroupOutput)
	if len(gs) > 0 {
		output.Group = gs[0]
	}

	return output, nil
}

func (s *ServiceOp) Update(ctx context.Context, input *UpdateGroupInput) (*UpdateGroupOutput, error) {
	path, err := uritemplates.Expand("/aws/ec2/group/{groupId}", uritemplates.Values{
		"groupId": spotinst.StringValue(input.Group.ID),
	})
	if err != nil {
		return nil, err
	}

	// We do not need the ID anymore so let's drop it.
	input.Group.ID = nil

	r := client.NewRequest(http.MethodPut, path)
	r.Obj = input

	if input.ShouldResumeStateful != nil {
		r.Params.Set("shouldResumeStateful",
			strconv.FormatBool(spotinst.BoolValue(input.ShouldResumeStateful)))
	}

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := groupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(UpdateGroupOutput)
	if len(gs) > 0 {
		output.Group = gs[0]
	}

	return output, nil
}

func (s *ServiceOp) Delete(ctx context.Context, input *DeleteGroupInput) (*DeleteGroupOutput, error) {
	path, err := uritemplates.Expand("/aws/ec2/group/{groupId}", uritemplates.Values{
		"groupId": spotinst.StringValue(input.GroupID),
	})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodDelete, path)

	if input.StatefulDeallocation != nil {
		r.Obj = &DeleteGroupInput{
			StatefulDeallocation: input.StatefulDeallocation,
		}
	}

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteGroupOutput{}, nil
}

func (s *ServiceOp) Status(ctx context.Context, input *StatusGroupInput) (*StatusGroupOutput, error) {
	path, err := uritemplates.Expand("/aws/ec2/group/{groupId}/status", uritemplates.Values{
		"groupId": spotinst.StringValue(input.GroupID),
	})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodGet, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	is, err := instancesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &StatusGroupOutput{Instances: is}, nil
}

func (s *ServiceOp) Detach(ctx context.Context, input *DetachGroupInput) (*DetachGroupOutput, error) {
	path, err := uritemplates.Expand("/aws/ec2/group/{groupId}/detachInstances", uritemplates.Values{
		"groupId": spotinst.StringValue(input.GroupID),
	})
	if err != nil {
		return nil, err
	}

	// We do not need the ID anymore so let's drop it.
	input.GroupID = nil

	r := client.NewRequest(http.MethodPut, path)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DetachGroupOutput{}, nil
}

func (s *ServiceOp) Roll(ctx context.Context, input *RollGroupInput) (*RollGroupOutput, error) {
	path, err := uritemplates.Expand("/aws/ec2/group/{groupId}/roll", uritemplates.Values{
		"groupId": spotinst.StringValue(input.GroupID),
	})
	if err != nil {
		return nil, err
	}

	// We do not need the ID anymore so let's drop it.
	input.GroupID = nil

	r := client.NewRequest(http.MethodPut, path)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &RollGroupOutput{}, nil
}

func (s *ServiceOp) ImportBeanstalkEnv(ctx context.Context, input *ImportBeanstalkInput) (*ImportBeanstalkOutput, error) {
	path := "/aws/ec2/group/beanstalk/import"
	r := client.NewRequest(http.MethodGet, path)

	r.Params["environmentName"] = []string{spotinst.StringValue(input.EnvironmentName)}
	r.Params["region"] = []string{spotinst.StringValue(input.Region)}

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := groupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ImportBeanstalkOutput)
	if len(gs) > 0 {
		output.Group = gs[0]
	}

	return output, nil
}

// region Group

func (o *Group) MarshalJSON() ([]byte, error) {
	type noMethod Group
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Group) SetId(v *string) *Group {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *Group) SetName(v *string) *Group {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *Group) SetDescription(v *string) *Group {
	if o.Description = v; o.Description == nil {
		o.nullFields = append(o.nullFields, "Description")
	}
	return o
}

func (o *Group) SetCapacity(v *Capacity) *Group {
	if o.Capacity = v; o.Capacity == nil {
		o.nullFields = append(o.nullFields, "Capacity")
	}
	return o
}

func (o *Group) SetCompute(v *Compute) *Group {
	if o.Compute = v; o.Compute == nil {
		o.nullFields = append(o.nullFields, "Compute")
	}
	return o
}

func (o *Group) SetStrategy(v *Strategy) *Group {
	if o.Strategy = v; o.Strategy == nil {
		o.nullFields = append(o.nullFields, "Strategy")
	}
	return o
}

func (o *Group) SetScaling(v *Scaling) *Group {
	if o.Scaling = v; o.Scaling == nil {
		o.nullFields = append(o.nullFields, "Scaling")
	}
	return o
}

func (o *Group) SetScheduling(v *Scheduling) *Group {
	if o.Scheduling = v; o.Scheduling == nil {
		o.nullFields = append(o.nullFields, "Scheduling")
	}
	return o
}

func (o *Group) SetIntegration(v *Integration) *Group {
	if o.Integration = v; o.Integration == nil {
		o.nullFields = append(o.nullFields, "Integration")
	}
	return o
}

func (o *Group) SetRegion(v *string) *Group {
	if o.Region = v; o.Region == nil {
		o.nullFields = append(o.nullFields, "Region")
	}
	return o
}

// endregion

// region Integration

func (o *Integration) MarshalJSON() ([]byte, error) {
	type noMethod Integration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Integration) SetRoute53(v *Route53Integration) *Integration {
	if o.Route53 = v; o.Route53 == nil {
		o.nullFields = append(o.nullFields, "Route53")
	}
	return o
}

func (o *Integration) SetEC2ContainerService(v *EC2ContainerServiceIntegration) *Integration {
	if o.EC2ContainerService = v; o.EC2ContainerService == nil {
		o.nullFields = append(o.nullFields, "EC2ContainerService")
	}
	return o
}

func (o *Integration) SetElasticBeanstalk(v *ElasticBeanstalkIntegration) *Integration {
	if o.ElasticBeanstalk = v; o.ElasticBeanstalk == nil {
		o.nullFields = append(o.nullFields, "ElasticBeanstalk")
	}
	return o
}

func (o *Integration) SetCodeDeploy(v *CodeDeployIntegration) *Integration {
	if o.CodeDeploy = v; o.CodeDeploy == nil {
		o.nullFields = append(o.nullFields, "CodeDeploy")
	}
	return o
}

func (o *Integration) SetOpsWorks(v *OpsWorksIntegration) *Integration {
	if o.OpsWorks = v; o.OpsWorks == nil {
		o.nullFields = append(o.nullFields, "OpsWorks")
	}
	return o
}

func (o *Integration) SetRancher(v *RancherIntegration) *Integration {
	if o.Rancher = v; o.Rancher == nil {
		o.nullFields = append(o.nullFields, "Rancher")
	}
	return o
}

func (o *Integration) SetKubernetes(v *KubernetesIntegration) *Integration {
	if o.Kubernetes = v; o.Kubernetes == nil {
		o.nullFields = append(o.nullFields, "Kubernetes")
	}
	return o
}

func (o *Integration) SetMesosphere(v *MesosphereIntegration) *Integration {
	if o.Mesosphere = v; o.Mesosphere == nil {
		o.nullFields = append(o.nullFields, "Mesosphere")
	}
	return o
}

func (o *Integration) SetMultai(v *MultaiIntegration) *Integration {
	if o.Multai = v; o.Multai == nil {
		o.nullFields = append(o.nullFields, "Multai")
	}
	return o
}

func (o *Integration) SetNomad(v *NomadIntegration) *Integration {
	if o.Nomad = v; o.Nomad == nil {
		o.nullFields = append(o.nullFields, "Nomad")
	}
	return o
}

func (o *Integration) SetChef(v *ChefIntegration) *Integration {
	if o.Chef = v; o.Chef == nil {
		o.nullFields = append(o.nullFields, "Chef")
	}
	return o
}

func (o *Integration) SetGitlab(v *GitlabIntegration) *Integration {
	if o.Gitlab = v; o.Gitlab == nil {
		o.nullFields = append(o.nullFields, "Gitlab")
	}
	return o
}

// endregion

// region RancherIntegration

func (o *RancherIntegration) MarshalJSON() ([]byte, error) {
	type noMethod RancherIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *RancherIntegration) SetMasterHost(v *string) *RancherIntegration {
	if o.MasterHost = v; o.MasterHost == nil {
		o.nullFields = append(o.nullFields, "MasterHost")
	}
	return o
}

func (o *RancherIntegration) SetAccessKey(v *string) *RancherIntegration {
	if o.AccessKey = v; o.AccessKey == nil {
		o.nullFields = append(o.nullFields, "AccessKey")
	}
	return o
}

func (o *RancherIntegration) SetSecretKey(v *string) *RancherIntegration {
	if o.SecretKey = v; o.SecretKey == nil {
		o.nullFields = append(o.nullFields, "SecretKey")
	}
	return o
}

// endregion

// region ElasticBeanstalkIntegration

func (o *ElasticBeanstalkIntegration) MarshalJSON() ([]byte, error) {
	type noMethod ElasticBeanstalkIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *ElasticBeanstalkIntegration) SetEnvironmentID(v *string) *ElasticBeanstalkIntegration {
	if o.EnvironmentID = v; o.EnvironmentID == nil {
		o.nullFields = append(o.nullFields, "EnvironmentID")
	}
	return o
}

func (o *ElasticBeanstalkIntegration) SetDeploymentPreferences(v *DeploymentPreferences) *ElasticBeanstalkIntegration {
	if o.DeploymentPreferences = v; o.DeploymentPreferences == nil {
		o.nullFields = append(o.nullFields, "DeploymentPreferences")
	}
	return o
}

// endregion

// region DeploymentPreferences

func (o *DeploymentPreferences) MarshalJSON() ([]byte, error) {
	type noMethod DeploymentPreferences
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *DeploymentPreferences) SetAutomaticRoll(v *bool) *DeploymentPreferences {
	if o.AutomaticRoll = v; o.AutomaticRoll == nil {
		o.nullFields = append(o.nullFields, "AutomaticRoll")
	}
	return o
}

func (o *DeploymentPreferences) SetBatchSizePercentage(v *int) *DeploymentPreferences {
	if o.BatchSizePercentage = v; o.BatchSizePercentage == nil {
		o.nullFields = append(o.nullFields, "BatchSizePercentage")
	}
	return o
}

func (o *DeploymentPreferences) SetGracePeriod(v *int) *DeploymentPreferences {
	if o.GracePeriod = v; o.GracePeriod == nil {
		o.nullFields = append(o.nullFields, "GracePeriod")
	}
	return o
}

func (o *DeploymentPreferences) SetBeanstalkStrategy(v *BeanstalkStrategy) *DeploymentPreferences {
	if o.Strategy = v; o.Strategy == nil {
		o.nullFields = append(o.nullFields, "Strategy")
	}
	return o
}

// endregion

// region BeanstalkStrategy

func (o *BeanstalkStrategy) MarshalJSON() ([]byte, error) {
	type noMethod BeanstalkStrategy
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *BeanstalkStrategy) SetAction(v *string) *BeanstalkStrategy {
	if o.Action = v; o.Action == nil {
		o.nullFields = append(o.nullFields, "Action")
	}
	return o
}

func (o *BeanstalkStrategy) SetShouldDrainInstances(v *bool) *BeanstalkStrategy {
	if o.ShouldDrainInstances = v; o.ShouldDrainInstances == nil {
		o.nullFields = append(o.nullFields, "ShouldDrainInstances")
	}
	return o
}

// endregion

// region EC2ContainerServiceIntegration

func (o *EC2ContainerServiceIntegration) MarshalJSON() ([]byte, error) {
	type noMethod EC2ContainerServiceIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *EC2ContainerServiceIntegration) SetClusterName(v *string) *EC2ContainerServiceIntegration {
	if o.ClusterName = v; o.ClusterName == nil {
		o.nullFields = append(o.nullFields, "ClusterName")
	}
	return o
}

func (o *AutoScaleECS) MarshalJSON() ([]byte, error) {
	type noMethod AutoScaleECS
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *EC2ContainerServiceIntegration) SetAutoScaleECS(v *AutoScaleECS) *EC2ContainerServiceIntegration {
	if o.AutoScaleECS = v; o.AutoScaleECS == nil {
		o.nullFields = append(o.nullFields, "AutoScaleECS")
	}
	return o
}

func (o *AutoScaleECS) SetAttributes(v []*AutoScaleAttributes) *AutoScaleECS {
	if o.Attributes = v; o.Attributes == nil {
		o.nullFields = append(o.nullFields, "Attributes")
	}
	return o
}

func (o *AutoScaleECS) SetShouldScaleDownNonServiceTasks(v *bool) *AutoScaleECS {
	if o.ShouldScaleDownNonServiceTasks = v; o.ShouldScaleDownNonServiceTasks == nil {
		o.nullFields = append(o.nullFields, "ShouldScaleDownNonServiceTasks")
	}
	return o
}

// endregion

// region Route53

func (o *Route53Integration) MarshalJSON() ([]byte, error) {
	type noMethod Route53Integration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Route53Integration) SetDomains(v []*Domain) *Route53Integration {
	if o.Domains = v; o.Domains == nil {
		o.nullFields = append(o.nullFields, "Domains")
	}
	return o
}

// endregion

// region Domain

func (o *Domain) MarshalJSON() ([]byte, error) {
	type noMethod Domain
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Domain) SetHostedZoneID(v *string) *Domain {
	if o.HostedZoneID = v; o.HostedZoneID == nil {
		o.nullFields = append(o.nullFields, "HostedZoneID")
	}
	return o
}

func (o *Domain) SetRecordSets(v []*RecordSet) *Domain {
	if o.RecordSets = v; o.RecordSets == nil {
		o.nullFields = append(o.nullFields, "RecordSets")
	}
	return o
}

// endregion

// region RecordSets

func (o *RecordSet) MarshalJSON() ([]byte, error) {
	type noMethod RecordSet
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *RecordSet) SetUsePublicIP(v *bool) *RecordSet {
	if o.UsePublicIP = v; o.UsePublicIP == nil {
		o.nullFields = append(o.nullFields, "UsePublicIP")
	}
	return o
}

func (o *RecordSet) SetName(v *string) *RecordSet {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

// endregion

// region AutoScale

func (o *AutoScale) MarshalJSON() ([]byte, error) {
	type noMethod AutoScale
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AutoScale) SetIsEnabled(v *bool) *AutoScale {
	if o.IsEnabled = v; o.IsEnabled == nil {
		o.nullFields = append(o.nullFields, "IsEnabled")
	}
	return o
}

func (o *AutoScale) SetIsAutoConfig(v *bool) *AutoScale {
	if o.IsAutoConfig = v; o.IsAutoConfig == nil {
		o.nullFields = append(o.nullFields, "IsAutoConfig")
	}
	return o
}

func (o *AutoScale) SetCooldown(v *int) *AutoScale {
	if o.Cooldown = v; o.Cooldown == nil {
		o.nullFields = append(o.nullFields, "Cooldown")
	}
	return o
}

func (o *AutoScale) SetHeadroom(v *AutoScaleHeadroom) *AutoScale {
	if o.Headroom = v; o.Headroom == nil {
		o.nullFields = append(o.nullFields, "Headroom")
	}
	return o
}

func (o *AutoScale) SetDown(v *AutoScaleDown) *AutoScale {
	if o.Down = v; o.Down == nil {
		o.nullFields = append(o.nullFields, "Down")
	}
	return o
}

// endregion

// region AutoScaleHeadroom

func (o *AutoScaleHeadroom) MarshalJSON() ([]byte, error) {
	type noMethod AutoScaleHeadroom
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AutoScaleHeadroom) SetCPUPerUnit(v *int) *AutoScaleHeadroom {
	if o.CPUPerUnit = v; o.CPUPerUnit == nil {
		o.nullFields = append(o.nullFields, "CPUPerUnit")
	}
	return o
}

func (o *AutoScaleHeadroom) SetMemoryPerUnit(v *int) *AutoScaleHeadroom {
	if o.MemoryPerUnit = v; o.MemoryPerUnit == nil {
		o.nullFields = append(o.nullFields, "MemoryPerUnit")
	}
	return o
}

func (o *AutoScaleHeadroom) SetNumOfUnits(v *int) *AutoScaleHeadroom {
	if o.NumOfUnits = v; o.NumOfUnits == nil {
		o.nullFields = append(o.nullFields, "NumOfUnits")
	}
	return o
}

// endregion

// region AutoScaleDown

func (o *AutoScaleDown) MarshalJSON() ([]byte, error) {
	type noMethod AutoScaleDown
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AutoScaleDown) SetEvaluationPeriods(v *int) *AutoScaleDown {
	if o.EvaluationPeriods = v; o.EvaluationPeriods == nil {
		o.nullFields = append(o.nullFields, "EvaluationPeriods")
	}
	return o
}

// endregion

// region AutoScaleConstraint

func (o *AutoScaleConstraint) MarshalJSON() ([]byte, error) {
	type noMethod AutoScaleConstraint
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AutoScaleConstraint) SetKey(v *string) *AutoScaleConstraint {
	if o.Key = v; o.Key == nil {
		o.nullFields = append(o.nullFields, "Key")
	}
	return o
}

func (o *AutoScaleConstraint) SetValue(v *string) *AutoScaleConstraint {
	if o.Value = v; o.Value == nil {
		o.nullFields = append(o.nullFields, "Value")
	}
	return o
}

// endregion

// region AutoScaleLabel

func (o *AutoScaleLabel) MarshalJSON() ([]byte, error) {
	type noMethod AutoScaleLabel
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AutoScaleLabel) SetKey(v *string) *AutoScaleLabel {
	if o.Key = v; o.Key == nil {
		o.nullFields = append(o.nullFields, "Key")
	}
	return o
}

func (o *AutoScaleLabel) SetValue(v *string) *AutoScaleLabel {
	if o.Value = v; o.Value == nil {
		o.nullFields = append(o.nullFields, "Value")
	}
	return o
}

// endregion

// region KubernetesIntegration

func (o *KubernetesIntegration) MarshalJSON() ([]byte, error) {
	type noMethod KubernetesIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *KubernetesIntegration) SetIntegrationMode(v *string) *KubernetesIntegration {
	if o.IntegrationMode = v; o.IntegrationMode == nil {
		o.nullFields = append(o.nullFields, "IntegrationMode")
	}
	return o
}

func (o *KubernetesIntegration) SetClusterIdentifier(v *string) *KubernetesIntegration {
	if o.ClusterIdentifier = v; o.ClusterIdentifier == nil {
		o.nullFields = append(o.nullFields, "ClusterIdentifier")
	}
	return o
}

func (o *KubernetesIntegration) SetServer(v *string) *KubernetesIntegration {
	if o.Server = v; o.Server == nil {
		o.nullFields = append(o.nullFields, "Server")
	}
	return o
}

func (o *KubernetesIntegration) SetToken(v *string) *KubernetesIntegration {
	if o.Token = v; o.Token == nil {
		o.nullFields = append(o.nullFields, "Token")
	}
	return o
}

func (o *KubernetesIntegration) SetAutoScaleKubernetes(v *AutoScaleKubernetes) *KubernetesIntegration {
	if o.AutoScaleKubernetes = v; o.AutoScaleKubernetes == nil {
		o.nullFields = append(o.nullFields, "AutoScaleKubernetes")
	}
	return o
}

func (o *AutoScaleKubernetes) MarshalJSON() ([]byte, error) {
	type noMethod AutoScaleKubernetes
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AutoScaleKubernetes) SetLabels(v []*AutoScaleLabel) *AutoScaleKubernetes {
	if o.Labels = v; o.Labels == nil {
		o.nullFields = append(o.nullFields, "Labels")
	}
	return o
}

// endregion

// region MesosphereIntegration

func (o *MesosphereIntegration) MarshalJSON() ([]byte, error) {
	type noMethod MesosphereIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *MesosphereIntegration) SetServer(v *string) *MesosphereIntegration {
	if o.Server = v; o.Server == nil {
		o.nullFields = append(o.nullFields, "Server")
	}
	return o
}

// endregion

// region MultaiIntegration

func (o *MultaiIntegration) MarshalJSON() ([]byte, error) {
	type noMethod MultaiIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *MultaiIntegration) SetDeploymentId(v *string) *MultaiIntegration {
	if o.DeploymentID = v; o.DeploymentID == nil {
		o.nullFields = append(o.nullFields, "DeploymentID")
	}
	return o
}

// endregion

// region NomadIntegration

func (o *NomadIntegration) MarshalJSON() ([]byte, error) {
	type noMethod NomadIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *NomadIntegration) SetMasterHost(v *string) *NomadIntegration {
	if o.MasterHost = v; o.MasterHost == nil {
		o.nullFields = append(o.nullFields, "MasterHost")
	}
	return o
}

func (o *NomadIntegration) SetMasterPort(v *int) *NomadIntegration {
	if o.MasterPort = v; o.MasterPort == nil {
		o.nullFields = append(o.nullFields, "MasterPort")
	}
	return o
}

func (o *NomadIntegration) SetAclToken(v *string) *NomadIntegration {
	if o.ACLToken = v; o.ACLToken == nil {
		o.nullFields = append(o.nullFields, "ACLToken")
	}
	return o
}

func (o *NomadIntegration) SetAutoScaleNomad(v *AutoScaleNomad) *NomadIntegration {
	if o.AutoScaleNomad = v; o.AutoScaleNomad == nil {
		o.nullFields = append(o.nullFields, "AutoScaleNomad")
	}
	return o
}

func (o *AutoScaleNomad) MarshalJSON() ([]byte, error) {
	type noMethod AutoScaleNomad
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AutoScaleNomad) SetConstraints(v []*AutoScaleConstraint) *AutoScaleNomad {
	if o.Constraints = v; o.Constraints == nil {
		o.nullFields = append(o.nullFields, "Constraints")
	}
	return o
}

// endregion

// region ChefIntegration

func (o *ChefIntegration) MarshalJSON() ([]byte, error) {
	type noMethod ChefIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *ChefIntegration) SetServer(v *string) *ChefIntegration {
	if o.Server = v; o.Server == nil {
		o.nullFields = append(o.nullFields, "Server")
	}
	return o
}

func (o *ChefIntegration) SetOrganization(v *string) *ChefIntegration {
	if o.Organization = v; o.Organization == nil {
		o.nullFields = append(o.nullFields, "Organization")
	}
	return o
}

func (o *ChefIntegration) SetUser(v *string) *ChefIntegration {
	if o.User = v; o.User == nil {
		o.nullFields = append(o.nullFields, "User")
	}
	return o
}

func (o *ChefIntegration) SetPEMKey(v *string) *ChefIntegration {
	if o.PEMKey = v; o.PEMKey == nil {
		o.nullFields = append(o.nullFields, "PEMKey")
	}
	return o
}

func (o *ChefIntegration) SetVersion(v *string) *ChefIntegration {
	if o.Version = v; o.Version == nil {
		o.nullFields = append(o.nullFields, "Version")
	}
	return o
}

// endregion

//region Gitlab
func (o *GitlabIntegration) MarshalJSON() ([]byte, error) {
	type noMethod GitlabIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *GitlabIntegration) SetRunner(v *GitlabRunner) *GitlabIntegration {
	if o.Runner = v; o.Runner == nil {
		o.nullFields = append(o.nullFields, "Runner")
	}
	return o
}

func (o *GitlabRunner) MarshalJSON() ([]byte, error) {
	type noMethod GitlabRunner
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *GitlabRunner) SetIsEnabled(v *bool) *GitlabRunner {
	if o.IsEnabled = v; o.IsEnabled == nil {
		o.nullFields = append(o.nullFields, "IsEnabled")
	}
	return o
}

//endregion

// region Scheduling

func (o *Scheduling) MarshalJSON() ([]byte, error) {
	type noMethod Scheduling
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Scheduling) SetTasks(v []*Task) *Scheduling {
	if o.Tasks = v; o.Tasks == nil {
		o.nullFields = append(o.nullFields, "Tasks")
	}
	return o
}

// endregion

// region Task

func (o *Task) MarshalJSON() ([]byte, error) {
	type noMethod Task
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Task) SetIsEnabled(v *bool) *Task {
	if o.IsEnabled = v; o.IsEnabled == nil {
		o.nullFields = append(o.nullFields, "IsEnabled")
	}
	return o
}

func (o *Task) SetType(v *string) *Task {
	if o.Type = v; o.Type == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

func (o *Task) SetFrequency(v *string) *Task {
	if o.Frequency = v; o.Frequency == nil {
		o.nullFields = append(o.nullFields, "Frequency")
	}
	return o
}

func (o *Task) SetCronExpression(v *string) *Task {
	if o.CronExpression = v; o.CronExpression == nil {
		o.nullFields = append(o.nullFields, "CronExpression")
	}
	return o
}

func (o *Task) SetStartTime(v *string) *Task {
	if o.StartTime = v; o.StartTime == nil {
		o.nullFields = append(o.nullFields, "StartTime")
	}
	return o
}

func (o *Task) SetScaleTargetCapacity(v *int) *Task {
	if o.ScaleTargetCapacity = v; o.ScaleTargetCapacity == nil {
		o.nullFields = append(o.nullFields, "ScaleTargetCapacity")
	}
	return o
}

func (o *Task) SetScaleMinCapacity(v *int) *Task {
	if o.ScaleMinCapacity = v; o.ScaleMinCapacity == nil {
		o.nullFields = append(o.nullFields, "ScaleMinCapacity")
	}
	return o
}

func (o *Task) SetScaleMaxCapacity(v *int) *Task {
	if o.ScaleMaxCapacity = v; o.ScaleMaxCapacity == nil {
		o.nullFields = append(o.nullFields, "ScaleMaxCapacity")
	}
	return o
}

func (o *Task) SetBatchSizePercentage(v *int) *Task {
	if o.BatchSizePercentage = v; o.BatchSizePercentage == nil {
		o.nullFields = append(o.nullFields, "BatchSizePercentage")
	}
	return o
}

func (o *Task) SetGracePeriod(v *int) *Task {
	if o.GracePeriod = v; o.GracePeriod == nil {
		o.nullFields = append(o.nullFields, "GracePeriod")
	}
	return o
}

func (o *Task) SetTargetCapacity(v *int) *Task {
	if o.TargetCapacity = v; o.TargetCapacity == nil {
		o.nullFields = append(o.nullFields, "TargetCapacity")
	}
	return o
}

func (o *Task) SetMinCapacity(v *int) *Task {
	if o.MinCapacity = v; o.MinCapacity == nil {
		o.nullFields = append(o.nullFields, "MinCapacity")
	}
	return o
}

func (o *Task) SetMaxCapacity(v *int) *Task {
	if o.MaxCapacity = v; o.MaxCapacity == nil {
		o.nullFields = append(o.nullFields, "MaxCapacity")
	}
	return o
}

// endregion

// region Scaling

func (o *Scaling) MarshalJSON() ([]byte, error) {
	type noMethod Scaling
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Scaling) SetUp(v []*ScalingPolicy) *Scaling {
	if o.Up = v; o.Up == nil {
		o.nullFields = append(o.nullFields, "Up")
	}
	return o
}

func (o *Scaling) SetDown(v []*ScalingPolicy) *Scaling {
	if o.Down = v; o.Down == nil {
		o.nullFields = append(o.nullFields, "Down")
	}
	return o
}

func (o *Scaling) SetTarget(v []*ScalingPolicy) *Scaling {
	if o.Target = v; o.Target == nil {
		o.nullFields = append(o.nullFields, "Target")
	}
	return o
}

// endregion

// region ScalingPolicy

func (o *ScalingPolicy) MarshalJSON() ([]byte, error) {
	type noMethod ScalingPolicy
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *ScalingPolicy) SetPolicyName(v *string) *ScalingPolicy {
	if o.PolicyName = v; o.PolicyName == nil {
		o.nullFields = append(o.nullFields, "PolicyName")
	}
	return o
}

func (o *ScalingPolicy) SetMetricName(v *string) *ScalingPolicy {
	if o.MetricName = v; o.MetricName == nil {
		o.nullFields = append(o.nullFields, "MetricName")
	}
	return o
}

func (o *ScalingPolicy) SetNamespace(v *string) *ScalingPolicy {
	if o.Namespace = v; o.Namespace == nil {
		o.nullFields = append(o.nullFields, "Namespace")
	}
	return o
}

func (o *ScalingPolicy) SetSource(v *string) *ScalingPolicy {
	if o.Source = v; o.Source == nil {
		o.nullFields = append(o.nullFields, "Source")
	}
	return o
}

func (o *ScalingPolicy) SetStatistic(v *string) *ScalingPolicy {
	if o.Statistic = v; o.Statistic == nil {
		o.nullFields = append(o.nullFields, "Statistic")
	}
	return o
}

func (o *ScalingPolicy) SetUnit(v *string) *ScalingPolicy {
	if o.Unit = v; o.Unit == nil {
		o.nullFields = append(o.nullFields, "Unit")
	}
	return o
}

func (o *ScalingPolicy) SetThreshold(v *float64) *ScalingPolicy {
	if o.Threshold = v; o.Threshold == nil {
		o.nullFields = append(o.nullFields, "Threshold")
	}
	return o
}

func (o *ScalingPolicy) SetAdjustment(v *int) *ScalingPolicy {
	if o.Adjustment = v; o.Adjustment == nil {
		o.nullFields = append(o.nullFields, "Adjustment")
	}
	return o
}

func (o *ScalingPolicy) SetMinTargetCapacity(v *int) *ScalingPolicy {
	if o.MinTargetCapacity = v; o.MinTargetCapacity == nil {
		o.nullFields = append(o.nullFields, "MinTargetCapacity")
	}
	return o
}

func (o *ScalingPolicy) SetMaxTargetCapacity(v *int) *ScalingPolicy {
	if o.MaxTargetCapacity = v; o.MaxTargetCapacity == nil {
		o.nullFields = append(o.nullFields, "MaxTargetCapacity")
	}
	return o
}

func (o *ScalingPolicy) SetEvaluationPeriods(v *int) *ScalingPolicy {
	if o.EvaluationPeriods = v; o.EvaluationPeriods == nil {
		o.nullFields = append(o.nullFields, "EvaluationPeriods")
	}
	return o
}

func (o *ScalingPolicy) SetPeriod(v *int) *ScalingPolicy {
	if o.Period = v; o.Period == nil {
		o.nullFields = append(o.nullFields, "Period")
	}
	return o
}

func (o *ScalingPolicy) SetCooldown(v *int) *ScalingPolicy {
	if o.Cooldown = v; o.Cooldown == nil {
		o.nullFields = append(o.nullFields, "Cooldown")
	}
	return o
}

func (o *ScalingPolicy) SetOperator(v *string) *ScalingPolicy {
	if o.Operator = v; o.Operator == nil {
		o.nullFields = append(o.nullFields, "Operator")
	}
	return o
}

func (o *ScalingPolicy) SetDimensions(v []*Dimension) *ScalingPolicy {
	if o.Dimensions = v; o.Dimensions == nil {
		o.nullFields = append(o.nullFields, "Dimensions")
	}
	return o
}

func (o *ScalingPolicy) SetAction(v *Action) *ScalingPolicy {
	if o.Action = v; o.Action == nil {
		o.nullFields = append(o.nullFields, "Action")
	}
	return o
}

func (o *ScalingPolicy) SetTarget(v *float64) *ScalingPolicy {
	if o.Target = v; o.Target == nil {
		o.nullFields = append(o.nullFields, "Target")
	}
	return o
}

// endregion

// region Action

func (o *Action) MarshalJSON() ([]byte, error) {
	type noMethod Action
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Action) SetType(v *string) *Action {
	if o.Type = v; o.Type == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

func (o *Action) SetAdjustment(v *string) *Action {
	if o.Adjustment = v; o.Adjustment == nil {
		o.nullFields = append(o.nullFields, "Adjustment")
	}
	return o
}

func (o *Action) SetMinTargetCapacity(v *string) *Action {
	if o.MinTargetCapacity = v; o.MinTargetCapacity == nil {
		o.nullFields = append(o.nullFields, "MinTargetCapacity")
	}
	return o
}

func (o *Action) SetMaxTargetCapacity(v *string) *Action {
	if o.MaxTargetCapacity = v; o.MaxTargetCapacity == nil {
		o.nullFields = append(o.nullFields, "MaxTargetCapacity")
	}
	return o
}

func (o *Action) SetMaximum(v *string) *Action {
	if o.Maximum = v; o.Maximum == nil {
		o.nullFields = append(o.nullFields, "Maximum")
	}
	return o
}

func (o *Action) SetMinimum(v *string) *Action {
	if o.Minimum = v; o.Minimum == nil {
		o.nullFields = append(o.nullFields, "Minimum")
	}
	return o
}

func (o *Action) SetTarget(v *string) *Action {
	if o.Target = v; o.Target == nil {
		o.nullFields = append(o.nullFields, "Target")
	}
	return o
}

// endregion

// region Dimension

func (o *Dimension) MarshalJSON() ([]byte, error) {
	type noMethod Dimension
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Dimension) SetName(v *string) *Dimension {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *Dimension) SetValue(v *string) *Dimension {
	if o.Value = v; o.Value == nil {
		o.nullFields = append(o.nullFields, "Value")
	}
	return o
}

// endregion

// region Strategy

func (o *Strategy) MarshalJSON() ([]byte, error) {
	type noMethod Strategy
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Strategy) SetRisk(v *float64) *Strategy {
	if o.Risk = v; o.Risk == nil {
		o.nullFields = append(o.nullFields, "Risk")
	}
	return o
}

func (o *Strategy) SetOnDemandCount(v *int) *Strategy {
	if o.OnDemandCount = v; o.OnDemandCount == nil {
		o.nullFields = append(o.nullFields, "OnDemandCount")
	}
	return o
}

func (o *Strategy) SetDrainingTimeout(v *int) *Strategy {
	if o.DrainingTimeout = v; o.DrainingTimeout == nil {
		o.nullFields = append(o.nullFields, "DrainingTimeout")
	}
	return o
}

func (o *Strategy) SetAvailabilityVsCost(v *string) *Strategy {
	if o.AvailabilityVsCost = v; o.AvailabilityVsCost == nil {
		o.nullFields = append(o.nullFields, "AvailabilityVsCost")
	}
	return o
}

func (o *Strategy) SetLifetimePeriod(v *string) *Strategy {
	if o.LifetimePeriod = v; o.LifetimePeriod == nil {
		o.nullFields = append(o.nullFields, "LifetimePeriod")
	}
	return o
}

func (o *Strategy) SetUtilizeReservedInstances(v *bool) *Strategy {
	if o.UtilizeReservedInstances = v; o.UtilizeReservedInstances == nil {
		o.nullFields = append(o.nullFields, "UtilizeReservedInstances")
	}
	return o
}

func (o *Strategy) SetFallbackToOnDemand(v *bool) *Strategy {
	if o.FallbackToOnDemand = v; o.FallbackToOnDemand == nil {
		o.nullFields = append(o.nullFields, "FallbackToOnDemand")
	}
	return o
}

func (o *Strategy) SetSpinUpTime(v *int) *Strategy {
	if o.SpinUpTime = v; o.SpinUpTime == nil {
		o.nullFields = append(o.nullFields, "SpinUpTime")
	}
	return o
}

func (o *Strategy) SetSignals(v []*Signal) *Strategy {
	if o.Signals = v; o.Signals == nil {
		o.nullFields = append(o.nullFields, "Signals")
	}
	return o
}

func (o *Strategy) SetPersistence(v *Persistence) *Strategy {
	if o.Persistence = v; o.Persistence == nil {
		o.nullFields = append(o.nullFields, "Persistence")
	}
	return o
}

func (o *Strategy) SetRevertToSpot(v *RevertToSpot) *Strategy {
	if o.RevertToSpot = v; o.RevertToSpot == nil {
		o.nullFields = append(o.nullFields, "RevertToSpot")
	}
	return o
}

// endregion

// region Persistence

func (o *Persistence) MarshalJSON() ([]byte, error) {
	type noMethod Persistence
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Persistence) SetShouldPersistPrivateIP(v *bool) *Persistence {
	if o.ShouldPersistPrivateIP = v; o.ShouldPersistPrivateIP == nil {
		o.nullFields = append(o.nullFields, "ShouldPersistPrivateIP")
	}
	return o
}

func (o *Persistence) SetShouldPersistBlockDevices(v *bool) *Persistence {
	if o.ShouldPersistBlockDevices = v; o.ShouldPersistBlockDevices == nil {
		o.nullFields = append(o.nullFields, "ShouldPersistBlockDevices")
	}
	return o
}

func (o *Persistence) SetShouldPersistRootDevice(v *bool) *Persistence {
	if o.ShouldPersistRootDevice = v; o.ShouldPersistRootDevice == nil {
		o.nullFields = append(o.nullFields, "ShouldPersistRootDevice")
	}
	return o
}

func (o *Persistence) SetBlockDevicesMode(v *string) *Persistence {
	if o.BlockDevicesMode = v; o.BlockDevicesMode == nil {
		o.nullFields = append(o.nullFields, "BlockDevicesMode")
	}
	return o
}

// endregion

// region RevertToSpot

func (o *RevertToSpot) MarshalJSON() ([]byte, error) {
	type noMethod RevertToSpot
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *RevertToSpot) SetPerformAt(v *string) *RevertToSpot {
	if o.PerformAt = v; o.PerformAt == nil {
		o.nullFields = append(o.nullFields, "PerformAt")
	}
	return o
}

func (o *RevertToSpot) SetTimeWindows(v []string) *RevertToSpot {
	if o.TimeWindows = v; o.TimeWindows == nil {
		o.nullFields = append(o.nullFields, "TimeWindows")
	}
	return o
}

// endregion

// region Signal

func (o *Signal) MarshalJSON() ([]byte, error) {
	type noMethod Signal
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Signal) SetName(v *string) *Signal {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *Signal) SetTimeout(v *int) *Signal {
	if o.Timeout = v; o.Timeout == nil {
		o.nullFields = append(o.nullFields, "Timeout")
	}
	return o
}

// endregion

// region Capacity

func (o *Capacity) MarshalJSON() ([]byte, error) {
	type noMethod Capacity
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Capacity) SetMinimum(v *int) *Capacity {
	if o.Minimum = v; o.Minimum == nil {
		o.nullFields = append(o.nullFields, "Minimum")
	}
	return o
}

func (o *Capacity) SetMaximum(v *int) *Capacity {
	if o.Maximum = v; o.Maximum == nil {
		o.nullFields = append(o.nullFields, "Maximum")
	}
	return o
}

func (o *Capacity) SetTarget(v *int) *Capacity {
	if o.Target = v; o.Target == nil {
		o.nullFields = append(o.nullFields, "Target")
	}
	return o
}

func (o *Capacity) SetUnit(v *string) *Capacity {
	if o.Unit = v; o.Unit == nil {
		o.nullFields = append(o.nullFields, "Unit")
	}
	return o
}

// endregion

// region Compute

func (o *Compute) MarshalJSON() ([]byte, error) {
	type noMethod Compute
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Compute) SetProduct(v *string) *Compute {
	if o.Product = v; o.Product == nil {
		o.nullFields = append(o.nullFields, "Product")
	}

	return o
}

func (o *Compute) SetPrivateIPs(v []string) *Compute {
	if o.PrivateIPs = v; o.PrivateIPs == nil {
		o.nullFields = append(o.nullFields, "PrivateIPs")
	}

	return o
}

func (o *Compute) SetInstanceTypes(v *InstanceTypes) *Compute {
	if o.InstanceTypes = v; o.InstanceTypes == nil {
		o.nullFields = append(o.nullFields, "InstanceTypes")
	}
	return o
}

func (o *Compute) SetLaunchSpecification(v *LaunchSpecification) *Compute {
	if o.LaunchSpecification = v; o.LaunchSpecification == nil {
		o.nullFields = append(o.nullFields, "LaunchSpecification")
	}
	return o
}

func (o *Compute) SetAvailabilityZones(v []*AvailabilityZone) *Compute {
	if o.AvailabilityZones = v; o.AvailabilityZones == nil {
		o.nullFields = append(o.nullFields, "AvailabilityZones")
	}
	return o
}

func (o *Compute) SetPreferredAvailabilityZones(v []string) *Compute {
	if o.PreferredAvailabilityZones = v; o.PreferredAvailabilityZones == nil {
		o.nullFields = append(o.nullFields, "PreferredAvailabilityZones")
	}
	return o
}

func (o *Compute) SetElasticIPs(v []string) *Compute {
	if o.ElasticIPs = v; o.ElasticIPs == nil {
		o.nullFields = append(o.nullFields, "ElasticIPs")
	}
	return o
}

func (o *Compute) SetEBSVolumePool(v []*EBSVolume) *Compute {
	if o.EBSVolumePool = v; o.EBSVolumePool == nil {
		o.nullFields = append(o.nullFields, "EBSVolumePool")
	}
	return o
}

func (o *Compute) SetSubnetIDs(v []string) *Compute {
	if o.SubnetIDs = v; o.SubnetIDs == nil {
		o.nullFields = append(o.nullFields, "SubnetIDs")
	}
	return o
}

// endregion

// region EBSVolume

func (o *EBSVolume) MarshalJSON() ([]byte, error) {
	type noMethod EBSVolume
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *EBSVolume) SetDeviceName(v *string) *EBSVolume {
	if o.DeviceName = v; o.DeviceName == nil {
		o.nullFields = append(o.nullFields, "DeviceName")
	}
	return o
}

func (o *EBSVolume) SetVolumeIDs(v []string) *EBSVolume {
	if o.VolumeIDs = v; o.VolumeIDs == nil {
		o.nullFields = append(o.nullFields, "VolumeIDs")
	}
	return o
}

// endregion

// region InstanceTypes

func (o *InstanceTypes) MarshalJSON() ([]byte, error) {
	type noMethod InstanceTypes
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *InstanceTypes) SetOnDemand(v *string) *InstanceTypes {
	if o.OnDemand = v; o.OnDemand == nil {
		o.nullFields = append(o.nullFields, "OnDemand")
	}
	return o
}

func (o *InstanceTypes) SetSpot(v []string) *InstanceTypes {
	if o.Spot = v; o.Spot == nil {
		o.nullFields = append(o.nullFields, "Spot")
	}
	return o
}

func (o *InstanceTypes) SetPreferredSpot(v []string) *InstanceTypes {
	if o.PreferredSpot = v; o.PreferredSpot == nil {
		o.nullFields = append(o.nullFields, "PreferredSpot")
	}
	return o
}

func (o *InstanceTypes) SetWeights(v []*InstanceTypeWeight) *InstanceTypes {
	if o.Weights = v; o.Weights == nil {
		o.nullFields = append(o.nullFields, "Weights")
	}
	return o
}

// endregion

// region InstanceTypeWeight

func (o *InstanceTypeWeight) MarshalJSON() ([]byte, error) {
	type noMethod InstanceTypeWeight
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *InstanceTypeWeight) SetInstanceType(v *string) *InstanceTypeWeight {
	if o.InstanceType = v; o.InstanceType == nil {
		o.nullFields = append(o.nullFields, "InstanceType")
	}
	return o
}

func (o *InstanceTypeWeight) SetWeight(v *int) *InstanceTypeWeight {
	if o.Weight = v; o.Weight == nil {
		o.nullFields = append(o.nullFields, "Weight")
	}
	return o
}

// endregion

// region AvailabilityZone

func (o *AvailabilityZone) MarshalJSON() ([]byte, error) {
	type noMethod AvailabilityZone
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AvailabilityZone) SetName(v *string) *AvailabilityZone {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *AvailabilityZone) SetSubnetId(v *string) *AvailabilityZone {
	if o.SubnetID = v; o.SubnetID == nil {
		o.nullFields = append(o.nullFields, "SubnetID")
	}
	return o
}

func (o *AvailabilityZone) SetPlacementGroupName(v *string) *AvailabilityZone {
	if o.PlacementGroupName = v; o.PlacementGroupName == nil {
		o.nullFields = append(o.nullFields, "PlacementGroupName")
	}
	return o
}

// endregion

// region LaunchSpecification

func (o *LaunchSpecification) MarshalJSON() ([]byte, error) {
	type noMethod LaunchSpecification
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *LaunchSpecification) SetLoadBalancerNames(v []string) *LaunchSpecification {
	if o.LoadBalancerNames = v; o.LoadBalancerNames == nil {
		o.nullFields = append(o.nullFields, "LoadBalancerNames")
	}
	return o
}

func (o *LaunchSpecification) SetLoadBalancersConfig(v *LoadBalancersConfig) *LaunchSpecification {
	if o.LoadBalancersConfig = v; o.LoadBalancersConfig == nil {
		o.nullFields = append(o.nullFields, "LoadBalancersConfig")
	}
	return o
}

func (o *LaunchSpecification) SetSecurityGroupIDs(v []string) *LaunchSpecification {
	if o.SecurityGroupIDs = v; o.SecurityGroupIDs == nil {
		o.nullFields = append(o.nullFields, "SecurityGroupIDs")
	}
	return o
}

func (o *LaunchSpecification) SetHealthCheckType(v *string) *LaunchSpecification {
	if o.HealthCheckType = v; o.HealthCheckType == nil {
		o.nullFields = append(o.nullFields, "HealthCheckType")
	}
	return o
}

func (o *LaunchSpecification) SetHealthCheckGracePeriod(v *int) *LaunchSpecification {
	if o.HealthCheckGracePeriod = v; o.HealthCheckGracePeriod == nil {
		o.nullFields = append(o.nullFields, "HealthCheckGracePeriod")
	}
	return o
}

func (o *LaunchSpecification) SetHealthCheckUnhealthyDurationBeforeReplacement(v *int) *LaunchSpecification {
	if o.HealthCheckUnhealthyDurationBeforeReplacement = v; o.HealthCheckUnhealthyDurationBeforeReplacement == nil {
		o.nullFields = append(o.nullFields, "HealthCheckUnhealthyDurationBeforeReplacement")
	}
	return o
}

func (o *LaunchSpecification) SetImageId(v *string) *LaunchSpecification {
	if o.ImageID = v; o.ImageID == nil {
		o.nullFields = append(o.nullFields, "ImageID")
	}
	return o
}

func (o *LaunchSpecification) SetKeyPair(v *string) *LaunchSpecification {
	if o.KeyPair = v; o.KeyPair == nil {
		o.nullFields = append(o.nullFields, "KeyPair")
	}
	return o
}

func (o *LaunchSpecification) SetUserData(v *string) *LaunchSpecification {
	if o.UserData = v; o.UserData == nil {
		o.nullFields = append(o.nullFields, "UserData")
	}
	return o
}

func (o *LaunchSpecification) SetShutdownScript(v *string) *LaunchSpecification {
	if o.ShutdownScript = v; o.ShutdownScript == nil {
		o.nullFields = append(o.nullFields, "ShutdownScript")
	}
	return o
}

func (o *LaunchSpecification) SetTenancy(v *string) *LaunchSpecification {
	if o.Tenancy = v; o.Tenancy == nil {
		o.nullFields = append(o.nullFields, "Tenancy")
	}
	return o
}

func (o *LaunchSpecification) SetMonitoring(v *bool) *LaunchSpecification {
	if o.Monitoring = v; o.Monitoring == nil {
		o.nullFields = append(o.nullFields, "Monitoring")
	}
	return o
}

func (o *LaunchSpecification) SetEBSOptimized(v *bool) *LaunchSpecification {
	if o.EBSOptimized = v; o.EBSOptimized == nil {
		o.nullFields = append(o.nullFields, "EBSOptimized")
	}
	return o
}

func (o *LaunchSpecification) SetIAMInstanceProfile(v *IAMInstanceProfile) *LaunchSpecification {
	if o.IAMInstanceProfile = v; o.IAMInstanceProfile == nil {
		o.nullFields = append(o.nullFields, "IAMInstanceProfile")
	}
	return o
}

func (o *LaunchSpecification) SetBlockDeviceMappings(v []*BlockDeviceMapping) *LaunchSpecification {
	if o.BlockDeviceMappings = v; o.BlockDeviceMappings == nil {
		o.nullFields = append(o.nullFields, "BlockDeviceMappings")
	}
	return o
}

func (o *LaunchSpecification) SetNetworkInterfaces(v []*NetworkInterface) *LaunchSpecification {
	if o.NetworkInterfaces = v; o.NetworkInterfaces == nil {
		o.nullFields = append(o.nullFields, "NetworkInterfaces")
	}
	return o
}

func (o *LaunchSpecification) SetTags(v []*Tag) *LaunchSpecification {
	if o.Tags = v; o.Tags == nil {
		o.nullFields = append(o.nullFields, "Tags")
	}
	return o
}

// endregion

// region LoadBalancersConfig

func (o *LoadBalancersConfig) MarshalJSON() ([]byte, error) {
	type noMethod LoadBalancersConfig
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *LoadBalancersConfig) SetLoadBalancers(v []*LoadBalancer) *LoadBalancersConfig {
	if o.LoadBalancers = v; o.LoadBalancers == nil {
		o.nullFields = append(o.nullFields, "LoadBalancers")
	}
	return o
}

// endregion

// region LoadBalancer

func (o *LoadBalancer) MarshalJSON() ([]byte, error) {
	type noMethod LoadBalancer
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *LoadBalancer) SetName(v *string) *LoadBalancer {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *LoadBalancer) SetArn(v *string) *LoadBalancer {
	if o.Arn = v; o.Arn == nil {
		o.nullFields = append(o.nullFields, "Arn")
	}
	return o
}

func (o *LoadBalancer) SetType(v *string) *LoadBalancer {
	if o.Type = v; o.Type == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

func (o *LoadBalancer) SetBalancerId(v *string) *LoadBalancer {
	if o.BalancerID = v; o.BalancerID == nil {
		o.nullFields = append(o.nullFields, "BalancerID")
	}
	return o
}

func (o *LoadBalancer) SetTargetSetId(v *string) *LoadBalancer {
	if o.TargetSetID = v; o.TargetSetID == nil {
		o.nullFields = append(o.nullFields, "TargetSetID")
	}
	return o
}

func (o *LoadBalancer) SetZoneAwareness(v *bool) *LoadBalancer {
	if o.ZoneAwareness = v; o.ZoneAwareness == nil {
		o.nullFields = append(o.nullFields, "ZoneAwareness")
	}
	return o
}

func (o *LoadBalancer) SetAutoWeight(v *bool) *LoadBalancer {
	if o.AutoWeight = v; o.AutoWeight == nil {
		o.nullFields = append(o.nullFields, "AutoWeight")
	}
	return o
}

// endregion

// region NetworkInterface

func (o *NetworkInterface) MarshalJSON() ([]byte, error) {
	type noMethod NetworkInterface
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *NetworkInterface) SetId(v *string) *NetworkInterface {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *NetworkInterface) SetDescription(v *string) *NetworkInterface {
	if o.Description = v; o.Description == nil {
		o.nullFields = append(o.nullFields, "Description")
	}
	return o
}

func (o *NetworkInterface) SetDeviceIndex(v *int) *NetworkInterface {
	if o.DeviceIndex = v; o.DeviceIndex == nil {
		o.nullFields = append(o.nullFields, "DeviceIndex")
	}
	return o
}

func (o *NetworkInterface) SetSecondaryPrivateIPAddressCount(v *int) *NetworkInterface {
	if o.SecondaryPrivateIPAddressCount = v; o.SecondaryPrivateIPAddressCount == nil {
		o.nullFields = append(o.nullFields, "SecondaryPrivateIPAddressCount")
	}
	return o
}

func (o *NetworkInterface) SetAssociatePublicIPAddress(v *bool) *NetworkInterface {
	if o.AssociatePublicIPAddress = v; o.AssociatePublicIPAddress == nil {
		o.nullFields = append(o.nullFields, "AssociatePublicIPAddress")
	}
	return o
}

func (o *NetworkInterface) SetDeleteOnTermination(v *bool) *NetworkInterface {
	if o.DeleteOnTermination = v; o.DeleteOnTermination == nil {
		o.nullFields = append(o.nullFields, "DeleteOnTermination")
	}
	return o
}

func (o *NetworkInterface) SetSecurityGroupsIDs(v []string) *NetworkInterface {
	if o.SecurityGroupsIDs = v; o.SecurityGroupsIDs == nil {
		o.nullFields = append(o.nullFields, "SecurityGroupsIDs")
	}
	return o
}

func (o *NetworkInterface) SetPrivateIPAddress(v *string) *NetworkInterface {
	if o.PrivateIPAddress = v; o.PrivateIPAddress == nil {
		o.nullFields = append(o.nullFields, "PrivateIPAddress")
	}
	return o
}

func (o *NetworkInterface) SetSubnetId(v *string) *NetworkInterface {
	if o.SubnetID = v; o.SubnetID == nil {
		o.nullFields = append(o.nullFields, "SubnetID")
	}
	return o
}

// endregion

// region BlockDeviceMapping

func (o *BlockDeviceMapping) MarshalJSON() ([]byte, error) {
	type noMethod BlockDeviceMapping
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *BlockDeviceMapping) SetDeviceName(v *string) *BlockDeviceMapping {
	if o.DeviceName = v; o.DeviceName == nil {
		o.nullFields = append(o.nullFields, "DeviceName")
	}
	return o
}

func (o *BlockDeviceMapping) SetVirtualName(v *string) *BlockDeviceMapping {
	if o.VirtualName = v; o.VirtualName == nil {
		o.nullFields = append(o.nullFields, "VirtualName")
	}
	return o
}

func (o *BlockDeviceMapping) SetEBS(v *EBS) *BlockDeviceMapping {
	if o.EBS = v; o.EBS == nil {
		o.nullFields = append(o.nullFields, "EBS")
	}
	return o
}

// endregion

// region EBS

func (o *EBS) MarshalJSON() ([]byte, error) {
	type noMethod EBS
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *EBS) SetDeleteOnTermination(v *bool) *EBS {
	if o.DeleteOnTermination = v; o.DeleteOnTermination == nil {
		o.nullFields = append(o.nullFields, "DeleteOnTermination")
	}
	return o
}

func (o *EBS) SetEncrypted(v *bool) *EBS {
	if o.Encrypted = v; o.Encrypted == nil {
		o.nullFields = append(o.nullFields, "Encrypted")
	}
	return o
}

func (o *EBS) SetKmsKeyId(v *string) *EBS {
	if o.KmsKeyId = v; o.KmsKeyId == nil {
		o.nullFields = append(o.nullFields, "KmsKeyId")
	}
	return o
}

func (o *EBS) SetSnapshotId(v *string) *EBS {
	if o.SnapshotID = v; o.SnapshotID == nil {
		o.nullFields = append(o.nullFields, "SnapshotID")
	}
	return o
}

func (o *EBS) SetVolumeType(v *string) *EBS {
	if o.VolumeType = v; o.VolumeType == nil {
		o.nullFields = append(o.nullFields, "VolumeType")
	}
	return o
}

func (o *EBS) SetVolumeSize(v *int) *EBS {
	if o.VolumeSize = v; o.VolumeSize == nil {
		o.nullFields = append(o.nullFields, "VolumeSize")
	}
	return o
}

func (o *EBS) SetIOPS(v *int) *EBS {
	if o.IOPS = v; o.IOPS == nil {
		o.nullFields = append(o.nullFields, "IOPS")
	}
	return o
}

// endregion

// region IAMInstanceProfile

func (o *IAMInstanceProfile) MarshalJSON() ([]byte, error) {
	type noMethod IAMInstanceProfile
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *IAMInstanceProfile) SetName(v *string) *IAMInstanceProfile {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *IAMInstanceProfile) SetArn(v *string) *IAMInstanceProfile {
	if o.Arn = v; o.Arn == nil {
		o.nullFields = append(o.nullFields, "Arn")
	}
	return o
}

// endregion

// region RollStrategy

func (o *RollStrategy) MarshalJSON() ([]byte, error) {
	type noMethod RollStrategy
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *RollStrategy) SetAction(v *string) *RollStrategy {
	if o.Action = v; o.Action == nil {
		o.nullFields = append(o.nullFields, "Action")
	}
	return o
}

func (o *RollStrategy) SetShouldDrainInstances(v *bool) *RollStrategy {
	if o.ShouldDrainInstances = v; o.ShouldDrainInstances == nil {
		o.nullFields = append(o.nullFields, "ShouldDrainInstances")
	}
	return o
}

// endregion

// region CodeDeployIntegration

func (o *CodeDeployIntegration) MarshalJSON() ([]byte, error) {
	type noMethod CodeDeployIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *CodeDeployIntegration) SetDeploymentGroups(v []*DeploymentGroup) *CodeDeployIntegration {
	if o.DeploymentGroups = v; o.DeploymentGroups == nil {
		o.nullFields = append(o.nullFields, "DeploymentGroups")
	}
	return o
}

func (o *CodeDeployIntegration) SetCleanUpOnFailure(v *bool) *CodeDeployIntegration {
	if o.CleanUpOnFailure = v; o.CleanUpOnFailure == nil {
		o.nullFields = append(o.nullFields, "CleanUpOnFailure")
	}
	return o
}

func (o *CodeDeployIntegration) SetTerminateInstanceOnFailure(v *bool) *CodeDeployIntegration {
	if o.TerminateInstanceOnFailure = v; o.TerminateInstanceOnFailure == nil {
		o.nullFields = append(o.nullFields, "TerminateInstanceOnFailure")
	}
	return o
}

// endregion

// region DeploymentGroup

func (o *DeploymentGroup) MarshalJSON() ([]byte, error) {
	type noMethod DeploymentGroup
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *DeploymentGroup) SetApplicationName(v *string) *DeploymentGroup {
	if o.ApplicationName = v; o.ApplicationName == nil {
		o.nullFields = append(o.nullFields, "ApplicationName")
	}
	return o
}

func (o *DeploymentGroup) SetDeploymentGroupName(v *string) *DeploymentGroup {
	if o.DeploymentGroupName = v; o.DeploymentGroupName == nil {
		o.nullFields = append(o.nullFields, "DeploymentGroupName")
	}
	return o
}

// endregion

// region OpsWorksIntegration

func (o *OpsWorksIntegration) MarshalJSON() ([]byte, error) {
	type noMethod OpsWorksIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *OpsWorksIntegration) SetLayerId(v *string) *OpsWorksIntegration {
	if o.LayerID = v; o.LayerID == nil {
		o.nullFields = append(o.nullFields, "LayerID")
	}
	return o
}

func (o *OpsWorksIntegration) SetStackType(v *string) *OpsWorksIntegration {
	if o.StackType = v; o.StackType == nil {
		o.nullFields = append(o.nullFields, "StackType")
	}
	return o
}

// endregion

// region Scale Request

type ScaleUpSpotItem struct {
	SpotInstanceRequestID *string `json:"spotInstanceRequestId,omitempty"`
	AvailabilityZone      *string `json:"availabilityZone,omitempty"`
	InstanceType          *string `json:"instanceType,omitempty"`
}

type ScaleUpOnDemandItem struct {
	InstanceID       *string `json:"instanceId,omitempty"`
	AvailabilityZone *string `json:"availabilityZone,omitempty"`
	InstanceType     *string `json:"instanceType,omitempty"`
}

type ScaleDownSpotItem struct {
	SpotInstanceRequestID *string `json:"spotInstanceRequestId,omitempty"`
}

type ScaleDownOnDemandItem struct {
	InstanceID *string `json:"instanceId,omitempty"`
}

type ScaleItem struct {
	NewSpotRequests    []*ScaleUpSpotItem       `json:"newSpotRequests,omitempty"`
	NewInstances       []*ScaleUpOnDemandItem   `json:"newInstances,omitempty"`
	VictimSpotRequests []*ScaleDownSpotItem     `json:"victimSpotRequests,omitempty"`
	VictimInstances    []*ScaleDownOnDemandItem `json:"victimInstances,omitempty"`
}

type ScaleGroupInput struct {
	GroupID    *string `json:"groupId,omitempty"`
	ScaleType  *string `json:"type,omitempty"`
	Adjustment *int    `json:"adjustment,omitempty"`
}

type ScaleGroupOutput struct {
	Items []*ScaleItem `json:"items"`
}

func scaleUpResponseFromJSON(in []byte) (*ScaleGroupOutput, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}

	var retVal ScaleGroupOutput
	retVal.Items = make([]*ScaleItem, len(rw.Response.Items))
	for i, rb := range rw.Response.Items {
		b, err := scaleUpItemFromJSON(rb)
		if err != nil {
			return nil, err
		}
		retVal.Items[i] = b
	}

	return &retVal, nil
}

func scaleUpItemFromJSON(in []byte) (*ScaleItem, error) {
	var rw *ScaleItem
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	return rw, nil
}

func scaleFromHttpResponse(resp *http.Response) (*ScaleGroupOutput, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return scaleUpResponseFromJSON(body)
}

func (s *ServiceOp) Scale(ctx context.Context, input *ScaleGroupInput) (*ScaleGroupOutput, error) {
	path, err := uritemplates.Expand("/aws/ec2/group/{groupId}/scale/{type}", uritemplates.Values{
		"groupId": spotinst.StringValue(input.GroupID),
		"type":    spotinst.StringValue(input.ScaleType),
	})
	if err != nil {
		return nil, err
	}

	// We do not need the ID anymore so let's drop it.
	input.GroupID = nil

	r := client.NewRequest(http.MethodPut, path)

	if input.Adjustment != nil {
		r.Params.Set("adjustment", strconv.Itoa(*input.Adjustment))
	}
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	output, err := scaleFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return output, err
}

//endregion
