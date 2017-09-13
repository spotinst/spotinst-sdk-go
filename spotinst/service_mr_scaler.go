package spotinst

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/jsonutil"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/uritemplates"
)
// AWSMrScalerService is an interface for interfacing with the AWSMrScaler
// endpoints of the Spotinst API.
type AWSMrScalerService interface {
	List(context.Context, *ListAWSMrScalerInput) (*ListAWSMrScalerOutput, error)
	Create(context.Context, *CreateAWSMrScalerInput) (*CreateAWSMrScalerOutput, error)
	Read(context.Context, *ReadAWSMrScalerInput) (*ReadAWSMrScalerOutput, error)
	Update(context.Context, *UpdateAWSMrScalerInput) (*UpdateAWSMrScalerOutput, error)
	Delete(context.Context, *DeleteAWSMrScalerInput) (*DeleteAWSMrScalerOutput, error)
}

// AWSMrScalerServiceOp handles communication with the balancer related methods
// of the Spotinst API.

type AWSMrScalerServiceOp struct {
	client *Client
}
//
var _ AWSMrScalerService = &AWSMrScalerServiceOp{}

type AWSMrScaler struct {
	ID          *string                 `json:"id,omitempty"`
	Name        *string                 `json:"name,omitempty"`
	Description *string                 `json:"description,omitempty"`
	Region      *string                 `json:"region,omitempty"`
	Strategy    *AWSMrScalerStrategy    `json:"strategy,omitempty"`
	Compute     *AWSMrScalerCompute     `json:"compute,omitempty"`
	Scaling     *AWSMrScalerScaling     `json:"scaling,omitempty"`
	CoreScaling *AWSMrScalerScaling     `json:"coreScaling,omitempty"`

	// forceSendFields is a list of field names (e.g. "Keys") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	forceSendFields []string `json:"-"`

	// nullFields is a list of field names (e.g. "Keys") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	nullFields []string `json:"-"`
}

