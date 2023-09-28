package aws

import (
	"context"
	"encoding/json"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/jsonutil"
	"io/ioutil"
	"net/http"
)

type Credentials struct {
	IamRole   *string `json:"iamRole,omitempty"`
	AccountId *string `json:"accountId,omitempty"`

	forceSendFields []string

	nullFields []string
}

func (o Credentials) MarshalJSON() ([]byte, error) {
	type noMethod Credentials
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Credentials) SetIamRole(v *string) *Credentials {
	if o.IamRole = v; o.IamRole == nil {
		o.nullFields = append(o.nullFields, "IamRole")
	}
	return o
}
func (o *Credentials) SetAccountId(v *string) *Credentials {
	if o.AccountId = v; o.AccountId == nil {
		o.nullFields = append(o.nullFields, "AccountId")
	}
	return o
}

type SetCredentialsInput struct {
	Credentials *Credentials `json:"credentials,omitempty"`
}
type SetCredentialsOutput struct {
	Credentials *Credentials `json:"Credentials,omitempty"`
}

func (s *ServiceOp) Credentials(ctx context.Context, input *SetCredentialsInput) (*SetCredentialsOutput, error) {
	r := client.NewRequest(http.MethodPost, "/setup/credentials/aws")

	if input != nil {
		r.Params.Set("accountId", spotinst.StringValue(input.Credentials.AccountId))
	}
	input.Credentials.AccountId = nil
	r.Obj = input

	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := credentialsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(SetCredentialsOutput)
	if len(gs) > 0 {
		output.Credentials = gs[0]
	}

	return output, nil
}

type ReadCredentialsInput struct {
	AccountId *string `json:"accountId,omitempty"`
}
type ReadCredentialsOutput struct {
	Credentials *Credentials `json:"Credentials,omitempty"`
}

func (s *ServiceOp) ReadCredentials(ctx context.Context, input *ReadCredentialsInput) (*ReadCredentialsOutput, error) {
	r := client.NewRequest(http.MethodGet, "/setup/credentials/aws")
	if input != nil {
		r.Params.Set("accountId", spotinst.StringValue(input.AccountId))
	}

	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := credentialsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadCredentialsOutput)
	if len(gs) > 0 {
		output.Credentials = gs[0]
	}

	return output, nil
}

func credentialsFromHttpResponse(resp *http.Response) ([]*Credentials, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return credentialsFromJSON(body)
}

func credentialsFromJSON(in []byte) ([]*Credentials, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Credentials, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := credentialFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func credentialFromJSON(in []byte) (*Credentials, error) {
	b := new(Credentials)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}
