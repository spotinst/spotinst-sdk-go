package credentials

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

type mockProvider struct {
	creds Value
}

func (m *mockProvider) Retrieve() (Value, error) {
	if m.creds.IsEmpty() {
		return Value{}, errors.New("spotinst: invalid credentials")
	}

	return m.creds, nil
}
func (m *mockProvider) String() string { return "mock" }

func TestChainCredentials(t *testing.T) {
	cases := map[string]struct {
		Providers []Provider
		Expected  Value
		Err       error
	}{
		"no_providers": {
			Providers: []Provider{},
			Err:       errors.New("spotinst: no valid credentials providers in chain"),
		},
		"single_provider": {
			Providers: []Provider{
				&mockProvider{
					creds: Value{
						Token:   "token",
						Account: "account",
					},
				},
			},
			Expected: Value{
				Token:   "token",
				Account: "account",
			},
		},
		"multiple_providers": {
			Providers: []Provider{
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
			Expected: Value{
				Token:   "token1",
				Account: "account1",
			},
		},
		"multiple_providers_first_no_token": {
			Providers: []Provider{
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
			Expected: Value{
				Token:   "token2",
				Account: "account1",
			},
		},
		"multiple_providers_first_no_account": {
			Providers: []Provider{
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
			Expected: Value{
				Token:   "token1",
				Account: "account2",
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			creds, err := NewChainCredentials(tc.Providers...).Get()
			if tc.Err != nil {
				if e, a := tc.Err, err; !reflect.DeepEqual(e, a) {
					t.Errorf("expect: %q to be in: %q", e, a)
				}
				return
			}
			if err != nil {
				t.Fatalf("expect: no error, got: %q", err)
			}
			if e, a := tc.Expected, creds; !reflect.DeepEqual(e, a) {
				t.Errorf("expect: %v, got: %v", e, a)
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

	cases := map[string]struct {
		Filename string
		Profile  string
		Expected Value
		Err      error
	}{
		"file_not_exist": {
			Filename: "file_not_exist",
			Profile:  "default",
			Err:      errors.New("spotinst: failed to load credentials file: open file_not_exist: no such file or directory"),
		},
		"invalid_ini": {
			Filename: filenameINIInvalid,
			Err:      errors.New("spotinst: failed to load credentials file: unclosed section: [profile_nam"),
		},
		"invalid_json": {
			Filename: filenameJSONInvalid,
			Err:      errors.New("spotinst: failed to load credentials file: key-value delimiter not found: {\"token"),
		},
		"profile_not_exist": {
			Filename: filenameINI,
			Profile:  "profile_not_exist",
			Err:      errors.New("spotinst: failed to load credentials file: section 'profile_not_exist' does not exist"),
		},
		"valid_ini_profile_default": {
			Filename: filenameINI,
			Expected: Value{
				ProviderName: FileCredentialsProviderName,
				Token:        "default_token",
			},
		},
		"valid_ini_profile_partial_credentials": {
			Filename: filenameINI,
			Profile:  "partial_credentials",
			Expected: Value{
				ProviderName: FileCredentialsProviderName,
				Token:        "partial_credentials_token",
			},
		},
		"valid_ini_profile_partial_credentials_with_default": {
			Filename: filenameINI,
			Profile:  "partial_credentials_with_default",
			Expected: Value{
				ProviderName: FileCredentialsProviderName,
				Token:        "default_token",
				Account:      "partial_credentials_with_default_account",
			},
		},
		"valid_ini_profile_complete_credentials": {
			Filename: filenameINI,
			Profile:  "complete_credentials",
			Expected: Value{
				ProviderName: FileCredentialsProviderName,
				Token:        "complete_credentials_token",
				Account:      "complete_credentials_account",
			},
		},
		"valid_json": {
			Filename: filenameJSON,
			Expected: Value{
				ProviderName: FileCredentialsProviderName,
				Token:        "token",
				Account:      "account",
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			creds, err := NewFileCredentials(tc.Profile, tc.Filename).Get()
			if tc.Err != nil {
				if e, a := tc.Err, err; !reflect.DeepEqual(e, a) {
					t.Errorf("expect: %q to be in: %q", e, a)
				}
				return
			}
			if err != nil {
				t.Fatalf("expect: no error, got: %q", err)
			}
			if e, a := tc.Expected, creds; !reflect.DeepEqual(e, a) {
				t.Errorf("expect: %v, got: %v", e, a)
			}
		})
	}
}

func TestEnvCredentials(t *testing.T) {
	env := os.Environ()
	defer func() {
		os.Clearenv()

		for _, e := range env {
			p := strings.SplitN(e, "=", 2)
			k, v := p[0], ""
			if len(p) > 1 {
				v = p[1]
			}
			os.Setenv(k, v)
		}
	}()

	cases := map[string]struct {
		Env      map[string]string
		Expected Value
		Err      error
	}{
		"no_variables": {
			Env: map[string]string{},
			Err: errors.New("spotinst: SPOTINST_TOKEN and SPOTINST_ACCOUNT not found in environment"),
		},
		"only_account": {
			Env: map[string]string{
				"SPOTINST_ACCOUNT": "account",
			},
			Err: errors.New("spotinst: token not found in \"EnvCredentialsProvider\""),
		},
		"only_token": {
			Env: map[string]string{
				"SPOTINST_TOKEN": "token",
			},
			Expected: Value{
				ProviderName: EnvCredentialsProviderName,
				Token:        "token",
			},
		},
		"all_variables": {
			Env: map[string]string{
				"SPOTINST_ACCOUNT": "account",
				"SPOTINST_TOKEN":   "token",
			},
			Expected: Value{
				ProviderName: EnvCredentialsProviderName,
				Account:      "account",
				Token:        "token",
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			os.Clearenv()

			for k, v := range tc.Env {
				os.Setenv(k, v)
			}

			creds, err := NewEnvCredentials().Get()
			if tc.Err != nil {
				if e, a := tc.Err, err; !reflect.DeepEqual(e, a) {
					t.Errorf("expect: %q to be in: %q", e, a)
				}
				return
			}
			if err != nil {
				t.Fatalf("expect: no error, got: %q", err)
			}
			if e, a := tc.Expected, creds; !reflect.DeepEqual(e, a) {
				t.Errorf("expect: %v, got: %v", e, a)
			}
		})
	}
}

func TestStaticCredentials(t *testing.T) {
	cases := map[string]struct {
		Token    string
		Account  string
		Expected Value
		Err      error
	}{
		"empty_credentials": {
			Account: "",
			Token:   "",
			Err:     errors.New("spotinst: static credentials are empty"),
		},
		"empty_token": {
			Account: "account",
			Token:   "",
			Err:     errors.New("spotinst: token not found in \"StaticCredentialsProvider\""),
		},
		"empty_account": {
			Token: "token",
			Expected: Value{
				ProviderName: StaticCredentialsProviderName,
				Token:        "token",
			},
		},
		"full_credentials": {
			Account: "account",
			Token:   "token",
			Expected: Value{
				ProviderName: StaticCredentialsProviderName,
				Account:      "account",
				Token:        "token",
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			creds, err := NewStaticCredentials(tc.Token, tc.Account).Get()
			if tc.Err != nil {
				if e, a := tc.Err, err; !reflect.DeepEqual(e, a) {
					t.Errorf("expect: %q to be in: %q", e, a)
				}
				return
			}
			if err != nil {
				t.Fatalf("expect: no error, got: %q", err)
			}
			if e, a := tc.Expected, creds; !reflect.DeepEqual(e, a) {
				t.Errorf("expect: %v, got: %v", e, a)
			}
		})
	}
}
