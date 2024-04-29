package aws

import (
	"context"
	"encoding/json"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/jsonutil"
	"io/ioutil"
	"net/http"

	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/uritemplates"
)

// ResourceSuggestion represents a single resource suggestion.
type ResourceSuggestion struct {
	ResourceName    *string                        `json:"resourceName,omitempty"`
	ResourceType    *string                        `json:"resourceType,omitempty"`
	DeploymentName  *string                        `json:"deploymentName,omitempty"`
	Namespace       *string                        `json:"namespace,omitempty"`
	SuggestedCPU    *float64                       `json:"suggestedCPU,omitempty"`
	RequestedCPU    *float64                       `json:"requestedCPU,omitempty"`
	SuggestedMemory *float64                       `json:"suggestedMemory,omitempty"`
	RequestedMemory *float64                       `json:"requestedMemory,omitempty"`
	Containers      []*ContainerResourceSuggestion `json:"containers,omitempty"`
}

// ContainerResourceSuggestion represents a resource suggestion for a
// single container.
type ContainerResourceSuggestion struct {
	Name            *string  `json:"name,omitempty"`
	SuggestedCPU    *float64 `json:"suggestedCpu,omitempty"`
	RequestedCPU    *float64 `json:"requestedCpu,omitempty"`
	SuggestedMemory *float64 `json:"suggestedMemory,omitempty"`
	RequestedMemory *float64 `json:"requestedMemory,omitempty"`
}