type AWSMrScalerStrategy struct {
	Cloning  *AWSMrScalerStrategyCloning  `json:"cloning,omitempty"`
	Wrapping *AWSMrScalerStrategyWrapping `json:"wrapping,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSMrScalerStrategyCloning struct {
	OriginClusterID *string `json:"originClusterId,omitempty"`
	NumberOfRetries *string `json:"numberOfRetries,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSMrScalerStrategyWrapping struct {
	SourceClusterID     *string `json:"sourceClusterId,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSMrScalerCompute struct {
	AvailabilityZones []*AWSMrScalerComputeAvailabilityZones `json:"availabilityZones,omitempty"`
	Tags              []*AWSMrScalerComputeTags              `json:"tags,omitempty"`
	InstanceGroups    *AWSMrScalerComputeInstanceGroups     `json:"instanceGroups,omitempty"`
	Configurations    *AWSMrScalerComputeConfigurations     `json:"configurations,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSMrScalerComputeAvailabilityZones struct {
	Name      *string `json:"name,omitempty"`
	SubnetID  *string `json:"subnetId,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSMrScalerComputeTags struct {
	TagKey      *string `json:"tagKey,omitempty"`
	TagValue    *string `json:"tagValue,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSMrScalerComputeInstanceGroups struct {
	MasterGroup   *AWSMrScalerInstanceGroup `json:"masterGroup,omitempty"`
	CoreGroup     *AWSMrScalerInstanceGroup `json:"coreGroup,omitempty"`
	TaskGroup     *AWSMrScalerInstanceGroup `json:"taskGroup,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSMrScalerInstanceGroup struct {
	InstanceTypes  []*string                       `json:"instanceTypes,omitempty"`
	Target           *int                          `json:"target,omitempty"`
	LifeCycle        *string                       `json:"lifeCycle,omitempty"`
	EbsConfiguration *AwsMrScalerEbsConfiguration  `json:"ebsConfiguration,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsMrScalerEbsConfiguration struct {
	EbsOptimized    *string                             `json:"ebsOptimized,omitempty"`
	EbsBlockDeviceConfigs []*AwsMrScalerBlockDeviceConfig `json:"ebsBlockDeviceConfigs,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsMrScalerBlockDeviceConfig struct {
	VolumesPerInstance    *int `json:"volumesPerInstance,omitempty"`
	VolumeSpecification   *AwsMrScalerVolumeSpecification `json:"volumeSpecification,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsMrScalerVolumeSpecification struct {
	VolumeType *string `json:"volumeType,omitempty"`
	SizeInGB *string   `json:"sizeInGB,omitempty"`
	Iops *string       `json:"iops,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSMrScalerScaling struct {
	Up   []*AWSMrScalerScalingPolicy `json:"up,omitempty"`
	Down []*AWSMrScalerScalingPolicy `json:"down,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSMrScalerScalingPolicy struct {
	PolicyName        *string                              `json:"policyName,omitempty"`
	Namespace         *string                              `json:"namespace,omitempty"`
	MetricName        *string                              `json:"metricName,omitempty"`
	Dimensions        []*AWSMrScalerScalingPolicyDimension `json:"dimensions,omitempty"`
	Statistic         *string                              `json:"statistic,omitempty"`
	Unit              *string                              `json:"unit,omitempty"`
	Threshold         *float64                             `json:"threshold,omitempty"`
	Adjustment        *int                                 `json:"adjustment,omitempty"`
	MinTargetCapacity *int                                 `json:"minTargetCapacity,omitempty"`
	MaxTargetCapacity *int                                 `json:"maxTargetCapacity,omitempty"`
	Period            *int                                 `json:"period,omitempty"`
	EvaluationPeriods *int                                 `json:"evaluationPeriods,omitempty"`
	Cooldown          *int                                 `json:"cooldown,omitempty"`
	Action            *AWSMrScalerScalingPolicyAction      `json:"action,omitempty"`
	Operator          *string                              `json:"operator,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSMrScalerScalingPolicyAction struct {
	Type              *string `json:"type,omitempty"`
	Adjustment        *string `json:"adjustment,omitempty"`
	MinTargetCapacity *string `json:"minTargetCapacity,omitempty"`
	MaxTargetCapacity *string `json:"maxTargetCapacity,omitempty"`
	Target            *string `json:"target,omitempty"`
	Minimum           *string `json:"minimum,omitempty"`
	Maximum           *string `json:"maximum,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSMrScalerScalingPolicyDimension struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSMrScalerComputeConfigurations struct {
	File      *AwsMrScalerComputeConfigurationFile `json:"file,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsMrScalerComputeConfigurationFile struct {
	Bucket    *string `json:"bucket,omitempty"`
	Key       *string `json:"key,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type ListAWSMrScalerInput struct{}

type ListAWSMrScalerOutput struct {
	MrScalers []*AWSMrScaler `json:"mrScalers,omitempty"`
}

type CreateAWSMrScalerInput struct {
	MrScaler *AWSMrScaler `json:"mrScaler,omitempty"`
}

type CreateAWSMrScalerOutput struct {
	MrScaler *AWSMrScaler `json:"mrScaler,omitempty"`
}

type ReadAWSMrScalerInput struct {
	MrScalerID *string `json:"mrScalerId,omitempty"`
}

type ReadAWSMrScalerOutput struct {
	MrScaler *AWSMrScaler `json:"mrScaler,omitempty"`
}

type UpdateAWSMrScalerInput struct {
	MrScaler *AWSMrScaler `json:"mrScaler,omitempty"`
}

type UpdateAWSMrScalerOutput struct {
	MrScaler *AWSMrScaler `json:"mrScaler,omitempty"`
}

type DeleteAWSMrScalerInput struct {
	MrScalerID *string `json:"mrScalerId,omitempty"`
}

type DeleteAWSMrScalerOutput struct{}

type StatusAWSMrScalerInput struct {
	MrScalerID *string `json:"mrScalerId,omitempty"`
}

type StatusAWSMrScalerOutput struct {
	Instances []*AWSInstance `json:"instances,omitempty"`
}


func awsMrScalerFromJSON(in []byte) (*AWSMrScaler, error) {
	b := new(AWSMrScaler)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func AWSMrScalersFromJSON(in []byte) ([]*AWSMrScaler, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*AWSMrScaler, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := awsMrScalerFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func AWSMrScalersFromHttpResponse(resp *http.Response) ([]*AWSMrScaler, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return AWSMrScalersFromJSON(body)
}



//region AWSMrScaler
func (o *AWSMrScaler) MarshalJSON() ([]byte, error) {
	type noMethod AWSMrScaler
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw,o.forceSendFields,o.nullFields)
}

func (o *AWSMrScaler) SetID(v *string) *AWSMrScaler{
	if o.ID = v; v == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}


func (o *AWSMrScaler) SetName(v *string) *AWSMrScaler{
	if o.Name = v; v == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}


func (o *AWSMrScaler) SetDescription(v *string) *AWSMrScaler{
	if o.Description = v; v == nil {
		o.nullFields = append(o.nullFields, "Description")
	}
	return o
}


func (o *AWSMrScaler) SetRegion(v *string) *AWSMrScaler{
	if o.Region = v; v == nil {
		o.nullFields = append(o.nullFields, "Region")
	}
	return o
}


func (o *AWSMrScaler) SetStrategy(v *AWSMrScalerStrategy) *AWSMrScaler{
	if o.Strategy = v; v == nil {
		o.nullFields = append(o.nullFields, "Strategy")
	}
	return o
}


func (o *AWSMrScaler) SetCompute(v *AWSMrScalerCompute) *AWSMrScaler{
	if o.Compute = v; v == nil {
		o.nullFields = append(o.nullFields, "Compute")
	}
	return o
}


func (o *AWSMrScaler) SetScaling(v *AWSMrScalerScaling) *AWSMrScaler{
	if o.Scaling = v; v == nil {
		o.nullFields = append(o.nullFields, "Scaling")
	}
	return o
}


func (o *AWSMrScaler) SetCoreScaling(v *AWSMrScalerScaling) *AWSMrScaler{
	if o.CoreScaling = v; v == nil {
		o.nullFields = append(o.nullFields, "CoreScaling")
	}
	return o
}

//endregion

//region AWSMrScalerStrategy
func (o *AWSMrScalerStrategy) MarshalJSON() ([]byte, error) {
	type noMethod AWSMrScalerStrategy
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw,o.forceSendFields,o.nullFields)
}

func (o *AWSMrScalerStrategy) SetCloning(v *AWSMrScalerStrategyCloning) *AWSMrScalerStrategy{
	if o.Cloning = v; v == nil {
		o.nullFields = append(o.nullFields, "Cloning")
	}
	return o
}


func (o *AWSMrScalerStrategy) SetWrapping(v *AWSMrScalerStrategyWrapping) *AWSMrScalerStrategy{
	if o.Wrapping = v; v == nil {
		o.nullFields = append(o.nullFields, "Wrapping")
	}
	return o
}

//endregion

//region AWSMrScalerStrategyCloning
func (o *AWSMrScalerStrategyCloning) MarshalJSON() ([]byte, error) {
	type noMethod AWSMrScalerStrategyCloning
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw,o.forceSendFields,o.nullFields)
}

