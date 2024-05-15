package oceancd

import (
	"context"
	"encoding/json"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/jsonutil"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/uritemplates"
	"io/ioutil"
	"net/http"
	"time"
)

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

type ListRolloutSpecsOutput struct {
	RolloutSpecs []*RolloutSpec `json:"rolloutSpec,omitempty"`
}

type CreateRolloutSpecInput struct {
	RolloutSpec *RolloutSpec `json:"rolloutSpec,omitempty"`
}

type CreateRolloutSpecOutput struct {
	RolloutSpec *RolloutSpec `json:"rolloutSpec,omitempty"`
}

type ReadRolloutSpecInput struct {
	RolloutSpecName *string `json:"rolloutSpecName,omitempty"`
}

type ReadRolloutSpecOutput struct {
	RolloutSpec *RolloutSpec `json:"rolloutSpec,omitempty"`
}

type UpdateRolloutSpecInput struct {
	RolloutSpec *RolloutSpec `json:"rolloutSpec,omitempty"`
}

type UpdateRolloutSpecOutput struct {
	RolloutSpec *RolloutSpec `json:"rolloutSpec,omitempty"`
}

type PatchRolloutSpecInput struct {
	RolloutSpec *RolloutSpec `json:"rolloutSpec,omitempty"`
}

type PatchRolloutSpecOutput struct {
	RolloutSpec *RolloutSpec `json:"rolloutSpec,omitempty"`
}

type DeleteRolloutSpecInput struct {
	RolloutSpecName *string `json:"rolloutSpecName,omitempty"`
}

type DeleteRolloutSpecOutput struct{}

