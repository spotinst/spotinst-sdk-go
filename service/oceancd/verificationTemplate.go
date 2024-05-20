package oceancd

import (
	"context"
	"encoding/json"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/jsonutil"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/uritemplates"
)

type VerificationTemplate struct {
	Args    []*Args    `json:"args,omitempty"`
	Metrics []*Metrics `json:"metrics,omitempty"`
	Name    *string    `json:"name,omitempty"`

	// Read-only fields.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Args struct {
	Name      *string    `json:"name,omitempty"`
	Value     *string    `json:"value,omitempty"`
	ValueFrom *ValueFrom `json:"valueFrom,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ValueFrom struct {
	SecretKeyRef *SecretKeyRef `json:"secretKeyRef,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type SecretKeyRef struct {
	Key  *string `json:"key,omitempty"`
	Name *string `json:"name,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Metrics struct {
	Baseline              *Baseline `json:"baseline,omitempty"`
	ConsecutiveErrorLimit *int      `json:"consecutiveErrorLimit,omitempty"`
	Count                 *int      `json:"count,omitempty"`
	DryRun                *bool     `json:"dryRun,omitempty"`
	FailureCondition      *string   `json:"failureCondition,omitempty"`
	FailureLimit          *int      `json:"failureLimit,omitempty"`
	InitialDelay          *string   `json:"initialDelay,omitempty"`
	Interval              *string   `json:"interval,omitempty"`
	Name                  *string   `json:"name,omitempty"`
	Provider              *Provider `json:"provider,omitempty"`
	SuccessCondition      *string   `json:"successCondition,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Baseline struct {
	MaxRange  *int      `json:"maxRange,omitempty"`
	MinRange  *int      `json:"minRange,omitempty"`
	Provider  *Provider `json:"provider,omitempty"`
	Threshold *string   `json:"threshold,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Provider struct {
	CloudWatch *CloudWatchProvider `json:"cloudWatch,omitempty"`
	Datadog    *DataDogProvider    `json:"datadog,omitempty"`
	Jenkins    *JenkinsProvider    `json:"jenkins,omitempty"`
	NewRelic   *NewRelicProvider   `json:"newRelic,omitempty"`
	Prometheus *PrometheusProvider `json:"prometheus,omitempty"`
	Job        *Job                `json:"job,omitempty"`
	Web        *Web                `json:"web,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type CloudWatchProvider struct {
	Duration          *string              `json:"duration,omitempty"`
	MetricDataQueries []*MetricDataQueries `json:"metricDataQueries,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type MetricDataQueries struct {
	Expression *string     `json:"expression,omitempty"`
	ID         *string     `json:"id,omitempty"`
	Label      *string     `json:"label,omitempty"`
	MetricStat *MetricStat `json:"metricStat,omitempty"`
	Period     *int        `json:"period,omitempty"`
	ReturnData *bool       `json:"returnData,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type MetricStat struct {
	Metric *Metric `json:"metric,omitempty"`
	Period *int    `json:"period,omitempty"`
	Stat   *string `json:"stat,omitempty"`
	Unit   *string `json:"unit,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Metric struct {
	Dimensions []*Dimensions `json:"dimensions,omitempty"`
	MetricName *string       `json:"metricName,omitempty"`
	Namespace  *string       `json:"namespace,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Dimensions struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type DataDogProvider struct {
	Duration *string `json:"duration,omitempty"`
	Query    *string `json:"query,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Job struct {
	Spec *Spec `json:"spec,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Spec struct {
	BackoffLimit *int      `json:"backoffLimit,omitempty"`
	Template     *Template `json:"template,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Template struct {
	Spec *TemplateSpec `json:"spec,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type TemplateSpec struct {
	RestartPolicy *string       `json:"restartPolicy,omitempty"`
	Containers    []*Containers `json:"containers,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Containers struct {
	Command []string `json:"command,omitempty"`
	Image   *string  `json:"image,omitempty"`
	Name    *string  `json:"name,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type JenkinsProvider struct {
	Interval        *string       `json:"interval,omitempty"`
	PipelineName    *string       `json:"pipelineName,omitempty"`
	Parameters      []*Parameters `json:"parameters,omitempty"`
	Timeout         *string       `json:"timeout,omitempty"`
	TLSVerification *bool         `json:"tlsVerification,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Parameters struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type NewRelicProvider struct {
	Profile *string `json:"profile,omitempty"`
	Query   *string `json:"query,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type PrometheusProvider struct {
	Query *string `json:"query,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Web struct {
	Body           *string    `json:"body,omitempty"`
	Insecure       *bool      `json:"insecure,omitempty"`
	Headers        []*Headers `json:"headers,omitempty"`
	JsonPath       *string    `json:"jsonPath,omitempty"`
	Method         *string    `json:"method,omitempty"`
	TimeoutSeconds *int       `json:"timeoutSeconds,omitempty"`
	Url            *string    `json:"url,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Headers struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ListVerificationTemplatesOutput struct {
	VerificationTemplate []*VerificationTemplate `json:"verificationTemplate,omitempty"`
}