func (o *AWSMrScalerStrategyCloning) SetOriginClusterID(v *string) *AWSMrScalerStrategyCloning{
	if o.OriginClusterID = v; v == nil {
		o.nullFields = append(o.nullFields, "OriginClusterID")
	}
	return o
}


func (o *AWSMrScalerStrategyCloning) SetNumberOfRetries(v *string) *AWSMrScalerStrategyCloning{
	if o.NumberOfRetries = v; v == nil {
		o.nullFields = append(o.nullFields, "NumberOfRetries")
	}
	return o
}

//endregion

//region AWSMrScalerStrategyWrapping
func (o *AWSMrScalerStrategyWrapping) MarshalJSON() ([]byte, error) {
	type noMethod AWSMrScalerStrategyWrapping
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw,o.forceSendFields,o.nullFields)
}

func (o *AWSMrScalerStrategyWrapping) SetSourceClusterID(v *string) *AWSMrScalerStrategyWrapping{
	if o.SourceClusterID = v; v == nil {
		o.nullFields = append(o.nullFields, "SourceClusterID")
	}
	return o
}

//endregion

//region AWSMrScalerCompute
func (o *AWSMrScalerCompute) MarshalJSON() ([]byte, error) {
	type noMethod AWSMrScalerCompute
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw,o.forceSendFields,o.nullFields)
}

