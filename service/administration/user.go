package administration

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/jsonutil"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/uritemplates"
)

type User struct {
	Email     *string       `json:"email,omitempty"`
	FirstName *string       `json:"firstName,omitempty"`
	LastName  *string       `json:"lastName,omitempty"`
	Password  *string       `json:"password,omitempty"`
	Role      *string       `json:"role,omitempty"`
	UserID    *string       `json:"userId,omitempty"`
	Username  *string       `json:"username,omitempty"`
	Type      *string       `json:"type,omitempty"`
	Mfa       *bool         `json:"mfa,omitempty"`
	Policies  []*UserPolicy `json:"policies,omitempty"`
	Tokens    []*Token      `json:"tokens,omitempty"`

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

type UserPolicy struct {
	PolicyId   *string `json:"policyId,omitempty"`
	PolicyName *string `json:"policyName,omitempty"`
	PolicyType *string `json:"policyType,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Token struct {
	Name            *string    `json:"name,omitempty"`
	CreatedAt       *time.Time `json:"createdAt,omitempty"`
	PolicyType      *string    `json:"policyType,omitempty"`
	TokenId         *int       `json:"tokenId,omitempty"`
	TokenLastDigits *string    `json:"tokenLastDigits,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ListUsersInput struct{}

type ListUsersOutput struct {
	Users []*User `json:"users,omitempty"`
}

type CreateUserInput struct {
}

type CreateUserOutput struct {
	User *User `json:"user,omitempty"`
}

type ReadUserInput struct {
	UserID *string `json:"userId,omitempty"`
}

type ReadUserOutput struct {
	User *User `json:"user,omitempty"`
}

type DeleteUserInput struct {
	UserID *string `json:"userId,omitempty"`
}

type DeleteUserOutput struct{}

func userFromJSON(in []byte) (*User, error) {
	b := new(User)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func usersFromJSON(in []byte) ([]*User, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*User, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := userFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func usersFromHttpResponse(resp *http.Response) ([]*User, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return usersFromJSON(body)
}

func (s *ServiceOp) ListUsers(ctx context.Context, input *ListUsersInput) (*ListUsersOutput, error) {
	r := client.NewRequest(http.MethodGet, "/setup/organization/user")
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := usersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListUsersOutput{Users: gs}, nil
}

func (s *ServiceOp) CreateUser(ctx context.Context, input *User, generateToken *bool) (*CreateUserOutput, error) {
	r := client.NewRequest(http.MethodPost, "/setup/user")
	genToken := strconv.FormatBool(spotinst.BoolValue(generateToken))
	r.Params.Set("generateToken", genToken)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ss, err := usersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateUserOutput)
	if len(ss) > 0 {
		output.User = ss[0]
	}

	return output, nil
}

func (s *ServiceOp) ReadUser(ctx context.Context, input *ReadUserInput) (*ReadUserOutput, error) {
	path, err := uritemplates.Expand("/setup/user/{userId}", uritemplates.Values{
		"userId": spotinst.StringValue(input.UserID),
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

	ss, err := usersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadUserOutput)
	if len(ss) > 0 {
		output.User = ss[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteUser(ctx context.Context, input *DeleteUserInput) (*DeleteUserOutput, error) {
	path, err := uritemplates.Expand("/setup/user/{userId}", uritemplates.Values{
		"userId": spotinst.StringValue(input.UserID),
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

	return &DeleteUserOutput{}, nil
}

// region User

func (o User) MarshalJSON() ([]byte, error) {
	type noMethod User
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *User) SetEmail(v *string) *User {
	if o.Email = v; o.Email == nil {
		o.nullFields = append(o.nullFields, "Email")
	}
	return o
}

func (o *User) SetFirstName(v *string) *User {
	if o.FirstName = v; o.FirstName == nil {
		o.nullFields = append(o.nullFields, "FirstName")
	}
	return o
}

func (o *User) SetLastName(v *string) *User {
	if o.LastName = v; o.LastName == nil {
		o.nullFields = append(o.nullFields, "LastName")
	}
	return o
}

func (o *User) SetPassword(v *string) *User {
	if o.Password = v; o.Password == nil {
		o.nullFields = append(o.nullFields, "Password")
	}
	return o
}

func (o *User) SetRole(v *string) *User {
	if o.Role = v; o.Role == nil {
		o.nullFields = append(o.nullFields, "Role")
	}
	return o
}

// endregion