type CreateVerificationTemplateInput struct {
	VerificationTemplate *VerificationTemplate `json:"verificationTemplate,omitempty"`
}

type CreateVerificationTemplateOutput struct {
	VerificationTemplate *VerificationTemplate `json:"verificationTemplate,omitempty"`
}

type ReadVerificationTemplateInput struct {
	Name *string `json:"name,omitempty"`
}

type ReadVerificationTemplateOutput struct {
	VerificationTemplate *VerificationTemplate `json:"verificationTemplate,omitempty"`
}

type UpdateVerificationTemplateInput struct {
	VerificationTemplate *VerificationTemplate `json:"verificationTemplate,omitempty"`
}

type UpdateVerificationTemplateOutput struct {
	VerificationTemplate *VerificationTemplate `json:"verificationTemplate,omitempty"`
}

type PatchVerificationTemplateInput struct {
	VerificationTemplate *VerificationTemplate `json:"verificationTemplate,omitempty"`
}

type PatchVerificationTemplateOutput struct {
	VerificationTemplate *VerificationTemplate `json:"verificationTemplate,omitempty"`
}

type DeleteVerificationTemplateInput struct {
	Name *string `json:"name,omitempty"`
}

type DeleteVerificationTemplateOutput struct{}

func verificationTemplateFromJSON(in []byte) (*VerificationTemplate, error) {
	b := new(VerificationTemplate)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func verificationTemplatesFromJSON(in []byte) ([]*VerificationTemplate, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*VerificationTemplate, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := verificationTemplateFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func verificationTemplatesFromHttpResponse(resp *http.Response) ([]*VerificationTemplate, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return verificationTemplatesFromJSON(body)
}

// endregion

// region API requests

func (s *ServiceOp) ListVerificationTemplates(ctx context.Context) (*ListVerificationTemplatesOutput, error) {
	r := client.NewRequest(http.MethodGet, "/ocean/cd/verificationTemplate")
	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	vt, err := verificationTemplatesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListVerificationTemplatesOutput{VerificationTemplate: vt}, nil
}

func (s *ServiceOp) CreateVerificationTemplate(ctx context.Context, input *CreateVerificationTemplateInput) (*CreateVerificationTemplateOutput, error) {
	r := client.NewRequest(http.MethodPost, "/ocean/cd/verificationTemplate")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	vt, err := verificationTemplatesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateVerificationTemplateOutput)
	if len(vt) > 0 {
		output.VerificationTemplate = vt[0]
	}

	return output, nil
}

func (s *ServiceOp) ReadVerificationTemplate(ctx context.Context, input *ReadVerificationTemplateInput) (*ReadVerificationTemplateOutput, error) {
	path, err := uritemplates.Expand("/ocean/cd/verificationTemplate/{name}", uritemplates.Values{
		"name": spotinst.StringValue(input.Name),
	})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodGet, path)
	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	vt, err := verificationTemplatesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadVerificationTemplateOutput)
	if len(vt) > 0 {
		output.VerificationTemplate = vt[0]
	}

	return output, nil
}

