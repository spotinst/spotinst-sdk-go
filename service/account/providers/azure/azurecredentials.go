package azure

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
	AccountId      *string `json:"accountId,omitempty"`
	ClientId       *string `json:"clientId,omitempty"`
	ClientSecret   *string `json:"clientSecret,omitempty"`
	TenantId       *string `json:"tenantId,omitempty"`
	SubscriptionId *string `json:"subscriptionId,omitempty"`

	forceSendFields []string

	nullFields []string
}

func (o Credentials) MarshalJSON() ([]byte, error) {
	type noMethod Credentials
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Credentials) SetClientId(v *string) *Credentials {
	if o.ClientId = v; o.ClientId == nil {
		o.nullFields = append(o.nullFields, "ClientId")
	}
	return o
}

func (o *Credentials) SetClientSecret(v *string) *Credentials {
	if o.ClientSecret = v; o.ClientSecret == nil {
		o.nullFields = append(o.nullFields, "ClientSecret")
	}
	return o
}

func (o *Credentials) SetTenantId(v *string) *Credentials {
	if o.TenantId = v; o.TenantId == nil {
		o.nullFields = append(o.nullFields, "TenantId")
	}
	return o
}

func (o *Credentials) SetSubscriptionId(v *string) *Credentials {
	if o.SubscriptionId = v; o.SubscriptionId == nil {
		o.nullFields = append(o.nullFields, "SubscriptionId")
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

func (s *ServiceOp) SetCredentials(ctx context.Context, input *SetCredentialsInput) (*SetCredentialsOutput, error) {
	r := client.NewRequest(http.MethodPost, "/azure/setup/credentials")

	if input != nil {
		r.Params.Set("accountId", spotinst.StringValue(input.Credentials.AccountId))
	}
	input.Credentials.AccountId = nil
	r.Obj = input.Credentials

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
	r := client.NewRequest(http.MethodGet, "/azure/setup/credentials")
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
