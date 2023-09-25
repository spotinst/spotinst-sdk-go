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
	"time"
)

type Account struct {
	ID                 *string `json:"id,omitempty"`
	Name               *string `json:"name,omitempty"`
	OrganizationId     *string `json:"organizationId,omitempty"`
	AccountId          *string `json:"accountId,omitempty"`
	CloudProvider      *string `json:"cloudProvider,omitempty"`
	ProviderExternalId *string `json:"providerExternalId,omitempty"`

	// Read-only fields.
	CreatedAt *time.Time `json:"createdAt,omitempty"`

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

func (o *Account) SetId(v *string) *Account {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}
func (o *Account) SetName(v *string) *Account {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o Account) MarshalJSON() ([]byte, error) {
	type noMethod Account
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

type CreateAccountInput struct {
	Account *Account `json:"account,omitempty"`
}
type CreateAccountOutput struct {
	Account *Account `json:"account,omitempty"`
}

func (s *ServiceOp) CreateAccount(ctx context.Context, input *CreateAccountInput) (*CreateAccountOutput, error) {
	r := client.NewRequest(http.MethodPost, "/setup/account")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := accountsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateAccountOutput)
	if len(gs) > 0 {
		output.Account = gs[0]
	}

	return output, nil
}

type ReadAccountInput struct {
	AccountID *string `json:"account,omitempty"`
}
type ReadAccountOutput struct {
	Account *Account `json:"account,omitempty"`
}

func (s *ServiceOp) ReadAccount(ctx context.Context, input *ReadAccountInput) (*ReadAccountOutput, error) {
	path, err := uritemplates.Expand("/setup/account/{acctId}", uritemplates.Values{"acctId": spotinst.StringValue(input.AccountID)})
	r := client.NewRequest(http.MethodGet, path)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.DoOrg(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := accountsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}
	output := new(ReadAccountOutput)
	if len(gs) > 0 {
		output.Account = gs[0]
	}

	return output, nil

}

func accountsFromHttpResponse(resp *http.Response) ([]*Account, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return accountsFromJSON(body)
}

func accountsFromJSON(in []byte) ([]*Account, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Account, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := accountFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func accountFromJSON(in []byte) (*Account, error) {
	b := new(Account)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

type DeleteAccountInput struct {
	AccountID *string `json:"accountId,omitempty"`
}

type DeleteAccountOutput struct{}

func (s *ServiceOp) DeleteAccount(ctx context.Context, input *DeleteAccountInput) (*DeleteAccountOutput, error) {
	path, err := uritemplates.Expand("/setup/account/{accountId}", uritemplates.Values{
		"accountId": spotinst.StringValue(input.AccountID),
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

	return &DeleteAccountOutput{}, nil
}
