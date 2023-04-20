package aws

import (
	"context"
	"encoding/json"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/jsonutil"
	"io/ioutil"
	"net/http"
	"time"
)

type Credential struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`

	// Read-only fields.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
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

func (o Credential) MarshalJSON() ([]byte, error) {
	type noMethod Credential
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

type CreateCredentialInput struct {
	Credential *Credential `json:"credential,omitempty"`
}
type CreateCredentialOutput struct {
	Credential *Credential `json:"Credential,omitempty"`
}

func (s *ServiceOp) CreateCredential(ctx context.Context, input *CreateCredentialInput) (*CreateCredentialOutput, error) {
	r := client.NewRequest(http.MethodPost, "/setup/account")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := credentialsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateCredentialOutput)
	if len(gs) > 0 {
		output.Credential = gs[0]
	}

	return output, nil
}

func credentialsFromHttpResponse(resp *http.Response) ([]*Credential, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return credentialsFromJSON(body)
}

func credentialsFromJSON(in []byte) ([]*Credential, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Credential, len(rw.Response.Items))
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

func credentialFromJSON(in []byte) (*Credential, error) {
	b := new(Credential)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}