func (o *AWSMrScalerCompute) SetAvailabilityZones(v []*AWSMrScalerComputeAvailabilityZones) *AWSMrScalerCompute{
	if o.AvailabilityZones = v; v == nil {
		o.nullFields = append(o.nullFields, "AvailabilityZones")
	}
	return o
}


func (o *AWSMrScalerCompute) SetTags(v []*AWSMrScalerComputeTags) *AWSMrScalerCompute{
	if o.Tags = v; v == nil {
		o.nullFields = append(o.nullFields, "Tags")
	}
	return o
}


func (o *AWSMrScalerCompute) SetInstanceGroups(v *AWSMrScalerComputeInstanceGroups) *AWSMrScalerCompute{
	if o.InstanceGroups = v; v == nil {
		o.nullFields = append(o.nullFields, "InstanceGroups")
	}
	return o
}


func (o *AWSMrScalerCompute) SetConfigurations(v *AWSMrScalerComputeConfigurations) *AWSMrScalerCompute{
	if o.Configurations = v; v == nil {
		o.nullFields = append(o.nullFields, "Configurations")
	}
	return o
}

//endregion

//region AWSMrScalerComputeAvailabilityZones
func (o *AWSMrScalerComputeAvailabilityZones) MarshalJSON() ([]byte, error) {
	type noMethod AWSMrScalerComputeAvailabilityZones
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw,o.forceSendFields,o.nullFields)
}

func (o *AWSMrScalerComputeAvailabilityZones) SetName(v *string) *AWSMrScalerComputeAvailabilityZones{
	if o.Name = v; v == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}


func (o *AWSMrScalerComputeAvailabilityZones) SetSubnetID(v *string) *AWSMrScalerComputeAvailabilityZones{
	if o.SubnetID = v; v == nil {
		o.nullFields = append(o.nullFields, "SubnetID")
	}
	return o
}

//endregion

//region AWSMrScalerComputeTags
func (o *AWSMrScalerComputeTags) MarshalJSON() ([]byte, error) {
	type noMethod AWSMrScalerComputeTags
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw,o.forceSendFields,o.nullFields)
}

func (o *AWSMrScalerComputeTags) SetTagKey(v *string) *AWSMrScalerComputeTags{
	if o.TagKey = v; v == nil {
		o.nullFields = append(o.nullFields, "TagKey")
	}
	return o
}


func (o *AWSMrScalerComputeTags) SetTagValue(v *string) *AWSMrScalerComputeTags{
	if o.TagValue = v; v == nil {
		o.nullFields = append(o.nullFields, "TagValue")
	}
	return o
}

//endregion

//region AWSMrScalerComputeInstanceGroups
func (o *AWSMrScalerComputeInstanceGroups) MarshalJSON() ([]byte, error) {
	type noMethod AWSMrScalerComputeInstanceGroups
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw,o.forceSendFields,o.nullFields)
}

func (o *AWSMrScalerComputeInstanceGroups) SetMasterGroup(v *AWSMrScalerInstanceGroup) *AWSMrScalerComputeInstanceGroups{
	if o.MasterGroup = v; v == nil {
		o.nullFields = append(o.nullFields, "MasterGroup")
	}
	return o
}


func (o *AWSMrScalerComputeInstanceGroups) SetCoreGroup(v *AWSMrScalerInstanceGroup) *AWSMrScalerComputeInstanceGroups{
	if o.CoreGroup = v; v == nil {
		o.nullFields = append(o.nullFields, "CoreGroup")
	}
	return o
}


func (o *AWSMrScalerComputeInstanceGroups) SetTaskGroup(v *AWSMrScalerInstanceGroup) *AWSMrScalerComputeInstanceGroups{
	if o.TaskGroup = v; v == nil {
		o.nullFields = append(o.nullFields, "TaskGroup")
	}
	return o
}

