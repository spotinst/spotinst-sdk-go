package aws

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"net/http"
	"time"

	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/jsonutil"
)

type Credentials struct {
	IamRole   *string `json:"iamRole,omitempty"`
	AccountId *string `json:"accountId,omitempty"`

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

func (o Credentials) MarshalJSON() ([]byte, error) {
	type noMethod Credentials
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

type SetCredentialInput struct {
	Credential *Credentials `json:"credentials,omitempty"`
}
type SetCredentialOutput struct {
	Credential *Credentials `json:"Credentials,omitempty"`
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

func (s *ServiceOp) SetCredential(ctx context.Context, input *SetCredentialInput) error {
	r := client.NewRequest(http.MethodPost, "/setup/credentials/aws")
	r.Params.Set("accountId", spotinst.StringValue(input.Credential.AccountId))
	input.Credential.AccountId = nil
	r.Obj = input
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
