package gcp

import (
	"context"
	"encoding/json"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/jsonutil"
	"io/ioutil"
	"net/http"
)

type ServiceAccounts struct {
	AccountId               *string `json:"accountId,omitempty"`
	Type                    *string `json:"type,omitempty"`
	ProjectId               *string `json:"project_id,omitempty"`
	PrivateKeyId            *string `json:"private_key_id,omitempty"`
	PrivateKey              *string `json:"private_key,omitempty"`
	ClientEmail             *string `json:"client_email,omitempty"`
	ClientId                *string `json:"client_id,omitempty"`
	AuthUri                 *string `json:"auth_uri,omitempty"`
	TokenUri                *string `json:"token_uri,omitempty"`
	AuthProviderX509CertUrl *string `json:"auth_provider_x509_cert_url,omitempty"`
	ClientX509CertUrl       *string `json:"client_x509_cert_url,omitempty"`
	forceSendFields         []string
	nullFields              []string
}

func (o ServiceAccounts) MarshalJSON() ([]byte, error) {
	type noMethod ServiceAccounts
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *ServiceAccounts) SetAccountId(v *string) *ServiceAccounts {
	if o.AccountId = v; o.AccountId == nil {
		o.nullFields = append(o.nullFields, "AccountId")
	}
	return o
}

func (o *ServiceAccounts) SetType(v *string) *ServiceAccounts {
	if o.Type = v; o.Type == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

func (o *ServiceAccounts) SetProjectId(v *string) *ServiceAccounts {
	if o.ProjectId = v; o.ProjectId == nil {
		o.nullFields = append(o.nullFields, "ProjectId")
	}
	return o
}

func (o *ServiceAccounts) SetPrivateKeyId(v *string) *ServiceAccounts {
	if o.PrivateKeyId = v; o.PrivateKeyId == nil {
		o.nullFields = append(o.nullFields, "PrivateKeyId")
	}
	return o
}

func (o *ServiceAccounts) SetPrivateKey(v *string) *ServiceAccounts {
	if o.PrivateKey = v; o.PrivateKey == nil {
		o.nullFields = append(o.nullFields, "PrivateKey")
	}
	return o
}

func (o *ServiceAccounts) SetClientEmail(v *string) *ServiceAccounts {
	if o.ClientEmail = v; o.ClientEmail == nil {
		o.nullFields = append(o.nullFields, "ClientEmail")
	}
	return o
}

func (o *ServiceAccounts) SetClientId(v *string) *ServiceAccounts {
	if o.ClientId = v; o.ClientId == nil {
		o.nullFields = append(o.nullFields, "ClientId")
	}
	return o
}

func (o *ServiceAccounts) SetAuthUri(v *string) *ServiceAccounts {
	if o.AuthUri = v; o.AuthUri == nil {
		o.nullFields = append(o.nullFields, "AuthUri")
	}
	return o
}

func (o *ServiceAccounts) SetTokenUri(v *string) *ServiceAccounts {
	if o.TokenUri = v; o.TokenUri == nil {
		o.nullFields = append(o.nullFields, "TokenUri")
	}
	return o
}

func (o *ServiceAccounts) SetAuthProviderX509CertUrl(v *string) *ServiceAccounts {
	if o.AuthProviderX509CertUrl = v; o.AuthProviderX509CertUrl == nil {
		o.nullFields = append(o.nullFields, "AuthProviderX509CertUrl")
	}
	return o
}

func (o *ServiceAccounts) SetClientX509CertUrl(v *string) *ServiceAccounts {
	if o.ClientX509CertUrl = v; o.ClientX509CertUrl == nil {
		o.nullFields = append(o.nullFields, "ClientX509CertUrl")
	}
	return o
}

type SetServiceAccountsInput struct {
	ServiceAccounts *ServiceAccounts `json:"serviceAccount,omitempty"`
}
type SetServiceAccountsOutput struct {
	ServiceAccounts *ServiceAccounts `json:"serviceAccount,omitempty"`
}

func (s *ServiceOp) SetServiceAccount(ctx context.Context, input *SetServiceAccountsInput) (*SetServiceAccountsOutput, error) {
	r := client.NewRequest(http.MethodPost, "/gcp/setup/credentials")

	if input != nil {
		r.Params.Set("accountId", spotinst.StringValue(input.ServiceAccounts.AccountId))
	}
	input.ServiceAccounts.AccountId = nil
	r.Obj = input

	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := serviceAccountsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(SetServiceAccountsOutput)
	if len(gs) > 0 {
		output.ServiceAccounts = gs[0]
	}

	return output, nil
}

type ReadServiceAccountsInput struct {
	AccountId *string `json:"accountId,omitempty"`
}
type ReadServiceAccountsOutput struct {
	ServiceAccounts *ServiceAccounts `json:"serviceAccount,omitempty"`
}

func (s *ServiceOp) ReadServiceAccount(ctx context.Context, input *ReadServiceAccountsInput) (*ReadServiceAccountsOutput, error) {
	r := client.NewRequest(http.MethodGet, "/gcp/setup/credentials")
	if input != nil {
		r.Params.Set("accountId", spotinst.StringValue(input.AccountId))
	}

	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := serviceAccountsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadServiceAccountsOutput)
	if len(gs) > 0 {
		output.ServiceAccounts = gs[0]
	}

	return output, nil
}

func serviceAccountsFromHttpResponse(resp *http.Response) ([]*ServiceAccounts, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return serviceAccountsFromJSON(body)
}

func serviceAccountsFromJSON(in []byte) ([]*ServiceAccounts, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*ServiceAccounts, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		var objmap map[string]json.RawMessage
		err := json.Unmarshal(rb, &objmap)
		b, err := serviceAccountFromJSON(objmap["serviceAccount"])
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func serviceAccountFromJSON(in []byte) (*ServiceAccounts, error) {
	b := new(ServiceAccounts)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}