type Filter struct {
	Attribute  *Attribute `json:"attribute,omitempty"`
	Namespaces []string   `json:"namespaces,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Attribute struct {
	Key      *string `json:"key,omitempty"`
	Operator *string `json:"operator,omitempty"`
	Type     *string `json:"type,omitempty"`
	Value    *string `json:"value,omitempty"`

	forceSendFields []string
	nullFields      []string
}

// Deprecated: Use ListOceanResourceSuggestionsOutput instead.
type ListResourceSuggestionsOutput struct {
	Suggestions []*ResourceSuggestion `json:"suggestions,omitempty"`
}

// ListOceanResourceSuggestionsInput represents the input of `ListOceanResourceSuggestions` function.
type ListOceanResourceSuggestionsInput struct {
	OceanID *string `json:"oceanId,omitempty"`
	Filter  *Filter `json:"filter,omitempty"`
}

// ListOceanResourceSuggestionsOutput represents the output of `ListOceanResourceSuggestions` function.
type ListOceanResourceSuggestionsOutput struct {
	Suggestions []*ResourceSuggestion `json:"suggestions,omitempty"`
}

// Deprecated: Use ListOceanResourceSuggestionsInput instead.
type ListResourceSuggestionsInput struct {
	OceanID   *string `json:"oceanId,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

func resourceSuggestionFromJSON(in []byte) (*ResourceSuggestion, error) {
	b := new(ResourceSuggestion)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func resourceSuggestionsFromJSON(in []byte) ([]*ResourceSuggestion, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*ResourceSuggestion, len(rw.Response.Items))
	for i, rb := range rw.Response.Items {
		b, err := resourceSuggestionFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func resourceSuggestionsFromHTTPResponse(resp *http.Response) ([]*ResourceSuggestion, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return resourceSuggestionsFromJSON(body)
}

// ListOceanResourceSuggestions returns a list of right-sizing resource suggestions
// for an Ocean cluster.
func (s *ServiceOp) ListOceanResourceSuggestions(ctx context.Context, input *ListOceanResourceSuggestionsInput) (*ListOceanResourceSuggestionsOutput, error) {
	path, err := uritemplates.Expand("/ocean/aws/k8s/cluster/{oceanId}/rightSizing/suggestion", uritemplates.Values{
		"oceanId": spotinst.StringValue(input.OceanID),
	})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodPost, path)

	// We do NOT need the ID anymore, so let's drop it.
	input.OceanID = nil
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rs, err := resourceSuggestionsFromHTTPResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListOceanResourceSuggestionsOutput{Suggestions: rs}, nil
}

// Deprecated: Use ListOceanResourceSuggestions instead.
func (s *ServiceOp) ListResourceSuggestions(ctx context.Context, input *ListResourceSuggestionsInput) (*ListResourceSuggestionsOutput, error) {
	path, err := uritemplates.Expand("/ocean/aws/k8s/cluster/{oceanId}/rightSizing/resourceSuggestion", uritemplates.Values{
		"oceanId": spotinst.StringValue(input.OceanID),
	})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodGet, path)

	if input.Namespace != nil {
		r.Params.Set("namespace", *input.Namespace)
	}
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := resourceSuggestionsFromHTTPResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListResourceSuggestionsOutput{Suggestions: gs}, nil
}

type RightSizingRule struct {
	Name                                  *string                                `json:"ruleName,omitempty"`
	OceanId                               *string                                `json:"oceanId,omitempty"`
	RestartPods                           *bool                                  `json:"restartPods,omitempty"`
	RecommendationApplicationIntervals    []*RecommendationApplicationInterval   `json:"recommendationApplicationIntervals,omitempty"`
	RecommendationApplicationMinThreshold *RecommendationApplicationMinThreshold `json:"recommendationApplicationMinThreshold,omitempty"`
	RecommendationApplicationBoundaries   *RecommendationApplicationBoundaries   `json:"recommendationApplicationBoundaries,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type RecommendationApplicationInterval struct {
	RepetitionBasis        *string                 `json:"repetitionBasis,omitempty"`
	WeeklyRepetitionBasis  *WeeklyRepetitionBasis  `json:"weeklyRepetitionBasis,omitempty"`
	MonthlyRepetitionBasis *MonthlyRepetitionBasis `json:"monthlyRepetitionBasis,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type WeeklyRepetitionBasis struct {
	IntervalDays  []string       `json:"intervalDays,omitempty"`
	IntervalHours *IntervalHours `json:"intervalHours,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type MonthlyRepetitionBasis struct {
	IntervalMonths        []int                  `json:"intervalMonths,omitempty"`
	WeekOfTheMonth        []string               `json:"weekOfTheMonth,omitempty"`
	WeeklyRepetitionBasis *WeeklyRepetitionBasis `json:"weeklyRepetitionBasis,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type IntervalHours struct {
	StartTime *string `json:"startTime,omitempty"`
	EndTime   *string `json:"endTime,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type RecommendationApplicationMinThreshold struct {
	CpuPercentage    *float64 `json:"cpuPercentage,omitempty"`
	MemoryPercentage *float64 `json:"memoryPercentage,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type RecommendationApplicationBoundaries struct {
	Cpu    *Cpu    `json:"cpu,omitempty"`
	Memory *Memory `json:"memory,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Cpu struct {
	Min *int `json:"min,omitempty"`
	Max *int `json:"max,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Memory struct {
	Min *int `json:"min,omitempty"`
	Max *int `json:"max,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type RightSizingWorkloads struct {
	Namespaces []*Namespace `json:"namespaces,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Namespace struct {
	NamespaceName *string     `json:"namespaceName,omitempty"`
	Workloads     []*Workload `json:"workloads,omitempty"`
	Labels        []*Label    `json:"labels,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Workload struct {
	Name         *string     `json:"name,omitempty"`
	WorkloadType []*Workload `json:"workloadType,omitempty"`
	RegexName    *string     `json:"regexName,omitempty"`
	Namespace    *string     `json:"Namespace,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ListRightSizingRulesInput struct {
	RightSizingRule *RightSizingRule `json:"rightSizingRule,omitempty"`
}

type ListRightSizingRulesOutput struct {
	RightSizingRules []*RightSizingRule `json:"rightSizingRules,omitempty"`
}

type CreateRightSizingRuleInput struct {
	RightSizingRule *RightSizingRule `json:"rightSizingRule,omitempty"`
}

type CreateRightSizingRuleOutput struct {
	RightSizingRule *RightSizingRule `json:"rightSizingRule,omitempty"`
}

type ReadRightSizingRuleInput struct {
	RuleName        *string          `json:"ruleName,omitempty"`
	RightSizingRule *RightSizingRule `json:"rightSizingRule,omitempty"`
}

type ReadRightSizingRuleOutput struct {
	RightSizingRule *RightSizingRule `json:"rightSizingRule,omitempty"`
}

type UpdateRightSizingRuleInput struct {
	RuleName        *string          `json:"ruleName,omitempty"`
	RightSizingRule *RightSizingRule `json:"rightSizingRule,omitempty"`
}

type UpdateRightSizingRuleOutput struct {
	RightSizingRule *RightSizingRule `json:"rightSizingRule,omitempty"`
}

type DeleteRightSizingRuleInput struct {
	RuleNames       []string         `json:"ruleNames,omitempty"`
	RightSizingRule *RightSizingRule `json:"rightSizingRule,omitempty"`
}

type DeleteRightSizingRuleOutput struct{}

func rightSizingRuleFromJSON(in []byte) (*RightSizingRule, error) {
	b := new(RightSizingRule)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func rightSizingRulesFromJSON(in []byte) ([]*RightSizingRule, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*RightSizingRule, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := rightSizingRuleFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func rightSizingRulesFromHttpResponse(resp *http.Response) ([]*RightSizingRule, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return rightSizingRulesFromJSON(body)
}

func (s *ServiceOp) ListRightSizingRules(ctx context.Context, input *ListRightSizingRulesInput) (*ListRightSizingRulesOutput, error) {
	path, err := uritemplates.Expand("/ocean/{oceanId}/rightSizing/rule", uritemplates.Values{
		"oceanId": spotinst.StringValue(input.RightSizingRule.OceanId),
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

	gs, err := rightSizingRulesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListRightSizingRulesOutput{RightSizingRules: gs}, nil
}

func (s *ServiceOp) CreateRightSizingRule(ctx context.Context, input *CreateRightSizingRuleInput) (*CreateRightSizingRuleOutput, error) {
	path, err := uritemplates.Expand("/ocean/{oceanId}/rightSizing/rule", uritemplates.Values{
		"oceanId": spotinst.StringValue(input.RightSizingRule.OceanId),
	})
	if err != nil {
		return nil, err
	}

	// We do NOT need the ID anymore, so let's drop it.
	input.RightSizingRule.OceanId = nil
	r := client.NewRequest(http.MethodPost, path)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := rightSizingRulesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateRightSizingRuleOutput)
	if len(gs) > 0 {
		output.RightSizingRule = gs[0]
	}

	return output, nil
}

func (s *ServiceOp) ReadRightSizingRule(ctx context.Context, input *ReadRightSizingRuleInput) (*ReadRightSizingRuleOutput, error) {
	path, err := uritemplates.Expand("/{oceanId}/rightSizing/rule/{ruleName}", uritemplates.Values{
		"oceanId":  spotinst.StringValue(input.RightSizingRule.OceanId),
		"ruleName": spotinst.StringValue(input.RuleName),
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

	gs, err := rightSizingRulesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadRightSizingRuleOutput)
	if len(gs) > 0 {
		output.RightSizingRule = gs[0]
	}

	return output, nil
}

func (s *ServiceOp) UpdateRightSizingRule(ctx context.Context, input *UpdateRightSizingRuleInput) (*UpdateRightSizingRuleOutput, error) {
	path, err := uritemplates.Expand("/{oceanId}/rightSizing/rule/{ruleName}", uritemplates.Values{
		"oceanId":  spotinst.StringValue(input.RightSizingRule.OceanId),
		"ruleName": spotinst.StringValue(input.RuleName),
	})

	input.RightSizingRule.OceanId = nil

	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodPut, path)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := rightSizingRulesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(UpdateRightSizingRuleOutput)
	if len(gs) > 0 {
		output.RightSizingRule = gs[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteRightSizingRules(ctx context.Context, input *DeleteRightSizingRuleInput) (*DeleteRightSizingRuleOutput, error) {
	path, err := uritemplates.Expand("/ocean/{oceanId}/rightSizing/rule", uritemplates.Values{
		"oceanId": spotinst.StringValue(input.RightSizingRule.OceanId),
	})
	if err != nil {
		return nil, err
	}

	// We do NOT need the ID anymore, so let's drop it.
	input.RightSizingRule = nil

	r := client.NewRequest(http.MethodDelete, path)
	r.Obj = input
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteRightSizingRuleOutput{}, nil
}

/*func (s *ServiceOp) AttachWorkloadsToRule(ctx context.Context, input *AttachWorkloadsToRuleInput) (*AttachWorkloadsToRuleOutput, error) {
	r := client.NewRequest(http.MethodPost, "/ocean/aws/k8s/cluster")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := rightSizingRulesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(AttachWorkloadsToRuleOutput)
	if len(gs) > 0 {
		output.RightSizingRule = gs[0]
	}

	return output, nil
}
*/

// region RightSizingRule

func (o RightSizingRule) MarshalJSON() ([]byte, error) {
	type noMethod RightSizingRule
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *RightSizingRule) SetName(v *string) *RightSizingRule {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *RightSizingRule) SetOceanId(v *string) *RightSizingRule {
	if o.OceanId = v; o.OceanId == nil {
		o.nullFields = append(o.nullFields, "oceanId")
	}
	return o
}

func (o *RightSizingRule) SetRestartPods(v *bool) *RightSizingRule {
	if o.RestartPods = v; o.RestartPods == nil {
		o.nullFields = append(o.nullFields, "RestartPods")
	}
	return o
}

func (o *RightSizingRule) SetRecommendationApplicationIntervals(v []*RecommendationApplicationInterval) *RightSizingRule {
	if o.RecommendationApplicationIntervals = v; o.RecommendationApplicationIntervals == nil {
		o.nullFields = append(o.nullFields, "RecommendationApplicationIntervals")
	}
	return o
}

func (o *RightSizingRule) SetRecommendationApplicationBoundaries(v *RecommendationApplicationBoundaries) *RightSizingRule {
	if o.RecommendationApplicationBoundaries = v; o.RecommendationApplicationBoundaries == nil {
		o.nullFields = append(o.nullFields, "RecommendationApplicationBoundaries")
	}
	return o
}

func (o *RightSizingRule) SetRecommendationApplicationMinThreshold(v *RecommendationApplicationMinThreshold) *RightSizingRule {
	if o.RecommendationApplicationMinThreshold = v; o.RecommendationApplicationMinThreshold == nil {
		o.nullFields = append(o.nullFields, "RecommendationApplicationMinThreshold")
	}
	return o
}

// region RecommendationApplicationInterval

func (o RecommendationApplicationInterval) MarshalJSON() ([]byte, error) {
	type noMethod RecommendationApplicationInterval
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *RecommendationApplicationInterval) SetRepetitionBasis(v *string) *RecommendationApplicationInterval {
	if o.RepetitionBasis = v; o.RepetitionBasis == nil {
		o.nullFields = append(o.nullFields, "RepetitionBasis")
	}
	return o
}

func (o *RecommendationApplicationInterval) SetWeeklyRepetitionBasis(v *WeeklyRepetitionBasis) *RecommendationApplicationInterval {
	if o.WeeklyRepetitionBasis = v; o.WeeklyRepetitionBasis == nil {
		o.nullFields = append(o.nullFields, "WeeklyRepetitionBasis")
	}
	return o
}

func (o *RecommendationApplicationInterval) SetMonthlyRepetitionBasis(v *MonthlyRepetitionBasis) *RecommendationApplicationInterval {
	if o.MonthlyRepetitionBasis = v; o.MonthlyRepetitionBasis == nil {
		o.nullFields = append(o.nullFields, "MonthlyRepetitionBasis")
	}
	return o
}

// region WeeklyRepetitionBasis

func (o *WeeklyRepetitionBasis) SetIntervalDays(v []string) *WeeklyRepetitionBasis {
	if o.IntervalDays = v; o.IntervalDays == nil {
		o.nullFields = append(o.nullFields, "IntervalDays")
	}
	return o
}

func (o *WeeklyRepetitionBasis) SetIntervalHours(v *IntervalHours) *WeeklyRepetitionBasis {
	if o.IntervalHours = v; o.IntervalHours == nil {
		o.nullFields = append(o.nullFields, "IntervalHours")
	}
	return o
}

// region IntervalHours

func (o *IntervalHours) SetStartTime(v *string) *IntervalHours {
	if o.StartTime = v; o.StartTime == nil {
		o.nullFields = append(o.nullFields, "StartTime")
	}
	return o
}

func (o *IntervalHours) SetEndTime(v *string) *IntervalHours {
	if o.EndTime = v; o.EndTime == nil {
		o.nullFields = append(o.nullFields, "EndTime")
	}
	return o
}

// region MonthlyRepetitionBasis

func (o *MonthlyRepetitionBasis) SetIntervalMonths(v []int) *MonthlyRepetitionBasis {
	if o.IntervalMonths = v; o.IntervalMonths == nil {
		o.nullFields = append(o.nullFields, "IntervalMonths")
	}
	return o
}

func (o *MonthlyRepetitionBasis) SetWeekOfTheMonth(v []string) *MonthlyRepetitionBasis {
	if o.WeekOfTheMonth = v; o.WeekOfTheMonth == nil {
		o.nullFields = append(o.nullFields, "WeekOfTheMonth")
	}
	return o
}

func (o *MonthlyRepetitionBasis) SetMonthlyWeeklyRepetitionBasis(v *WeeklyRepetitionBasis) *MonthlyRepetitionBasis {
	if o.WeeklyRepetitionBasis = v; o.WeeklyRepetitionBasis == nil {
		o.nullFields = append(o.nullFields, "WeeklyRepetitionBasis")
	}
	return o
}

// region RecommendationApplicationBoundaries

func (o *RecommendationApplicationBoundaries) SetCpu(v *Cpu) *RecommendationApplicationBoundaries {
	if o.Cpu = v; o.Cpu == nil {
		o.nullFields = append(o.nullFields, "Cpu")
	}
	return o
}

func (o *RecommendationApplicationBoundaries) SetMemory(v *Memory) *RecommendationApplicationBoundaries {
	if o.Memory = v; o.Memory == nil {
		o.nullFields = append(o.nullFields, "Memory")
	}
	return o
}

// region Cpu

func (o *Cpu) SetMin(v *int) *Cpu {
	if o.Min = v; o.Min == nil {
		o.nullFields = append(o.nullFields, "Cpu")
	}
	return o
}

func (o *Cpu) SetMax(v *int) *Cpu {
	if o.Max = v; o.Min == nil {
		o.nullFields = append(o.nullFields, "Cpu")
	}
	return o
}

// region Memory

func (o *Memory) SetMin(v *int) *Memory {
	if o.Min = v; o.Min == nil {
		o.nullFields = append(o.nullFields, "Memory")
	}
	return o
}

func (o *Memory) SetMax(v *int) *Memory {
	if o.Max = v; o.Max == nil {
		o.nullFields = append(o.nullFields, "Memory")
	}
	return o
}

// region RecommendationApplicationMinThreshold

func (o *RecommendationApplicationMinThreshold) SetCpuPercentage(v *float64) *RecommendationApplicationMinThreshold {
	if o.CpuPercentage = v; o.CpuPercentage == nil {
		o.nullFields = append(o.nullFields, "CpuPercentage")
	}
	return o
}

func (o *RecommendationApplicationMinThreshold) SetMemoryPercentage(v *float64) *RecommendationApplicationMinThreshold {
	if o.MemoryPercentage = v; o.MemoryPercentage == nil {
		o.nullFields = append(o.nullFields, "MemoryPercentage")
	}
	return o
}
