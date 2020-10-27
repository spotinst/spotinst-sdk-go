package credentials

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/spotinst/spotinst-sdk-go/spotinst/featureflag"
)

type mockProvider struct {
	creds Value
}

func (m *mockProvider) Retrieve() (Value, error) {
	if m.creds.IsEmpty() {
		return m.creds, errors.New("spotinst: invalid credentials")
	}

	return m.creds, nil
}

func (m *mockProvider) String() string { return "mock" }

func TestChainCredentials(t *testing.T) {
	tests := map[string]struct {
		providers []Provider
		features  string
		want      Value
		err       error
	}{
		"no_providers": {
			providers: []Provider{},
			err:       ErrNoValidProvidersFoundInChain,
		},
		"single_provider_valid": {
			providers: []Provider{
				&mockProvider{
					creds: Value{
						Token:   "token",
						Account: "account",
					},
				},
			},
			want: Value{
				Token:   "token",
				Account: "account",
			},
		},
		"single_provider_invalid": {
			providers: []Provider{
				&mockProvider{
					creds: Value{
						Token:   "",
						Account: "",
					},
				},
			},
			err: errorList{
				errors.New("spotinst: invalid credentials"),
			},
		},
		"multiple_providers_valid": {
			providers: []Provider{
				&mockProvider{
					creds: Value{
						Token:   "token1",
						Account: "account1",
					},
				},
				&mockProvider{
					creds: Value{
						Token:   "token2",
						Account: "account2",
					},
				},
			},
			want: Value{
				Token:   "token1",
				Account: "account1",
			},
		},
		"multiple_providers_invalid": {
			providers: []Provider{
				&mockProvider{
					creds: Value{
						Token:   "",
						Account: "",
					},
				},
				&mockProvider{
					creds: Value{
						Token:   "",
						Account: "",
					},
				},
			},
			err: errorList{
				errors.New("spotinst: invalid credentials"),
				errors.New("spotinst: invalid credentials"),
			},
		},
		"partial_providers_first_no_token": {
			providers: []Provider{
				&mockProvider{
					creds: Value{
						Token:   "",
						Account: "account1",
					},
				},
				&mockProvider{
					creds: Value{
						Token:   "token2",
						Account: "account2",
					},
				},
			},
			features: "MergeCredentialsChain=false",
			err: errorList{
				ErrNoValidProvidersFoundInChain,
			},
		},
		"partial_providers_first_no_account": {
			providers: []Provider{
				&mockProvider{
					creds: Value{
						Token:   "token1",
						Account: "",
					},
				},
				&mockProvider{
					creds: Value{
						Token:   "token2",
						Account: "account2",
					},
				},
			},
			features: "MergeCredentialsChain=false",
			want: Value{
				Token:   "token1",
				Account: "",
			},
		},
		"partial_providers_first_no_token_with_merge": {
			providers: []Provider{
				&mockProvider{
					creds: Value{
						Token:   "",
						Account: "account1",
					},
				},
				&mockProvider{
					creds: Value{
						Token:   "token2",
						Account: "account2",
					},
				},
			},
			features: "MergeCredentialsChain=true",
			want: Value{
				Token:   "token2",
				Account: "account1",
			},
		},
		"partial_providers_first_no_account_with_merge": {
			providers: []Provider{
				&mockProvider{
					creds: Value{
						Token:   "token1",
						Account: "",
					},
				},
				&mockProvider{
					creds: Value{
						Token:   "token2",
						Account: "account2",
					},
				},
			},
			features: "MergeCredentialsChain=true",
			want: Value{
				Token:   "token1",
				Account: "account2",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.features != "" {
				origFlags := featureflag.All()
				defer func() { featureflag.Set(origFlags.String()) }() // restore
				featureflag.Set(test.features)
			}
			creds, err := NewChainCredentials(test.providers...).Get()
			if err != nil {
				if test.err != nil {
					if !strings.Contains(err.Error(), test.err.Error()) {
						t.Fatalf("want: %v, got: %v", test.err, err)
					}
					return // want failure
				} else {
					t.Fatalf("want: nil, got: %v", err)
				}
			}
			if e, a := test.want, creds; !reflect.DeepEqual(e, a) {
				t.Errorf("want: %v, got: %v", e, a)
			}
		})
	}
}