//endregion

//region AWSMrScalerInstanceGroup
func (o *AWSMrScalerInstanceGroup) MarshalJSON() ([]byte, error) {
	type noMethod AWSMrScalerInstanceGroup
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw,o.forceSendFields,o.nullFields)
}

func (o *AWSMrScalerInstanceGroup) SetInstanceTypes(v []*string) *AWSMrScalerInstanceGroup{
	if o.InstanceTypes = v; v == nil {
		o.nullFields = append(o.nullFields, "InstanceTypes")
	}
	return o
}


func (o *AWSMrScalerInstanceGroup) SetTarget(v *int) *AWSMrScalerInstanceGroup{
	if o.Target = v; v == nil {
		o.nullFields = append(o.nullFields, "Target")
	}
	return o
}


func (o *AWSMrScalerInstanceGroup) SetLifeCycle(v *string) *AWSMrScalerInstanceGroup{
	if o.LifeCycle = v; v == nil {
		o.nullFields = append(o.nullFields, "LifeCycle")
	}
	return o
}


func (o *AWSMrScalerInstanceGroup) SetEbsConfiguration(v *AwsMrScalerEbsConfiguration) *AWSMrScalerInstanceGroup{
	if o.EbsConfiguration = v; v == nil {
		o.nullFields = append(o.nullFields, "EbsConfiguration")
	}
	return o
}

//endregion

//region AWSMrScalerScaling
func (o *AWSMrScalerScaling) MarshalJSON() ([]byte, error) {
	type noMethod AWSMrScalerScaling
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw,o.forceSendFields,o.nullFields)
}

func (o *AWSMrScalerScaling) SetUp(v []*AWSMrScalerScalingPolicy) *AWSMrScalerScaling{
	if o.Up = v; v == nil {
		o.nullFields = append(o.nullFields, "Up")
	}
	return o
}


func (o *AWSMrScalerScaling) SetDown(v []*AWSMrScalerScalingPolicy) *AWSMrScalerScaling{
	if o.Down = v; v == nil {
		o.nullFields = append(o.nullFields, "Down")
	}
	return o
}

//endregion

//region AWSMrScalerScalingPolicy
func (o *AWSMrScalerScalingPolicy) MarshalJSON() ([]byte, error) {
	type noMethod AWSMrScalerScalingPolicy
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw,o.forceSendFields,o.nullFields)
}

func (o *AWSMrScalerScalingPolicy) SetPolicyName(v *string) *AWSMrScalerScalingPolicy{
	if o.PolicyName = v; v == nil {
		o.nullFields = append(o.nullFields, "PolicyName")
	}
	return o
}


func (o *AWSMrScalerScalingPolicy) SetNamespace(v *string) *AWSMrScalerScalingPolicy{
	if o.Namespace = v; v == nil {
		o.nullFields = append(o.nullFields, "Namespace")
	}
	return o
}


func (o *AWSMrScalerScalingPolicy) SetMetricName(v *string) *AWSMrScalerScalingPolicy{
	if o.MetricName = v; v == nil {
		o.nullFields = append(o.nullFields, "MetricName")
	}
	return o
}


func (o *AWSMrScalerScalingPolicy) SetDimensions(v []*AWSMrScalerScalingPolicyDimension) *AWSMrScalerScalingPolicy{
	if o.Dimensions = v; v == nil {
		o.nullFields = append(o.nullFields, "Dimensions")
	}
	return o
}


func (o *AWSMrScalerScalingPolicy) SetStatistic(v *string) *AWSMrScalerScalingPolicy{
	if o.Statistic = v; v == nil {
		o.nullFields = append(o.nullFields, "Statistic")
	}
	return o
}


func (o *AWSMrScalerScalingPolicy) SetUnit(v *string) *AWSMrScalerScalingPolicy{
	if o.Unit = v; v == nil {
		o.nullFields = append(o.nullFields, "Unit")
	}
	return o
}


func (o *AWSMrScalerScalingPolicy) SetThreshold(v *float64) *AWSMrScalerScalingPolicy{
	if o.Threshold = v; v == nil {
		o.nullFields = append(o.nullFields, "Threshold")
	}
	return o
}


