package administration

import (
	"context"
	"encoding/json"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/jsonutil"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/uritemplates"
	"io/ioutil"
	"net/http"
	"strconv"
)

type User struct {
	Email               *string       `json:"email,omitempty"`
	FirstName           *string       `json:"firstName,omitempty"`
	LastName            *string       `json:"lastName,omitempty"`
	Password            *string       `json:"password,omitempty"`
	Role                *string       `json:"role,omitempty"`
	UserID              *string       `json:"userId,omitempty"`
	Username            *string       `json:"username,omitempty"`
	Type                *string       `json:"type,omitempty"`
	Mfa                 *bool         `json:"mfa,omitempty"`
	Policies            []*UserPolicy `json:"policies,omitempty"`
	Tokens              []*Token      `json:"tokens,omitempty"`
	PersonalAccessToken *string       `json:"personalAccessToken,omitempty"`
	Id                  *int          `json:"id,omitempty"`
	GroupNames          []*string     `json:"groupNames,omitempty"`
	Groups              []*Group      `json:"groups,omitempty"`

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
	PolicyId   *string   `json:"policyId,omitempty"`
	PolicyName *string   `json:"policyName,omitempty"`
	PolicyType *string   `json:"policyType,omitempty"`
	AccountIds []*string `json:"accountIds,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ProgrammaticUser struct {
	Name        *string       `json:"name,omitempty"`
	Description *string       `json:"description,omitempty"`
	Policies    []*ProgPolicy `json:"policies,omitempty"`
	Accounts    []*Account    `json:"accounts,omitempty"`
	Token       *string       `json:"token,omitempty"`
	ProgUserId  *string       `json:"id,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ProgPolicy struct {
	PolicyId   *string  `json:"policyId,omitempty"`
	AccountIds []string `json:"accountIds,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Account struct {
	Id   *string `json:"id,omitempty"`
	Role *string `json:"role,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Token struct {
	Name            *string `json:"name,omitempty"`
	CreatedAt       *string `json:"createdAt,omitempty"`
	TokenId         *int    `json:"tokenId,omitempty"`
	TokenLastDigits *string `json:"tokenLastDigits,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Group struct {
	Id          *string   `json:"id,omitempty"`
	Name        *string   `json:"name,omitempty"`
	PolicyNames []*string `json:"policyNames,omitempty"`

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

type CreateProgrammaticUserInput struct {
}

type CreateProgrammaticUserOutput struct {
	ProgrammaticUser *ProgrammaticUser `json:"user,omitempty"`
}

type ReadUserInput struct {
	UserID *string `json:"userId,omitempty"`
}

type ReadUserOutput struct {
	User *User `json:"user,omitempty"`
}

type ReadProgUserOutput struct {
	ProgUser *ProgrammaticUser `json:"user,omitempty"`
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

func progUserFromJSON(in []byte) (*ProgrammaticUser, error) {
	b := new(ProgrammaticUser)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func progUsersFromJSON(in []byte) ([]*ProgrammaticUser, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*ProgrammaticUser, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := progUserFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func progUsersFromHttpResponse(resp *http.Response) ([]*ProgrammaticUser, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return progUsersFromJSON(body)
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

func (s *ServiceOp) CreateProgUser(ctx context.Context, input *ProgrammaticUser) (*CreateProgrammaticUserOutput, error) {
	r := client.NewRequest(http.MethodPost, "/setup/user/programmatic")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ss, err := progUsersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateProgrammaticUserOutput)
	if len(ss) > 0 {
		output.ProgrammaticUser = ss[0]
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

func (s *ServiceOp) ReadProgUser(ctx context.Context, input *ReadUserInput) (*ReadProgUserOutput, error) {
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

	ss, err := progUsersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadProgUserOutput)
	if len(ss) > 0 {
		output.ProgUser = ss[0]
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

// region ProgrammaticUser

func (o ProgrammaticUser) MarshalJSON() ([]byte, error) {
	type noMethod ProgrammaticUser
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *ProgrammaticUser) SetName(v *string) *ProgrammaticUser {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *ProgrammaticUser) SetDescription(v *string) *ProgrammaticUser {
	if o.Description = v; o.Description == nil {
		o.nullFields = append(o.nullFields, "Description")
	}
	return o
}

func (o *ProgrammaticUser) SetPolicies(v []*ProgPolicy) *ProgrammaticUser {
	if o.Policies = v; o.Policies == nil {
		o.nullFields = append(o.nullFields, "Policies")
	}
	return o
}

func (o *ProgrammaticUser) SetAccounts(v []*Account) *ProgrammaticUser {
	if o.Accounts = v; o.Accounts == nil {
		o.nullFields = append(o.nullFields, "Accounts")
	}
	return o
}

// endregion

func (o ProgPolicy) MarshalJSON() ([]byte, error) {
	type noMethod ProgPolicy
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *ProgPolicy) SetAccountIds(v []string) *ProgPolicy {
	if o.AccountIds = v; o.AccountIds == nil {
		o.nullFields = append(o.nullFields, "AccountIds")
	}
	return o
}

func (o *ProgPolicy) SetPolicyId(v *string) *ProgPolicy {
	if o.PolicyId = v; o.PolicyId == nil {
		o.nullFields = append(o.nullFields, "PolicyId")
	}
	return o
}

//end region

func (o Account) MarshalJSON() ([]byte, error) {
	type noMethod Account
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Account) SetAccountId(v *string) *Account {
	if o.Id = v; o.Id == nil {
		o.nullFields = append(o.nullFields, "Id")
	}
	return o
}

func (o *Account) SetRole(v *string) *Account {
	if o.Role = v; o.Role == nil {
		o.nullFields = append(o.nullFields, "Role")
	}
	return o
}
