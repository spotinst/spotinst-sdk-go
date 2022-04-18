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
)

type DataIntegration struct {
	ID              *string `json:"id,omitempty"`
	Config          *Config `json:"config,omitempty"`
	Name            *string `json:"name,omitempty"`
	Vendor          *string `json:"vendor,omitempty"`
	Status          *string `json:"status,omitempty"`
	Health          *string `json:"health,omitempty"`
	LastHealthCheck *string `json:"lastHealthCheck,omitempty"`

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

type Config struct {
	BucketName *string `json:"bucketName,omitempty"`
	SubDir     *string `json:"subdir,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ListDataIntegrationsInput struct{}

type ListDataIntegrationsOutput struct {
	DataIntegrations []*DataIntegration `json:"dataIntegrations,omitempty"`
}

type CreateDataIntegrationInput struct {
	DataIntegration *DataIntegration `json:"dataIntegration,omitempty"`
}

type CreateDataIntegrationOutput struct {
	DataIntegration *DataIntegration `json:"dataIntegration,omitempty"`
}

type ReadDataIntegrationInput struct {
	DataIntegrationId *string `json:"dataIntegrationId,omitempty"`
}

type ReadDataIntegrationOutput struct {
	DataIntegration *DataIntegration `json:"dataIntegration,omitempty"`
}

type UpdateDataIntegrationInput struct {
	DataIntegration *DataIntegration `json:"dataIntegration,omitempty"`
}

type UpdateDataIntegrationOutput struct {
	DataIntegration *DataIntegration `json:"dataIntegration,omitempty"`
}

type DeleteDataIntegrationInput struct {
	DataIntegrationId *string `json:"dataIntegrationId,omitempty"`
}

type DeleteDataIntegrationOutput struct{}

func dataIntegrationFromJSON(in []byte) (*DataIntegration, error) {
	b := new(DataIntegration)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func dataIntegrationsFromJSON(in []byte) ([]*DataIntegration, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*DataIntegration, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := dataIntegrationFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func dataIntegrationsFromHttpResponse(resp *http.Response) ([]*DataIntegration, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return dataIntegrationsFromJSON(body)
}

func (s *ServiceOp) ListDataIntegration(ctx context.Context, input *ListDataIntegrationsInput) (*ListDataIntegrationsOutput, error) {
	r := client.NewRequest(http.MethodGet, "/insights/dataIntegration")
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	di, err := dataIntegrationsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListDataIntegrationsOutput{DataIntegrations: di}, nil
}

func (s *ServiceOp) CreateDataIntegration(ctx context.Context, input *CreateDataIntegrationInput) (*CreateDataIntegrationOutput, error) {
	r := client.NewRequest(http.MethodPost, "/insights/dataIntegration")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	di, err := dataIntegrationsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateDataIntegrationOutput)
	if len(di) > 0 {
		output.DataIntegration = di[0]
	}

	return output, nil
}

func (s *ServiceOp) ReadDataIntegration(ctx context.Context, input *ReadDataIntegrationInput) (*ReadDataIntegrationOutput, error) {
	path, err := uritemplates.Expand("/insights/dataIntegration/{dataIntegrationId}", uritemplates.Values{
		"dataIntegrationId": spotinst.StringValue(input.DataIntegrationId),
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

	di, err := dataIntegrationsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadDataIntegrationOutput)
	if len(di) > 0 {
		output.DataIntegration = di[0]
	}

	return output, nil
}

func (s *ServiceOp) UpdateDataIntegration(ctx context.Context, input *UpdateDataIntegrationInput) (*UpdateDataIntegrationOutput, error) {
	path, err := uritemplates.Expand("/insights/dataIntegration/{dataIntegrationId}", uritemplates.Values{
		"dataIntegrationId": spotinst.StringValue(input.DataIntegration.ID),
	})
	if err != nil {
		return nil, err
	}

	// We do NOT need the ID anymore, so let's drop it.
	input.DataIntegration.ID = nil

	r := client.NewRequest(http.MethodPut, path)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	di, err := dataIntegrationsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(UpdateDataIntegrationOutput)
	if len(di) > 0 {
		output.DataIntegration = di[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteDataIntegration(ctx context.Context, input *DeleteDataIntegrationInput) (*DeleteDataIntegrationOutput, error) {
	path, err := uritemplates.Expand("/insights/dataIntegration/{dataIntegrationId}", uritemplates.Values{
		"dataIntegrationId": spotinst.StringValue(input.DataIntegrationId),
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

	return &DeleteDataIntegrationOutput{}, nil
}

// region DataIntegration

func (o DataIntegration) MarshalJSON() ([]byte, error) {
	type noMethod DataIntegration
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *DataIntegration) SetID(v *string) *DataIntegration {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "Id")
	}
	return o
}

func (o *DataIntegration) SetConfig(v *Config) *DataIntegration {
	if o.Config = v; o.Config == nil {
		o.nullFields = append(o.nullFields, "Config")
	}
	return o
}

func (o *DataIntegration) SetName(v *string) *DataIntegration {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *DataIntegration) SetVendor(v *string) *DataIntegration {
	if o.Vendor = v; o.Vendor == nil {
		o.nullFields = append(o.nullFields, "Vendor")
	}
	return o
}

func (o *DataIntegration) SetStatus(v *string) *DataIntegration {
	if o.Status = v; o.Status == nil {
		o.nullFields = append(o.nullFields, "Status")
	}
	return o
}

func (o *DataIntegration) SetHealth(v *string) *DataIntegration {
	if o.Health = v; o.Health == nil {
		o.nullFields = append(o.nullFields, "Health")
	}
	return o
}

func (o *DataIntegration) SetLastHealthCheck(v *string) *DataIntegration {
	if o.LastHealthCheck = v; o.LastHealthCheck == nil {
		o.nullFields = append(o.nullFields, "LastHealthCheck")
	}
	return o
}

// endregion

// region Config

func (o Config) MarshalJSON() ([]byte, error) {
	type noMethod Config
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Config) SetBucketName(v *string) *Config {
	if o.BucketName = v; o.BucketName == nil {
		o.nullFields = append(o.nullFields, "BucketName")
	}
	return o
}

func (o *Config) SetSubDir(v *string) *Config {
	if o.SubDir = v; o.SubDir == nil {
		o.nullFields = append(o.nullFields, "Subdir")
	}
	return o
}

// endregion