func TestFileCredentials(t *testing.T) {
	var (
		filenameINI         = filepath.Join("testdata", "credentials_ini")
		filenameJSON        = filepath.Join("testdata", "credentials_json")
		filenameINIInvalid  = filepath.Join("testdata", "credentials_ini_invalid")
		filenameJSONInvalid = filepath.Join("testdata", "credentials_json_invalid")
	)

	tests := map[string]struct {
		filename string
		profile  string
		want     Value
		err      error
	}{
		"file_not_exist": {
			filename: "file_not_exist",
			profile:  "default",
			err:      errors.New("spotinst: failed to load credentials file: open file_not_exist: no such file or directory"),
		},
		"invalid_ini": {
			filename: filenameINIInvalid,
			err:      errors.New("spotinst: failed to load credentials file: unclosed section: [profile_nam"),
		},
		"invalid_json": {
			filename: filenameJSONInvalid,
			err:      errors.New("spotinst: failed to load credentials file: key-value delimiter not found: {\"token"),
		},
		"profile_not_exist": {
			filename: filenameINI,
			profile:  "profile_not_exist",
			err:      errors.New("spotinst: failed to load credentials file: section \"profile_not_exist\" does not exist"),
		},
		"valid_ini_profile_default": {
			filename: filenameINI,
			want: Value{
				ProviderName: FileCredentialsProviderName,
				Token:        "default_token",
			},
		},
		"valid_ini_profile_partial_credentials": {
			filename: filenameINI,
			profile:  "partial_credentials",
			want: Value{
				ProviderName: FileCredentialsProviderName,
				Token:        "partial_credentials_token",
			},
		},
		"valid_ini_profile_partial_credentials_with_default": {
			filename: filenameINI,
			profile:  "partial_credentials_with_default",
			want: Value{
				ProviderName: FileCredentialsProviderName,
				Token:        "default_token",
				Account:      "partial_credentials_with_default_account",
			},
		},
		"valid_ini_profile_complete_credentials": {
			filename: filenameINI,
			profile:  "complete_credentials",
			want: Value{
				ProviderName: FileCredentialsProviderName,
				Token:        "complete_credentials_token",
				Account:      "complete_credentials_account",
			},
		},
		"valid_json": {
			filename: filenameJSON,
			want: Value{
				ProviderName: FileCredentialsProviderName,
				Token:        "token",
				Account:      "account",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			creds, err := NewFileCredentials(test.profile, test.filename).Get()
			if err != nil {
				if test.err != nil {
					if !strings.Contains(err.Error(), test.err.Error()) {
						t.Fatalf("want: %v, got: %v", test.err, err)
					}
					return // want failure
				} else {
					t.Fatalf("want: nil, got: %v", err)
				}
			}
			if e, a := test.want, creds; !reflect.DeepEqual(e, a) {
				t.Errorf("want: %v, got: %v", e, a)
			}
		})
	}
}

func TestEnvCredentials(t *testing.T) {
	origEnv := os.Environ()
	defer func() { // restore env
		os.Clearenv()

		for _, kv := range origEnv {
			p := strings.SplitN(kv, "=", 2)
			k, v := p[0], ""
			if len(p) > 1 {
				v = p[1]
			}
			os.Setenv(k, v)
		}
	}()

	tests := map[string]struct {
		env  map[string]string
		want Value
		err  error
	}{
		"no_variables": {
			env: map[string]string{},
			err: errors.New("spotinst: SPOTINST_TOKEN and SPOTINST_ACCOUNT not found in environment"),
		},
		"only_account": {
			env: map[string]string{
				"SPOTINST_ACCOUNT": "account",
			},
			err: ErrNoValidTokenFound,
		},
		"only_token": {
			env: map[string]string{
				"SPOTINST_TOKEN": "token",
			},
			want: Value{
				ProviderName: EnvCredentialsProviderName,
				Token:        "token",
			},
		},
		"all_variables": {
			env: map[string]string{
				"SPOTINST_ACCOUNT": "account",
				"SPOTINST_TOKEN":   "token",
			},
			want: Value{
				ProviderName: EnvCredentialsProviderName,
				Account:      "account",
				Token:        "token",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			os.Clearenv()
			for k, v := range test.env {
				os.Setenv(k, v)
			}

			creds, err := NewEnvCredentials().Get()
			if err != nil {
				if test.err != nil {
					if !strings.Contains(err.Error(), test.err.Error()) {
						t.Fatalf("want: %v, got: %v", test.err, err)
					}
					return // want failure
				} else {
					t.Fatalf("want: nil, got: %v", err)
				}
			}
			if e, a := test.want, creds; !reflect.DeepEqual(e, a) {
				t.Errorf("want: %v, got: %v", e, a)
			}
		})
	}
}

func TestStaticCredentials(t *testing.T) {
	tests := map[string]struct {
		token   string
		account string
		want    Value
		err     error
	}{
		"empty_credentials": {
			account: "",
			token:   "",
			err:     errors.New("spotinst: static credentials are empty"),
		},
		"empty_token": {
			account: "account",
			token:   "",
			err:     ErrNoValidTokenFound,
		},
		"empty_account": {
			token: "token",
			want: Value{
				ProviderName: StaticCredentialsProviderName,
				Token:        "token",
			},
		},
		"full_credentials": {
			account: "account",
			token:   "token",
			want: Value{
				ProviderName: StaticCredentialsProviderName,
				Account:      "account",
				Token:        "token",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			creds, err := NewStaticCredentials(test.token, test.account).Get()
			if err != nil {
				if test.err != nil {
					if !strings.Contains(err.Error(), test.err.Error()) {
						t.Fatalf("want: %v, got: %v", test.err, err)
					}
					return // want failure
				} else {
					t.Fatalf("want: nil, got: %v", err)
				}
			}
			if e, a := test.want, creds; !reflect.DeepEqual(e, a) {
				t.Errorf("want: %v, got: %v", e, a)
			}
		})
	}
}
