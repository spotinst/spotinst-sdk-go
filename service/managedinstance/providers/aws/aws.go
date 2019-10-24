package aws

import (
	"context"
	"encoding/json"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/jsonutil"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/uritemplates"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
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

//func (p Product) String() string {
//	return ProductName[p]
//}

type MangedInstance struct {
	ID          *string      `json:"id,omitempty"`
	Name        *string      `json:"name,omitempty"`
	Description *string      `json:"description,omitempty"`
	Region      *string      `json:"region,omitempty"`
	Strategy    *Strategy    `json:"strategy,omitempty"`
	Persistence *Persistence `json:"persistence,omitempty"`
	HealthCheck *healthCheck `json:"healthCheck,omitempty"`
	Compute     *Compute     `json:"compute,omitempty"`
	Scheduling  *Scheduling  `json:"scheduling,omitempty"`
	Integration *Integration `json:"thirdPartiesIntegration,omitempty"`

	// Read-only fields.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

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

type Compute struct {
	LaunchSpecification *LaunchSpecification `json:"launchSpecification,omitempty"`
	ElasticIP           *string              `json:"elasticIps,omitempty"`
	PrivateIP           *string              `json:"privateIps,omitempty"`
	SubnetIDs           []string             `json:"subnetIds,omitempty"`
	vpcId               *string              `json:"vpcId,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type LaunchSpecification struct {
	InstanceTypes       *InstanceTypes       `json:"instanceTypes,omitempty"`
	SecurityGroupIDs    []string             `json:"securityGroupIds,omitempty"`
	ImageID             *string              `json:"imageId,omitempty"`
	KeyPair             *string              `json:"keyPair,omitempty"`
	UserData            *string              `json:"userData,omitempty"`
	ShutdownScript      *string              `json:"shutdownScript,omitempty"`
	Tenancy             *string              `json:"tenancy,omitempty"`
	Monitoring          *bool                `json:"monitoring,omitempty"`
	EBSOptimized        *bool                `json:"ebsOptimized,omitempty"`
	CreditSpecification *CreditSpecification `json:"creditSpecification,omitempty"`
	NetworkInterfaces   []*NetworkInterface  `json:"networkInterfaces,omitempty"`
	Tags                []*Tag               `json:"tags,omitempty"`
	IAMInstanceProfile  *IAMInstanceProfile  `json:"iamRole,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type CreditSpecification struct {
	CPUCredits *string `json:"cpuCredits,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type NetworkInterface struct {
	ID                       *string `json:"networkInterfaceId,omitempty"`
	DeviceIndex              *int    `json:"deviceIndex,omitempty"`
	AssociatePublicIPAddress *bool   `json:"associatePublicIpAddress,omitempty"`
	AssociateIPV6Address     *bool   `json:"associateIpv6Address,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type IAMInstanceProfile struct {
	Name *string `json:"name,omitempty"`
	Arn  *string `json:"arn,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type InstanceTypes struct {
	preferredType *string  `json:"preferredType,omitempty"`
	types         []string `json:"types,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Strategy struct {
	lifeCycle                *string       `json:"lifeCycle,omitempty"`
	Orientation              *string       `json:"orientation,omitempty"`
	DrainingTimeout          *int          `json:"drainingTimeout,omitempty"`
	FallbackToOnDemand       *bool         `json:"fallbackToOd,omitempty"`
	UtilizeReservedInstances *bool         `json:"utilizeReservedInstances,omitempty"`
	OptimizationWindows      []string      `json:"optimizationWindows,omitempty"`
	RevertToSpot             *RevertToSpot `json:"revertToSpot,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type RevertToSpot struct {
	PerformAt *string `json:"performAt,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Scheduling struct {
	Tasks []*Task `json:"tasks,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Task struct {
	IsEnabled      *bool   `json:"isEnabled,omitempty"`
	Type           *string `json:"taskType,omitempty"`
	Frequency      *string `json:"frequency,omitempty"`
	CronExpression *string `json:"cronExpression,omitempty"`
	StartTime      *string `json:"startTime,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Persistence struct {
	PersistBlockDevices *bool   `json:"shouldPersistBlockDevices,omitempty"`
	PersistRootDevice   *bool   `json:"shouldPersistRootDevice,omitempty"`
	PersistPrivateIp    *bool   `json:"shouldPersistPrivateIp,omitempty"`
	BlockDevicesMode    *string `json:"blockDevicesMode,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type healthCheck struct {
	HealthCheckType                               *string `json:"healthCheckType,omitempty"`
	HealthCheckGracePeriod                        *int    `json:"healthCheckGracePeriod,omitempty"`
	HealthCheckUnhealthyDurationBeforeReplacement *int    `json:"healthCheckUnhealthyDurationBeforeReplacement,omitempty"`
	AutoHealing                                   *bool   `json:"autoHealing,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Integration struct {
	LoadBalancersConfig *LoadBalancersConfig `json:"loadBalancersConfig,omitempty"`
	Route53             *Route53Integration  `json:"route53,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Route53Integration struct {
	Domains []*Domain `json:"domains,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Domain struct {
	HostedZoneID      *string      `json:"hostedZoneId,omitempty"`
	SpotinstAccountID *string      `json:"spotinstAccountId,omitempty"`
	RecordSets        []*RecordSet `json:"recordSets,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type RecordSet struct {
	UsePublicIP *bool   `json:"usePublicIp,omitempty"`
	Name        *string `json:"name,omitempty"`

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

///////////
type ListMangedInstancesInput struct{}

type ListMangedInstancesOutput struct {
	MangedInstances []*MangedInstance `json:"mangedInstances,omitempty"`
}
type CreateMangedInstanceInput struct {
	MangedInstance *MangedInstance `json:"mangedInstance,omitempty"`
}
type CreateMangedInstanceOutput struct {
	MangedInstance *MangedInstance `json:"mangedInstance,omitempty"`
}
type ReadMangedInstanceInput struct {
	MangedInstanceID *string `json:"mangedInstanceId,omitempty"`
}
type ReadMangedInstanceOutput struct {
	MangedInstance *MangedInstance `json:"mangedInstance,omitempty"`
}
type UpdateMangedInstanceInput struct {
	MangedInstance       *MangedInstance `json:"mangedInstance,omitempty"`
	ShouldResumeStateful *bool           `json:"-"`
	AutoApplyTags        *bool           `json:"-"`
}
type UpdateMangedInstanceOutput struct {
	MangedInstance *MangedInstance `json:"mangedInstance,omitempty"`
}
type DeleteMangedInstanceInput struct {
	MangedInstanceID *string `json:"mangedInstanceId,omitempty"`
}

type DeleteMangedInstanceOutput struct{}

func MangedInstancesFromHttpResponse(resp *http.Response) ([]*MangedInstance, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return MangedInstancesFromJSON(body)
}

func MangedInstanceFromJSON(in []byte) (*MangedInstance, error) {
	b := new(MangedInstance)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}
func MangedInstancesFromJSON(in []byte) ([]*MangedInstance, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*MangedInstance, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := MangedInstanceFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func (s *ServiceOp) List(ctx context.Context, input *ListMangedInstancesInput) (*ListMangedInstancesOutput, error) {
	r := client.NewRequest(http.MethodGet, "/aws/ec2/managedInstance") //todo change the path
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := MangedInstancesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListMangedInstancesOutput{MangedInstances: gs}, nil
}

func (s *ServiceOp) Create(ctx context.Context, input *CreateMangedInstanceInput) (*CreateMangedInstanceOutput, error) {
	r := client.NewRequest(http.MethodPost, "/aws/ec2/managedInstance") //todo change the path
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := MangedInstancesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateMangedInstanceOutput)
	if len(gs) > 0 {
		output.MangedInstance = gs[0]
	}

	return output, nil
}

func (s *ServiceOp) Read(ctx context.Context, input *ReadMangedInstanceInput) (*ReadMangedInstanceOutput, error) {
	path, err := uritemplates.Expand("/aws/ec2/managedInstance{managedInstance}", uritemplates.Values{ ///todo change the path
		"managedInstance": spotinst.StringValue(input.MangedInstanceID),
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

	gs, err := MangedInstancesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadMangedInstanceOutput)
	if len(gs) > 0 {
		output.MangedInstance = gs[0]
	}

	return output, nil
}

func (s *ServiceOp) Update(ctx context.Context, input *UpdateMangedInstanceInput) (*UpdateMangedInstanceOutput, error) {
	path, err := uritemplates.Expand("/aws/ec2/managedInstance/{managedInstance}", uritemplates.Values{ //todo change the path
		"managedInstance": spotinst.StringValue(input.MangedInstance.ID),
	})
	if err != nil {
		return nil, err
	}

	// We do NOT need the ID anymore, so let's drop it.
	input.MangedInstance.ID = nil

	r := client.NewRequest(http.MethodPut, path)
	r.Obj = input

	if input.AutoApplyTags != nil {
		r.Params.Set("autoApplyTags",
			strconv.FormatBool(spotinst.BoolValue(input.AutoApplyTags)))
	}

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := MangedInstancesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(UpdateMangedInstanceOutput)
	if len(gs) > 0 {
		output.MangedInstance = gs[0]
	}

	return output, nil
}
func (s *ServiceOp) Delete(ctx context.Context, input *DeleteMangedInstanceInput) (*DeleteMangedInstanceOutput, error) {
	path, err := uritemplates.Expand("/aws/ec2/managedInstance/{managedInstance}", uritemplates.Values{ // todo need to change the path
		"managedInstance": spotinst.StringValue(input.MangedInstanceID),
	})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodDelete, path)

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteMangedInstanceOutput{}, nil
}

///////////
// region Group

func (o MangedInstance) MarshalJSON() ([]byte, error) {
	type noMethod MangedInstance
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *MangedInstance) SetId(v *string) *MangedInstance {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *MangedInstance) SetName(v *string) *MangedInstance {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *MangedInstance) SetDescription(v *string) *MangedInstance {
	if o.Description = v; o.Description == nil {
		o.nullFields = append(o.nullFields, "Description")
	}
	return o
}

func (o *MangedInstance) SetRegion(v *string) *MangedInstance {
	if o.Region = v; o.Region == nil {
		o.nullFields = append(o.nullFields, "Region")
	}
	return o
}

func (o *MangedInstance) SetCompute(v *Compute) *MangedInstance {
	if o.Compute = v; o.Compute == nil {
		o.nullFields = append(o.nullFields, "Compute")
	}
	return o
}

func (o *MangedInstance) SetStrategy(v *Strategy) *MangedInstance {
	if o.Strategy = v; o.Strategy == nil {
		o.nullFields = append(o.nullFields, "Strategy")
	}
	return o
}
func (o *MangedInstance) SetScheduling(v *Scheduling) *MangedInstance {
	if o.Scheduling = v; o.Scheduling == nil {
		o.nullFields = append(o.nullFields, "Scheduling")
	}
	return o
}

func (o *MangedInstance) SetIntegration(v *Integration) *MangedInstance {
	if o.Integration = v; o.Integration == nil {
		o.nullFields = append(o.nullFields, "Integration")
	}
	return o
}

func (o *MangedInstance) SetPersistence(v *Persistence) *MangedInstance {
	if o.Persistence = v; o.Persistence == nil {
		o.nullFields = append(o.nullFields, "persistence")
	}
	return o
}

func (o *MangedInstance) SetHealthCheck(v *healthCheck) *MangedInstance {
	if o.HealthCheck = v; o.HealthCheck == nil {
		o.nullFields = append(o.nullFields, "healthCheck")
	}
	return o
}

// endregion
// region Integration

func (o Integration) MarshalJSON() ([]byte, error) {
	type noMethod Integration
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Integration) SetRoute53(v *Route53Integration) *Integration {
	if o.Route53 = v; o.Route53 == nil {
		o.nullFields = append(o.nullFields, "Route53")
	}
	return o
}

func (o *Integration) SetLoadBalancersConfig(v *LoadBalancersConfig) *Integration {
	if o.LoadBalancersConfig = v; o.LoadBalancersConfig == nil {
		o.nullFields = append(o.nullFields, "loadBalancersConfig")
	}
	return o
}

// endregion
// region Route53

func (o Route53Integration) MarshalJSON() ([]byte, error) {
	type noMethod Route53Integration
	raw := noMethod(o)
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

func (o Domain) MarshalJSON() ([]byte, error) {
	type noMethod Domain
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Domain) SetHostedZoneID(v *string) *Domain {
	if o.HostedZoneID = v; o.HostedZoneID == nil {
		o.nullFields = append(o.nullFields, "HostedZoneID")
	}
	return o
}

func (o *Domain) SetSpotinstAccountID(v *string) *Domain {
	if o.SpotinstAccountID = v; o.SpotinstAccountID == nil {
		o.nullFields = append(o.nullFields, "SpotinstAccountID")
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

func (o RecordSet) MarshalJSON() ([]byte, error) {
	type noMethod RecordSet
	raw := noMethod(o)
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
// region loadBalancersConfig

func (o LoadBalancersConfig) MarshalJSON() ([]byte, error) {
	type noMethod LoadBalancersConfig
	raw := noMethod(o)
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

func (o LoadBalancer) MarshalJSON() ([]byte, error) {
	type noMethod LoadBalancer
	raw := noMethod(o)
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

// region Scheduling

func (o Scheduling) MarshalJSON() ([]byte, error) {
	type noMethod Scheduling
	raw := noMethod(o)
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

func (o Task) MarshalJSON() ([]byte, error) {
	type noMethod Task
	raw := noMethod(o)
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

// endregion

// region Compute

func (o Compute) MarshalJSON() ([]byte, error) {
	type noMethod Compute
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Compute) SetLaunchSpecification(v *LaunchSpecification) *Compute {
	if o.LaunchSpecification = v; o.LaunchSpecification == nil {
		o.nullFields = append(o.nullFields, "LaunchSpecification")
	}
	return o
}

func (o *Compute) SetSubnetIDs(v []string) *Compute {
	if o.SubnetIDs = v; o.SubnetIDs == nil {
		o.nullFields = append(o.nullFields, "SubnetIDs")
	}
	return o
}

func (o *Compute) SetPrivateIP(v *string) *Compute {
	if o.PrivateIP = v; o.PrivateIP == nil {
		o.nullFields = append(o.nullFields, "PrivateIP")
	}

	return o
}

func (o *Compute) SetElasticIP(v *string) *Compute {
	if o.ElasticIP = v; o.ElasticIP == nil {
		o.nullFields = append(o.nullFields, "ElasticIPs")
	}
	return o
}

func (o *Compute) SetVpcId(v *string) *Compute {
	if o.vpcId = v; o.vpcId == nil {
		o.nullFields = append(o.nullFields, "vpcId")
	}
	return o
}

// endregion
// region LaunchSpecification

func (o LaunchSpecification) MarshalJSON() ([]byte, error) {
	type noMethod LaunchSpecification
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
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

func (o *LaunchSpecification) SetInstanceTypes(v *InstanceTypes) *LaunchSpecification {
	if o.InstanceTypes = v; o.InstanceTypes == nil {
		o.nullFields = append(o.nullFields, "InstanceTypes")
	}
	return o
}

func (o *LaunchSpecification) SetTenancy(v *string) *LaunchSpecification {
	if o.Tenancy = v; o.Tenancy == nil {
		o.nullFields = append(o.nullFields, "Tenancy")
	}
	return o
}

func (o *LaunchSpecification) SetIAMInstanceProfile(v *IAMInstanceProfile) *LaunchSpecification {
	if o.IAMInstanceProfile = v; o.IAMInstanceProfile == nil {
		o.nullFields = append(o.nullFields, "IAMInstanceProfile")
	}
	return o
}

func (o *LaunchSpecification) SetSecurityGroupIDs(v []string) *LaunchSpecification {
	if o.SecurityGroupIDs = v; o.SecurityGroupIDs == nil {
		o.nullFields = append(o.nullFields, "SecurityGroupIDs")
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

func (o *LaunchSpecification) SetCreditSpecification(v *CreditSpecification) *LaunchSpecification {
	if o.CreditSpecification = v; o.CreditSpecification == nil {
		o.nullFields = append(o.nullFields, "CreditSpecification")
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
// region NetworkInterface

func (o NetworkInterface) MarshalJSON() ([]byte, error) {
	type noMethod NetworkInterface
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *NetworkInterface) SetDeviceIndex(v *int) *NetworkInterface {
	if o.DeviceIndex = v; o.DeviceIndex == nil {
		o.nullFields = append(o.nullFields, "DeviceIndex")
	}
	return o
}

func (o *NetworkInterface) SetAssociatePublicIPAddress(v *bool) *NetworkInterface {
	if o.AssociatePublicIPAddress = v; o.AssociatePublicIPAddress == nil {
		o.nullFields = append(o.nullFields, "AssociatePublicIPAddress")
	}
	return o
}

func (o *NetworkInterface) SetAssociateIPV6Address(v *bool) *NetworkInterface {
	if o.AssociateIPV6Address = v; o.AssociateIPV6Address == nil {
		o.nullFields = append(o.nullFields, "AssociateIPV6Address")
	}
	return o
}

// endregion

// region CreditSpecification

func (o CreditSpecification) MarshalJSON() ([]byte, error) {
	type noMethod CreditSpecification
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *CreditSpecification) SetCPUCredits(v *string) *CreditSpecification {
	if o.CPUCredits = v; o.CPUCredits == nil {
		o.nullFields = append(o.nullFields, "CPUCredits")
	}
	return o
}

// endregion
// region iamRole

func (o IAMInstanceProfile) MarshalJSON() ([]byte, error) {
	type noMethod IAMInstanceProfile
	raw := noMethod(o)
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
// region iamRole

func (o InstanceTypes) MarshalJSON() ([]byte, error) {
	type noMethod InstanceTypes
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *InstanceTypes) SetPreferredType(v *string) *InstanceTypes {
	if o.preferredType = v; o.preferredType == nil {
		o.nullFields = append(o.nullFields, "preferredType")
	}
	return o
}

func (o *InstanceTypes) SetTypesn(v []string) *InstanceTypes {
	if o.types = v; o.types == nil {
		o.nullFields = append(o.nullFields, "types")
	}
	return o
}

// endregion
// region healthCheck

func (o healthCheck) MarshalJSON() ([]byte, error) {
	type noMethod healthCheck
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *healthCheck) SetGracePeriod(v *int) *healthCheck {
	if o.HealthCheckGracePeriod = v; o.HealthCheckGracePeriod == nil {
		o.nullFields = append(o.nullFields, "gracePeriod")
	}
	return o
}

func (o *healthCheck) SetunhealthyDuration(v *int) *healthCheck {
	if o.HealthCheckUnhealthyDurationBeforeReplacement = v; o.HealthCheckUnhealthyDurationBeforeReplacement == nil {
		o.nullFields = append(o.nullFields, "unhealthyDuration")
	}
	return o
}

func (o *healthCheck) SetType(v *string) *healthCheck {
	if o.HealthCheckType = v; o.HealthCheckType == nil {
		o.nullFields = append(o.nullFields, "type")
	}
	return o
}

func (o *healthCheck) SetAutoHealing(v *bool) *healthCheck {
	if o.AutoHealing = v; o.AutoHealing == nil {
		o.nullFields = append(o.nullFields, "autoHealing")
	}
	return o
}

// region Persistence

func (o Persistence) MarshalJSON() ([]byte, error) {
	type noMethod Persistence
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Persistence) SetBlockDevicesMode(v *string) *Persistence {
	if o.BlockDevicesMode = v; o.BlockDevicesMode == nil {
		o.nullFields = append(o.nullFields, "BlockDevicesMode")
	}
	return o
}

func (o *Persistence) SetPersistPrivateIP(v *bool) *Persistence {
	if o.PersistPrivateIp = v; o.PersistPrivateIp == nil {
		o.nullFields = append(o.nullFields, "persistPrivateIp")
	}
	return o
}

func (o *Persistence) SetShouldPersistRootDevice(v *bool) *Persistence {
	if o.PersistRootDevice = v; o.PersistRootDevice == nil {
		o.nullFields = append(o.nullFields, "persistRootDevice")
	}
	return o
}

func (o *Persistence) SetShouldPersistBlockDevices(v *bool) *Persistence {
	if o.PersistBlockDevices = v; o.PersistBlockDevices == nil {
		o.nullFields = append(o.nullFields, "persistBlockDevices")
	}
	return o
}

// endregion

// region Strategy

func (o Strategy) MarshalJSON() ([]byte, error) {
	type noMethod Strategy
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Strategy) SetDrainingTimeout(v *int) *Strategy {
	if o.DrainingTimeout = v; o.DrainingTimeout == nil {
		o.nullFields = append(o.nullFields, "DrainingTimeout")
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

func (o *Strategy) SetRevertToSpot(v *RevertToSpot) *Strategy {
	if o.RevertToSpot = v; o.RevertToSpot == nil {
		o.nullFields = append(o.nullFields, "RevertToSpot")
	}
	return o
}

func (o *Strategy) SetOptimizationWindows(v []string) *Strategy {
	if o.OptimizationWindows = v; o.OptimizationWindows == nil {
		o.nullFields = append(o.nullFields, "optimizationWindows")
	}
	return o
}

func (o *Strategy) SetOrientation(v *string) *Strategy {
	if o.Orientation = v; o.Orientation == nil {
		o.nullFields = append(o.nullFields, "orientation")
	}
	return o
}

func (o *Strategy) SetLifeCycle(v *string) *Strategy {
	if o.lifeCycle = v; o.lifeCycle == nil {
		o.nullFields = append(o.nullFields, "lifeCycle")
	}
	return o
}

// endregion
// region RevertToSpot

func (o RevertToSpot) MarshalJSON() ([]byte, error) {
	type noMethod RevertToSpot
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *RevertToSpot) SetPerformAt(v *string) *RevertToSpot {
	if o.PerformAt = v; o.PerformAt == nil {
		o.nullFields = append(o.nullFields, "PerformAt")
	}
	return o
}

// endregion