func (o *AWSMrScalerScalingPolicy) SetAdjustment(v *int) *AWSMrScalerScalingPolicy{
	if o.Adjustment = v; v == nil {
		o.nullFields = append(o.nullFields, "Adjustment")
	}
	return o
}


func (o *AWSMrScalerScalingPolicy) SetMinTargetCapacity(v *int) *AWSMrScalerScalingPolicy{
	if o.MinTargetCapacity = v; v == nil {
		o.nullFields = append(o.nullFields, "MinTargetCapacity")
	}
	return o
}


func (o *AWSMrScalerScalingPolicy) SetMaxTargetCapacity(v *int) *AWSMrScalerScalingPolicy{
	if o.MaxTargetCapacity = v; v == nil {
		o.nullFields = append(o.nullFields, "MaxTargetCapacity")
	}
	return o
}


func (o *AWSMrScalerScalingPolicy) SetPeriod(v *int) *AWSMrScalerScalingPolicy{
	if o.Period = v; v == nil {
		o.nullFields = append(o.nullFields, "Period")
	}
	return o
}


func (o *AWSMrScalerScalingPolicy) SetEvaluationPeriods(v *int) *AWSMrScalerScalingPolicy{
	if o.EvaluationPeriods = v; v == nil {
		o.nullFields = append(o.nullFields, "EvaluationPeriods")
	}
	return o
}


func (o *AWSMrScalerScalingPolicy) SetCooldown(v *int) *AWSMrScalerScalingPolicy{
	if o.Cooldown = v; v == nil {
		o.nullFields = append(o.nullFields, "Cooldown")
	}
	return o
}


func (o *AWSMrScalerScalingPolicy) SetAction(v *AWSMrScalerScalingPolicyAction) *AWSMrScalerScalingPolicy{
	if o.Action = v; v == nil {
		o.nullFields = append(o.nullFields, "Action")
	}
	return o
}


func (o *AWSMrScalerScalingPolicy) SetOperator(v *string) *AWSMrScalerScalingPolicy{
	if o.Operator = v; v == nil {
		o.nullFields = append(o.nullFields, "Operator")
	}
	return o
}

//endregion

//region AWSMrScalerScalingPolicyAction
func (o *AWSMrScalerScalingPolicyAction) MarshalJSON() ([]byte, error) {
	type noMethod AWSMrScalerScalingPolicyAction
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw,o.forceSendFields,o.nullFields)
}

func (o *AWSMrScalerScalingPolicyAction) SetType(v *string) *AWSMrScalerScalingPolicyAction{
	if o.Type = v; v == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}


func (o *AWSMrScalerScalingPolicyAction) SetAdjustment(v *string) *AWSMrScalerScalingPolicyAction{
	if o.Adjustment = v; v == nil {
		o.nullFields = append(o.nullFields, "Adjustment")
	}
	return o
}


func (o *AWSMrScalerScalingPolicyAction) SetMinTargetCapacity(v *string) *AWSMrScalerScalingPolicyAction{
	if o.MinTargetCapacity = v; v == nil {
		o.nullFields = append(o.nullFields, "MinTargetCapacity")
	}
	return o
}


func (o *AWSMrScalerScalingPolicyAction) SetMaxTargetCapacity(v *string) *AWSMrScalerScalingPolicyAction{
	if o.MaxTargetCapacity = v; v == nil {
		o.nullFields = append(o.nullFields, "MaxTargetCapacity")
	}
	return o
}


func (o *AWSMrScalerScalingPolicyAction) SetTarget(v *string) *AWSMrScalerScalingPolicyAction{
	if o.Target = v; v == nil {
		o.nullFields = append(o.nullFields, "Target")
	}
	return o
}


func (o *AWSMrScalerScalingPolicyAction) SetMinimum(v *string) *AWSMrScalerScalingPolicyAction{
	if o.Minimum = v; v == nil {
		o.nullFields = append(o.nullFields, "Minimum")
	}
	return o
}


