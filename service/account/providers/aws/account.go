package aws

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/jsonutil"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/uritemplates"
)

type Account struct {
	ID                 *string `json:"id,omitempty"`
	Name               *string `json:"name,omitempty"`
	OrganizationId     *string `json:"organizationId,omitempty"`
	AccountId          *string `json:"accountId,omitempty"`
	ProviderExternalId *string `json:"providerExternalId,omitempty"`
	CloudAccountId     *string `json:"cloudAccountId,omitempty"`
	ExternalId         *string `json:"externalId,omitempty"`

	// Read-only fields.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

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

type ListAccounts struct {
	AccountId          *string `json:"accountId,omitempty"`
	Name               *string `json:"name,omitempty"`
	OrganizationId     *string `json:"organizationId,omitempty"`
	ProviderExternalId *string `json:"providerExternalId,omitempty"`
	CloudAccountId     *string `json:"cloudAccountId,omitempty"`

	forceSendFields []string
	nullFields      []string
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

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
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

	output1, err := s.CreateAwsAccountExternalId(ctx, output.Account.ID)
	output.Account.ExternalId = output1.AWSAccountExternalId.ExternalId

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

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteAccountOutput{}, nil
}

type ReadAccountInput struct {
	AccountID *string `json:"accountId,omitempty"`
}
type ReadAccountOutput struct {
	Account *Account `json:"account,omitempty"`
}

func (s *ServiceOp) ReadAccount(ctx context.Context, input *ReadAccountInput) (*ReadAccountOutput, error) {
	r := client.NewRequest(http.MethodGet, "/setup/account")
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
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
		for i, value := range gs {
			if spotinst.StringValue(input.AccountID) == spotinst.StringValue(value.AccountId) {
				output.Account = gs[i]
				break
			}
		}
	}

	return output, nil
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

func (o ListAccounts) MarshalJSON() ([]byte, error) {
	type noMethod ListAccounts
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *ListAccounts) SetCloudAccountId(v *string) *ListAccounts {
	if o.CloudAccountId = v; o.CloudAccountId == nil {
		o.nullFields = append(o.nullFields, "CloudAccountId")
	}
	return o
}

type ListAccountsInput struct {
	ListAccounts *ListAccounts `json:"account,omitempty"`
}
type ListAccountsOutput struct {
	ListAccounts []*ListAccounts `json:"account,omitempty"`
}

func (s *ServiceOp) ListAccounts(ctx context.Context, input *ListAccountsInput) (*ListAccountsOutput, error) {
	r := client.NewRequest(http.MethodGet, "/setup/account")
	r.Obj = input

	if input.ListAccounts.CloudAccountId != nil {
		r.Params.Set("cloudAccountId", spotinst.StringValue(input.ListAccounts.CloudAccountId))
	}

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := getAccountsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListAccountsOutput{ListAccounts: gs}, nil
}

func getAccountsFromHttpResponse(resp *http.Response) ([]*ListAccounts, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return getAccountsFromJSON(body)
}

func getAccountsFromJSON(in []byte) ([]*ListAccounts, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*ListAccounts, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := getAccountFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func getAccountFromJSON(in []byte) (*ListAccounts, error) {
	b := new(ListAccounts)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}