func (s *ServiceOp) UpdateVerificationTemplate(ctx context.Context, input *UpdateVerificationTemplateInput) (*UpdateVerificationTemplateOutput, error) {
	path, err := uritemplates.Expand("/ocean/cd/verificationTemplate/{name}", uritemplates.Values{
		"name": spotinst.StringValue(input.VerificationTemplate.Name),
	})
	if err != nil {
		return nil, err
	}

	// We do NOT need the Name anymore, so let's drop it.
	input.VerificationTemplate.Name = nil

	r := client.NewRequest(http.MethodPut, path)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	vt, err := verificationTemplatesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(UpdateVerificationTemplateOutput)
	if len(vt) > 0 {
		output.VerificationTemplate = vt[0]
	}

	return output, nil
}

func (s *ServiceOp) PatchVerificationTemplate(ctx context.Context, input *PatchVerificationTemplateInput) (*PatchVerificationTemplateOutput, error) {
	path, err := uritemplates.Expand("/ocean/cd/verificationTemplate/{name}", uritemplates.Values{
		"name": spotinst.StringValue(input.VerificationTemplate.Name),
	})
	if err != nil {
		return nil, err
	}

	// We do NOT need the Name anymore, so let's drop it.
	input.VerificationTemplate.Name = nil

	r := client.NewRequest(http.MethodPatch, path)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	vt, err := verificationTemplatesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(PatchVerificationTemplateOutput)
	if len(vt) > 0 {
		output.VerificationTemplate = vt[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteVerificationTemplate(ctx context.Context, input *DeleteVerificationTemplateInput) (*DeleteVerificationTemplateOutput, error) {
	path, err := uritemplates.Expand("/ocean/cd/verificationTemplate/{name}", uritemplates.Values{
		"name": spotinst.StringValue(input.Name),
	})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodDelete, path)
	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteVerificationTemplateOutput{}, nil
}

//end region

// region Verification Template

func (o VerificationTemplate) MarshalJSON() ([]byte, error) {
	type noMethod VerificationTemplate
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *VerificationTemplate) SetArgs(v []*Args) *VerificationTemplate {
	if o.Args = v; o.Args == nil {
		o.nullFields = append(o.nullFields, "Args")
	}
	return o
}

func (o *VerificationTemplate) SetMetrics(v []*Metrics) *VerificationTemplate {
	if o.Metrics = v; o.Metrics == nil {
		o.nullFields = append(o.nullFields, "Metrics")
	}
	return o
}

func (o *VerificationTemplate) SetName(v *string) *VerificationTemplate {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

// end region

//region Args

func (o Args) MarshalJSON() ([]byte, error) {
	type noMethod Args
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Args) SetName(v *string) *Args {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *Args) SetValue(v *string) *Args {
	if o.Value = v; o.Value == nil {
		o.nullFields = append(o.nullFields, "Value")
	}
	return o
}

func (o *Args) SetValueFrom(v *ValueFrom) *Args {
	if o.ValueFrom = v; o.ValueFrom == nil {
		o.nullFields = append(o.nullFields, "ValueFrom")
	}
	return o
}

//end region

//region ValueFrom

func (o ValueFrom) MarshalJSON() ([]byte, error) {
	type noMethod ValueFrom
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *ValueFrom) SetSecretKeyRef(v *SecretKeyRef) *ValueFrom {
	if o.SecretKeyRef = v; o.SecretKeyRef == nil {
		o.nullFields = append(o.nullFields, "SecretKeyRef")
	}
	return o
}

//end region

//region SecretKeyRef

func (o SecretKeyRef) MarshalJSON() ([]byte, error) {
	type noMethod SecretKeyRef
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *SecretKeyRef) SetKey(v *string) *SecretKeyRef {
	if o.Key = v; o.Key == nil {
		o.nullFields = append(o.nullFields, "Key")
	}
	return o
}

func (o *SecretKeyRef) SetName(v *string) *SecretKeyRef {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

//end region

//region Metrics

func (o Metrics) MarshalJSON() ([]byte, error) {
	type noMethod Metrics
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Metrics) SetBaseLine(v *Baseline) *Metrics {
	if o.Baseline = v; o.Baseline == nil {
		o.nullFields = append(o.nullFields, "Baseline")
	}
	return o
}

func (o *Metrics) SetConsecutiveErrorLimit(v *int) *Metrics {
	if o.ConsecutiveErrorLimit = v; o.ConsecutiveErrorLimit == nil {
		o.nullFields = append(o.nullFields, "ConsecutiveErrorLimit")
	}
	return o
}

func (o *Metrics) SetCount(v *int) *Metrics {
	if o.Count = v; o.Count == nil {
		o.nullFields = append(o.nullFields, "Count")
	}
	return o
}

func (o *Metrics) SetDryRun(v *bool) *Metrics {
	if o.DryRun = v; o.DryRun == nil {
		o.nullFields = append(o.nullFields, "DryRun")
	}
	return o
}

func (o *Metrics) SetFailureCondition(v *string) *Metrics {
	if o.FailureCondition = v; o.FailureCondition == nil {
		o.nullFields = append(o.nullFields, "FailureCondition")
	}
	return o
}

func (o *Metrics) SetFailureLimit(v *int) *Metrics {
	if o.FailureLimit = v; o.FailureLimit == nil {
		o.nullFields = append(o.nullFields, "FailureLimit")
	}
	return o
}

func (o *Metrics) SetInitialDelay(v *string) *Metrics {
	if o.InitialDelay = v; o.InitialDelay == nil {
		o.nullFields = append(o.nullFields, "InitialDelay")
	}
	return o
}

func (o *Metrics) SetInterval(v *string) *Metrics {
	if o.Interval = v; o.Interval == nil {
		o.nullFields = append(o.nullFields, "Interval")
	}
	return o
}

func (o *Metrics) SetName(v *string) *Metrics {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *Metrics) SetProvider(v *Provider) *Metrics {
	if o.Provider = v; o.Provider == nil {
		o.nullFields = append(o.nullFields, "Provider")
	}
	return o
}

func (o *Metrics) SetSuccessCondition(v *string) *Metrics {
	if o.SuccessCondition = v; o.SuccessCondition == nil {
		o.nullFields = append(o.nullFields, "SuccessCondition")
	}
	return o
}

//end region

//region Baseline

func (o Baseline) MarshalJSON() ([]byte, error) {
	type noMethod Baseline
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Baseline) SetMinRange(v *int) *Baseline {
	if o.MinRange = v; o.MinRange == nil {
		o.nullFields = append(o.nullFields, "MinRange")
	}
	return o
}

func (o *Baseline) SetMaxRange(v *int) *Baseline {
	if o.MaxRange = v; o.MaxRange == nil {
		o.nullFields = append(o.nullFields, "MaxRange")
	}
	return o
}

func (o *Baseline) SetProvider(v *Provider) *Baseline {
	if o.Provider = v; o.Provider == nil {
		o.nullFields = append(o.nullFields, "Provider")
	}
	return o
}

func (o *Baseline) SetThreshold(v *string) *Baseline {
	if o.Threshold = v; o.Threshold == nil {
		o.nullFields = append(o.nullFields, "Threshold")
	}
	return o
}

//end region

//region Provider

func (o Provider) MarshalJSON() ([]byte, error) {
	type noMethod Provider
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Provider) SetCloudWatch(v *CloudWatchProvider) *Provider {
	if o.CloudWatch = v; o.CloudWatch == nil {
		o.nullFields = append(o.nullFields, "CloudWatch")
	}
	return o
}

func (o *Provider) SetDataDog(v *DataDogProvider) *Provider {
	if o.Datadog = v; o.Datadog == nil {
		o.nullFields = append(o.nullFields, "Datadog")
	}
	return o
}

func (o *Provider) SetJenkins(v *JenkinsProvider) *Provider {
	if o.Jenkins = v; o.Jenkins == nil {
		o.nullFields = append(o.nullFields, "Jenkins")
	}
	return o
}

func (o *Provider) SetNewRelic(v *NewRelicProvider) *Provider {
	if o.NewRelic = v; o.NewRelic == nil {
		o.nullFields = append(o.nullFields, "NewRelic")
	}
	return o
}

func (o *Provider) SetPrometheus(v *PrometheusProvider) *Provider {
	if o.Prometheus = v; o.Prometheus == nil {
		o.nullFields = append(o.nullFields, "Prometheus")
	}
	return o
}

func (o *Provider) SetWeb(v *Web) *Provider {
	if o.Web = v; o.Web == nil {
		o.nullFields = append(o.nullFields, "Web")
	}
	return o
}

func (o *Provider) SetJob(v *Job) *Provider {
	if o.Job = v; o.Job == nil {
		o.nullFields = append(o.nullFields, "Job")
	}
	return o
}

//end region

//region Cloud Watch Provider

func (o CloudWatchProvider) MarshalJSON() ([]byte, error) {
	type noMethod CloudWatchProvider
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *CloudWatchProvider) SetDuration(v *string) *CloudWatchProvider {
	if o.Duration = v; o.Duration == nil {
		o.nullFields = append(o.nullFields, "Duration")
	}
	return o
}

func (o *CloudWatchProvider) SetMetricDataQueries(v []*MetricDataQueries) *CloudWatchProvider {
	if o.MetricDataQueries = v; o.MetricDataQueries == nil {
		o.nullFields = append(o.nullFields, "MetricDataQueries")
	}
	return o
}

//end region

//region Metric Data Queries

func (o MetricDataQueries) MarshalJSON() ([]byte, error) {
	type noMethod MetricDataQueries
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *MetricDataQueries) SetExpression(v *string) *MetricDataQueries {
	if o.Expression = v; o.Expression == nil {
		o.nullFields = append(o.nullFields, "Expression")
	}
	return o
}

func (o *MetricDataQueries) SetID(v *string) *MetricDataQueries {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *MetricDataQueries) SetLabel(v *string) *MetricDataQueries {
	if o.Label = v; o.Label == nil {
		o.nullFields = append(o.nullFields, "Label")
	}
	return o
}

func (o *MetricDataQueries) SetMetricStat(v *MetricStat) *MetricDataQueries {
	if o.MetricStat = v; o.MetricStat == nil {
		o.nullFields = append(o.nullFields, "MetricStat")
	}
	return o
}

func (o *MetricDataQueries) SetPeriod(v *int) *MetricDataQueries {
	if o.Period = v; o.Period == nil {
		o.nullFields = append(o.nullFields, "Period")
	}
	return o
}

func (o *MetricDataQueries) SetReturnData(v *bool) *MetricDataQueries {
	if o.ReturnData = v; o.ReturnData == nil {
		o.nullFields = append(o.nullFields, "ReturnData")
	}
	return o
}

//end region

//region Metric Stat

func (o MetricStat) MarshalJSON() ([]byte, error) {
	type noMethod MetricStat
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *MetricStat) SetMetric(v *Metric) *MetricStat {
	if o.Metric = v; o.Metric == nil {
		o.nullFields = append(o.nullFields, "Metric")
	}
	return o
}

func (o *MetricStat) SetPeriod(v *int) *MetricStat {
	if o.Period = v; o.Period == nil {
		o.nullFields = append(o.nullFields, "Period")
	}
	return o
}

func (o *MetricStat) SetStat(v *string) *MetricStat {
	if o.Stat = v; o.Stat == nil {
		o.nullFields = append(o.nullFields, "Stat")
	}
	return o
}

func (o *MetricStat) SetUnit(v *string) *MetricStat {
	if o.Unit = v; o.Unit == nil {
		o.nullFields = append(o.nullFields, "Unit")
	}
	return o
}

//end region

//region Metric

func (o Metric) MarshalJSON() ([]byte, error) {
	type noMethod Metric
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Metric) SetDimensions(v []*Dimensions) *Metric {
	if o.Dimensions = v; o.Dimensions == nil {
		o.nullFields = append(o.nullFields, "Dimensions")
	}
	return o
}

func (o *Metric) SetMetricName(v *string) *Metric {
	if o.MetricName = v; o.MetricName == nil {
		o.nullFields = append(o.nullFields, "MetricName")
	}
	return o
}

func (o *Metric) SetNamespace(v *string) *Metric {
	if o.Namespace = v; o.Namespace == nil {
		o.nullFields = append(o.nullFields, "Namespace")
	}
	return o
}

//end region

//region Dimensions

func (o Dimensions) MarshalJSON() ([]byte, error) {
	type noMethod Dimensions
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Dimensions) SetName(v *string) *Dimensions {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *Dimensions) SetValue(v *string) *Dimensions {
	if o.Value = v; o.Value == nil {
		o.nullFields = append(o.nullFields, "Value")
	}
	return o
}

//end region

//region Data Dog Provider

func (o DataDogProvider) MarshalJSON() ([]byte, error) {
	type noMethod DataDogProvider
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *DataDogProvider) SetDuration(v *string) *DataDogProvider {
	if o.Duration = v; o.Duration == nil {
		o.nullFields = append(o.nullFields, "Duration")
	}
	return o
}

func (o *DataDogProvider) SetQuery(v *string) *DataDogProvider {
	if o.Query = v; o.Query == nil {
		o.nullFields = append(o.nullFields, "Query")
	}
	return o
}

//end region

//region Jenkins Provider

func (o JenkinsProvider) MarshalJSON() ([]byte, error) {
	type noMethod JenkinsProvider
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *JenkinsProvider) SetInterval(v *string) *JenkinsProvider {
	if o.Interval = v; o.Interval == nil {
		o.nullFields = append(o.nullFields, "Interval")
	}
	return o
}

func (o *JenkinsProvider) SetPipelineName(v *string) *JenkinsProvider {
	if o.PipelineName = v; o.PipelineName == nil {
		o.nullFields = append(o.nullFields, "PipelineName")
	}
	return o
}

func (o *JenkinsProvider) SetParameters(v []*Parameters) *JenkinsProvider {
	if o.Parameters = v; o.Parameters == nil {
		o.nullFields = append(o.nullFields, "Parameters")
	}
	return o
}

func (o *JenkinsProvider) SetTimeout(v *string) *JenkinsProvider {
	if o.Timeout = v; o.Timeout == nil {
		o.nullFields = append(o.nullFields, "Timeout")
	}
	return o
}

func (o *JenkinsProvider) SetTLSVerification(v *bool) *JenkinsProvider {
	if o.TLSVerification = v; o.TLSVerification == nil {
		o.nullFields = append(o.nullFields, "TLSVerification")
	}
	return o
}

//end region

//region Parameters

func (o Parameters) MarshalJSON() ([]byte, error) {
	type noMethod Parameters
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Parameters) SetKey(v *string) *Parameters {
	if o.Key = v; o.Key == nil {
		o.nullFields = append(o.nullFields, "Key")
	}
	return o
}

func (o *Parameters) SetValue(v *string) *Parameters {
	if o.Value = v; o.Value == nil {
		o.nullFields = append(o.nullFields, "Value")
	}
	return o
}

//end region

//region New Relic Provider

func (o NewRelicProvider) MarshalJSON() ([]byte, error) {
	type noMethod NewRelicProvider
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *NewRelicProvider) SetProfile(v *string) *NewRelicProvider {
	if o.Profile = v; o.Profile == nil {
		o.nullFields = append(o.nullFields, "Profile")
	}
	return o
}

func (o *NewRelicProvider) SetQuery(v *string) *NewRelicProvider {
	if o.Query = v; o.Query == nil {
		o.nullFields = append(o.nullFields, "Query")
	}
	return o
}

//end region

//region Prometheus Provider

func (o PrometheusProvider) MarshalJSON() ([]byte, error) {
	type noMethod PrometheusProvider
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *PrometheusProvider) SetQuery(v *string) *PrometheusProvider {
	if o.Query = v; o.Query == nil {
		o.nullFields = append(o.nullFields, "Query")
	}
	return o
}

//end region

//region Web

func (o Web) MarshalJSON() ([]byte, error) {
	type noMethod Web
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Web) SetBody(v *string) *Web {
	if o.Body = v; o.Body == nil {
		o.nullFields = append(o.nullFields, "Body")
	}
	return o
}

func (o *Web) SetInsecure(v *bool) *Web {
	if o.Insecure = v; o.Insecure == nil {
		o.nullFields = append(o.nullFields, "Insecure")
	}
	return o
}

func (o *Web) SetHeaders(v []*Headers) *Web {
	if o.Headers = v; o.Headers == nil {
		o.nullFields = append(o.nullFields, "Headers")
	}
	return o
}

func (o *Web) SetJsonPath(v *string) *Web {
	if o.JsonPath = v; o.JsonPath == nil {
		o.nullFields = append(o.nullFields, "Jsonpath")
	}
	return o
}

func (o *Web) SetMethod(v *string) *Web {
	if o.Method = v; o.Method == nil {
		o.nullFields = append(o.nullFields, "Method")
	}
	return o
}

func (o *Web) SetTimeoutSeconds(v *int) *Web {
	if o.TimeoutSeconds = v; o.TimeoutSeconds == nil {
		o.nullFields = append(o.nullFields, "TimeoutSeconds")
	}
	return o
}

func (o *Web) SetUrl(v *string) *Web {
	if o.Url = v; o.Url == nil {
		o.nullFields = append(o.nullFields, "Url")
	}
	return o
}

//end region

//region Headers

func (o Headers) MarshalJSON() ([]byte, error) {
	type noMethod Headers
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Headers) SetKey(v *string) *Headers {
	if o.Key = v; o.Key == nil {
		o.nullFields = append(o.nullFields, "Key")
	}
	return o
}

func (o *Headers) SetValue(v *string) *Headers {
	if o.Value = v; o.Value == nil {
		o.nullFields = append(o.nullFields, "Value")
	}
	return o
}

//end region

//region Job

func (o Job) MarshalJSON() ([]byte, error) {
	type noMethod Job
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Job) SetSpec(v *Spec) *Job {
	if o.Spec = v; o.Spec == nil {
		o.nullFields = append(o.nullFields, "Spec")
	}
	return o
}

//end region

//region Spec

func (o Spec) MarshalJSON() ([]byte, error) {
	type noMethod Spec
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Spec) SetBackoffLimit(v *int) *Spec {
	if o.BackoffLimit = v; o.BackoffLimit == nil {
		o.nullFields = append(o.nullFields, "BackoffLimit")
	}
	return o
}

func (o *Spec) SetTemplate(v *Template) *Spec {
	if o.Template = v; o.Template == nil {
		o.nullFields = append(o.nullFields, "Template")
	}
	return o
}

//end region

// region Template
func (o Template) MarshalJSON() ([]byte, error) {
	type noMethod Template
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Template) SetSpec(v *TemplateSpec) *Template {
	if o.Spec = v; o.Spec == nil {
		o.nullFields = append(o.nullFields, "Spec")
	}
	return o
}

//end region

//region TemplateSpec

func (o TemplateSpec) MarshalJSON() ([]byte, error) {
	type noMethod TemplateSpec
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *TemplateSpec) SetRestartPolicy(v *string) *TemplateSpec {
	if o.RestartPolicy = v; o.RestartPolicy == nil {
		o.nullFields = append(o.nullFields, "RestartPolicy")
	}
	return o
}

func (o *TemplateSpec) SetContainers(v []*Containers) *TemplateSpec {
	if o.Containers = v; o.Containers == nil {
		o.nullFields = append(o.nullFields, "Containers")
	}
	return o
}

//end region

//region containers

func (o Containers) MarshalJSON() ([]byte, error) {
	type noMethod Containers
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Containers) SetCommand(v []string) *Containers {
	if o.Command = v; o.Command == nil {
		o.nullFields = append(o.nullFields, "Command")
	}
	return o
}

func (o *Containers) SetImage(v *string) *Containers {
	if o.Image = v; o.Image == nil {
		o.nullFields = append(o.nullFields, "Image")
	}
	return o
}

func (o *Containers) SetName(v *string) *Containers {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}