func (o *AWSMrScalerScalingPolicyAction) SetMaximum(v *string) *AWSMrScalerScalingPolicyAction{
	if o.Maximum = v; v == nil {
		o.nullFields = append(o.nullFields, "Maximum")
	}
	return o
}

//endregion

//region AWSMrScalerScalingPolicyDimension
func (o *AWSMrScalerScalingPolicyDimension) MarshalJSON() ([]byte, error) {
	type noMethod AWSMrScalerScalingPolicyDimension
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw,o.forceSendFields,o.nullFields)
}

func (o *AWSMrScalerScalingPolicyDimension) SetName(v *string) *AWSMrScalerScalingPolicyDimension{
	if o.Name = v; v == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}


func (o *AWSMrScalerScalingPolicyDimension) SetValue(v *string) *AWSMrScalerScalingPolicyDimension{
	if o.Value = v; v == nil {
		o.nullFields = append(o.nullFields, "Value")
	}
	return o
}

//endregion

//region AWSMrScalerComputeConfigurations
func (o *AWSMrScalerComputeConfigurations) MarshalJSON() ([]byte, error) {
	type noMethod AWSMrScalerComputeConfigurations
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw,o.forceSendFields,o.nullFields)
}

func (o *AWSMrScalerComputeConfigurations) SetFile(v *AwsMrScalerComputeConfigurationFile) *AWSMrScalerComputeConfigurations{
	if o.File = v; v == nil {
		o.nullFields = append(o.nullFields, "File")
	}
	return o
}

//endregion

func (s *AWSMrScalerServiceOp) List(ctx context.Context, input *ListAWSMrScalerInput) (*ListAWSMrScalerOutput, error) {
	r := s.client.newRequest(ctx, "GET", "/aws/emr/mrScaler")
	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := AWSMrScalersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListAWSMrScalerOutput{MrScalers: gs}, nil
}

func (s *AWSMrScalerServiceOp) Create(ctx context.Context, input *CreateAWSMrScalerInput) (*CreateAWSMrScalerOutput, error) {
	r := s.client.newRequest(ctx, "POST", "/aws/emr/mrScaler")
	r.obj = input

	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := AWSMrScalersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateAWSMrScalerOutput)
	if len(gs) > 0 {
		output.MrScaler = gs[0]
	}

	return output, nil
}

func (s *AWSMrScalerServiceOp) Read(ctx context.Context, input *ReadAWSMrScalerInput) (*ReadAWSMrScalerOutput, error) {
	path, err := uritemplates.Expand("/aws/emr/mrScaler/{mrScalerId}", map[string]string{
		"mrScalerId": StringValue(input.MrScalerID),
	})
	if err != nil {
		return nil, err
	}

	r := s.client.newRequest(ctx, "GET", path)
	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := AWSMrScalersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadAWSMrScalerOutput)
	if len(gs) > 0 {
		output.MrScaler = gs[0]
	}

	return output, nil
}

func (s *AWSMrScalerServiceOp) Update(ctx context.Context, input *UpdateAWSMrScalerInput) (*UpdateAWSMrScalerOutput, error) {
	path, err := uritemplates.Expand("/aws/emr/mrScaler/{mrScalerId}", map[string]string{
		"mrScalerId": StringValue(input.MrScaler.ID),
	})
	if err != nil {
		return nil, err
	}

	// We do not need the ID anymore so let's drop it.
	input.MrScaler.ID = nil

	r := s.client.newRequest(ctx, "PUT", path)
	r.obj = input

	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := AWSMrScalersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(UpdateAWSMrScalerOutput)
	if len(gs) > 0 {
		output.MrScaler = gs[0]
	}

	return output, nil
}

func (s *AWSMrScalerServiceOp) Delete(ctx context.Context, input *DeleteAWSMrScalerInput) (*DeleteAWSMrScalerOutput, error) {
	path, err := uritemplates.Expand("/aws/emr/mrScaler/{mrScalerId}", map[string]string{
		"mrScalerId": StringValue(input.MrScalerID),
	})
	if err != nil {
		return nil, err
	}

	r := s.client.newRequest(ctx, "DELETE", path)
	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteAWSMrScalerOutput{}, nil
}

