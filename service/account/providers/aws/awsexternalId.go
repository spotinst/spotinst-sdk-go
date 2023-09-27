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

type AwsAccountExternalId struct {
	AccountId  *string `json:"accountId,omitempty"`
	ExternalId *string `json:"externalId,omitempty"`

	forceSendFields []string
	
	nullFields []string
}

func (o *AwsAccountExternalId) SetAccountId(v *string) *AwsAccountExternalId {
	if o.AccountId = v; o.AccountId == nil {
		o.nullFields = append(o.nullFields, "AccountId")
	}
	return o
}
func (o *AwsAccountExternalId) SetExternalId(v *string) *AwsAccountExternalId {
	if o.ExternalId = v; o.ExternalId == nil {
		o.nullFields = append(o.nullFields, "ExternalId")
	}
	return o
}

func (o AwsAccountExternalId) MarshalJSON() ([]byte, error) {
	type noMethod AwsAccountExternalId
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

type CreateAWSAccountExternalIdInput struct {
	AccountID *string `json:"accountId,omitempty"`
}
type CreateAWSAccountExternalIdOutput struct {
	AWSAccountExternalId *AwsAccountExternalId `json:"AWSAccountExternalId,omitempty"`
}

func (s *ServiceOp) CreateAWSAccountExternalId(ctx context.Context, input *CreateAWSAccountExternalIdInput) (*CreateAWSAccountExternalIdOutput, error) {
	r := client.NewRequest(http.MethodPost, "/setup/credentials/aws/externalId")
	if input != nil {
		r.Params.Set("accountId", spotinst.StringValue(input.AccountID))
	}

	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	gs, err := awsAccountExternalIdFromHttpResponse(resp)
	gs[0].AccountId = input.AccountID
	if err != nil {
		return nil, err
	}
	output := new(CreateAWSAccountExternalIdOutput)
	if len(gs) > 0 {
		output.AWSAccountExternalId = gs[0]
	}

	return output, nil
}

type ReadAWSAccountExternalIdInput struct {
	AccountID *string `json:"account,omitempty"`
}
type ReadAWSAccountExternalIdOutput struct {
	AwsAccountExternalId *AwsAccountExternalId `json:"externalId,omitempty"`
}

func (s *ServiceOp) ReadAWSAccountExternalId(ctx context.Context, input *ReadAWSAccountExternalIdInput) (*ReadAWSAccountExternalIdOutput, error) {
	path, err := uritemplates.Expand("/setup/credentials/aws/externalId/{acctId}", uritemplates.Values{"acctId": spotinst.StringValue(input.AccountID)})
	r := client.NewRequest(http.MethodGet, path)
	r.Obj = input
	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := awsAccountExternalIdFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadAWSAccountExternalIdOutput)
	if len(gs) > 0 {
		output.AwsAccountExternalId = gs[0]
	}

	return output, nil

}

func awsAccountExternalIdFromHttpResponse(resp *http.Response) ([]*AwsAccountExternalId, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return awsAccountExternalIdsFromJSON(body)
}

func awsAccountExternalIdsFromJSON(in []byte) ([]*AwsAccountExternalId, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*AwsAccountExternalId, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := awsAccountExternalIdFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func awsAccountExternalIdFromJSON(in []byte) (*AwsAccountExternalId, error) {
	b := new(AwsAccountExternalId)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}
