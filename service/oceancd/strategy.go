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

type Strategy struct {
	Canary  *Canary  `json:"canary,omitempty"`
	Name    *string  `json:"name,omitempty"`
	Rolling *Rolling `json:"rolling,omitempty"`

	// Read-only fields.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Canary struct {
	BackgroundVerification *BackgroundVerification `json:"backgroundVerification,omitempty"`
	Steps                  []*CanarySteps          `json:"steps,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type BackgroundVerification struct {
	TemplateNames []string `json:"templateNames,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type CanarySteps struct {
	Name           *string         `json:"name,omitempty"`
	Pause          *Pause          `json:"pause,omitempty"`
	SetCanaryScale *SetCanaryScale `json:"setCanaryScale,omitempty"`
	SetHeaderRoute *SetHeaderRoute `json:"setHeaderRoute,omitempty"`
	SetWeight      *int            `json:"setWeight,omitempty"`
	Verification   *Verification   `json:"verification,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Pause struct {
	Duration *string `json:"duration,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type SetHeaderRoute struct {
	Match []*Match `json:"match,omitempty"`
	Name  *string  `json:"name,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Match struct {
	HeaderName  *string      `json:"headerName,omitempty"`
	HeaderValue *HeaderValue `json:"headerValue,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type HeaderValue struct {
	Exact  *string `json:"exact,omitempty"`
	Prefix *string `json:"prefix,omitempty"`
	Regex  *string `json:"regex,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type SetCanaryScale struct {
	MatchTrafficWeight *bool `json:"matchTrafficWeight,omitempty"`
	Replicas           *int  `json:"replicas,omitempty"`
	Weight             *int  `json:"weight,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Verification struct {
	TemplateNames []string `json:"templateNames,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Rolling struct {
	Steps []*RollingSteps `json:"steps,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type RollingSteps struct {
	Name         *string       `json:"name,omitempty"`
	Pause        *Pause        `json:"pause,omitempty"`
	Verification *Verification `json:"verification,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ListStrategiesOutput struct {
	Strategies []*Strategy `json:"strategy,omitempty"`
}

type CreateStrategyInput struct {
	Strategy *Strategy `json:"strategy,omitempty"`
}

type CreateStrategyOutput struct {
	Strategy *Strategy `json:"strategy,omitempty"`
}

type ReadStrategyInput struct {
	StrategyName *string `json:"strategyName,omitempty"`
}

type ReadStrategyOutput struct {
	Strategy *Strategy `json:"strategy,omitempty"`
}

type UpdateStrategyInput struct {
	Strategy *Strategy `json:"strategy,omitempty"`
}

type UpdateStrategyOutput struct {
	Strategy *Strategy `json:"strategy,omitempty"`
}

type PatchStrategyInput struct {
	Strategy *Strategy `json:"strategy,omitempty"`
}

type PatchStrategyOutput struct {
	Strategy *Strategy `json:"strategy,omitempty"`
}

type DeleteStrategyInput struct {
	StrategyName *string `json:"strategyName,omitempty"`
}

type DeleteStrategyOutput struct{}

func strategyFromJSON(in []byte) (*Strategy, error) {
	b := new(Strategy)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func strategiesFromJSON(in []byte) ([]*Strategy, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Strategy, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := strategyFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func strategiesFromHttpResponse(resp *http.Response) ([]*Strategy, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return strategiesFromJSON(body)
}

// endregion

// region API requests

func (s *ServiceOp) ListStrategies(ctx context.Context) (*ListStrategiesOutput, error) {
	r := client.NewRequest(http.MethodGet, "/ocean/cd/strategy")
	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	strategy, err := strategiesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListStrategiesOutput{Strategies: strategy}, nil
}

func (s *ServiceOp) CreateStrategy(ctx context.Context, input *CreateStrategyInput) (*CreateStrategyOutput, error) {
	r := client.NewRequest(http.MethodPost, "/ocean/cd/strategy")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	strategy, err := strategiesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateStrategyOutput)
	if len(strategy) > 0 {
		output.Strategy = strategy[0]
	}

	return output, nil
}

func (s *ServiceOp) ReadStrategy(ctx context.Context, input *ReadStrategyInput) (*ReadStrategyOutput, error) {
	path, err := uritemplates.Expand("/ocean/cd/strategy/{strategyName}", uritemplates.Values{
		"strategyName": spotinst.StringValue(input.StrategyName),
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

	strategy, err := strategiesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadStrategyOutput)
	if len(strategy) > 0 {
		output.Strategy = strategy[0]
	}

	return output, nil
}

func (s *ServiceOp) UpdateStrategy(ctx context.Context, input *UpdateStrategyInput) (*UpdateStrategyOutput, error) {
	path, err := uritemplates.Expand("/ocean/cd/strategy/{strategyName}", uritemplates.Values{
		"strategyName": spotinst.StringValue(input.Strategy.Name),
	})
	if err != nil {
		return nil, err
	}

	// We do NOT need the Name anymore, so let's drop it.
	input.Strategy.Name = nil

	r := client.NewRequest(http.MethodPut, path)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	strategy, err := strategiesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(UpdateStrategyOutput)
	if len(strategy) > 0 {
		output.Strategy = strategy[0]
	}

	return output, nil
}

func (s *ServiceOp) PatchStrategy(ctx context.Context, input *PatchStrategyInput) (*PatchStrategyOutput, error) {
	path, err := uritemplates.Expand("/ocean/cd/strategy/{strategyName}", uritemplates.Values{
		"strategyName": spotinst.StringValue(input.Strategy.Name),
	})
	if err != nil {
		return nil, err
	}

	// We do NOT need the Name anymore, so let's drop it.
	input.Strategy.Name = nil

	r := client.NewRequest(http.MethodPatch, path)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	strategy, err := strategiesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(PatchStrategyOutput)
	if len(strategy) > 0 {
		output.Strategy = strategy[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteStrategy(ctx context.Context, input *DeleteStrategyInput) (*DeleteStrategyOutput, error) {
	path, err := uritemplates.Expand("/ocean/cd/strategy/{strategyName}", uritemplates.Values{
		"strategyName": spotinst.StringValue(input.StrategyName),
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

	return &DeleteStrategyOutput{}, nil
}

//end region

// region Strategy

func (o Strategy) MarshalJSON() ([]byte, error) {
	type noMethod Strategy
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Strategy) SetCanary(v *Canary) *Strategy {
	if o.Canary = v; o.Canary == nil {
		o.nullFields = append(o.nullFields, "Canary")
	}
	return o
}

func (o *Strategy) SetName(v *string) *Strategy {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *Strategy) SetRolling(v *Rolling) *Strategy {
	if o.Rolling = v; o.Rolling == nil {
		o.nullFields = append(o.nullFields, "Rolling")
	}
	return o
}

// end region

//region Canary

func (o Canary) MarshalJSON() ([]byte, error) {
	type noMethod Canary
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Canary) SetBackgroundVerification(v *BackgroundVerification) *Canary {
	if o.BackgroundVerification = v; o.BackgroundVerification == nil {
		o.nullFields = append(o.nullFields, "BackgroundVerification")
	}
	return o
}

func (o *Canary) SetSteps(v []*CanarySteps) *Canary {
	if o.Steps = v; o.Steps == nil {
		o.nullFields = append(o.nullFields, "Steps")
	}
	return o
}

//end region

//region BackgroundVerification

func (o BackgroundVerification) MarshalJSON() ([]byte, error) {
	type noMethod BackgroundVerification
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *BackgroundVerification) SetTemplateNames(v []string) *BackgroundVerification {
	if o.TemplateNames = v; o.TemplateNames == nil {
		o.nullFields = append(o.nullFields, "TemplateNames")
	}
	return o
}

//end region

//region Canary Steps

func (o CanarySteps) MarshalJSON() ([]byte, error) {
	type noMethod CanarySteps
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *CanarySteps) SetName(v *string) *CanarySteps {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *CanarySteps) SetPause(v *Pause) *CanarySteps {
	if o.Pause = v; o.Pause == nil {
		o.nullFields = append(o.nullFields, "Pause")
	}
	return o
}

func (o *CanarySteps) SetSetCanaryScale(v *SetCanaryScale) *CanarySteps {
	if o.SetCanaryScale = v; o.SetCanaryScale == nil {
		o.nullFields = append(o.nullFields, "SetCanaryScale")
	}
	return o
}

func (o *CanarySteps) SetSetHeaderRoute(v *SetHeaderRoute) *CanarySteps {
	if o.SetHeaderRoute = v; o.SetHeaderRoute == nil {
		o.nullFields = append(o.nullFields, "SetHeaderRoute")
	}
	return o
}

func (o *CanarySteps) SetSetWeight(v *int) *CanarySteps {
	if o.SetWeight = v; o.SetWeight == nil {
		o.nullFields = append(o.nullFields, "SetWeight")
	}
	return o
}

func (o *CanarySteps) SetVerification(v *Verification) *CanarySteps {
	if o.Verification = v; o.Verification == nil {
		o.nullFields = append(o.nullFields, "Verification")
	}
	return o
}

//end region

//region Pause

func (o Pause) MarshalJSON() ([]byte, error) {
	type noMethod Pause
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Pause) SetDuration(v *string) *Pause {
	if o.Duration = v; o.Duration == nil {
		o.nullFields = append(o.nullFields, "Duration")
	}
	return o
}

//end region

//region SetCanaryScale

func (o SetCanaryScale) MarshalJSON() ([]byte, error) {
	type noMethod SetCanaryScale
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *SetCanaryScale) SetMatchTrafficWeight(v *bool) *SetCanaryScale {
	if o.MatchTrafficWeight = v; o.MatchTrafficWeight == nil {
		o.nullFields = append(o.nullFields, "MatchTrafficWeight")
	}
	return o
}

func (o *SetCanaryScale) SetReplicas(v *int) *SetCanaryScale {
	if o.Replicas = v; o.Replicas == nil {
		o.nullFields = append(o.nullFields, "Replicas")
	}
	return o
}

func (o *SetCanaryScale) SetWeight(v *int) *SetCanaryScale {
	if o.Weight = v; o.Weight == nil {
		o.nullFields = append(o.nullFields, "Weight")
	}
	return o
}

//end region

//region Set Header Route

func (o SetHeaderRoute) MarshalJSON() ([]byte, error) {
	type noMethod SetHeaderRoute
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *SetHeaderRoute) SetMatch(v []*Match) *SetHeaderRoute {
	if o.Match = v; o.Match == nil {
		o.nullFields = append(o.nullFields, "Match")
	}
	return o
}

func (o *SetHeaderRoute) SetName(v *string) *SetHeaderRoute {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

//end region

//region Match

func (o Match) MarshalJSON() ([]byte, error) {
	type noMethod Match
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Match) SetHeaderName(v *string) *Match {
	if o.HeaderName = v; o.HeaderName == nil {
		o.nullFields = append(o.nullFields, "HeaderName")
	}
	return o
}

func (o *Match) SetHeaderValue(v *HeaderValue) *Match {
	if o.HeaderValue = v; o.HeaderValue == nil {
		o.nullFields = append(o.nullFields, "HeaderValue")
	}
	return o
}

//end region

//region HeaderValue

func (o HeaderValue) MarshalJSON() ([]byte, error) {
	type noMethod HeaderValue
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *HeaderValue) SetExact(v *string) *HeaderValue {
	if o.Exact = v; o.Exact == nil {
		o.nullFields = append(o.nullFields, "Exact")
	}
	return o
}

func (o *HeaderValue) SetPrefix(v *string) *HeaderValue {
	if o.Prefix = v; o.Prefix == nil {
		o.nullFields = append(o.nullFields, "Prefix")
	}
	return o
}

func (o *HeaderValue) SetRegex(v *string) *HeaderValue {
	if o.Regex = v; o.Regex == nil {
		o.nullFields = append(o.nullFields, "Regex")
	}
	return o
}

//end region

//region Verification

func (o Verification) MarshalJSON() ([]byte, error) {
	type noMethod Verification
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Verification) SetTemplateNames(v []string) *Verification {
	if o.TemplateNames = v; o.TemplateNames == nil {
		o.nullFields = append(o.nullFields, "TemplateNames")
	}
	return o
}

//end region

//region Rolling

func (o Rolling) MarshalJSON() ([]byte, error) {
	type noMethod Rolling
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Rolling) SetSteps(v []*RollingSteps) *Rolling {
	if o.Steps = v; o.Steps == nil {
		o.nullFields = append(o.nullFields, "Steps")
	}
	return o
}

//end region

//region RollingSteps

func (o RollingSteps) MarshalJSON() ([]byte, error) {
	type noMethod RollingSteps
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *RollingSteps) SetName(v *string) *RollingSteps {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *RollingSteps) SetPause(v *Pause) *RollingSteps {
	if o.Pause = v; o.Pause == nil {
		o.nullFields = append(o.nullFields, "Pause")
	}
	return o
}

func (o *RollingSteps) SetVerification(v *Verification) *RollingSteps {
	if o.Verification = v; o.Verification == nil {
		o.nullFields = append(o.nullFields, "Verification")
	}
	return o
}

//end region