func rolloutSpecFromJSON(in []byte) (*RolloutSpec, error) {
	b := new(RolloutSpec)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func rolloutSpecsFromJSON(in []byte) ([]*RolloutSpec, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*RolloutSpec, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := rolloutSpecFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func rolloutSpecsFromHttpResponse(resp *http.Response) ([]*RolloutSpec, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return rolloutSpecsFromJSON(body)
}

// endregion

// region API requests

func (s *ServiceOp) ListRolloutSpecs(ctx context.Context) (*ListRolloutSpecsOutput, error) {
	r := client.NewRequest(http.MethodGet, "/ocean/cd/rolloutSpec")
	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	RolloutSpec, err := rolloutSpecsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListRolloutSpecsOutput{RolloutSpecs: RolloutSpec}, nil
}

func (s *ServiceOp) CreateRolloutSpec(ctx context.Context, input *CreateRolloutSpecInput) (*CreateRolloutSpecOutput, error) {
	r := client.NewRequest(http.MethodPost, "/ocean/cd/rolloutSpec")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rs, err := rolloutSpecsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateRolloutSpecOutput)
	if len(rs) > 0 {
		output.RolloutSpec = rs[0]
	}

	return output, nil
}

func (s *ServiceOp) ReadRolloutSpec(ctx context.Context, input *ReadRolloutSpecInput) (*ReadRolloutSpecOutput, error) {
	path, err := uritemplates.Expand("/ocean/cd/rolloutSpec/{rolloutSpecName}", uritemplates.Values{
		"rolloutSpecName": spotinst.StringValue(input.RolloutSpecName),
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

	rolloutSpec, err := rolloutSpecsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadRolloutSpecOutput)
	if len(rolloutSpec) > 0 {
		output.RolloutSpec = rolloutSpec[0]
	}

	return output, nil
}

func (s *ServiceOp) UpdateRolloutSpec(ctx context.Context, input *UpdateRolloutSpecInput) (*UpdateRolloutSpecOutput, error) {
	path, err := uritemplates.Expand("/ocean/cd/rolloutSpec/{rolloutSpecName}", uritemplates.Values{
		"rolloutSpecName": spotinst.StringValue(input.RolloutSpec.Name),
	})
	if err != nil {
		return nil, err
	}

	// We do NOT need the Name anymore, so let's drop it.
	input.RolloutSpec.Name = nil

	r := client.NewRequest(http.MethodPut, path)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rolloutSpec, err := rolloutSpecsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(UpdateRolloutSpecOutput)
	if len(rolloutSpec) > 0 {
		output.RolloutSpec = rolloutSpec[0]
	}

	return output, nil
}

func (s *ServiceOp) PatchRolloutSpec(ctx context.Context, input *PatchRolloutSpecInput) (*PatchRolloutSpecOutput, error) {
	path, err := uritemplates.Expand("/ocean/cd/rolloutSpec/{rolloutSpecName}", uritemplates.Values{
		"rolloutSpecName": spotinst.StringValue(input.RolloutSpec.Name),
	})
	if err != nil {
		return nil, err
	}

	// We do NOT need the Name anymore, so let's drop it.
	input.RolloutSpec.Name = nil

	r := client.NewRequest(http.MethodPatch, path)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rolloutSpec, err := rolloutSpecsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(PatchRolloutSpecOutput)
	if len(rolloutSpec) > 0 {
		output.RolloutSpec = rolloutSpec[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteRolloutSpec(ctx context.Context, input *DeleteRolloutSpecInput) (*DeleteRolloutSpecOutput, error) {
	path, err := uritemplates.Expand("/ocean/cd/rolloutSpec/{rolloutSpecName}", uritemplates.Values{
		"rolloutSpecName": spotinst.StringValue(input.RolloutSpecName),
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

	return &DeleteRolloutSpecOutput{}, nil
}

//end region

// region RolloutSpec

func (o RolloutSpec) MarshalJSON() ([]byte, error) {
	type noMethod RolloutSpec
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *RolloutSpec) SetName(v *string) *RolloutSpec {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *RolloutSpec) SetFailurePolicy(v *FailurePolicy) *RolloutSpec {
	if o.FailurePolicy = v; o.FailurePolicy == nil {
		o.nullFields = append(o.nullFields, "FailurePolicy")
	}
	return o
}

func (o *RolloutSpec) SetSpotDeployment(v *SpotDeployment) *RolloutSpec {
	if o.SpotDeployment = v; o.SpotDeployment == nil {
		o.nullFields = append(o.nullFields, "SpotDeployment")
	}
	return o
}

func (o *RolloutSpec) SetSpotDeployments(v []*SpotDeployment) *RolloutSpec {
	if o.SpotDeployments = v; o.SpotDeployments == nil {
		o.nullFields = append(o.nullFields, "SpotDeployments")
	}
	return o
}

func (o *RolloutSpec) SetStrategy(v *RolloutSpecStrategy) *RolloutSpec {
	if o.Strategy = v; o.Strategy == nil {
		o.nullFields = append(o.nullFields, "Strategy")
	}
	return o
}

func (o *RolloutSpec) SetTraffic(v *Traffic) *RolloutSpec {
	if o.Traffic = v; o.Traffic == nil {
		o.nullFields = append(o.nullFields, "Traffic")
	}
	return o
}

// end region

//region FailurePolicy

func (o FailurePolicy) MarshalJSON() ([]byte, error) {
	type noMethod FailurePolicy
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *FailurePolicy) SetAction(v *string) *FailurePolicy {
	if o.Action = v; o.Action == nil {
		o.nullFields = append(o.nullFields, "Action")
	}
	return o
}

//end region

//region SpotDeployment

func (o SpotDeployment) MarshalJSON() ([]byte, error) {
	type noMethod SpotDeployment
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *SpotDeployment) SetClusterId(v *string) *SpotDeployment {
	if o.ClusterId = v; o.ClusterId == nil {
		o.nullFields = append(o.nullFields, "ClusterId")
	}
	return o
}

func (o *SpotDeployment) SetName(v *string) *SpotDeployment {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *SpotDeployment) SetNamespace(v *string) *SpotDeployment {
	if o.Namespace = v; o.Namespace == nil {
		o.nullFields = append(o.nullFields, "Namespace")
	}
	return o
}

//end region

//region Strategy

func (o RolloutSpecStrategy) MarshalJSON() ([]byte, error) {
	type noMethod RolloutSpecStrategy
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *RolloutSpecStrategy) SetArgs(v []*RolloutSpecArgs) *RolloutSpecStrategy {
	if o.Args = v; o.Args == nil {
		o.nullFields = append(o.nullFields, "Args")
	}
	return o
}

func (o *RolloutSpecStrategy) SetName(v *string) *RolloutSpecStrategy {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

//end region

//region Args

func (o RolloutSpecArgs) MarshalJSON() ([]byte, error) {
	type noMethod RolloutSpecArgs
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *RolloutSpecArgs) SetName(v *string) *RolloutSpecArgs {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *RolloutSpecArgs) SetValue(v *string) *RolloutSpecArgs {
	if o.Value = v; o.Value == nil {
		o.nullFields = append(o.nullFields, "Value")
	}
	return o
}

func (o *RolloutSpecArgs) SetValueFrom(v *RolloutSpecValueFrom) *RolloutSpecArgs {
	if o.ValueFrom = v; o.ValueFrom == nil {
		o.nullFields = append(o.nullFields, "ValueFrom")
	}
	return o
}

//end region

//region ValueFrom

func (o RolloutSpecValueFrom) MarshalJSON() ([]byte, error) {
	type noMethod RolloutSpecValueFrom
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *RolloutSpecValueFrom) SetFieldRef(v *FieldRef) *RolloutSpecValueFrom {
	if o.FieldRef = v; o.FieldRef == nil {
		o.nullFields = append(o.nullFields, "FieldRef")
	}
	return o
}

//end region

//region FieldRef

func (o FieldRef) MarshalJSON() ([]byte, error) {
	type noMethod FieldRef
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *FieldRef) SetFieldPath(v *string) *FieldRef {
	if o.FieldPath = v; o.FieldPath == nil {
		o.nullFields = append(o.nullFields, "FieldPath")
	}
	return o
}

//end region

//region Traffic

func (o Traffic) MarshalJSON() ([]byte, error) {
	type noMethod Traffic
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Traffic) SetAlb(v *Alb) *Traffic {
	if o.Alb = v; o.Alb == nil {
		o.nullFields = append(o.nullFields, "Alb")
	}
	return o
}

func (o *Traffic) SetAmbassador(v *Ambassador) *Traffic {
	if o.Ambassador = v; o.Ambassador == nil {
		o.nullFields = append(o.nullFields, "Ambassador")
	}
	return o
}

func (o *Traffic) SetCanaryService(v *string) *Traffic {
	if o.CanaryService = v; o.CanaryService == nil {
		o.nullFields = append(o.nullFields, "CanaryService")
	}
	return o
}

func (o *Traffic) SetIstio(v *Istio) *Traffic {
	if o.Istio = v; o.Istio == nil {
		o.nullFields = append(o.nullFields, "Istio")
	}
	return o
}

func (o *Traffic) SetNginx(v *Nginx) *Traffic {
	if o.Nginx = v; o.Nginx == nil {
		o.nullFields = append(o.nullFields, "Nginx")
	}
	return o
}

func (o *Traffic) SetPingPong(v *PingPong) *Traffic {
	if o.PingPong = v; o.PingPong == nil {
		o.nullFields = append(o.nullFields, "PingPong")
	}
	return o
}

func (o *Traffic) SetSmi(v *Smi) *Traffic {
	if o.Smi = v; o.Smi == nil {
		o.nullFields = append(o.nullFields, "Smi")
	}
	return o
}

func (o *Traffic) SetStableService(v *string) *Traffic {
	if o.StableService = v; o.StableService == nil {
		o.nullFields = append(o.nullFields, "StableService")
	}
	return o
}

//end region

//region Alb

func (o Alb) MarshalJSON() ([]byte, error) {
	type noMethod Alb
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Alb) SetAnnotationPrefix(v *string) *Alb {
	if o.AnnotationPrefix = v; o.AnnotationPrefix == nil {
		o.nullFields = append(o.nullFields, "AnnotationPrefix")
	}
	return o
}

func (o *Alb) SetIngress(v *string) *Alb {
	if o.Ingress = v; o.Ingress == nil {
		o.nullFields = append(o.nullFields, "Ingress")
	}
	return o
}

func (o *Alb) SetRootService(v *string) *Alb {
	if o.RootService = v; o.RootService == nil {
		o.nullFields = append(o.nullFields, "RootService")
	}
	return o
}

func (o *Alb) SetServicePort(v *int) *Alb {
	if o.ServicePort = v; o.ServicePort == nil {
		o.nullFields = append(o.nullFields, "ServicePort")
	}
	return o
}

func (o *Alb) SetStickinessConfig(v *StickinessConfig) *Alb {
	if o.StickinessConfig = v; o.StickinessConfig == nil {
		o.nullFields = append(o.nullFields, "StickinessConfig")
	}
	return o
}

//end region

//region StickinessConfig

func (o StickinessConfig) MarshalJSON() ([]byte, error) {
	type noMethod StickinessConfig
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *StickinessConfig) SetDurationSeconds(v *int) *StickinessConfig {
	if o.DurationSeconds = v; o.DurationSeconds == nil {
		o.nullFields = append(o.nullFields, "DurationSeconds")
	}
	return o
}

func (o *StickinessConfig) SetEnabled(v *bool) *StickinessConfig {
	if o.Enabled = v; o.Enabled == nil {
		o.nullFields = append(o.nullFields, "Enabled")
	}
	return o
}

//end region

//region Ambassador

func (o Ambassador) MarshalJSON() ([]byte, error) {
	type noMethod Ambassador
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Ambassador) SetSteps(v []string) *Ambassador {
	if o.Mappings = v; o.Mappings == nil {
		o.nullFields = append(o.nullFields, "Mappings")
	}
	return o
}

//end region

//region Istio

func (o Istio) MarshalJSON() ([]byte, error) {
	type noMethod Istio
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Istio) SetDestinationRune(v *DestinationRule) *Istio {
	if o.DestinationRule = v; o.DestinationRule == nil {
		o.nullFields = append(o.nullFields, "DestinationRule")
	}
	return o
}

func (o *Istio) SetVirtualServices(v []*VirtualServices) *Istio {
	if o.VirtualServices = v; o.VirtualServices == nil {
		o.nullFields = append(o.nullFields, "VirtualServices")
	}
	return o
}

//end region

//region DestinationRule

func (o DestinationRule) MarshalJSON() ([]byte, error) {
	type noMethod DestinationRule
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *DestinationRule) SetCanarySubsetName(v *string) *DestinationRule {
	if o.CanarySubsetName = v; o.CanarySubsetName == nil {
		o.nullFields = append(o.nullFields, "CanarySubsetName")
	}
	return o
}

func (o *DestinationRule) SetName(v *string) *DestinationRule {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *DestinationRule) SetStableSubsetName(v *string) *DestinationRule {
	if o.StableSubsetName = v; o.StableSubsetName == nil {
		o.nullFields = append(o.nullFields, "StableSubsetName")
	}
	return o
}

//end region

//region VirtualServices

func (o VirtualServices) MarshalJSON() ([]byte, error) {
	type noMethod VirtualServices
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *VirtualServices) SetName(v *string) *VirtualServices {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *VirtualServices) SetRoutes(v []string) *VirtualServices {
	if o.Routes = v; o.Routes == nil {
		o.nullFields = append(o.nullFields, "Routes")
	}
	return o
}

func (o *VirtualServices) SetTlsRoutes(v []*TlsRoutes) *VirtualServices {
	if o.TlsRoutes = v; o.TlsRoutes == nil {
		o.nullFields = append(o.nullFields, "TlsRoutes")
	}
	return o
}

//end region

//region TlsRoutes

func (o TlsRoutes) MarshalJSON() ([]byte, error) {
	type noMethod TlsRoutes
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *TlsRoutes) SetPort(v *int) *TlsRoutes {
	if o.Port = v; o.Port == nil {
		o.nullFields = append(o.nullFields, "Port")
	}
	return o
}

func (o *TlsRoutes) SetSniHosts(v []string) *TlsRoutes {
	if o.SniHosts = v; o.SniHosts == nil {
		o.nullFields = append(o.nullFields, "SniHosts")
	}
	return o
}

//end region

//region Nginx

func (o Nginx) MarshalJSON() ([]byte, error) {
	type noMethod Nginx
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Nginx) SetAdditionalIngressAnnotation(v *AdditionalIngressAnnotations) *Nginx {
	if o.AdditionalIngressAnnotations = v; o.AdditionalIngressAnnotations == nil {
		o.nullFields = append(o.nullFields, "AdditionalIngressAnnotations")
	}
	return o
}

func (o *Nginx) SetAnnotationPreffix(v *string) *Nginx {
	if o.AnnotationPrefix = v; o.AnnotationPrefix == nil {
		o.nullFields = append(o.nullFields, "AnnotationPrefix")
	}
	return o
}

func (o *Nginx) SetStableIngress(v *string) *Nginx {
	if o.StableIngress = v; o.StableIngress == nil {
		o.nullFields = append(o.nullFields, "StableIngress")
	}
	return o
}

//end region

//region AdditionalIngressAnnotations

func (o AdditionalIngressAnnotations) MarshalJSON() ([]byte, error) {
	type noMethod AdditionalIngressAnnotations
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AdditionalIngressAnnotations) SetCanaryByHeader(v *string) *AdditionalIngressAnnotations {
	if o.CanaryByHeader = v; o.CanaryByHeader == nil {
		o.nullFields = append(o.nullFields, "CanaryByHeader")
	}
	return o
}

func (o *AdditionalIngressAnnotations) SetKey1(v *string) *AdditionalIngressAnnotations {
	if o.Key1 = v; o.Key1 == nil {
		o.nullFields = append(o.nullFields, "Key1")
	}
	return o
}

//end region

//region PingPong

func (o PingPong) MarshalJSON() ([]byte, error) {
	type noMethod PingPong
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *PingPong) SetPingService(v *string) *PingPong {
	if o.PingService = v; o.PingService == nil {
		o.nullFields = append(o.nullFields, "PingService")
	}
	return o
}

func (o *PingPong) SetPongService(v *string) *PingPong {
	if o.PongService = v; o.PongService == nil {
		o.nullFields = append(o.nullFields, "PongService")
	}
	return o
}

//end region

//region Smi

func (o Smi) MarshalJSON() ([]byte, error) {
	type noMethod Smi
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Smi) SetRootService(v *string) *Smi {
	if o.RootService = v; o.RootService == nil {
		o.nullFields = append(o.nullFields, "RootService")
	}
	return o
}

func (o *Smi) SetTrafficSplitName(v *string) *Smi {
	if o.TrafficSplitName = v; o.TrafficSplitName == nil {
		o.nullFields = append(o.nullFields, "TrafficSplitName")
	}
	return o
}

//end region
