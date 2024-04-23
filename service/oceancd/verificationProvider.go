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

type VerificationProvider struct {
	CloudWatch *CloudWatch `json:"cloudWatch,omitempty"`
	ClusterIDs []string    `json:"clusterIds,omitempty"`
	DataDog    *DataDog    `json:"datadog,omitempty"`
	Jenkins    *Jenkins    `json:"jenkins,omitempty"`
	Name       *string     `json:"name,omitempty"`
	NewRelic   *NewRelic   `json:"newRelic,omitempty"`
	Prometheus *Prometheus `json:"prometheus,omitempty"`

	// Read-only fields.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type CloudWatch struct {
	IAmArn *string `json:"iamArn,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type DataDog struct {
	Address *string `json:"address,omitempty"`
	ApiKey  *string `json:"apiKey,omitempty"`
	AppKey  *string `json:"appKey,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Jenkins struct {
	ApiToken *string `json:"apiToken,omitempty"`
	BaseUrl  *string `json:"baseUrl,omitempty"`
	UserName *string `json:"username,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type NewRelic struct {
	AccountId        *string `json:"accountId,omitempty"`
	BaseUrlNerdGraph *string `json:"baseUrlNerdGraph,omitempty"`
	BaseUrlRest      *string `json:"baseUrlRest,omitempty"`
	PersonalApiKey   *string `json:"personalApiKey,omitempty"`
	Region           *string `json:"region,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Prometheus struct {
	Address *string `json:"address,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ListVerificationProvidersOutput struct {
	VerificationProviders []*VerificationProvider `json:"verificationProvider,omitempty"`
}

type CreateVerificationProviderInput struct {
	VerificationProvider *VerificationProvider `json:"verificationProvider,omitempty"`
}

type CreateVerificationProviderOutput struct {
	VerificationProvider *VerificationProvider `json:"verificationProvider,omitempty"`
}

type ReadVerificationProviderInput struct {
	Name *string `json:"name,omitempty"`
}

type ReadVerificationProviderOutput struct {
	VerificationProvider *VerificationProvider `json:"verificationProvider,omitempty"`
}

type UpdateVerificationProviderInput struct {
	VerificationProvider *VerificationProvider `json:"verificationProvider,omitempty"`
}

type UpdateVerificationProviderOutput struct {
	VerificationProvider *VerificationProvider `json:"verificationProvider,omitempty"`
}

type PatchVerificationProviderInput struct {
	VerificationProvider *VerificationProvider `json:"verificationProvider,omitempty"`
}

type PatchVerificationProviderOutput struct {
	VerificationProvider *VerificationProvider `json:"verificationProvider,omitempty"`
}

type DeleteVerificationProviderInput struct {
	Name *string `json:"name,omitempty"`
}

type DeleteVerificationProviderOutput struct{}

func verificationProviderFromJSON(in []byte) (*VerificationProvider, error) {
	b := new(VerificationProvider)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func verificationProvidersFromJSON(in []byte) ([]*VerificationProvider, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*VerificationProvider, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := verificationProviderFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func verificationProvidersFromHttpResponse(resp *http.Response) ([]*VerificationProvider, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return verificationProvidersFromJSON(body)
}

// endregion

// region API requests

func (s *ServiceOp) ListVerificationProviders(ctx context.Context) (*ListVerificationProvidersOutput, error) {
	r := client.NewRequest(http.MethodGet, "/ocean/cd/verificationProvider")
	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	vp, err := verificationProvidersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListVerificationProvidersOutput{VerificationProviders: vp}, nil
}

func (s *ServiceOp) CreateVerificationProvider(ctx context.Context, input *CreateVerificationProviderInput) (*CreateVerificationProviderOutput, error) {
	r := client.NewRequest(http.MethodPost, "/ocean/cd/verificationProvider")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	vp, err := verificationProvidersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateVerificationProviderOutput)
	if len(vp) > 0 {
		output.VerificationProvider = vp[0]
	}

	return output, nil
}

func (s *ServiceOp) ReadVerificationProvider(ctx context.Context, input *ReadVerificationProviderInput) (*ReadVerificationProviderOutput, error) {
	path, err := uritemplates.Expand("/ocean/cd/verificationProvider/{name}", uritemplates.Values{
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

	vp, err := verificationProvidersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadVerificationProviderOutput)
	if len(vp) > 0 {
		output.VerificationProvider = vp[0]
	}

	return output, nil
}

func (s *ServiceOp) UpdateVerificationProvider(ctx context.Context, input *UpdateVerificationProviderInput) (*UpdateVerificationProviderOutput, error) {
	path, err := uritemplates.Expand("/ocean/cd/verificationProvider/{name}", uritemplates.Values{
		"name": spotinst.StringValue(input.VerificationProvider.Name),
	})
	if err != nil {
		return nil, err
	}

	// We do NOT need the Name anymore, so let's drop it.
	input.VerificationProvider.Name = nil

	r := client.NewRequest(http.MethodPut, path)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	vp, err := verificationProvidersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(UpdateVerificationProviderOutput)
	if len(vp) > 0 {
		output.VerificationProvider = vp[0]
	}

	return output, nil
}

func (s *ServiceOp) PatchVerificationProvider(ctx context.Context, input *PatchVerificationProviderInput) (*PatchVerificationProviderOutput, error) {
	path, err := uritemplates.Expand("/ocean/cd/verificationProvider/{name}", uritemplates.Values{
		"name": spotinst.StringValue(input.VerificationProvider.Name),
	})
	if err != nil {
		return nil, err
	}

	// We do NOT need the Name anymore, so let's drop it.
	input.VerificationProvider.Name = nil

	r := client.NewRequest(http.MethodPatch, path)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	vp, err := verificationProvidersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(PatchVerificationProviderOutput)
	if len(vp) > 0 {
		output.VerificationProvider = vp[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteVerificationProvider(ctx context.Context, input *DeleteVerificationProviderInput) (*DeleteVerificationProviderOutput, error) {
	path, err := uritemplates.Expand("/ocean/cd/verificationProvider/{name}", uritemplates.Values{
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

	return &DeleteVerificationProviderOutput{}, nil
}

//end region

// region Verification Provider

func (o VerificationProvider) MarshalJSON() ([]byte, error) {
	type noMethod VerificationProvider
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *VerificationProvider) SetCloudWatch(v *CloudWatch) *VerificationProvider {
	if o.CloudWatch = v; o.CloudWatch == nil {
		o.nullFields = append(o.nullFields, "CloudWatch")
	}
	return o
}

func (o *VerificationProvider) SetClusterIDs(v []string) *VerificationProvider {
	if o.ClusterIDs = v; o.ClusterIDs == nil {
		o.nullFields = append(o.nullFields, "ClusterIDs")
	}
	return o
}

func (o *VerificationProvider) SetDataDog(v *DataDog) *VerificationProvider {
	if o.DataDog = v; o.DataDog == nil {
		o.nullFields = append(o.nullFields, "DataDog")
	}
	return o
}

func (o *VerificationProvider) SetJenkins(v *Jenkins) *VerificationProvider {
	if o.Jenkins = v; o.Jenkins == nil {
		o.nullFields = append(o.nullFields, "Jenkins")
	}
	return o
}

func (o *VerificationProvider) SetName(v *string) *VerificationProvider {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *VerificationProvider) SetNewRelic(v *NewRelic) *VerificationProvider {
	if o.NewRelic = v; o.NewRelic == nil {
		o.nullFields = append(o.nullFields, "NewRelic")
	}
	return o
}

func (o *VerificationProvider) SetPrometheus(v *Prometheus) *VerificationProvider {
	if o.Prometheus = v; o.Prometheus == nil {
		o.nullFields = append(o.nullFields, "Prometheus")
	}
	return o
}

// end region

//region CloudWatch

func (o CloudWatch) MarshalJSON() ([]byte, error) {
	type noMethod CloudWatch
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *CloudWatch) SetIAmArn(v *string) *CloudWatch {
	if o.IAmArn = v; o.IAmArn == nil {
		o.nullFields = append(o.nullFields, "IAmArn")
	}
	return o
}

//end region

//region DataDog

func (o DataDog) MarshalJSON() ([]byte, error) {
	type noMethod DataDog
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *DataDog) SetAddress(v *string) *DataDog {
	if o.Address = v; o.Address == nil {
		o.nullFields = append(o.nullFields, "Address")
	}
	return o
}

func (o *DataDog) SetApiKey(v *string) *DataDog {
	if o.ApiKey = v; o.ApiKey == nil {
		o.nullFields = append(o.nullFields, "ApiKey")
	}
	return o
}

func (o *DataDog) SetAppKey(v *string) *DataDog {
	if o.AppKey = v; o.AppKey == nil {
		o.nullFields = append(o.nullFields, "AppKey")
	}
	return o
}

//end region

//region Jenkins

func (o Jenkins) MarshalJSON() ([]byte, error) {
	type noMethod Jenkins
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Jenkins) SetApiToken(v *string) *Jenkins {
	if o.ApiToken = v; o.ApiToken == nil {
		o.nullFields = append(o.nullFields, "ApiToken")
	}
	return o
}

func (o *Jenkins) SetBaseUrl(v *string) *Jenkins {
	if o.BaseUrl = v; o.BaseUrl == nil {
		o.nullFields = append(o.nullFields, "BaseUrl")
	}
	return o
}

func (o *Jenkins) SetUserName(v *string) *Jenkins {
	if o.UserName = v; o.UserName == nil {
		o.nullFields = append(o.nullFields, "UserName")
	}
	return o
}

//end region

//region NewRelic

func (o NewRelic) MarshalJSON() ([]byte, error) {
	type noMethod NewRelic
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *NewRelic) SetAccountId(v *string) *NewRelic {
	if o.AccountId = v; o.AccountId == nil {
		o.nullFields = append(o.nullFields, "AccountId")
	}
	return o
}

func (o *NewRelic) SetBaseUrlNerdGraph(v *string) *NewRelic {
	if o.BaseUrlNerdGraph = v; o.BaseUrlNerdGraph == nil {
		o.nullFields = append(o.nullFields, "BaseUrlNerdGraph")
	}
	return o
}

func (o *NewRelic) SetBaseUrlRest(v *string) *NewRelic {
	if o.BaseUrlRest = v; o.BaseUrlRest == nil {
		o.nullFields = append(o.nullFields, "BaseUrlRest")
	}
	return o
}

func (o *NewRelic) SetPersonalApiKey(v *string) *NewRelic {
	if o.PersonalApiKey = v; o.PersonalApiKey == nil {
		o.nullFields = append(o.nullFields, "PersonalApiKey")
	}
	return o
}

func (o *NewRelic) SetRegion(v *string) *NewRelic {
	if o.Region = v; o.Region == nil {
		o.nullFields = append(o.nullFields, "Region")
	}
	return o
}

//end region

//region Prometheus

func (o Prometheus) MarshalJSON() ([]byte, error) {
	type noMethod Prometheus
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Prometheus) SetAddress(v *string) *Prometheus {
	if o.Address = v; o.Address == nil {
		o.nullFields = append(o.nullFields, "Address")
	}
	return o
}

//end region
