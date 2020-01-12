package useragent

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	cases := map[string]struct {
		Product  string
		Version  string
		Comment  []string
		Expected UserAgent
	}{
		"simple": {
			Product: "spotinst-sdk-go",
			Version: "1.0.0",
			Expected: UserAgent{
				Product: "spotinst-sdk-go",
				Version: "1.0.0",
				Comment: nil,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			ua := New(tc.Product, tc.Version)
			if e, a := tc.Expected, ua; !reflect.DeepEqual(e, a) {
				t.Errorf("expect: %q, got: %q", e, a)
			}
		})
	}
}

func TestStringerSingle(t *testing.T) {
	cases := map[string]struct {
		UserAgent UserAgent
		Expected  string
	}{
		"simple": {
			UserAgent: UserAgent{
				Product: "spotinst-sdk-go",
				Version: "1.0.0",
			},
			Expected: "spotinst-sdk-go/1.0.0",
		},
		"comment": {
			UserAgent: UserAgent{
				Product: "spotinst-sdk-go",
				Version: "1.0.0",
				Comment: []string{"test"},
			},
			Expected: "spotinst-sdk-go/1.0.0 (test)",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if e, a := tc.Expected, tc.UserAgent.String(); e != a {
				t.Errorf("expect: %q, got: %q", e, a)
			}
		})
	}
}

func TestStringerMulti(t *testing.T) {
	cases := map[string]struct {
		UserAgents UserAgents
		Expected   string
	}{
		"single": {
			UserAgents: UserAgents{
				{
					Product: "spotinst-sdk-go",
					Version: "1.0.0",
				},
			},
			Expected: "spotinst-sdk-go/1.0.0",
		},
		"multi": {
			UserAgents: UserAgents{
				{
					Product: "spotinst-sdk-go",
					Version: "1.2.3",
				},
				{
					Product: "spotctl",
					Version: "4.5.6",
					Comment: []string{"test"},
				},
			},
			Expected: "spotinst-sdk-go/1.2.3 spotctl/4.5.6 (test)",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if e, a := tc.Expected, tc.UserAgents.String(); e != a {
				t.Errorf("expect: %q, got: %q", e, a)
			}
		})
	}
}
