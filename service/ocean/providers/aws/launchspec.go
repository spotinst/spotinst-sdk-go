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

type LaunchSpec struct {
	ID                       *string                            `json:"id,omitempty"`
	Name                     *string                            `json:"name,omitempty"`
	OceanID                  *string                            `json:"oceanId,omitempty"`
	ImageID                  *string                            `json:"imageId,omitempty"`
	UserData                 *string                            `json:"userData,omitempty"`
	RootVolumeSize           *int                               `json:"rootVolumeSize,omitempty"`
	SecurityGroupIDs         []string                           `json:"securityGroupIds,omitempty"`
	SubnetIDs                []string                           `json:"subnetIds,omitempty"`
	InstanceTypes            []string                           `json:"instanceTypes,omitempty"`
	PreferredSpotTypes       []string                           `json:"preferredSpotTypes,omitempty"`
	Strategy                 *LaunchSpecStrategy                `json:"strategy,omitempty"`
	ResourceLimits           *ResourceLimits                    `json:"resourceLimits,omitempty"`
	IAMInstanceProfile       *IAMInstanceProfile                `json:"iamInstanceProfile,omitempty"`
	AutoScale                *AutoScale                         `json:"autoScale,omitempty"`
	ElasticIPPool            *ElasticIPPool                     `json:"elasticIpPool,omitempty"`
	BlockDeviceMappings      []*BlockDeviceMapping              `json:"blockDeviceMappings,omitempty"`
	EphemeralStorage         *EphemeralStorage                  `json:"ephemeralStorage,omitempty"`
	Labels                   []*Label                           `json:"labels,omitempty"`
	Taints                   []*Taint                           `json:"taints,omitempty"`
	Tags                     []*Tag                             `json:"tags,omitempty"`
	AssociatePublicIPAddress *bool                              `json:"associatePublicIpAddress,omitempty"`
	RestrictScaleDown        *bool                              `json:"restrictScaleDown,omitempty"`
	LaunchSpecScheduling     *LaunchSpecScheduling              `json:"scheduling,omitempty"`
	InstanceMetadataOptions  *LaunchspecInstanceMetadataOptions `json:"instanceMetadataOptions,omitempty"`
	Images                   []*Images                          `json:"images,omitempty"`
	InstanceTypesFilters     *InstanceTypesFilters              `json:"instanceTypesFilters,omitempty"`
	PreferredOnDemandTypes   []string                           `json:"preferredOnDemandTypes,omitempty"`
	ReservedENIs             *int                               `json:"reservedENIs,omitempty"`
	InstanceStorePolicy      *InstanceStorePolicy               `json:"instanceStorePolicy,omitempty"`
	StartupTaints            []*StartupTaints                   `json:"startupTaints,omitempty"`

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

// region InstanceMetadataOptions

type LaunchspecInstanceMetadataOptions struct {
	HTTPTokens              *string `json:"httpTokens,omitempty"`
	HTTPPutResponseHopLimit *int    `json:"httpPutResponseHopLimit,omitempty"`

	forceSendFields []string
	nullFields      []string
}

func (o *LaunchspecInstanceMetadataOptions) SetHTTPTokens(v *string) *LaunchspecInstanceMetadataOptions {
	if o.HTTPTokens = v; o.HTTPTokens == nil {
		o.nullFields = append(o.nullFields, "HTTPTokens")
	}
	return o
}

func (o *LaunchspecInstanceMetadataOptions) SetHTTPPutResponseHopLimit(v *int) *LaunchspecInstanceMetadataOptions {
	if o.HTTPPutResponseHopLimit = v; o.HTTPPutResponseHopLimit == nil {
		o.nullFields = append(o.nullFields, "HTTPPutResponseHopLimit")
	}
	return o
}

func (o *LaunchSpec) SetLaunchspecInstanceMetadataOptions(v *LaunchspecInstanceMetadataOptions) *LaunchSpec {
	if o.InstanceMetadataOptions = v; o.InstanceMetadataOptions == nil {
		o.nullFields = append(o.nullFields, "InstanceMetadataOptions")
	}
	return o
}
func (o LaunchspecInstanceMetadataOptions) MarshalJSON() ([]byte, error) {
	type noMethod LaunchspecInstanceMetadataOptions
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

// endregion

type ResourceLimits struct {
	MaxInstanceCount *int `json:"maxInstanceCount,omitempty"`
	MinInstanceCount *int `json:"minInstanceCount,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type BlockDeviceMapping struct {
	DeviceName  *string `json:"deviceName,omitempty"`
	NoDevice    *string `json:"noDevice,omitempty"`
	VirtualName *string `json:"virtualName,omitempty"`
	EBS         *EBS    `json:"ebs,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type EBS struct {
	DeleteOnTermination *bool              `json:"deleteOnTermination,omitempty"`
	Encrypted           *bool              `json:"encrypted,omitempty"`
	KMSKeyID            *string            `json:"kmsKeyId,omitempty"`
	SnapshotID          *string            `json:"snapshotId,omitempty"`
	VolumeType          *string            `json:"volumeType,omitempty"`
	IOPS                *int               `json:"iops,omitempty"`
	VolumeSize          *int               `json:"volumeSize,omitempty"`
	Throughput          *int               `json:"throughput,omitempty"`
	DynamicVolumeSize   *DynamicVolumeSize `json:"dynamicVolumeSize,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type DynamicVolumeSize struct {
	BaseSize            *int    `json:"baseSize,omitempty"`
	SizePerResourceUnit *int    `json:"sizePerResourceUnit,omitempty"`
	Resource            *string `json:"resource,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Label struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Taint struct {
	Key    *string `json:"key,omitempty"`
	Value  *string `json:"value,omitempty"`
	Effect *string `json:"effect,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type AutoScale struct {
	Headrooms              []*AutoScaleHeadroom `json:"headrooms,omitempty"`
	AutoHeadroomPercentage *int                 `json:"autoHeadroomPercentage,omitempty"`
	Down                   *AutoScalerDownVNG   `json:"down,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type AutoScaleHeadroom struct {
	CPUPerUnit    *int `json:"cpuPerUnit,omitempty"`
	GPUPerUnit    *int `json:"gpuPerUnit,omitempty"`
	MemoryPerUnit *int `json:"memoryPerUnit,omitempty"`
	NumOfUnits    *int `json:"numOfUnits,omitempty"`

	forceSendFields []string
	nullFields      []string
}
type AutoScalerDownVNG struct {
	MaxScaleDownPercentage *float64 `json:"maxScaleDownPercentage,omitempty"`
	forceSendFields        []string
	nullFields             []string
}

func (o *AutoScalerDownVNG) SetMaxScaleDownPercentage(v *float64) *AutoScalerDownVNG {
	if o.MaxScaleDownPercentage = v; o.MaxScaleDownPercentage == nil {
		o.nullFields = append(o.nullFields, "MaxScaleDownPercentage")
	}
	return o
}

func (o *AutoScale) SetAutoScalerDownVNG(v *AutoScalerDownVNG) *AutoScale {
	if o.Down = v; o.Down == nil {
		o.nullFields = append(o.nullFields, "Down")
	}
	return o
}

func (o AutoScalerDownVNG) MarshalJSON() ([]byte, error) {
	type noMethod AutoScalerDownVNG
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

type ElasticIPPool struct {
	TagSelector *TagSelector `json:"tagSelector,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type TagSelector struct {
	Key   *string `json:"tagKey,omitempty"`
	Value *string `json:"tagValue,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type LaunchSpecStrategy struct {
	SpotPercentage           *int                   `json:"spotPercentage,omitempty"`
	DrainingTimeout          *int                   `json:"drainingTimeout,omitempty"`
	UtilizeCommitments       *bool                  `json:"utilizeCommitments,omitempty"`
	UtilizeReservedInstances *bool                  `json:"utilizeReservedInstances,omitempty"`
	Orientation              *LaunchSpecOrientation `json:"orientation,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type LaunchSpecOrientation struct {
	AvailabilityVsCost *string `json:"availabilityVsCost,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type LaunchSpecScheduling struct {
	Tasks         []*LaunchSpecTask        `json:"tasks,omitempty"`
	ShutdownHours *LaunchSpecShutdownHours `json:"shutdownHours,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type LaunchSpecTask struct {
	IsEnabled      *bool       `json:"isEnabled,omitempty"`
	CronExpression *string     `json:"cronExpression,omitempty"`
	TaskType       *string     `json:"taskType,omitempty"`
	Config         *TaskConfig `json:"config,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type LaunchSpecShutdownHours struct {
	IsEnabled   *bool    `json:"isEnabled,omitempty"`
	TimeWindows []string `json:"timeWindows,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type TaskConfig struct {
	TaskHeadrooms []*LaunchSpecTaskHeadroom `json:"headrooms,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type LaunchSpecTaskHeadroom struct {
	CPUPerUnit    *int `json:"cpuPerUnit,omitempty"`
	GPUPerUnit    *int `json:"gpuPerUnit,omitempty"`
	MemoryPerUnit *int `json:"memoryPerUnit,omitempty"`
	NumOfUnits    *int `json:"numOfUnits,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Images struct {
	ImageId *string `json:"id,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type InstanceTypesFilters struct {
	Categories            []string `json:"categories,omitempty"`
	DiskTypes             []string `json:"diskTypes,omitempty"`
	ExcludeFamilies       []string `json:"excludeFamilies,omitempty"`
	ExcludeMetal          *bool    `json:"excludeMetal,omitempty"`
	Hypervisor            []string `json:"hypervisor,omitempty"`
	IncludeFamilies       []string `json:"includeFamilies,omitempty"`
	IsEnaSupported        *bool    `json:"isEnaSupported,omitempty"`
	MaxGpu                *int     `json:"maxGpu,omitempty"`
	MaxMemoryGiB          *float64 `json:"maxMemoryGiB,omitempty"`
	MaxNetworkPerformance *int     `json:"maxNetworkPerformance,omitempty"`
	MaxVcpu               *int     `json:"maxVcpu,omitempty"`
	MinEnis               *int     `json:"minEnis,omitempty"`
	MinGpu                *int     `json:"minGpu,omitempty"`
	MinMemoryGiB          *float64 `json:"minMemoryGiB,omitempty"`
	MinNetworkPerformance *int     `json:"minNetworkPerformance,omitempty"`
	MinVcpu               *int     `json:"minVcpu,omitempty"`
	RootDeviceTypes       []string `json:"rootDeviceTypes,omitempty"`
	VirtualizationTypes   []string `json:"virtualizationTypes,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type EphemeralStorage struct {
	DeviceName *string `json:"deviceName,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ListLaunchSpecsInput struct {
	OceanID *string `json:"oceanId,omitempty"`
}

type ListLaunchSpecsOutput struct {
	LaunchSpecs []*LaunchSpec `json:"launchSpecs,omitempty"`
}

type CreateLaunchSpecInput struct {
	LaunchSpec   *LaunchSpec `json:"launchSpec,omitempty"`
	InitialNodes *int        `json:"-"`
}

type CreateLaunchSpecOutput struct {
	LaunchSpec *LaunchSpec `json:"launchSpec,omitempty"`
}

type ReadLaunchSpecInput struct {
	LaunchSpecID *string `json:"launchSpecId,omitempty"`
}

type ReadLaunchSpecOutput struct {
	LaunchSpec *LaunchSpec `json:"launchSpec,omitempty"`
}

type UpdateLaunchSpecInput struct {
	LaunchSpec *LaunchSpec `json:"launchSpec,omitempty"`
}

type UpdateLaunchSpecOutput struct {
	LaunchSpec *LaunchSpec `json:"launchSpec,omitempty"`
}

type DeleteLaunchSpecInput struct {
	LaunchSpecID *string `json:"launchSpecId,omitempty"`
	ForceDelete  *bool   `json:"-"`
	DeleteNodes  *bool   `json:"-"`
}

type DeleteLaunchSpecOutput struct{}

func launchSpecFromJSON(in []byte) (*LaunchSpec, error) {
	b := new(LaunchSpec)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func launchSpecsFromJSON(in []byte) ([]*LaunchSpec, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*LaunchSpec, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := launchSpecFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func launchSpecsFromHttpResponse(resp *http.Response) ([]*LaunchSpec, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return launchSpecsFromJSON(body)
}

func (s *ServiceOp) ListLaunchSpecs(ctx context.Context, input *ListLaunchSpecsInput) (*ListLaunchSpecsOutput, error) {
	r := client.NewRequest(http.MethodGet, "/ocean/aws/k8s/launchSpec")

	if input.OceanID != nil {
		r.Params.Set("oceanId", spotinst.StringValue(input.OceanID))
	}

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := launchSpecsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListLaunchSpecsOutput{LaunchSpecs: gs}, nil
}

func (s *ServiceOp) CreateLaunchSpec(ctx context.Context, input *CreateLaunchSpecInput) (*CreateLaunchSpecOutput, error) {
	r := client.NewRequest(http.MethodPost, "/ocean/aws/k8s/launchSpec")
	r.Obj = input

	if input.InitialNodes != nil {
		r.Params.Set("initialNodes", strconv.Itoa(spotinst.IntValue(input.InitialNodes)))
	}

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := launchSpecsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateLaunchSpecOutput)
	if len(gs) > 0 {
		output.LaunchSpec = gs[0]
	}

	return output, nil
}

func (s *ServiceOp) ReadLaunchSpec(ctx context.Context, input *ReadLaunchSpecInput) (*ReadLaunchSpecOutput, error) {
	path, err := uritemplates.Expand("/ocean/aws/k8s/launchSpec/{launchSpecId}", uritemplates.Values{
		"launchSpecId": spotinst.StringValue(input.LaunchSpecID),
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

	gs, err := launchSpecsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadLaunchSpecOutput)
	if len(gs) > 0 {
		output.LaunchSpec = gs[0]
	}

	return output, nil
}

func (s *ServiceOp) UpdateLaunchSpec(ctx context.Context, input *UpdateLaunchSpecInput) (*UpdateLaunchSpecOutput, error) {
	path, err := uritemplates.Expand("/ocean/aws/k8s/launchSpec/{launchSpecId}", uritemplates.Values{
		"launchSpecId": spotinst.StringValue(input.LaunchSpec.ID),
	})
	if err != nil {
		return nil, err
	}

	// We do not need the ID anymore so let's drop it.
	input.LaunchSpec.ID = nil

	r := client.NewRequest(http.MethodPut, path)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := launchSpecsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(UpdateLaunchSpecOutput)
	if len(gs) > 0 {
		output.LaunchSpec = gs[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteLaunchSpec(ctx context.Context, input *DeleteLaunchSpecInput) (*DeleteLaunchSpecOutput, error) {
	path, err := uritemplates.Expand("/ocean/aws/k8s/launchSpec/{launchSpecId}", uritemplates.Values{
		"launchSpecId": spotinst.StringValue(input.LaunchSpecID),
	})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodDelete, path)

	if input.ForceDelete != nil {
		r.Params.Set("forceDelete", strconv.FormatBool(spotinst.BoolValue(input.ForceDelete)))
	}
	if input.DeleteNodes != nil {
		r.Params.Set("deleteNodes", strconv.FormatBool(spotinst.BoolValue(input.DeleteNodes)))
	}

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteLaunchSpecOutput{}, nil
}

// endregion

// region LaunchSpec

func (o LaunchSpec) MarshalJSON() ([]byte, error) {
	type noMethod LaunchSpec
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *LaunchSpec) SetId(v *string) *LaunchSpec {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *LaunchSpec) SetName(v *string) *LaunchSpec {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *LaunchSpec) SetOceanId(v *string) *LaunchSpec {
	if o.OceanID = v; o.OceanID == nil {
		o.nullFields = append(o.nullFields, "OceanID")
	}
	return o
}

func (o *LaunchSpec) SetImageId(v *string) *LaunchSpec {
	if o.ImageID = v; o.ImageID == nil {
		o.nullFields = append(o.nullFields, "ImageID")
	}
	return o
}

func (o *LaunchSpec) SetUserData(v *string) *LaunchSpec {
	if o.UserData = v; o.UserData == nil {
		o.nullFields = append(o.nullFields, "UserData")
	}
	return o
}

func (o *LaunchSpec) SetSecurityGroupIDs(v []string) *LaunchSpec {
	if o.SecurityGroupIDs = v; o.SecurityGroupIDs == nil {
		o.nullFields = append(o.nullFields, "SecurityGroupIDs")
	}
	return o
}

func (o *LaunchSpec) SetSubnetIDs(v []string) *LaunchSpec {
	if o.SubnetIDs = v; o.SubnetIDs == nil {
		o.nullFields = append(o.nullFields, "SubnetIDs")
	}
	return o
}

func (o *LaunchSpec) SetInstanceTypes(v []string) *LaunchSpec {
	if o.InstanceTypes = v; o.InstanceTypes == nil {
		o.nullFields = append(o.nullFields, "InstanceTypes")
	}
	return o
}

func (o *LaunchSpec) SetPreferredSpotTypes(v []string) *LaunchSpec {
	if o.PreferredSpotTypes = v; o.PreferredSpotTypes == nil {
		o.nullFields = append(o.nullFields, "PreferredSpotTypes")
	}
	return o
}

func (o *LaunchSpec) SetPreferredOnDemandTypes(v []string) *LaunchSpec {
	if o.PreferredOnDemandTypes = v; o.PreferredOnDemandTypes == nil {
		o.nullFields = append(o.nullFields, "PreferredOnDemandTypes")
	}
	return o
}

func (o *LaunchSpec) SetRootVolumeSize(v *int) *LaunchSpec {
	if o.RootVolumeSize = v; o.RootVolumeSize == nil {
		o.nullFields = append(o.nullFields, "RootVolumeSize")
	}
	return o
}

func (o *LaunchSpec) SetIAMInstanceProfile(v *IAMInstanceProfile) *LaunchSpec {
	if o.IAMInstanceProfile = v; o.IAMInstanceProfile == nil {
		o.nullFields = append(o.nullFields, "IAMInstanceProfile")
	}
	return o
}

func (o *LaunchSpec) SetLabels(v []*Label) *LaunchSpec {
	if o.Labels = v; o.Labels == nil {
		o.nullFields = append(o.nullFields, "Labels")
	}
	return o
}

func (o *LaunchSpec) SetTaints(v []*Taint) *LaunchSpec {
	if o.Taints = v; o.Taints == nil {
		o.nullFields = append(o.nullFields, "Taints")
	}
	return o
}

func (o *LaunchSpec) SetAutoScale(v *AutoScale) *LaunchSpec {
	if o.AutoScale = v; o.AutoScale == nil {
		o.nullFields = append(o.nullFields, "AutoScale")
	}
	return o
}

func (o *LaunchSpec) SetElasticIPPool(v *ElasticIPPool) *LaunchSpec {
	if o.ElasticIPPool = v; o.ElasticIPPool == nil {
		o.nullFields = append(o.nullFields, "ElasticIPPool")
	}
	return o
}

func (o *LaunchSpec) SetBlockDeviceMappings(v []*BlockDeviceMapping) *LaunchSpec {
	if o.BlockDeviceMappings = v; o.BlockDeviceMappings == nil {
		o.nullFields = append(o.nullFields, "BlockDeviceMappings")
	}
	return o
}

func (o *LaunchSpec) SetTags(v []*Tag) *LaunchSpec {
	if o.Tags = v; o.Tags == nil {
		o.nullFields = append(o.nullFields, "Tags")
	}
	return o
}

func (o *LaunchSpec) SetResourceLimits(v *ResourceLimits) *LaunchSpec {
	if o.ResourceLimits = v; o.ResourceLimits == nil {
		o.nullFields = append(o.nullFields, "ResourceLimits")
	}
	return o
}

func (o *LaunchSpec) SetStrategy(v *LaunchSpecStrategy) *LaunchSpec {
	if o.Strategy = v; o.Strategy == nil {
		o.nullFields = append(o.nullFields, "Strategy")
	}
	return o
}

func (o *LaunchSpec) SetAssociatePublicIPAddress(v *bool) *LaunchSpec {
	if o.AssociatePublicIPAddress = v; o.AssociatePublicIPAddress == nil {
		o.nullFields = append(o.nullFields, "AssociatePublicIPAddress")
	}
	return o
}

func (o *LaunchSpec) SetRestrictScaleDown(v *bool) *LaunchSpec {
	if o.RestrictScaleDown = v; o.RestrictScaleDown == nil {
		o.nullFields = append(o.nullFields, "RestrictScaleDown")
	}
	return o
}

func (o *LaunchSpec) SetScheduling(v *LaunchSpecScheduling) *LaunchSpec {
	if o.LaunchSpecScheduling = v; o.LaunchSpecScheduling == nil {
		o.nullFields = append(o.nullFields, "scheduling")
	}
	return o
}
func (o *LaunchSpec) SetInstanceTypesFilters(v *InstanceTypesFilters) *LaunchSpec {
	if o.InstanceTypesFilters = v; o.InstanceTypesFilters == nil {
		o.nullFields = append(o.nullFields, "InstanceTypesFilters")
	}
	return o
}
func (o *LaunchSpec) SetEphemeralStorage(v *EphemeralStorage) *LaunchSpec {
	if o.EphemeralStorage = v; o.EphemeralStorage == nil {
		o.nullFields = append(o.nullFields, "EphemeralStorage")
	}
	return o
}

func (o *LaunchSpec) SetReservedENIs(v *int) *LaunchSpec {
	if o.ReservedENIs = v; o.ReservedENIs == nil {
		o.nullFields = append(o.nullFields, "ReservedENIs")
	}
	return o
}

func (o *LaunchSpec) SetInstanceStorePolicy(v *InstanceStorePolicy) *LaunchSpec {
	if o.InstanceStorePolicy = v; o.InstanceStorePolicy == nil {
		o.nullFields = append(o.nullFields, "InstanceStorePolicy")
	}
	return o
}

func (o *LaunchSpec) SetStartupTaints(v []*StartupTaints) *LaunchSpec {
	if o.StartupTaints = v; o.StartupTaints == nil {
		o.nullFields = append(o.nullFields, "StartupTaints")
	}
	return o
}

// endregion

// region BlockDeviceMapping

func (o BlockDeviceMapping) MarshalJSON() ([]byte, error) {
	type noMethod BlockDeviceMapping
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *BlockDeviceMapping) SetDeviceName(v *string) *BlockDeviceMapping {
	if o.DeviceName = v; o.DeviceName == nil {
		o.nullFields = append(o.nullFields, "DeviceName")
	}
	return o
}

func (o *BlockDeviceMapping) SetNoDevice(v *string) *BlockDeviceMapping {
	if o.NoDevice = v; o.NoDevice == nil {
		o.nullFields = append(o.nullFields, "NoDevice")
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

func (o EBS) MarshalJSON() ([]byte, error) {
	type noMethod EBS
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *EBS) SetEncrypted(v *bool) *EBS {
	if o.Encrypted = v; o.Encrypted == nil {
		o.nullFields = append(o.nullFields, "Encrypted")
	}
	return o
}

func (o *EBS) SetIOPS(v *int) *EBS {
	if o.IOPS = v; o.IOPS == nil {
		o.nullFields = append(o.nullFields, "IOPS")
	}
	return o
}

func (o *EBS) SetKMSKeyId(v *string) *EBS {
	if o.KMSKeyID = v; o.KMSKeyID == nil {
		o.nullFields = append(o.nullFields, "KMSKeyID")
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

func (o *EBS) SetDeleteOnTermination(v *bool) *EBS {
	if o.DeleteOnTermination = v; o.DeleteOnTermination == nil {
		o.nullFields = append(o.nullFields, "DeleteOnTermination")
	}
	return o
}

func (o *EBS) SetVolumeSize(v *int) *EBS {
	if o.VolumeSize = v; o.VolumeSize == nil {
		o.nullFields = append(o.nullFields, "VolumeSize")
	}
	return o
}

func (o *EBS) SetDynamicVolumeSize(v *DynamicVolumeSize) *EBS {
	if o.DynamicVolumeSize = v; o.DynamicVolumeSize == nil {
		o.nullFields = append(o.nullFields, "DynamicVolumeSize")
	}
	return o
}

func (o *EBS) SetThroughput(v *int) *EBS {
	if o.Throughput = v; o.Throughput == nil {
		o.nullFields = append(o.nullFields, "Throughput")
	}
	return o
}

// endregion

// region DynamicVolumeSize

func (o DynamicVolumeSize) MarshalJSON() ([]byte, error) {
	type noMethod DynamicVolumeSize
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *DynamicVolumeSize) SetBaseSize(v *int) *DynamicVolumeSize {
	if o.BaseSize = v; o.BaseSize == nil {
		o.nullFields = append(o.nullFields, "BaseSize")
	}
	return o
}

func (o *DynamicVolumeSize) SetResource(v *string) *DynamicVolumeSize {
	if o.Resource = v; o.Resource == nil {
		o.nullFields = append(o.nullFields, "Resource")
	}
	return o
}

func (o *DynamicVolumeSize) SetSizePerResourceUnit(v *int) *DynamicVolumeSize {
	if o.SizePerResourceUnit = v; o.SizePerResourceUnit == nil {
		o.nullFields = append(o.nullFields, "SizePerResourceUnit")
	}
	return o
}

// endregion

// region ResourceLimits

func (o ResourceLimits) MarshalJSON() ([]byte, error) {
	type noMethod ResourceLimits
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *ResourceLimits) SetMaxInstanceCount(v *int) *ResourceLimits {
	if o.MaxInstanceCount = v; o.MaxInstanceCount == nil {
		o.nullFields = append(o.nullFields, "MaxInstanceCount")
	}
	return o
}

func (o *ResourceLimits) SetMinInstanceCount(v *int) *ResourceLimits {
	if o.MinInstanceCount = v; o.MinInstanceCount == nil {
		o.nullFields = append(o.nullFields, "MinInstanceCount")
	}
	return o
}

// endregion

// region Label

func (o Label) MarshalJSON() ([]byte, error) {
	type noMethod Label
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Label) SetKey(v *string) *Label {
	if o.Key = v; o.Key == nil {
		o.nullFields = append(o.nullFields, "Key")
	}
	return o
}

func (o *Label) SetValue(v *string) *Label {
	if o.Value = v; o.Value == nil {
		o.nullFields = append(o.nullFields, "Value")
	}
	return o
}

// endregion

// region Taints

func (o Taint) MarshalJSON() ([]byte, error) {
	type noMethod Taint
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Taint) SetKey(v *string) *Taint {
	if o.Key = v; o.Key == nil {
		o.nullFields = append(o.nullFields, "Key")
	}
	return o
}

func (o *Taint) SetValue(v *string) *Taint {
	if o.Value = v; o.Value == nil {
		o.nullFields = append(o.nullFields, "Value")
	}
	return o
}

func (o *Taint) SetEffect(v *string) *Taint {
	if o.Effect = v; o.Effect == nil {
		o.nullFields = append(o.nullFields, "Effect")
	}
	return o
}

// endregion

// region AutoScale

func (o AutoScale) MarshalJSON() ([]byte, error) {
	type noMethod AutoScale
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AutoScale) SetHeadrooms(v []*AutoScaleHeadroom) *AutoScale {
	if o.Headrooms = v; o.Headrooms == nil {
		o.nullFields = append(o.nullFields, "Headrooms")
	}
	return o
}

func (o *AutoScale) SetAutoHeadroomPercentage(v *int) *AutoScale {
	if o.AutoHeadroomPercentage = v; o.AutoHeadroomPercentage == nil {
		o.nullFields = append(o.nullFields, "AutoHeadroomPercentage")
	}
	return o
}

// endregion

// region AutoScaleHeadroom

func (o AutoScaleHeadroom) MarshalJSON() ([]byte, error) {
	type noMethod AutoScaleHeadroom
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AutoScaleHeadroom) SetCPUPerUnit(v *int) *AutoScaleHeadroom {
	if o.CPUPerUnit = v; o.CPUPerUnit == nil {
		o.nullFields = append(o.nullFields, "CPUPerUnit")
	}
	return o
}

func (o *AutoScaleHeadroom) SetGPUPerUnit(v *int) *AutoScaleHeadroom {
	if o.GPUPerUnit = v; o.GPUPerUnit == nil {
		o.nullFields = append(o.nullFields, "GPUPerUnit")
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

// region ElasticIPPool

func (o ElasticIPPool) MarshalJSON() ([]byte, error) {
	type noMethod ElasticIPPool
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *ElasticIPPool) SetTagSelector(v *TagSelector) *ElasticIPPool {
	if o.TagSelector = v; o.TagSelector == nil {
		o.nullFields = append(o.nullFields, "TagSelector")
	}
	return o
}

// endregion

// region TagSelector

func (o TagSelector) MarshalJSON() ([]byte, error) {
	type noMethod TagSelector
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *TagSelector) SetTagKey(v *string) *TagSelector {
	if o.Key = v; o.Key == nil {
		o.nullFields = append(o.nullFields, "Key")
	}
	return o
}

func (o *TagSelector) SetTagValue(v *string) *TagSelector {
	if o.Value = v; o.Value == nil {
		o.nullFields = append(o.nullFields, "Value")
	}
	return o
}

// endregion

// region Strategy

func (o LaunchSpecStrategy) MarshalJSON() ([]byte, error) {
	type noMethod LaunchSpecStrategy
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *LaunchSpecStrategy) SetSpotPercentage(v *int) *LaunchSpecStrategy {
	if o.SpotPercentage = v; o.SpotPercentage == nil {
		o.nullFields = append(o.nullFields, "SpotPercentage")
	}
	return o
}

func (o *LaunchSpecStrategy) SetDrainingTimeout(v *int) *LaunchSpecStrategy {
	if o.DrainingTimeout = v; o.DrainingTimeout == nil {
		o.nullFields = append(o.nullFields, "DrainingTimeout")
	}
	return o
}

func (o *LaunchSpecStrategy) SetUtilizeCommitments(v *bool) *LaunchSpecStrategy {
	if o.UtilizeCommitments = v; o.UtilizeCommitments == nil {
		o.nullFields = append(o.nullFields, "UtilizeCommitments")
	}
	return o
}

func (o *LaunchSpecStrategy) SetUtilizeReservedInstances(v *bool) *LaunchSpecStrategy {
	if o.UtilizeReservedInstances = v; o.UtilizeReservedInstances == nil {
		o.nullFields = append(o.nullFields, "UtilizeReservedInstances")
	}
	return o
}

func (o *LaunchSpecStrategy) SetOrientation(v *LaunchSpecOrientation) *LaunchSpecStrategy {
	if o.Orientation = v; o.Orientation == nil {
		o.nullFields = append(o.nullFields, "Orientation")
	}
	return o
}

// endregion

func (o LaunchSpecOrientation) MarshalJSON() ([]byte, error) {
	type noMethod LaunchSpecOrientation
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *LaunchSpecOrientation) SetAvailabilityVsCost(v *string) *LaunchSpecOrientation {
	if o.AvailabilityVsCost = v; o.AvailabilityVsCost == nil {
		o.nullFields = append(o.nullFields, "AvailabilityVsCost")
	}
	return o
}

//region Scheduling

func (o LaunchSpecScheduling) MarshalJSON() ([]byte, error) {
	type noMethod LaunchSpecScheduling
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *LaunchSpecScheduling) SetTasks(v []*LaunchSpecTask) *LaunchSpecScheduling {
	if o.Tasks = v; o.Tasks == nil {
		o.nullFields = append(o.nullFields, "Tasks")
	}
	return o
}

func (o *LaunchSpecScheduling) SetShutdownHours(v *LaunchSpecShutdownHours) *LaunchSpecScheduling {
	if o.ShutdownHours = v; o.ShutdownHours == nil {
		o.nullFields = append(o.nullFields, "ShutdownHours")
	}
	return o
}

// endregion

//region LaunchSpecTask

func (o LaunchSpecTask) MarshalJSON() ([]byte, error) {
	type noMethod LaunchSpecTask
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *LaunchSpecTask) SetIsEnabled(v *bool) *LaunchSpecTask {
	if o.IsEnabled = v; o.IsEnabled == nil {
		o.nullFields = append(o.nullFields, "IsEnabled")
	}
	return o
}

func (o *LaunchSpecTask) SetCronExpression(v *string) *LaunchSpecTask {
	if o.CronExpression = v; o.CronExpression == nil {
		o.nullFields = append(o.nullFields, "CronExpression")
	}
	return o
}

func (o *LaunchSpecTask) SetTaskType(v *string) *LaunchSpecTask {
	if o.TaskType = v; o.TaskType == nil {
		o.nullFields = append(o.nullFields, "TaskType")
	}
	return o
}

func (o *LaunchSpecTask) SetTaskConfig(v *TaskConfig) *LaunchSpecTask {
	if o.Config = v; o.Config == nil {
		o.nullFields = append(o.nullFields, "Config")
	}
	return o
}

// endregion

//region LaunchSpecShutdownHours

func (o LaunchSpecShutdownHours) MarshalJSON() ([]byte, error) {
	type noMethod LaunchSpecShutdownHours
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *LaunchSpecShutdownHours) SetIsEnabled(v *bool) *LaunchSpecShutdownHours {
	if o.IsEnabled = v; o.IsEnabled == nil {
		o.nullFields = append(o.nullFields, "IsEnabled")
	}
	return o
}

func (o *LaunchSpecShutdownHours) SetTimeWindows(v []string) *LaunchSpecShutdownHours {
	if o.TimeWindows = v; o.TimeWindows == nil {
		o.nullFields = append(o.nullFields, "TimeWindows")
	}
	return o
}

//region TaskConfig

func (o TaskConfig) MarshalJSON() ([]byte, error) {
	type noMethod TaskConfig
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *TaskConfig) SetHeadrooms(v []*LaunchSpecTaskHeadroom) *TaskConfig {
	if o.TaskHeadrooms = v; o.TaskHeadrooms == nil {
		o.nullFields = append(o.nullFields, "TaskHeadrooms")
	}
	return o
}

// endregion

// region LaunchSpecTaskHeadroom

func (o LaunchSpecTaskHeadroom) MarshalJSON() ([]byte, error) {
	type noMethod LaunchSpecTaskHeadroom
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *LaunchSpecTaskHeadroom) SetCPUPerUnit(v *int) *LaunchSpecTaskHeadroom {
	if o.CPUPerUnit = v; o.CPUPerUnit == nil {
		o.nullFields = append(o.nullFields, "CPUPerUnit")
	}
	return o
}

func (o *LaunchSpecTaskHeadroom) SetGPUPerUnit(v *int) *LaunchSpecTaskHeadroom {
	if o.GPUPerUnit = v; o.GPUPerUnit == nil {
		o.nullFields = append(o.nullFields, "GPUPerUnit")
	}
	return o
}

func (o *LaunchSpecTaskHeadroom) SetMemoryPerUnit(v *int) *LaunchSpecTaskHeadroom {
	if o.MemoryPerUnit = v; o.MemoryPerUnit == nil {
		o.nullFields = append(o.nullFields, "MemoryPerUnit")
	}
	return o
}

func (o *LaunchSpecTaskHeadroom) SetNumOfUnits(v *int) *LaunchSpecTaskHeadroom {
	if o.NumOfUnits = v; o.NumOfUnits == nil {
		o.nullFields = append(o.nullFields, "NumOfUnits")
	}
	return o
}

func (o *LaunchSpec) SetImages(v []*Images) *LaunchSpec {
	if o.Images = v; o.Images == nil {
		o.nullFields = append(o.nullFields, "Images")
	}
	return o
}

func (o Images) MarshalJSON() ([]byte, error) {
	type noMethod Images
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Images) SetImageId(v *string) *Images {
	if o.ImageId = v; o.ImageId == nil {
		o.nullFields = append(o.nullFields, "ImageId")
	}
	return o
}

// endregion

// region InstanceTypesFilters

func (o InstanceTypesFilters) MarshalJSON() ([]byte, error) {
	type noMethod InstanceTypesFilters
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *InstanceTypesFilters) SetCategories(v []string) *InstanceTypesFilters {
	if o.Categories = v; o.Categories == nil {
		o.nullFields = append(o.nullFields, "Categories")
	}
	return o
}

func (o *InstanceTypesFilters) SetDiskTypes(v []string) *InstanceTypesFilters {
	if o.DiskTypes = v; o.DiskTypes == nil {
		o.nullFields = append(o.nullFields, "DiskTypes")
	}
	return o
}

func (o *InstanceTypesFilters) SetExcludeFamilies(v []string) *InstanceTypesFilters {
	if o.ExcludeFamilies = v; o.ExcludeFamilies == nil {
		o.nullFields = append(o.nullFields, "ExcludeFamilies")
	}
	return o
}

func (o *InstanceTypesFilters) SetExcludeMetal(v *bool) *InstanceTypesFilters {
	if o.ExcludeMetal = v; o.ExcludeMetal == nil {
		o.nullFields = append(o.nullFields, "ExcludeMetal")
	}
	return o
}

func (o *InstanceTypesFilters) SetHypervisor(v []string) *InstanceTypesFilters {
	if o.Hypervisor = v; o.Hypervisor == nil {
		o.nullFields = append(o.nullFields, "Hypervisor")
	}
	return o
}

func (o *InstanceTypesFilters) SetIncludeFamilies(v []string) *InstanceTypesFilters {
	if o.IncludeFamilies = v; o.IncludeFamilies == nil {
		o.nullFields = append(o.nullFields, "IncludeFamilies")
	}
	return o
}

func (o *InstanceTypesFilters) SetIsEnaSupported(v *bool) *InstanceTypesFilters {
	if o.IsEnaSupported = v; o.IsEnaSupported == nil {
		o.nullFields = append(o.nullFields, "IsEnaSupported")
	}
	return o
}

func (o *InstanceTypesFilters) SetMaxGpu(v *int) *InstanceTypesFilters {
	if o.MaxGpu = v; o.MaxGpu == nil {
		o.nullFields = append(o.nullFields, "MaxGpu")
	}
	return o
}

func (o *InstanceTypesFilters) SetMaxMemoryGiB(v *float64) *InstanceTypesFilters {
	if o.MaxMemoryGiB = v; o.MaxMemoryGiB == nil {
		o.nullFields = append(o.nullFields, "MaxMemoryGiB")
	}
	return o
}

func (o *InstanceTypesFilters) SetMaxNetworkPerformance(v *int) *InstanceTypesFilters {
	if o.MaxNetworkPerformance = v; o.MaxNetworkPerformance == nil {
		o.nullFields = append(o.nullFields, "MaxNetworkPerformance")
	}
	return o
}

func (o *InstanceTypesFilters) SetMaxVcpu(v *int) *InstanceTypesFilters {
	if o.MaxVcpu = v; o.MaxVcpu == nil {
		o.nullFields = append(o.nullFields, "MaxVcpu")
	}
	return o
}

func (o *InstanceTypesFilters) SetMinEnis(v *int) *InstanceTypesFilters {
	if o.MinEnis = v; o.MinEnis == nil {
		o.nullFields = append(o.nullFields, "MinEnis")
	}
	return o
}

func (o *InstanceTypesFilters) SetMinGpu(v *int) *InstanceTypesFilters {
	if o.MinGpu = v; o.MinGpu == nil {
		o.nullFields = append(o.nullFields, "MinGpu")
	}
	return o
}

func (o *InstanceTypesFilters) SetMinMemoryGiB(v *float64) *InstanceTypesFilters {
	if o.MinMemoryGiB = v; o.MinMemoryGiB == nil {
		o.nullFields = append(o.nullFields, "MinMemoryGiB")
	}
	return o
}

func (o *InstanceTypesFilters) SetMinNetworkPerformance(v *int) *InstanceTypesFilters {
	if o.MinNetworkPerformance = v; o.MinNetworkPerformance == nil {
		o.nullFields = append(o.nullFields, "MinNetworkPerformance")
	}
	return o
}

func (o *InstanceTypesFilters) SetMinVcpu(v *int) *InstanceTypesFilters {
	if o.MinVcpu = v; o.MinVcpu == nil {
		o.nullFields = append(o.nullFields, "MinVcpu")
	}
	return o
}

func (o *InstanceTypesFilters) SetRootDeviceTypes(v []string) *InstanceTypesFilters {
	if o.RootDeviceTypes = v; o.RootDeviceTypes == nil {
		o.nullFields = append(o.nullFields, "RootDeviceTypes")
	}
	return o
}

func (o *InstanceTypesFilters) SetVirtualizationTypes(v []string) *InstanceTypesFilters {
	if o.VirtualizationTypes = v; o.VirtualizationTypes == nil {
		o.nullFields = append(o.nullFields, "VirtualizationTypes")
	}
	return o
}

// endregion

// region EphemeralStorage

func (o *EphemeralStorage) SetDeviceName(v *string) *EphemeralStorage {
	if o.DeviceName = v; o.DeviceName == nil {
		o.nullFields = append(o.nullFields, "DeviceName")
	}
	return o
}

func (o EphemeralStorage) MarshalJSON() ([]byte, error) {
	type noMethod EphemeralStorage
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

// endregion
